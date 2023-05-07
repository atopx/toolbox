use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ActiveModelTrait, ColumnTrait, DatabaseConnection, DbErr, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, sea_query};
use sea_orm::ActiveValue::Set;

use domain::{novel_service, public};
use domain::public::{BooleanScope, ECode, Operation};
use model::book::{ActiveModel, Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn decode(&self, dto: novel_service::Book) -> Model {
        let current_time = common::utils::current_timestamp();
        Model {
            id: 0,
            src: dto.src,
            url: dto.url,
            name: dto.name,
            author: dto.author,
            status: dto.status,
            message: dto.message,
            creator: self.header.operator,
            updater: self.header.operator,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        }
    }

    fn encode(&self, model: Model) -> novel_service::Book {
        novel_service::Book {
            id: model.id,
            src: model.src,
            url: model.url,
            name: model.name,
            author: model.author,
            status: model.status,
            message: model.message,
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
        filter: Option<novel_service::BookFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> novel_service::ListBookReply {
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
            if filter.names.is_empty().not() {
                query = query.filter(Column::Name.is_in(filter.names))
            }
            if filter.authors.is_empty().not() {
                query = query.filter(Column::Author.is_in(filter.authors))
            }
            if filter.states.is_empty().not() {
                query = query.filter(Column::Author.is_in(filter.states))
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
                if keywords.keyword.is_empty().not() {
                    query = query.filter(
                        Column::Name.contains(&keywords.keyword).or(
                            Column::Author.contains(&keywords.keyword)
                        )
                    )
                }
                if keywords.name.is_empty().not() {
                    query = query.filter(Column::Name.contains(&keywords.name))
                }
                if keywords.author.is_empty().not() {
                    query = query.filter(Column::Author.contains(&keywords.author))
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
                    ECode::NovelServiceErrorListBook,
                    format!("list book failed: {}", e),
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

        novel_service::ListBookReply { header, pager: Some(pager), data: dtos }
    }

    pub async fn operate(&self, operate: Operation, data: Option<novel_service::Book>) -> novel_service::OperateBookReply {
        let (header, data) = match data {
            None => {
                (common::header::err_reply(
                    self.header.trace_id,
                    ECode::NovelServiceErrorOperateBook,
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
                            ECode::NovelServiceErrorOperateBook,
                            format!("{} book failed: {}", operate.as_str_name(), e),
                        );
                        (header, None)
                    }
                }
            }
        };
        novel_service::OperateBookReply { header, data }
    }

    pub async fn batch_operate(&self, operate: Operation, data: Vec<novel_service::Book>) -> novel_service::BatchOperateBookReply {
        let header = if data.is_empty() {
            common::header::err_reply(
                self.header.trace_id,
                ECode::NovelServiceErrorOperateBook,
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
                    ECode::NovelServiceErrorBatchOperateBook,
                    format!("batch {} book failed: {}", operate.as_str_name(), e),
                )
            }
        };
        novel_service::BatchOperateBookReply { header }
    }


    // 插入
    async fn insert(&self, dto: novel_service::Book) -> Result<Option<novel_service::Book>, DbErr> {
        let active: ActiveModel = self.decode(dto).into();
        let model = active.insert(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 更新
    async fn update(&self, dto: novel_service::Book) -> Result<Option<novel_service::Book>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.name = Set(dto.name);
        active.author = Set(dto.author);
        active.src = Set(dto.src);
        active.url = Set(dto.url);
        active.status = Set(dto.status);
        active.message = Set(dto.message);
        active.updater = Set(self.header.operator);
        active.update_time = Set(common::utils::current_timestamp());
        let model = active.update(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 逻辑删除
    async fn delete(&self, id: i32) -> Result<Option<novel_service::Book>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.delete_time = Set(common::utils::current_timestamp());
        active.update(self.db).await?;
        Ok(None)
    }

    // 物理删除
    async fn delete_real(&self, id: i32) -> Result<Option<novel_service::Book>, DbErr> {
        Entity::delete_by_id(id).exec(self.db).await?;
        Ok(None)
    }

    // 保存并更新
    async fn save(&self, mut data: novel_service::Book) -> Result<Option<novel_service::Book>, DbErr> {
        let active: ActiveModel = self.decode(data.clone()).into();
        let result = Entity::insert(active).on_conflict(
            sea_query::OnConflict::column(Column::Src)
                .update_columns(
                    [
                        Column::Url,
                        Column::Name,
                        Column::Author,
                        Column::Status,
                        Column::Message,
                        Column::Updater,
                        Column::UpdateTime
                    ],
                )
                .to_owned()
        ).exec(self.db).await?;
        data.id = result.last_insert_id;
        Ok(Some(data))
    }

    // 批量写入，有冲突则更新
    async fn batch_save(
        &self,
        dtos: Vec<novel_service::Book>,
    ) -> Result<Option<i32>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let active: ActiveModel = self.decode(dto).into();
            actives.push(active);
        }
        Entity::insert_many(actives).on_conflict(
            sea_query::OnConflict::column(Column::Src)
                .update_columns(
                    [
                        Column::Url,
                        Column::Name,
                        Column::Author,
                        Column::Status,
                        Column::Message,
                        Column::Updater,
                        Column::UpdateTime
                    ],
                )
                .to_owned()
        ).exec(self.db).await?;
        Ok(None)
    }

    // 批量删除
    async fn batch_delete(&self, dtos: Vec<novel_service::Book>) -> Result<Option<i32>, DbErr> {
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
