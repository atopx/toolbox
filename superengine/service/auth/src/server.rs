use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Result, Status};

use domain::auth_service;
use domain::public;

use super::business;

pub struct AuthService {
    db: DatabaseConnection,
}

impl AuthService {
    pub fn new(db: DatabaseConnection) -> Self {
        return Self { db };
    }
}

#[tonic::async_trait]
impl auth_service::auth_service_server::AuthService for AuthService {
    async fn list_user(
        &self,
        request: Request<auth_service::ListUserParams>,
    ) -> Result<Response<auth_service::ListUserReply>, Status> {
        let params = request.into_inner();
        let dao = business::user::Dao::new(&self.db);
        let (data, pager) = dao.list(params).await.unwrap();
        return Ok(Response::new(auth_service::ListUserReply {
            header: common::header::reply(public::ECode::Success),
            pager,
            data,
        }));
    }

    async fn operate_user(
        &self,
        request: Request<auth_service::OperateUserParams>,
    ) -> Result<Response<auth_service::OperateUserReply>, Status> {
        let params = request.into_inner();
        let _header = params.header.unwrap();
        return match params.data {
            None => {
                let reply = auth_service::OperateUserReply {
                    header: common::header::reply(public::ECode::BadRequest),
                    data: None,
                };
                Ok(Response::new(reply))
            }
            Some(user) => {
                return match public::Operation::from_i32(params.operate).unwrap() {
                    public::Operation::Create => {
                        let dao = business::user::Dao::new(&self.db);
                        let result = dao.insert(user).await.unwrap();
                        Ok(Response::new(auth_service::OperateUserReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(auth_service::User {
                                id: result.id,
                                ..Default::default()
                            }),
                        }))
                    }
                    public::Operation::Update => {
                        let dao = business::user::Dao::new(&self.db);
                        dao.update(user.clone()).await.unwrap();
                        let reply = auth_service::OperateUserReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(user),
                        };
                        Ok(Response::new(reply))
                    }
                    public::Operation::Delete => {
                        let dao = business::user::Dao::new(&self.db);
                        dao.delete(user.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateUserReply::default()))
                    }
                    public::Operation::RealDelete => {
                        let dao = business::user::Dao::new(&self.db);
                        dao.delete_real(user.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateUserReply::default()))
                    }
                    public::Operation::Upsert => {
                        let code = tonic::Code::Unimplemented;
                        Err(Status::new(code, code.description()))
                    }
                };
            }
        };
    }

    async fn batch_operate_user(
        &self,
        request: Request<auth_service::BatchOperateUserParams>,
    ) -> Result<Response<auth_service::BatchOperateUserReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateUserReply::default();
        return Ok(Response::new(reply));
    }

    async fn list_role(
        &self,
        request: Request<auth_service::ListRoleParams>,
    ) -> Result<Response<auth_service::ListRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListRoleReply::default();
        return Ok(Response::new(reply));
    }

    async fn operate_role(
        &self,
        request: Request<auth_service::OperateRoleParams>,
    ) -> Result<Response<auth_service::OperateRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperateRoleReply::default();
        return Ok(Response::new(reply));
    }

    async fn batch_operate_role(
        &self,
        request: Request<auth_service::BatchOperateRoleParams>,
    ) -> Result<Response<auth_service::BatchOperateRoleReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateRoleReply::default();
        return Ok(Response::new(reply));
    }

    async fn list_access(
        &self,
        request: Request<auth_service::ListAccessParams>,
    ) -> Result<Response<auth_service::ListAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListAccessReply::default();
        return Ok(Response::new(reply));
    }

    async fn operate_access(
        &self,
        request: Request<auth_service::OperateAccessParams>,
    ) -> Result<Response<auth_service::OperateAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperateAccessReply::default();
        return Ok(Response::new(reply));
    }

    async fn batch_operate_access(
        &self,
        request: Request<auth_service::BatchOperateAccessParams>,
    ) -> Result<Response<auth_service::BatchOperateAccessReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperateAccessReply::default();
        return Ok(Response::new(reply));
    }

    async fn list_permission(
        &self,
        request: Request<auth_service::ListPermissionParams>,
    ) -> Result<Response<auth_service::ListPermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::ListPermissionReply::default();
        return Ok(Response::new(reply));
    }

    async fn operate_permission(
        &self,
        request: Request<auth_service::OperatePermissionParams>,
    ) -> Result<Response<auth_service::OperatePermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::OperatePermissionReply::default();
        return Ok(Response::new(reply));
    }

    async fn batch_operate_permission(
        &self,
        request: Request<auth_service::BatchOperatePermissionParams>,
    ) -> Result<Response<auth_service::BatchOperatePermissionReply>, Status> {
        println!("{:?}", request);
        let reply = auth_service::BatchOperatePermissionReply::default();
        return Ok(Response::new(reply));
    }

    async fn authorization(
        &self,
        request: Request<auth_service::AuthorizationParams>,
    ) -> Result<Response<auth_service::AuthorizationReply>, Status> {
        println!("extensions: {:#?}", request.extensions());
        println!("metadata: {:?}", request.metadata());
        println!("message: {:?}", request.get_ref());
        let reply = auth_service::AuthorizationReply::default();
        return Ok(Response::new(reply));
    }
}
