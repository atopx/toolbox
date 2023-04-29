use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ActiveModelTrait, ColumnTrait, DatabaseConnection, DbErr, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, sea_query};
use sea_orm::ActiveValue::Set;

use domain::{note_service, public};
use domain::public::{BooleanScope, ECode, Operation};
use model::note_label::{ActiveModel, Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn decode(&self, dto: note_service::NoteLabel) -> Model {
        let current_time = common::utils::current_timestamp();
        Model {
            id: 0,
            note_id: dto.note_id,
            label_id: dto.label_id,
            creator: self.header.operator,
            updater: self.header.operator,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        }
    }

    fn encode(&self, model: Model) -> note_service::NoteLabel {
        note_service::NoteLabel {
            id: model.id,
            note_id: model.note_id,
            label_id: model.label_id,
            creator: self.header.operator,
            updater: self.header.operator,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
        }
    }

    // 列表筛选
    pub async fn list(
        &self,
        filter: Option<note_service::NoteLabelFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> note_service::ListNoteLabelReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            query = match filter.deleted() {
                BooleanScope::BoolAll => query,
                BooleanScope::BoolFalse => query.filter(Column::DeleteTime.eq(0)),
                BooleanScope::BoolTrue => query.filter(Column::DeleteTime.gt(0)),
            };
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.label_ids.is_empty().not() {
                query = query.filter(Column::LabelId.is_in(filter.label_ids))
            }
            if filter.note_ids.is_empty().not() {
                query = query.filter(Column::NoteId.is_in(filter.note_ids))
            }
            if filter.creators.is_empty().not() {
                query = query.filter(Column::Creator.is_in(filter.creators));
            }
            if filter.updaters.is_empty().not() {
                query = query.filter(Column::Updater.is_in(filter.updaters));
            }
            if let Some(range) = filter.create_time_range {
                query = query.filter(Column::CreateTime.between(range.left, range.right))
            }
            if let Some(range) = filter.update_time_range {
                query = query.filter(Column::UpdateTime.between(range.left, range.right))
            }
            if let Some(range) = filter.delete_time_range {
                query = query.filter(Column::DeleteTime.between(range.left, range.right))
            }
        }

        // 排序
        if sorts.is_empty() {
            // add a default sort value
            sorts.push(public::Sort {
                field: "id".to_string(),
                direction: public::SortDirection::SortAsc.into(),
            });
        }
        for sort in sorts {
            if let Ok(col) = Column::from_str(sort.field.as_str()) {
                query = match sort.direction() {
                    public::SortDirection::SortAsc => query.order_by_asc(col),
                    public::SortDirection::SortDesc => query.order_by_desc(col),
                }
            }
        }
        let mut pager = pager.unwrap_or(public::Pager { disabled: true, ..Default::default() });


        let mut dtos = vec![];

        let header = if pager.disabled {
            // 不分页
            match query.all(self.db).await {
                Ok(models) => {
                    for model in models {
                        dtos.push(self.encode(model));
                    };
                    common::header::reply(self.header.trace_id)
                }
                Err(e) => common::header::err_reply(
                    self.header.trace_id,
                    ECode::NoteServiceErrorListNoteLabel,
                    format!("list note_label failed: {}", e),
                ),
            }
        } else {
            // 分页
            let paginator = query.paginate(self.db, pager.size);
            pager.count = paginator.num_items().await.unwrap_or_default();
            for model in paginator.fetch_page(pager.index - 1).await.unwrap_or_default() {
                dtos.push(self.encode(model))
            };
            common::header::reply(self.header.trace_id)
        };

        note_service::ListNoteLabelReply { header, pager: Some(pager), data: dtos }
    }

    pub async fn operate(&self, operate: Operation, data: Option<note_service::NoteLabel>) -> note_service::OperateNoteLabelReply {
        let (header, data) = match data {
            None => {
                (common::header::err_reply(
                    self.header.trace_id,
                    ECode::NoteServiceErrorOperateNoteLabel,
                    "missing operate data".to_string(),
                ), None)
            }
            Some(data) => {
                let result = match operate {
                    Operation::Create => self.insert(data).await,
                    Operation::Update => self.update(data).await,
                    Operation::Upsert => self.save(data).await,
                    Operation::Delete => self.delete(data.id).await,
                    Operation::RealDelete => self.delete_real(data.id).await,
                };
                match result {
                    Ok(data) => {
                        let header = common::header::reply(self.header.trace_id);
                        (header, data)
                    }
                    Err(e) => {
                        let header = common::header::err_reply(
                            self.header.trace_id,
                            ECode::NoteServiceErrorOperateNoteLabel,
                            format!("{} access failed: {}", operate.as_str_name(), e),
                        );
                        (header, None)
                    }
                }
            }
        };
        note_service::OperateNoteLabelReply { header, data }
    }

    pub async fn batch_operate(&self, operate: Operation, data: Vec<note_service::NoteLabel>) -> note_service::BatchOperateNoteLabelReply {
        let header = if data.is_empty() {
            common::header::err_reply(
                self.header.trace_id,
                ECode::NoteServiceErrorOperateNoteLabel,
                "missing operate data".to_string(),
            )
        } else {
            let result = match operate {
                Operation::Upsert => self.batch_save(data).await,
                Operation::Delete => self.batch_delete(data).await,
                _ => Ok(None)
            };
            match result {
                Ok(_) => common::header::reply(self.header.trace_id),
                Err(e) => common::header::err_reply(
                    self.header.trace_id,
                    ECode::NoteServiceErrorBatchOperateNoteLabel,
                    format!("batch {} access failed: {}", operate.as_str_name(), e),
                )
            }
        };
        note_service::BatchOperateNoteLabelReply { header }
    }


    // 插入
    async fn insert(&self, dto: note_service::NoteLabel) -> Result<Option<note_service::NoteLabel>, DbErr> {
        let active: ActiveModel = self.decode(dto).into();
        let model = active.insert(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 更新
    async fn update(&self, dto: note_service::NoteLabel) -> Result<Option<note_service::NoteLabel>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.note_id = Set(dto.note_id);
        active.label_id = Set(dto.label_id);
        active.updater = Set(self.header.operator);
        active.update_time = Set(common::utils::current_timestamp());
        let model = active.update(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 逻辑删除
    async fn delete(&self, id: i32) -> Result<Option<note_service::NoteLabel>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.delete_time = Set(common::utils::current_timestamp());
        active.update(self.db).await?;
        Ok(None)
    }

    // 物理删除
    async fn delete_real(&self, id: i32) -> Result<Option<note_service::NoteLabel>, DbErr> {
        Entity::delete_by_id(id).exec(self.db).await?;
        Ok(None)
    }

    // 保存并更新
    async fn save(&self, mut data: note_service::NoteLabel) -> Result<Option<note_service::NoteLabel>, DbErr> {
        let active: ActiveModel = self.decode(data.clone()).into();
        let result = Entity::insert(active).on_conflict(
            sea_query::OnConflict::columns([Column::LabelId, Column::NoteId])
                .do_nothing().to_owned()
        ).exec(self.db).await?;
        data.id = result.last_insert_id;
        Ok(Some(data))
    }

    // 批量写入，有冲突则更新
    async fn batch_save(
        &self,
        dtos: Vec<note_service::NoteLabel>,
    ) -> Result<Option<i32>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let active: ActiveModel = self.decode(dto).into();
            actives.push(active);
        }
        Entity::insert_many(actives).on_conflict(
            sea_query::OnConflict::columns([Column::LabelId, Column::NoteId])
                .do_nothing().to_owned()
        ).exec(self.db).await?;
        Ok(None)
    }

    // 批量删除
    async fn batch_delete(&self, dtos: Vec<note_service::NoteLabel>) -> Result<Option<i32>, DbErr> {
        let mut ids = vec![];
        for dto in dtos {
            ids.push(dto.id);
        }
        let active = ActiveModel {
            delete_time: Set(common::utils::current_timestamp()),
            ..Default::default()
        };
        Entity::update_many()
            .set(active)
            .filter(Column::Id.is_in(ids))
            .exec(self.db)
            .await?;
        Ok(None)
    }
}
