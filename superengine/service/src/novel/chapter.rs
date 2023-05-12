use std::ops::Not;
use std::str::FromStr;

use sea_orm::{ColumnTrait, DatabaseConnection, EntityTrait, PaginatorTrait, QueryFilter, QueryOrder};

use domain::{novel_service, public};
use domain::ecode::ECode;
use model::chapter::{Column, Entity, Model};

pub struct Business<'a> {
    db: &'a DatabaseConnection,
    header: public::Header,
}

impl<'a> Business<'a> {
    pub fn new(db: &'a DatabaseConnection, header: public::Header) -> Self {
        Self { db, header }
    }

    fn encode(&self, model: Model) -> novel_service::Chapter {
        novel_service::Chapter {
            id: model.id,
            book_id: model.book_id,
            code: model.code,
            src: model.src,
            title: model.title,
            content: model.content,
            state: model.state,
            delete_time: model.delete_time,
            create_time: model.create_time,
            update_time: model.update_time,
        }
    }

    // 列表筛选
    pub async fn list(
        &self,
        filter: Option<novel_service::ChapterFilter>,
        mut sorts: Vec<public::Sort>,
        pager: Option<public::Pager>,
    ) -> novel_service::ListChapterReply {
        let mut query = Entity::find();

        // 筛选
        if let Some(filter) = filter {
            if filter.ids.is_empty().not() {
                query = query.filter(Column::Id.is_in(filter.ids));
            }
            if filter.book_ids.is_empty().not() {
                query = query.filter(Column::BookId.is_in(filter.book_ids))
            }

            if filter.states.is_empty().not() {
                query = query.filter(Column::State.is_in(filter.states))
            }

            if filter.srcs.is_empty().not() {
                query = query.filter(Column::Src.is_in(filter.srcs))
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
                    ECode::NovelServiceErrorListChapter,
                    format!("list book chapter failed: {}", e),
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

        novel_service::ListChapterReply { header, pager: Some(pager), data: dtos }
    }
}