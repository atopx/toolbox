use sea_orm::{ActiveModelTrait, DbConn, DbErr, EntityTrait};
use sea_orm::ActiveValue::Set;

use domain::auth_service;
use model::user::{ActiveModel, Model};
use model::user::Entity as UserEntity;

pub struct Dao;

impl Dao {
    pub async fn insert(db: &DbConn, user: auth_service::User) -> Result<Model, DbErr> {
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
        return active.insert(db).await;
    }

    pub async fn update(db: &DbConn, user: auth_service::User) -> Result<Model, DbErr> {
        let value: Option<Model> = UserEntity::find_by_id(user.id).one(db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.username = Set(user.username);
        active.password = Set(user.password);
        active.status = Set(user.status);
        active.update_time = Set(common::utils::current_timestamp());
        return active.update(db).await;
    }

    pub async fn delete(db: &DbConn, id: i32) -> Result<Model, DbErr> {
        let value: Option<Model> = UserEntity::find_by_id(id).one(db).await?;
        let mut active: ActiveModel = value.unwrap().into();
        active.delete_time = Set(common::utils::current_timestamp());
        return active.update(db).await;
    }
}

