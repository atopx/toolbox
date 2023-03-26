use std::str::FromStr;

use sea_orm::ActiveValue::Set;
use sea_orm::{
    ActiveModelTrait, ColumnTrait, DbConn, DbErr, EntityTrait, PaginatorTrait, QueryFilter,
    QueryOrder,
};

use domain::{auth_service, public};
use model::auth_token::{ActiveModel, Column, Entity, Model};

pub struct Dao<'a> {
    /// 我实例化的对象(所有权), 借给你(智能指针), 你需要声明用多久(生命周期);
    db: &'a DbConn,
}

impl<'a> Dao<'a> {
    pub fn new(db: &'a DbConn) -> Dao<'a> {
        return Dao { db };
    }

    async fn from_dto(&self, dto: auth_service::AuthToken) -> Model {
        let current_time = common::utils::current_timestamp();
        return Model {
            id: 0,
            user_id: dto.user_id,
            access_token: dto.access_token,
            refresh_token: dto.refresh_token,
            issued_time: dto.issued_time,
            expire_time: dto.expire_time,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        };
    }

    async fn transformed(&self, model: Model) -> auth_service::AuthToken {
        return auth_service::AuthToken {
            id: model.id,
            user_id: model.user_id,
            access_token: model.access_token,
            refresh_token: model.refresh_token,
            issued_time: model.issued_time,
            expire_time: model.expire_time,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
        };
    }

    // 插入
    pub async fn insert(&self, dto: auth_service::AuthToken) -> Result<Model, DbErr> {
        let active: ActiveModel = self.from_dto(dto).await.into();
        return active.insert(self.db).await;
    }

    // 更新
    pub async fn update(&self, dto: auth_service::AuthToken) -> Result<Model, DbErr> {
        let value: Option<Model> = Entity::find_by_id(dto.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.access_token = Set(dto.access_token);
        active.refresh_token = Set(dto.refresh_token);
        active.issued_time = Set(dto.issued_time);
        active.expire_time = Set(dto.expire_time);
        active.update_time = Set(common::utils::current_timestamp());
        return active.update(self.db).await;
    }

    // 逻辑删除
    pub async fn delete(&self, id: i32) -> Result<Model, DbErr> {
        let value: Option<Model> = Entity::find_by_id(id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.delete_time = Set(common::utils::current_timestamp());
        return active.update(self.db).await;
    }

    // 物理删除
    pub async fn delete_real(&self, id: i32) -> Result<sea_orm::DeleteResult, DbErr> {
        return Entity::delete_by_id(id).exec(self.db).await;
    }

    // 列表筛选
    pub async fn list(
        &self,
        params: auth_service::ListAuthTokenParams,
    ) -> Result<(Vec<auth_service::AuthToken>, Option<public::Pager>), DbErr> {
        let mut query = Entity::find().filter(Column::DeleteTime.eq(0));

        // 筛选
        if let Some(filter) = params.filter {
            if !filter.ids.is_empty() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if !filter.user_ids.is_empty() {
                query = query.filter(Column::UserId.is_in(filter.user_ids));
            }
            if !filter.access_tokens.is_empty() {
                query = query.filter(Column::AccessToken.is_in(filter.access_tokens));
            }
            if !&filter.refresh_tokens.is_empty() {
                query = query.filter(Column::RefreshToken.is_in(filter.refresh_tokens));
            }
            if let Some(range) = filter.issued_time_range {
                query = query.filter(Column::CreateTime.between(range.left, range.right))
            }
            if let Some(range) = filter.expire_time_range {
                query = query.filter(Column::CreateTime.between(range.left, range.right))
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
        let mut sorts = params.sorts;
        if sorts.is_empty() {
            // add a default sort value
            sorts.push(public::Sort {
                field: "update_time".to_string(),
                direction: public::SortDirection::SortDesc.into(),
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
                let mut dtos = vec![];
                for model in models {
                    dtos.push(self.transformed(model).await);
                }
                Ok((dtos, None))
            }
            Err(err) => Err(err),
        };
    }
}
