use sea_orm::DatabaseConnection;
use tonic::{Request, Response, Result, Status};

use domain::auth_service;
use domain::public;
use tracing::info;

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
        let header = params.header.unwrap();
        info!("operate_user {}", header.trace_id);
        return match params.data {
            None => {
                let reply = auth_service::OperateUserReply {
                    header: common::header::reply(public::ECode::BadRequest),
                    data: None,
                };
                Ok(Response::new(reply))
            }
            Some(dto) => {
                return match public::Operation::from_i32(params.operate).unwrap() {
                    public::Operation::Create => {
                        let dao = business::user::Dao::new(&self.db);
                        let result = dao.insert(dto).await.unwrap();
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
                        dao.update(dto.clone()).await.unwrap();
                        let reply = auth_service::OperateUserReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(dto),
                        };
                        Ok(Response::new(reply))
                    }
                    public::Operation::Delete => {
                        let dao = business::user::Dao::new(&self.db);
                        dao.delete(dto.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateUserReply::default()))
                    }
                    public::Operation::RealDelete => {
                        let dao = business::user::Dao::new(&self.db);
                        dao.delete_real(dto.id).await.unwrap();
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
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        info!("batch_operate_user {}", header.trace_id);
        if params.data.is_empty() {
            return Ok(Response::new(auth_service::BatchOperateUserReply {
                header: common::header::reply(public::ECode::BadRequest),
            }));
        }
        return match public::Operation::from_i32(params.operate).unwrap() {
            public::Operation::Upsert => {
                let dao = business::user::Dao::new(&self.db);
                match dao.save(params.data).await {
                    Ok(_) => Ok(Response::new(auth_service::BatchOperateUserReply {
                        header: common::header::reply(public::ECode::BadRequest),
                    })),
                    Err(err) => Ok(Response::new(auth_service::BatchOperateUserReply {
                        header: common::header::err_reply(
                            public::ECode::SystemError,
                            err.to_string(),
                        ),
                    })),
                }
            }
            _ => {
                let code = tonic::Code::Unimplemented;
                Err(Status::new(code, code.description()))
            }
        };
    }

    async fn list_role(
        &self,
        request: Request<auth_service::ListRoleParams>,
    ) -> Result<Response<auth_service::ListRoleReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        let dao = business::role::Dao::new(&self.db, header.operator);
        let (data, pager) = dao.list(params).await.unwrap();
        return Ok(Response::new(auth_service::ListRoleReply {
            header: common::header::reply(public::ECode::Success),
            pager,
            data,
        }));
    }

    async fn operate_role(
        &self,
        request: Request<auth_service::OperateRoleParams>,
    ) -> Result<Response<auth_service::OperateRoleReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        return match params.data {
            None => {
                let reply = auth_service::OperateRoleReply {
                    header: common::header::reply(public::ECode::BadRequest),
                    data: None,
                };
                Ok(Response::new(reply))
            }
            Some(dto) => {
                return match public::Operation::from_i32(params.operate).unwrap() {
                    public::Operation::Create => {
                        let dao = business::role::Dao::new(&self.db, header.operator);
                        let result = dao.insert(dto).await.unwrap();
                        Ok(Response::new(auth_service::OperateRoleReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(auth_service::Role {
                                id: result.id,
                                ..Default::default()
                            }),
                        }))
                    }
                    public::Operation::Update => {
                        let dao = business::role::Dao::new(&self.db, header.operator);
                        dao.update(dto.clone()).await.unwrap();
                        let reply = auth_service::OperateRoleReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(dto),
                        };
                        Ok(Response::new(reply))
                    }
                    public::Operation::Delete => {
                        let dao = business::role::Dao::new(&self.db, header.operator);
                        dao.delete(dto.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateRoleReply::default()))
                    }
                    public::Operation::RealDelete => {
                        let dao = business::role::Dao::new(&self.db, header.operator);
                        dao.delete_real(dto.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateRoleReply::default()))
                    }
                    public::Operation::Upsert => {
                        let code = tonic::Code::Unimplemented;
                        Err(Status::new(code, code.description()))
                    }
                };
            }
        };
    }

    async fn batch_operate_role(
        &self,
        request: Request<auth_service::BatchOperateRoleParams>,
    ) -> Result<Response<auth_service::BatchOperateRoleReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        if params.data.is_empty() {
            return Ok(Response::new(auth_service::BatchOperateRoleReply {
                header: common::header::reply(public::ECode::BadRequest),
            }));
        }
        return match public::Operation::from_i32(params.operate).unwrap() {
            public::Operation::Upsert => {
                let dao = business::role::Dao::new(&self.db, header.operator);
                match dao.save(params.data).await {
                    Ok(_) => Ok(Response::new(auth_service::BatchOperateRoleReply {
                        header: common::header::reply(public::ECode::BadRequest),
                    })),
                    Err(err) => Ok(Response::new(auth_service::BatchOperateRoleReply {
                        header: common::header::err_reply(
                            public::ECode::SystemError,
                            err.to_string(),
                        ),
                    })),
                }
            }
            _ => {
                let code = tonic::Code::Unimplemented;
                Err(Status::new(code, code.description()))
            }
        };
    }

    async fn list_access(
        &self,
        request: Request<auth_service::ListAccessParams>,
    ) -> Result<Response<auth_service::ListAccessReply>, Status> {
        let params = request.into_inner();
        let dao = business::access::Dao::new(&self.db);
        let (data, pager) = dao.list(params).await.unwrap();
        return Ok(Response::new(auth_service::ListAccessReply {
            header: common::header::reply(public::ECode::Success),
            pager,
            data,
        }));
    }

    async fn operate_access(
        &self,
        request: Request<auth_service::OperateAccessParams>,
    ) -> Result<Response<auth_service::OperateAccessReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        info!("operate_access {}", header.trace_id);
        return match params.data {
            None => {
                let reply = auth_service::OperateAccessReply {
                    header: common::header::reply(public::ECode::BadRequest),
                    data: None,
                };
                Ok(Response::new(reply))
            }
            Some(dto) => {
                return match public::Operation::from_i32(params.operate).unwrap() {
                    public::Operation::Create => {
                        let dao = business::access::Dao::new(&self.db);
                        let result = dao.insert(dto).await.unwrap();
                        Ok(Response::new(auth_service::OperateAccessReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(auth_service::Access {
                                id: result.id,
                                ..Default::default()
                            }),
                        }))
                    }
                    public::Operation::Update => {
                        let dao = business::access::Dao::new(&self.db);
                        dao.update(dto.clone()).await.unwrap();
                        let reply = auth_service::OperateAccessReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(dto),
                        };
                        Ok(Response::new(reply))
                    }
                    public::Operation::Delete | public::Operation::RealDelete => {
                        let dao = business::access::Dao::new(&self.db);
                        dao.delete_real(dto.id).await.unwrap();
                        Ok(Response::new(auth_service::OperateAccessReply::default()))
                    }
                    public::Operation::Upsert => {
                        let code = tonic::Code::Unimplemented;
                        Err(Status::new(code, code.description()))
                    }
                };
            }
        };
    }

    async fn batch_operate_access(
        &self,
        request: Request<auth_service::BatchOperateAccessParams>,
    ) -> Result<Response<auth_service::BatchOperateAccessReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        info!("batch_operate_access {}", header.trace_id);
        if params.data.is_empty() {
            return Ok(Response::new(auth_service::BatchOperateAccessReply {
                header: common::header::reply(public::ECode::BadRequest),
            }));
        }
        return match public::Operation::from_i32(params.operate).unwrap() {
            public::Operation::Upsert => {
                let dao = business::access::Dao::new(&self.db);
                match dao.save(params.data).await {
                    Ok(_) => Ok(Response::new(auth_service::BatchOperateAccessReply {
                        header: common::header::reply(public::ECode::BadRequest),
                    })),
                    Err(err) => Ok(Response::new(auth_service::BatchOperateAccessReply {
                        header: common::header::err_reply(
                            public::ECode::SystemError,
                            err.to_string(),
                        ),
                    })),
                }
            }
            _ => {
                let code = tonic::Code::Unimplemented;
                Err(Status::new(code, code.description()))
            }
        };
    }

    async fn list_permission(
        &self,
        request: Request<auth_service::ListPermissionParams>,
    ) -> Result<Response<auth_service::ListPermissionReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        let dao = business::permission::Dao::new(&self.db, header.operator);
        let (data, pager) = dao.list(params).await.unwrap();
        return Ok(Response::new(auth_service::ListPermissionReply {
            header: common::header::reply(public::ECode::Success),
            pager,
            data,
        }));
    }

    async fn operate_permission(
        &self,
        request: Request<auth_service::OperatePermissionParams>,
    ) -> Result<Response<auth_service::OperatePermissionReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        return match params.data {
            None => {
                let reply = auth_service::OperatePermissionReply {
                    header: common::header::reply(public::ECode::BadRequest),
                    data: None,
                };
                Ok(Response::new(reply))
            }
            Some(dto) => {
                return match public::Operation::from_i32(params.operate).unwrap() {
                    public::Operation::Create => {
                        let dao = business::permission::Dao::new(&self.db, header.operator);
                        let result = dao.insert(dto).await.unwrap();
                        Ok(Response::new(auth_service::OperatePermissionReply {
                            header: common::header::reply(public::ECode::Success),
                            data: Some(auth_service::Permission {
                                id: result.id,
                                ..Default::default()
                            }),
                        }))
                    }
                    public::Operation::Delete | public::Operation::RealDelete => {
                        let dao = business::permission::Dao::new(&self.db, header.operator);
                        dao.delete(dto.id).await.unwrap();
                        Ok(Response::new(
                            auth_service::OperatePermissionReply::default(),
                        ))
                    }
                    public::Operation::Update | public::Operation::Upsert => {
                        let code = tonic::Code::Unimplemented;
                        Err(Status::new(code, code.description()))
                    }
                };
            }
        };
    }

    async fn batch_operate_permission(
        &self,
        request: Request<auth_service::BatchOperatePermissionParams>,
    ) -> Result<Response<auth_service::BatchOperatePermissionReply>, Status> {
        let params = request.into_inner();
        let header = params.header.as_ref().unwrap();
        info!("batch_operate_permission {}", header.trace_id);
        if params.data.is_empty() {
            return Ok(Response::new(auth_service::BatchOperatePermissionReply {
                header: common::header::reply(public::ECode::BadRequest),
            }));
        }
        return match public::Operation::from_i32(params.operate).unwrap() {
            public::Operation::Upsert => {
                let dao = business::permission::Dao::new(&self.db, header.operator);
                match dao.save(params.data).await {
                    Ok(_) => Ok(Response::new(auth_service::BatchOperatePermissionReply {
                        header: common::header::reply(public::ECode::BadRequest),
                    })),
                    Err(err) => Ok(Response::new(auth_service::BatchOperatePermissionReply {
                        header: common::header::err_reply(
                            public::ECode::SystemError,
                            err.to_string(),
                        ),
                    })),
                }
            }
            _ => {
                let code = tonic::Code::Unimplemented;
                Err(Status::new(code, code.description()))
            }
        };
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
