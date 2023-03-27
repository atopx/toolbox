use std::str::FromStr;

use sea_orm::{
    sea_query, ActiveModelTrait, ColumnTrait, DbConn, DbErr, EntityTrait, PaginatorTrait,
    QueryFilter, QueryOrder,
};

use model::computer::{ActiveModel, Column, Entity, Model};
use domain::{mainframe_service, public};


pub struct Dao<'a> {
    /// 我实例化的对象(所有权), 借给你(智能指针), 你需要声明用多久(生命周期);
    db: &'a DbConn,
    operator: i32,
}

impl<'a> Dao<'a> {
    pub fn new(db: &'a DbConn, operator: i32) -> Dao<'a> {
        return Dao { db, operator };
    }

    async fn from_dto(&self, dto: mainframe_service::Computer) -> Model {
        let current_time = common::utils::current_timestamp();
        return Model {
            id: 0,
            name: dto.name,
            username: dto.username,
            password: dto.password,
            lan_hostname: dto.lan_hostname,
            wan_hostname: dto.wan_hostname,
            mac_address: dto.mac_address,
            power_status: dto.power_status,
            scan_time: dto.scan_time,
            creator: self.operator,
            updater: self.operator,
            create_time: current_time,
            update_time: current_time,
            delete_time: 0,
        };
    }

    async fn transformed(&self, model: Model) -> mainframe_service::Computer {
        return mainframe_service::Computer {
            id: model.id,
            name: model.name,
            username: model.username,
            password: model.password,
            lan_hostname: model.lan_hostname,
            wan_hostname: model.wan_hostname,
            mac_address: model.mac_address,
            power_status: model.power_status,
            scan_time: model.scan_time,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
            creator: self.operator,
            updater: self.operator,
        };
    }

    // 插入
    pub async fn insert(&self, dto: mainframe_service::Computer) -> Result<Model, DbErr> {
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
        dtos: Vec<mainframe_service::Computer>,
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
                sea_query::OnConflict::column(Column::MacAddress)
                    // 冲突后更新的字段
                    .update_columns([
                        Column::Name,
                        Column::Username,
                        Column::Password,
                        Column::LanHostname,
                        Column::WanHostname,
                        Column::PowerStatus,
                        Column::ScanTime,
                        Column::UpdateTime,
                        Column::Updater,
                    ])
                    .to_owned(),
            )
            .exec(self.db)
            .await;
    }

    // 列表筛选
    pub async fn list(
        &self,
        params: mainframe_service::ListComputerParams,
    ) -> Result<(Vec<mainframe_service::Computer>, Option<public::Pager>), DbErr> {
        let mut query = Entity::find().filter(Column::DeleteTime.eq(0));

        // 筛选
        if let Some(filter) = params.filter {
            if !filter.ids.is_empty() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if !filter.names.is_empty() {
                query = query.filter(Column::Name.is_in(filter.names));
            }
            if !filter.usernames.is_empty() {
                query = query.filter(Column::Username.is_in(filter.usernames));
            }
            if !filter.passwords.is_empty() {
                query = query.filter(Column::Password.is_in(filter.passwords));
            }
            if !filter.mac_address.is_empty() {
                query = query.filter(Column::MacAddress.is_in(filter.mac_address));
            }
            if !filter.lan_hostnames.is_empty() {
                query = query.filter(Column::LanHostname.is_in(filter.lan_hostnames));
            }
            if !filter.wan_hostnames.is_empty() {
                query = query.filter(Column::LanHostname.is_in(filter.wan_hostnames));
            }
            if !filter.power_states.is_empty() {
                query = query.filter(Column::PowerStatus.is_in(filter.power_states));
            }
            if !filter.creators.is_empty() {
                query = query.filter(Column::Creator.is_in(filter.creators));
            }
            if !filter.updaters.is_empty() {
                query = query.filter(Column::Updater.is_in(filter.updaters));
            }
            if let Some(range) = filter.scan_time_range {
                query = query.filter(Column::ScanTime.between(range.left, range.right))
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
                if !keywords.keyword.is_empty() {
                    query = query.filter(
                        Column::LanHostname
                            .contains(&keywords.keyword)
                            .or(Column::WanHostname.contains(&keywords.keyword))
                            .or(Column::Name.contains(&keywords.keyword))
                            .or(Column::Username.contains(&keywords.keyword))
                            .or(Column::Password.contains(&keywords.keyword))
                            .or(Column::MacAddress.contains(&keywords.keyword)),
                    )
                }

                if !keywords.hostname.is_empty() {
                    query = query.filter(
                        Column::LanHostname
                            .contains(&keywords.hostname)
                            .or(Column::WanHostname.contains(&keywords.hostname)),
                    )
                }

                if !keywords.name_or_uname.is_empty() {
                    query = query.filter(
                        Column::Name
                            .contains(&keywords.name_or_uname)
                            .or(Column::Username.contains(&keywords.name_or_uname)),
                    )
                }

                if !keywords.name_or_uname_or_hostname.is_empty() {
                    query = query.filter(
                        Column::Name
                            .contains(&keywords.name_or_uname_or_hostname)
                            .or(Column::Username.contains(&keywords.name_or_uname_or_hostname))
                            .or(Column::LanHostname.contains(&keywords.keyword))
                            .or(Column::WanHostname.contains(&keywords.keyword)),
                    )
                }
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
                query = match sort.direction() {
                    public::SortDirection::SortAsc => query.order_by_asc(col),
                    public::SortDirection::SortDesc => query.order_by_desc(col),
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

    // 执行动作
    pub async fn deal(&self, params: mainframe_service::DealComputerParams) {}
}

