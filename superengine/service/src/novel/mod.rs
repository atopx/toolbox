use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Status};
use domain::novel_service;

mod book;
mod chapter;


pub struct NovelService {
    db: DatabaseConnection,
}

impl NovelService {
    pub async fn new(db_url: &str) -> Self {
        Self { db: common::db::new_client(db_url).await.unwrap() }
    }
}

#[tonic::async_trait]
impl novel_service::novel_service_server::NovelService for NovelService {
    async fn list_book(
        &self,
        request: Request<novel_service::ListBookParams>,
    ) -> Result<Response<novel_service::ListBookReply>, Status> {
        let params = request.into_inner();
        let bis = book::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn list_chapter(
        &self,
        request: Request<novel_service::ListChapterParams>,
    ) -> Result<Response<novel_service::ListChapterReply>, Status> {
        let params = request.into_inner();
        let bis = chapter::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }
}