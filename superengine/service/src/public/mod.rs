use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Status};

use domain::public_service;
use domain::public_service::{ListEnumParams, ListEnumReply};

mod label;
mod folder;
mod enums;

pub struct PublicService {
    db: DatabaseConnection,
}

impl PublicService {
    pub async fn new(db_url: &str) -> Self {
        Self { db: common::db::new_client(db_url).await.unwrap() }
    }
}

#[tonic::async_trait]
impl public_service::public_service_server::PublicService for PublicService {
    async fn list_label(
        &self,
        request: Request<public_service::ListLabelParams>,
    ) -> Result<Response<public_service::ListLabelReply>, Status> {
        let params = request.into_inner();
        let bis = label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_label(
        &self,
        request: Request<public_service::OperateLabelParams>,
    ) -> Result<Response<public_service::OperateLabelReply>, Status> {
        let params = request.into_inner();
        let bis = label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_label(
        &self,
        request: Request<public_service::BatchOperateLabelParams>,
    ) -> Result<Response<public_service::BatchOperateLabelReply>, Status> {
        let params = request.into_inner();
        let bis = label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_folder(
        &self,
        request: Request<public_service::ListFolderParams>,
    ) -> Result<Response<public_service::ListFolderReply>, Status> {
        let params = request.into_inner();
        let bis = folder::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_folder(
        &self,
        request: Request<public_service::OperateFolderParams>,
    ) -> Result<Response<public_service::OperateFolderReply>, Status> {
        let params = request.into_inner();
        let bis = folder::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_folder(
        &self,
        request: Request<public_service::BatchOperateFolderParams>,
    ) -> Result<Response<public_service::BatchOperateFolderReply>, Status> {
        let params = request.into_inner();
        let bis = folder::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_enum(&self, request: Request<ListEnumParams>) -> Result<Response<ListEnumReply>, Status> {
        let params = request.into_inner();
        Ok(Response::new(enums::list_enum(params.header.unwrap().trace_id)))
    }
}
