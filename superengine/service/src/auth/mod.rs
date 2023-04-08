use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Status};

use domain::auth_service;

mod user;
mod access;
mod auth_token;
mod role;
mod user_role_ref;
mod permission;

pub struct AuthService {
    db: DatabaseConnection,
}


impl AuthService {
    pub async fn new(db_url: &str) -> Self {
        Self { db: common::db::new_client(db_url).await.unwrap() }
    }
}


#[tonic::async_trait]
impl auth_service::auth_service_server::AuthService for AuthService {
    async fn list_user(
        &self,
        request: Request<auth_service::ListUserParams>,
    ) -> Result<Response<auth_service::ListUserReply>, Status> {
        let params = request.into_inner();
        let bis = user::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_user(
        &self,
        request: Request<auth_service::OperateUserParams>,
    ) -> Result<Response<auth_service::OperateUserReply>, Status> {
        let params = request.into_inner();
        let bis = user::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_user(
        &self,
        request: Request<auth_service::BatchOperateUserParams>,
    ) -> Result<Response<auth_service::BatchOperateUserReply>, Status> {
        let params = request.into_inner();
        let bis = user::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_role(
        &self,
        request: Request<auth_service::ListRoleParams>,
    ) -> Result<Response<auth_service::ListRoleReply>, Status> {
        let params = request.into_inner();
        let bis = role::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_role(
        &self,
        request: Request<auth_service::OperateRoleParams>,
    ) -> Result<Response<auth_service::OperateRoleReply>, Status> {
        let params = request.into_inner();
        let bis = role::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_role(
        &self,
        request: Request<auth_service::BatchOperateRoleParams>,
    ) -> Result<Response<auth_service::BatchOperateRoleReply>, Status> {
        let params = request.into_inner();
        let bis = role::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_user_role_ref(
        &self,
        request: Request<auth_service::ListUserRoleRefParams>,
    ) -> Result<Response<auth_service::ListUserRoleRefReply>, Status> {
        let params = request.into_inner();
        let bis = user_role_ref::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_user_role_ref(
        &self,
        request: Request<auth_service::OperateUserRoleRefParams>,
    ) -> Result<Response<auth_service::OperateUserRoleRefReply>, Status> {
        let params = request.into_inner();
        let bis = user_role_ref::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_user_role_ref(
        &self,
        request: Request<auth_service::BatchOperateUserRoleRefParams>,
    ) -> Result<Response<auth_service::BatchOperateUserRoleRefReply>, Status> {
        let params = request.into_inner();
        let bis = user_role_ref::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_access(
        &self,
        request: Request<auth_service::ListAccessParams>,
    ) -> Result<Response<auth_service::ListAccessReply>, Status> {
        let params = request.into_inner();
        let bis = access::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_access(
        &self,
        request: Request<auth_service::OperateAccessParams>,
    ) -> Result<Response<auth_service::OperateAccessReply>, Status> {
        let params = request.into_inner();
        let bis = access::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_access(
        &self,
        request: Request<auth_service::BatchOperateAccessParams>,
    ) -> Result<Response<auth_service::BatchOperateAccessReply>, Status> {
        let params = request.into_inner();
        let bis = access::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_permission(
        &self,
        request: Request<auth_service::ListPermissionParams>,
    ) -> Result<Response<auth_service::ListPermissionReply>, Status> {
        let params = request.into_inner();
        let bis = permission::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_permission(
        &self,
        request: Request<auth_service::OperatePermissionParams>,
    ) -> Result<Response<auth_service::OperatePermissionReply>, Status> {
        let params = request.into_inner();
        let bis = permission::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_permission(
        &self,
        request: Request<auth_service::BatchOperatePermissionParams>,
    ) -> Result<Response<auth_service::BatchOperatePermissionReply>, Status> {
        let params = request.into_inner();
        let bis = permission::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn list_auth_token(
        &self,
        request: Request<auth_service::ListAuthTokenParams>,
    ) -> Result<Response<auth_service::ListAuthTokenReply>, Status> {
        let params = request.into_inner();
        let bis = auth_token::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.list(params.filter, params.sorts, params.pager).await;
        Ok(Response::new(reply))
    }

    async fn operate_auth_token(
        &self,
        request: Request<auth_service::OperateAuthTokenParams>,
    ) -> Result<Response<auth_service::OperateAuthTokenReply>, Status> {
        let params = request.into_inner();
        let bis = auth_token::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }

    async fn batch_operate_auth_token(
        &self,
        request: Request<auth_service::BatchOperateAuthTokenParams>,
    ) -> Result<Response<auth_service::BatchOperateAuthTokenReply>, Status> {
        let params = request.into_inner();
        let bis = auth_token::Business::new(&self.db, params.header.to_owned().unwrap());
        let reply = bis.batch_operate(params.operate(), params.data).await;
        Ok(Response::new(reply))
    }
}
