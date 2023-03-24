use domain::auth_service;
use tonic::{Request, Response, Result, Status};

pub struct AuthService {
    db: sqlx::MySqlPool,
}

impl AuthService {
    pub fn new(db: sqlx::MySqlPool) -> Self {
        println!("this");
        return Self { db };
    }
}

#[tonic::async_trait]
impl auth_service::auth_service_server::AuthService for AuthService {
    async fn list_user(
        &self,
        request: Request<auth_service::ListUserParams>,
    ) -> Result<Response<auth_service::ListUserReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListUserReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn operate_user(
        &self,
        request: Request<auth_service::OperateUserParams>,
    ) -> Result<Response<auth_service::OperateUserReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperateUserReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn batch_operate_user(
        &self,
        request: Request<auth_service::BatchOperateUserParams>,
    ) -> Result<Response<auth_service::BatchOperateUserReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateUserReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn list_role(
        &self,
        request: Request<auth_service::ListRoleParams>,
    ) -> Result<Response<auth_service::ListRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListRoleReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn operate_role(
        &self,
        request: Request<auth_service::OperateRoleParams>,
    ) -> Result<Response<auth_service::OperateRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperateRoleReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn batch_operate_role(
        &self,
        request: Request<auth_service::BatchOperateRoleParams>,
    ) -> Result<Response<auth_service::BatchOperateRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateRoleReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn list_access(
        &self,
        request: Request<auth_service::ListAccessParams>,
    ) -> Result<Response<auth_service::ListAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListAccessReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn operate_access(
        &self,
        request: Request<auth_service::OperateAccessParams>,
    ) -> Result<Response<auth_service::OperateAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperateAccessReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn batch_operate_access(
        &self,
        request: Request<auth_service::BatchOperateAccessParams>,
    ) -> Result<Response<auth_service::BatchOperateAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateAccessReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn list_permission(
        &self,
        request: Request<auth_service::ListPermissionParams>,
    ) -> Result<Response<auth_service::ListPermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListPermissionReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn operate_permission(
        &self,
        request: Request<auth_service::OperatePermissionParams>,
    ) -> Result<Response<auth_service::OperatePermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperatePermissionReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn batch_operate_permission(
        &self,
        request: Request<auth_service::BatchOperatePermissionParams>,
    ) -> Result<Response<auth_service::BatchOperatePermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperatePermissionReply::default();
        return Ok(tonic::Response::new(reply));
    }

    async fn authorization(
        &self,
        request: Request<auth_service::AuthorizationParams>,
    ) -> Result<Response<auth_service::AuthorizationReply>, Status> {
        println!("extensions: {:#?}", request.extensions());
        println!("metadata: {:?}", request.metadata());
        println!("message: {:?}", request.get_ref());
        let reply = auth_service::AuthorizationReply::default();
        return Ok(tonic::Response::new(reply));
    }
}
