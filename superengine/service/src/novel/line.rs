use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ActiveModelTrait, ColumnTrait, DatabaseConnection, DbErr, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder};

use domain::{novel_service, public};
use domain::public::{ECode, Operation};
use model::line::{ActiveModel, Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn decode(&self, dto: novel_service::Line) -> Model {
        Model {
            id: 0,
            book_id: dto.book_id,
            code: dto.code,
            value: Some(dto.value),
        }
    }

    fn encode(&self, model: Model) -> novel_service::Line {
        novel_service::Line {
            id: model.id,
            book_id: model.book_id,
            code: model.code,
            value: model.value.unwrap_or_default(),
        }
    }

    // 列表筛选
    pub async fn list(
        &self,
        filter: Option<novel_service::LineFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> novel_service::ListLineReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.book_ids.is_empty().not() {
                query = query.filter(Column::BookId.is_in(filter.book_ids))
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
                    ECode::NovelServiceErrorListLine,
                    format!("list book line failed: {}", e),
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

        novel_service::ListLineReply { header, pager: Some(pager), data: dtos }
    }

    pub async fn operate(&self, operate: Operation, data: Option<novel_service::Line>) -> novel_service::OperateLineReply {
        let (header, data) = match data {
            None => {
                (common::header::err_reply(
                    self.header.trace_id,
                    ECode::NovelServiceErrorOperateLine,
                    "missing operate data".to_string(),
                ), None)
            }
            Some(data) => {
                let result = match operate {
                    Operation::Create => self.insert(data).await,
                    Operation::Update => Ok(None),
                    Operation::Upsert => Ok(None),
                    Operation::Delete => Ok(None),
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
                            ECode::NovelServiceErrorOperateLine,
                            format!("{} book line failed: {}", operate.as_str_name(), e),
                        );
                        (header, None)
                    }
                }
            }
        };
        novel_service::OperateLineReply { header, data }
    }

    pub async fn batch_operate(&self, _: Operation, _: Vec<novel_service::Line>) -> novel_service::BatchOperateLineReply {
        let header = common::header::err_reply(
            self.header.trace_id,
            ECode::SystemErrorUnimplemented,
            "SystemErrorUnimplemented".to_string(),
        );
        novel_service::BatchOperateLineReply { header }
    }

    // 插入
    async fn insert(&self, dto: novel_service::Line) -> Result<Option<novel_service::Line>, DbErr> {
        let active: ActiveModel = self.decode(dto).into();
        let model = active.insert(self.db).await?;
        Ok(Some(self.encode(model)))
    }

    // 物理删除
    async fn delete_real(&self, id: i32) -> Result<Option<novel_service::Line>, DbErr> {
        Entity::delete_by_id(id).exec(self.db).await?;
        Ok(None)
    }
}