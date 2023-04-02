//! `SeaORM` Entity. Generated by sea-orm-codegen 0.11.1

use sea_orm::entity::prelude::*;

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Eq)]
#[sea_orm(table_name = "note")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    pub folder_id: i32,
    pub topic_id: i32,
    pub sign: String,
    pub title: String,
    pub public: i8,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    pub creator: i32,
    pub updater: i32,
    pub create_time: i64,
    pub update_time: i64,
    pub delete_time: i64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}