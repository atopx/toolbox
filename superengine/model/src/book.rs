//! `SeaORM` Entity. Generated by sea-orm-codegen 0.11.1

use sea_orm::entity::prelude::*;

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Eq)]
#[sea_orm(table_name = "book")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    pub name: String,
    pub author: String,
    pub src: String,
    pub cover: String,
    pub label: String,
    #[sea_orm(column_type = "Text")]
    pub intro: String,
    pub state: i32,
    pub last_modify: i64,
    pub create_time: i64,
    pub update_time: i64,
    pub delete_time: i64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
