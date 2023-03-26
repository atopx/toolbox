use std::str::FromStr;

use sea_orm::ActiveValue::Set;
use sea_orm::{
    sea_query, ActiveModelTrait, ColumnTrait, DbConn, DbErr, EntityTrait, PaginatorTrait,
    QueryFilter, QueryOrder,
};

use domain::{auth_service, public};
use model::user::{ActiveModel, Column, Entity, Model};

pub struct Dao<'a> {
    /// 我实例化的对象(所有权), 借给你(智能指针), 你需要声明用多久(生命周期);
    db: &'a DbConn,
}

impl<'a> Dao<'a> {
    pub fn new(db: &'a DbConn) -> Dao<'a> {
        return Dao { db };
    }

    async fn transformed(&self, model: Model) -> auth_service::User {
        return auth_service::User {
            id: model.id,
            username: model.username,
            password: model.password,
            status: model.status,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
        };
    }

    // 插入
    pub async fn insert(&self, user: auth_service::User) -> Result<Model, DbErr> {
        let current_time = common::utils::current_timestamp();
        let value = Model {
            id: 0,
            username: user.username,
            password: user.password,
            status: user.status,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        };
        let active: ActiveModel = value.into();
        return active.insert(self.db).await;
    }

    // 更新
    pub async fn update(&self, user: auth_service::User) -> Result<Model, DbErr> {
        let value: Option<Model> = Entity::find_by_id(user.id).one(self.db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.username = Set(user.username);
        active.password = Set(user.password);
        active.status = Set(user.status);
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

    // 批量写入，有冲突则更新
    pub async fn save(
        &self,
        users: Vec<auth_service::User>,
    ) -> Result<sea_orm::InsertResult<ActiveModel>, DbErr> {
        let current_time = common::utils::current_timestamp();
        let mut actives = vec![];
        for user in users {
            let value = Model {
                id: 0,
                username: user.username,
                password: user.password,
                status: user.status,
                create_time: current_time,
                update_time: current_time,
                delete_time: 0,
            };
            let active: ActiveModel = value.into();
            actives.push(active);
        }
        return Entity::insert_many(actives)
            .on_conflict(
                // 定义冲突
                sea_query::OnConflict::column(Column::Username)
                    // 冲突后更新的字段
                    .update_columns([Column::Password, Column::Status, Column::UpdateTime])
                    .to_owned(),
            )
            .exec(self.db)
            .await;
    }

    // 列表筛选
    pub async fn list(
        &self,
        params: auth_service::ListUserParams,
    ) -> Result<(Vec<auth_service::User>, Option<public::Pager>), DbErr> {
        let mut query = Entity::find().filter(Column::DeleteTime.eq(0));

        // 筛选
        if let Some(filter) = params.filter {
            if filter.ids.len() > 0 {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.usernames.len() > 0 {
                query = query.filter(Column::Username.is_in(filter.usernames));
            }
            if filter.states.len() > 0 {
                query = query.filter(Column::Status.is_in(filter.states));
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
                let mut users = vec![];
                for model in paginator.fetch_page(value.index-1).await? {
                    println!("1");
                    users.push(self.transformed(model).await)
                }
                return Ok((users, Some(value)));
            }
        }

        // 不分页
        return match query.all(self.db).await {
            Ok(models) => {
                let mut users = vec![];
                for model in models {
                    println!("1");
                    users.push(self.transformed(model).await);
                }
                Ok((users, None))
            }
            Err(err) => Err(err),
        };
    }
}

mod tests {

    #[test]
    fn test_insert_many() {}
}
