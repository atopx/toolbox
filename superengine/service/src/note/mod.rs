use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Status};

use domain::note_service;

mod note;
mod note_label;

pub struct NoteService {
    db: DatabaseConnection,
}

impl NoteService {
    pub async fn new(db_url: &str) -> Self {
        Self { db: common::db::new_client(db_url).await.unwrap() }
    }
}

#[tonic::async_trait]
impl note_service::note_service_server::NoteService for NoteService {
    async fn list_note(
        &self,
        request: Request<note_service::ListNoteParams>,
    ) -> Result<Response<note_service::ListNoteReply>, Status> {
        let params = request.into_inner();
        let bis = note::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_note(
        &self,
        request: Request<note_service::OperateNoteParams>,
    ) -> Result<Response<note_service::OperateNoteReply>, Status> {
        let params = request.into_inner();
        let bis = note::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_note(
        &self,
        request: Request<note_service::BatchOperateNoteParams>,
    ) -> Result<Response<note_service::BatchOperateNoteReply>, Status> {
        let params = request.into_inner();
        let bis = note::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_note_label(
        &self,
        request: Request<note_service::ListNoteLabelParams>,
    ) -> Result<Response<note_service::ListNoteLabelReply>, Status> {
        let params = request.into_inner();
        let bis = note_label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_note_label(
        &self,
        request: Request<note_service::OperateNoteLabelParams>,
    ) -> Result<Response<note_service::OperateNoteLabelReply>, Status> {
        let params = request.into_inner();
        let bis = note_label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_note_label(
        &self,
        request: Request<note_service::BatchOperateNoteLabelParams>,
    ) -> Result<Response<note_service::BatchOperateNoteLabelReply>, Status> {
        let params = request.into_inner();
        let bis = note_label::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }
}