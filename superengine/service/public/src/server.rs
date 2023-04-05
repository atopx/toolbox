use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Result, Status};

use domain::public_service;
use domain::public;
use tracing::info;
use super::business;

pub struct PublicService {
    db: DatabaseConnection,
}

impl PublicService {
    pub fn new(db: DatabaseConnection) -> Self {
        return Self { db };
    }
}

#[tonic::async_trait]
impl public_service::public_service_server::PublicService for PublicService {
    async fn list_label(
        &self,
        request: Request<public_service::ListLabelParams>,
    ) -> Result<Response<public_service::ListLabelReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn operate_label(
        &self,
        request: Request<public_service::OperateLabelParams>,
    ) -> Result<Response<public_service::OperateLabelReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn batch_operate_label(
        &self,
        request: Request<public_service::BatchOperateLabelParams>,
    ) -> Result<Response<public_service::BatchOperateLabelReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn list_folder(
        &self,
        request: Request<public_service::ListFolderParams>,
    ) -> Result<Response<public_service::ListFolderReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn operate_folder(
        &self,
        request: Request<public_service::OperateFolderParams>,
    ) -> Result<Response<public_service::OperateFolderReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn batch_operate_folder(
        &self,
        request: Request<public_service::BatchOperateFolderParams>,
    ) -> Result<Response<public_service::BatchOperateFolderReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn transform_curl(
        &self,
        request: Request<public_service::TransformCurlParams>,
    ) -> Result<Response<public_service::TransformCurlReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);
        let curl = business::curl::Curl::with_input(params.src.as_str());
        let reply = match curl {
            Ok(curl) => {
                public_service::TransformCurlReply {
                    header: common::header::reply(header.trace_id, public::ECode::Success),
                    data: business::curl::transform_curl(curl),
                }
            }
            Err(e) => {
                public_service::TransformCurlReply {
                    header: common::header::err_reply(header.trace_id, public::ECode::PublicServiceErrorTransCurl, e.to_string()),
                    data: None,
                }
            }
        };
        return Ok(Response::new(reply));
    }
}