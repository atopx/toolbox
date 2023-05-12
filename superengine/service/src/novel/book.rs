use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ColumnTrait, DatabaseConnection, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder};

use domain::{novel_service, public};
use domain::ecode::ECode;
use domain::public::BooleanScope;
use model::book::{Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }


    fn encode(&self, model: Model) -> novel_service::Book {
        novel_service::Book {
            id: model.id,
            src: model.src,
            name: model.name,
            author: model.author,
            label: model.label,
            cover: model.cover,
            intro: model.intro,
            state: model.state,
            last_modify: model.last_modify,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
        }
    }

    // 列表筛选
    pub async fn list(
        &self,
        filter: Option<novel_service::BookFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> novel_service::ListBookReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            query = match filter.deleted() {
                BooleanScope::BoolAll => query,
                BooleanScope::BoolFalse => query.filter(Column::DeleteTime.eq(0)),
                BooleanScope::BoolTrue => query.filter(Column::DeleteTime.gt(0)),
            };
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.names.is_empty().not() {
                query = query.filter(Column::Name.is_in(filter.names))
            }
            if filter.authors.is_empty().not() {
                query = query.filter(Column::Author.is_in(filter.authors))
            }
            if filter.states.is_empty().not() {
                query = query.filter(Column::Author.is_in(filter.states))
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
                if keywords.keyword.is_empty().not() {
                    query = query.filter(
                        Column::Name.contains(&keywords.keyword).or(
                            Column::Author.contains(&keywords.keyword)
                        )
                    )
                }
                if keywords.name.is_empty().not() {
                    query = query.filter(Column::Name.contains(&keywords.name))
                }
                if keywords.author.is_empty().not() {
                    query = query.filter(Column::Author.contains(&keywords.author))
                }
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
                    ECode::NovelServiceErrorListBook,
                    format!("list book failed: {}", e),
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

        novel_service::ListBookReply { header, pager: Some(pager), data: dtos }
    }
}
