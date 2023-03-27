use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Result, Status};

use domain::mainframe_service;
use domain::public;
use tracing::info;

// use super::business;

pub struct MainframeService {
    db: DatabaseConnection,
}

impl MainframeService {
    pub fn new(db: DatabaseConnection) -> Self {
        return Self { db };
    }
}

#[tonic::async_trait]
impl mainframe_service::mainframe_service_server::MainframeService for MainframeService {
    async fn list_computer(
        &self,
        request: Request<mainframe_service::ListComputerParams>,
    ) -> Result<Response<mainframe_service::ListComputerReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn deal_computer(
        &self,
        request: Request<mainframe_service::DealComputerParams>,
    ) -> Result<Response<mainframe_service::DealComputerReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn operate_computer(
        &self,
        request: Request<mainframe_service::OperateComputerParams>,
    ) -> Result<Response<mainframe_service::OperateComputerReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);

        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }

    async fn batch_operate_computer(
        &self,
        request: Request<mainframe_service::BatchOperateComputerParams>,
    ) -> Result<Response<mainframe_service::BatchOperateComputerReply>, Status> {
        let params = request.into_inner();
        let header = params.header.to_owned().unwrap();
        info!("operate_user {}", header.trace_id);
        
        let code = tonic::Code::Unimplemented;
        return Err(Status::new(code, code.description()));
    }
}
