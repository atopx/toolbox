use std::str::FromStr;

use sea_orm::ActiveValue::Set;
use sea_orm::{
    sea_query, ActiveModelTrait, ColumnTrait, DbConn, DbErr, EntityTrait, PaginatorTrait,
    QueryFilter, QueryOrder,
};

use domain::{auth_service, public};
use model::access::{ActiveModel, Column, Entity, Model};

pub struct Dao<'a> {
    /// 我实例化的对象(所有权), 借给你(智能指针), 你需要声明用多久(生命周期);
    db: &'a DbConn,
}

impl<'a> Dao<'a> {
    pub fn new(db: &'a DbConn) -> Dao<'a> {
        return Dao { db };
    }

    async fn from_dto(&self, dto: auth_service::Access) -> Model {
        return Model {
            id: 0,
            status: dto.status,
            path: dto.path,
            method: dto.method,
        };
    }

    async fn transformed(&self, model: Model) -> auth_service::Access {
        return auth_service::Access {
            id: model.id,
            path: model.path,
            method: model.method,
            status: model.status,
        };
    }

    // 插入
    pub async fn insert(&self, dto: auth_service::Access) -> Result<Model, DbErr> {
        let value = self.from_dto(dto).await;
        let active: ActiveModel = value.into();
        return active.insert(self.db).await;
    }

    // 更新
    pub async fn update(&self, dto: auth_service::Access) -> Result<Model, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.status = Set(dto.status);
        active.method = Set(dto.method);
        return active.update(self.db).await;
    }

    // 物理删除
    pub async fn delete_real(&self, id: i32) -> Result<sea_orm::DeleteResult, DbErr> {
        return Entity::delete_by_id(id).exec(self.db).await;
    }

    // 批量写入，有冲突则更新
    pub async fn save(
        &self,
        dtos: Vec<auth_service::Access>,
    ) -> Result<sea_orm::InsertResult<ActiveModel>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let active: ActiveModel = self.from_dto(dto).await.into();
            actives.push(active);
        }
        return Entity::insert_many(actives)
            .on_conflict(
                // 定义冲突
                sea_query::OnConflict::column(Column::Path)
                    // 冲突后更新的字段
                    .update_columns([Column::Status, Column::Method])
                    .to_owned(),
            )
            .exec(self.db)
            .await;
    }

    // 列表筛选
    pub async fn list(
        &self,
        params: auth_service::ListAccessParams,
    ) -> Result<(Vec<auth_service::Access>, Option<public::Pager>), DbErr> {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = params.filter {
            if !filter.ids.is_empty() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if !filter.paths.is_empty() {
                query = query.filter(Column::Path.is_in(filter.paths));
            }
            if !filter.methods.is_empty() {
                query = query.filter(Column::Method.is_in(filter.methods))
            }
            if !filter.states.is_empty() {
                query = query.filter(Column::Status.is_in(filter.states));
            }
        }

        // 排序
        let mut sorts = params.sorts;
        if sorts.is_empty() {
            // add a default sort value
            sorts.push(public::Sort {
                field: "id".to_string(),
                direction: public::SortDirection::SortAsc.into(),
            });
        }
        for sort in sorts {
            if let Ok(col) = Column::from_str(sort.field.as_str()) {
                query = match public::SortDirection::from_i32(sort.direction) {
                    Some(public::SortDirection::SortAsc) => query.order_by_asc(col),
                    Some(public::SortDirection::SortDesc) => query.order_by_desc(col),
                    None => query,
                }
            }
        }

        // 分页
        if let Some(mut value) = params.pager {
            if !value.disabled {
                let paginator = query.paginate(self.db, value.size);
                value.count = paginator.num_items().await?;
                let mut dtos = vec![];
                for model in paginator.fetch_page(value.index - 1).await? {
                    dtos.push(self.transformed(model).await)
                }
                return Ok((dtos, Some(value)));
            }
        }

        // 不分页
        return match query.all(self.db).await {
            Ok(models) => {
                let mut users = vec![];
                for model in models {
                    users.push(self.transformed(model).await);
                }
                Ok((users, None))
            }
            Err(err) => Err(err),
        };
    }
}
