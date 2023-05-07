use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Status};
use domain::novel_service;

mod book;
mod line;


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

    async fn operate_book(
        &self,
        request: Request<novel_service::OperateBookParams>,
    ) -> Result<Response<novel_service::OperateBookReply>, Status> {
        let params = request.into_inner();
        let bis = book::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_book(
        &self,
        request: Request<novel_service::BatchOperateBookParams>,
    ) -> Result<Response<novel_service::BatchOperateBookReply>, Status> {
        let params = request.into_inner();
        let bis = book::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_line(
        &self,
        request: Request<novel_service::ListLineParams>,
    ) -> Result<Response<novel_service::ListLineReply>, Status> {
        let params = request.into_inner();
        let bis = line::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_line(
        &self,
        request: Request<novel_service::OperateLineParams>,
    ) -> Result<Response<novel_service::OperateLineReply>, Status> {
        let params = request.into_inner();
        let bis = line::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_line(
        &self,
        request: Request<novel_service::BatchOperateLineParams>,
    ) -> Result<Response<novel_service::BatchOperateLineReply>, Status> {
        let params = request.into_inner();
        let bis = line::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }
}