use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ActiveModelTrait, ColumnTrait, DatabaseConnection, DbErr, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, sea_query};
use sea_orm::ActiveValue::Set;

use domain::{public, public_service};
use domain::public::{BooleanScope, ECode, Operation};
use model::label::{ActiveModel, Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn decode(&self, dto: public_service::Label) -> Model {
        let current_time = common::utils::current_timestamp();
        Model {
            id: 0,
            source: dto.source,
            name: dto.name,
            creator: self.header.operator,
            updater: self.header.operator,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        }
    }

    fn encode(&self, model: Model) -> public_service::Label {
        public_service::Label {
            id: model.id,
            source: model.source,
            name: model.name,
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
        filter: Option<public_service::LabelFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> public_service::ListLabelReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            query = match filter.deleted() {
                BooleanScope::BoolAll => query,
                BooleanScope::BoolFalse => query.filter(Column::DeleteTime.gt(0)),
                BooleanScope::BoolTrue => query.filter(Column::DeleteTime.eq(0)),
            };
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.names.is_empty().not() {
                query = query.filter(Column::Name.is_in(filter.names))
            }
            if filter.sources.is_empty().not() {
                query = query.filter(Column::Source.is_in(filter.sources))
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

            if let Some(keywords) = filter.keywords {
                if keywords.name.is_empty().not() {
                    query = query.filter(
                        Column::Name.contains(&keywords.name),
                    )
                }
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
                    ECode::PublicServiceErrorListLabel,
                    format!("list label failed: {}", e),
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

        public_service::ListLabelReply { header, pager: Some(pager), data: dtos }
    }

    pub async fn operate(&self, operate: Operation, data: Option<public_service::Label>) -> public_service::OperateLabelReply {
        let (header, data) = match data {
            None => {
                (common::header::err_reply(
                    self.header.trace_id,
                    ECode::PublicServiceErrorOperateLabel,
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
                            ECode::PublicServiceErrorOperateLabel,
                            format!("{} access failed: {}", operate.as_str_name(), e),
                        );
                        (header, None)
                    }
                }
            }
        };
        public_service::OperateLabelReply { header, data }
    }

    pub async fn batch_operate(&self, operate: Operation, data: Vec<public_service::Label>) -> public_service::BatchOperateLabelReply {
        let header = if data.is_empty() {
            common::header::err_reply(
                self.header.trace_id,
                ECode::PublicServiceErrorOperateLabel,
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
                    ECode::PublicServiceErrorBatchOperateLabel,
                    format!("batch {} access failed: {}", operate.as_str_name(), e),
                )
            }
        };
        public_service::BatchOperateLabelReply { header }
    }


    // 插入
    async fn insert(&self, dto: public_service::Label) -> Result<Option<public_service::Label>, DbErr> {
        let active: ActiveModel = self.decode(dto).into();
        let model = active.insert(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 更新
    async fn update(&self, dto: public_service::Label) -> Result<Option<public_service::Label>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.name = Set(dto.name);
        active.updater = Set(self.header.operator);
        active.update_time = Set(common::utils::current_timestamp());
        let model = active.update(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 逻辑删除
    async fn delete(&self, id: i32) -> Result<Option<public_service::Label>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.delete_time = Set(common::utils::current_timestamp());
        active.update(self.db).await?;
        Ok(None)
    }

    // 物理删除
    async fn delete_real(&self, id: i32) -> Result<Option<public_service::Label>, DbErr> {
        Entity::delete_by_id(id).exec(self.db).await?;
        Ok(None)
    }

    // 保存并更新
    async fn save(&self, mut data: public_service::Label) -> Result<Option<public_service::Label>, DbErr> {
        let active: ActiveModel = self.decode(data.clone()).into();
        let result = Entity::insert(active).on_conflict(
            sea_query::OnConflict::columns([Column::Name, Column::Source])
                .update_columns([Column::Name, Column::Updater, Column::UpdateTime])
                .to_owned()
        ).exec(self.db).await?;
        data.id = result.last_insert_id;
        Ok(Some(data))
    }

    // 批量写入，有冲突则更新
    async fn batch_save(
        &self,
        dtos: Vec<public_service::Label>,
    ) -> Result<Option<i32>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let active: ActiveModel = self.decode(dto).into();
            actives.push(active);
        }
        Entity::insert_many(actives).on_conflict(
            // 定义冲突
            sea_query::OnConflict::columns([Column::Name, Column::Source])
                // 冲突后更新的字段
                .update_columns([Column::Name]).to_owned(),
        ).exec(self.db).await?;
        Ok(None)
    }

    // 批量删除
    async fn batch_delete(&self, dtos: Vec<public_service::Label>) -> Result<Option<i32>, DbErr> {
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
