use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ActiveModelTrait, ColumnTrait, DatabaseConnection, DbErr, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder, sea_query};
use sea_orm::ActiveValue::Set;

use domain::{public, auth_service};
use domain::public::{ECode, Operation};
use model::access::{ActiveModel, Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn decode(&self, dto: auth_service::Access) -> Model {
        Model {
            id: 0,
            status: dto.status,
            path: dto.path,
            method: dto.method,
        }
    }

    fn encode(&self, model: Model) -> auth_service::Access {
        auth_service::Access {
            id: model.id,
            path: model.path,
            method: model.method,
            status: model.status,

        }
    }

    // 列表筛选
    pub async fn list(
        &self,
        filter: Option<auth_service::AccessFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> auth_service::ListAccessReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.paths.is_empty().not() {
                query = query.filter(Column::Path.is_in(filter.paths));
            }
            if filter.methods.is_empty().not() {
                query = query.filter(Column::Method.is_in(filter.methods));
            }
            if filter.states.is_empty().not() {
                query = query.filter(Column::Status.is_in(filter.states));
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
                    ECode::AuthServiceErrorListAccess,
                    format!("list access failed: {}", e),
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

        auth_service::ListAccessReply { header, pager: Some(pager), data: dtos }
    }

    pub async fn operate(&self, operate: Operation, data: Option<auth_service::Access>) -> auth_service::OperateAccessReply {
        let (header, data) = match data {
            None => {
                (common::header::err_reply(
                    self.header.trace_id,
                    ECode::AuthServiceErrorOperateAccess,
                    "missing operate data".to_string(),
                ), None)
            }
            Some(data) => {
                let result = match operate {
                    Operation::Create => self.insert(data).await,
                    Operation::Update => self.update(data).await,
                    Operation::Upsert => self.save(data).await,
                    Operation::Delete => self.delete_real(data.id).await,
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
                            ECode::AuthServiceErrorOperateAccess,
                            format!("{} access failed: {}", operate.as_str_name(), e),
                        );
                        (header, None)
                    }
                }
            }
        };
        auth_service::OperateAccessReply { header, data }
    }

    pub async fn batch_operate(&self, operate: Operation, data: Vec<auth_service::Access>) -> auth_service::BatchOperateAccessReply {
        let header = if data.is_empty() {
            common::header::err_reply(
                self.header.trace_id,
                ECode::AuthServiceErrorOperateAccess,
                "missing operate data".to_string(),
            )
        } else {
            let result = match operate {
                Operation::Upsert => self.batch_save(data).await,
                _ => Ok(None)
            };
            match result {
                Ok(_) => common::header::reply(self.header.trace_id),
                Err(e) => common::header::err_reply(
                    self.header.trace_id,
                    ECode::AuthServiceErrorBatchOperateAccess,
                    format!("batch {} access failed: {}", operate.as_str_name(), e),
                )
            }
        };
        auth_service::BatchOperateAccessReply { header }
    }


    // 插入
    async fn insert(&self, dto: auth_service::Access) -> Result<Option<auth_service::Access>, DbErr> {
        let active: ActiveModel = self.decode(dto).into();
        let model = active.insert(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 更新
    async fn update(&self, dto: auth_service::Access) -> Result<Option<auth_service::Access>, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.status = Set(dto.status);
        active.method = Set(dto.method);
        let model = active.update(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 物理删除
    async fn delete_real(&self, id: i32) -> Result<Option<auth_service::Access>, DbErr> {
        Entity::delete_by_id(id).exec(self.db).await?;
        Ok(None)
    }

    // 保存并更新
    async fn save(&self, mut data: auth_service::Access) -> Result<Option<auth_service::Access>, DbErr> {
        let active: ActiveModel = self.decode(data.clone()).into();
        let result = Entity::insert(active).on_conflict(
            // 定义冲突
            sea_query::OnConflict::column(Column::Path)
                // 冲突后更新的字段
                .update_columns([Column::Status, Column::Method])
                .to_owned(),
        ).exec(self.db).await?;
        data.id = result.last_insert_id;
        Ok(Some(data))
    }

    // 批量写入，有冲突则更新
    async fn batch_save(
        &self,
        dtos: Vec<auth_service::Access>,
    ) -> Result<Option<i32>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let active: ActiveModel = self.decode(dto).into();
            actives.push(active);
        }
        Entity::insert_many(actives).on_conflict(
            // 定义冲突
            sea_query::OnConflict::column(Column::Path)
                // 冲突后更新的字段
                .update_columns([Column::Status, Column::Method])
                .to_owned(),
        ).exec(self.db).await?;
        Ok(None)
    }
}
