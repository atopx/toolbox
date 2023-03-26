use std::str::FromStr;

use sea_orm::{
    sea_query, ActiveModelTrait, ColumnTrait, DbConn, DbErr, EntityTrait, PaginatorTrait,
    QueryFilter, QueryOrder,
};

use domain::{auth_service, public};
use model::user_role_ref::{ActiveModel, Column, Entity, Model};

pub struct Dao<'a> {
    /// 我实例化的对象(所有权), 借给你(智能指针), 你需要声明用多久(生命周期);
    db: &'a DbConn,
    operator: i32,
}

impl<'a> Dao<'a> {
    pub fn new(db: &'a DbConn, operator: i32) -> Dao<'a> {
        return Dao { db, operator };
    }

    async fn from_dto(&self, dto: auth_service::UserRoleRef) -> Model {
        let current_time = common::utils::current_timestamp();
        return Model {
            id: 0,
            user_id: dto.user_id,
            role_id: dto.role_id,
            creator: self.operator,
            updater: self.operator,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        };
    }

    async fn transformed(&self, model: Model) -> auth_service::UserRoleRef {
        return auth_service::UserRoleRef {
            id: model.id,
            user_id: model.user_id,
            role_id: model.role_id,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
            creator: self.operator,
            updater: self.operator,
        };
    }

    // 插入
    pub async fn insert(&self, dto: auth_service::UserRoleRef) -> Result<Model, DbErr> {
        let value = self.from_dto(dto).await;
        let active: ActiveModel = value.into();
        return active.insert(self.db).await;
    }

    // 物理删除
    pub async fn delete(&self, id: i32) -> Result<sea_orm::DeleteResult, DbErr> {
        return Entity::delete_by_id(id).exec(self.db).await;
    }

    // 批量写入，有冲突则更新
    pub async fn save(
        &self,
        dtos: Vec<auth_service::UserRoleRef>,
    ) -> Result<sea_orm::InsertResult<ActiveModel>, DbErr> {
        let mut actives = vec![];
        for dto in dtos {
            let value = self.from_dto(dto).await;
            let active: ActiveModel = value.into();
            actives.push(active);
        }
        return Entity::insert_many(actives)
            .on_conflict(
                // 定义冲突
                sea_query::OnConflict::columns([Column::UserId, Column::RoleId])
                    // 冲突后更新的字段
                    .do_nothing()
                    .to_owned(),
            )
            .exec(self.db)
            .await;
    }

    // 列表筛选
    pub async fn list(
        &self,
        params: auth_service::ListUserRoleRefParams,
    ) -> Result<(Vec<auth_service::UserRoleRef>, Option<public::Pager>), DbErr> {
        let mut query = Entity::find().filter(Column::DeleteTime.eq(0));

        // 筛选
        if let Some(filter) = params.filter {
            if !filter.ids.is_empty() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if !filter.role_ids.is_empty() {
                query = query.filter(Column::RoleId.is_in(filter.role_ids));
            }
            if !filter.user_id.is_empty() {
                query = query.filter(Column::UserId.is_in(filter.user_id));
            }
            if !filter.creators.is_empty() {
                query = query.filter(Column::Creator.is_in(filter.creators));
            }
            if !filter.updaters.is_empty() {
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
