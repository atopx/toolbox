#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Access {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(string, tag = "2")]
    pub path: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub method: ::prost::alloc::string::String,
    #[prost(enumeration = "AccessStatus", tag = "4")]
    pub status: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AccessFilter {
    #[prost(int64, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(string, repeated, tag = "2")]
    pub paths: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(string, repeated, tag = "3")]
    pub methods: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(enumeration = "AccessStatus", repeated, tag = "4")]
    pub status: ::prost::alloc::vec::Vec<i32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListAccessParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<AccessFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListAccessReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Access>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateAccessParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<Access>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateAccessReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<Access>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateAccessParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Access>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateAccessReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum AccessStatus {
    /// 默认状态, 需要验证权限
    AccessDefault = 0,
    /// 可匿名访问，无需验证权限
    AccessAnonymous = 1,
    /// 已禁用，禁止访问
    AccessDisabled = 2,
}
impl AccessStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            AccessStatus::AccessDefault => "ACCESS_DEFAULT",
            AccessStatus::AccessAnonymous => "ACCESS_ANONYMOUS",
            AccessStatus::AccessDisabled => "ACCESS_DISABLED",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "ACCESS_DEFAULT" => Some(Self::AccessDefault),
            "ACCESS_ANONYMOUS" => Some(Self::AccessAnonymous),
            "ACCESS_DISABLED" => Some(Self::AccessDisabled),
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Permission {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(int64, tag = "2")]
    pub access_id: i64,
    #[prost(int64, tag = "3")]
    pub role_id: i64,
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
    #[prost(int32, tag = "1004")]
    pub creator: i32,
    #[prost(int32, tag = "1005")]
    pub updater: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PermissionFilter {
    #[prost(int64, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(int64, repeated, tag = "2")]
    pub access_ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(int64, repeated, tag = "3")]
    pub role_ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(int32, repeated, tag = "104")]
    pub creators: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "105")]
    pub updaters: ::prost::alloc::vec::Vec<i32>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListPermissionParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<PermissionFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListPermissionReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Permission>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperatePermissionParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<Permission>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperatePermissionReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<Permission>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperatePermissionParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Permission>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperatePermissionReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Role {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(enumeration = "RoleNature", tag = "3")]
    pub nature: i32,
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
    #[prost(int32, tag = "1004")]
    pub creator: i32,
    #[prost(int32, tag = "1005")]
    pub updater: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RoleFilter {
    #[prost(int64, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(string, repeated, tag = "2")]
    pub names: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(enumeration = "RoleNature", repeated, tag = "3")]
    pub natures: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(int32, repeated, tag = "104")]
    pub creators: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "105")]
    pub updaters: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "201")]
    pub keywords: ::core::option::Option<role_filter::Keywords>,
}
/// Nested message and enum types in `RoleFilter`.
pub mod role_filter {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Keywords {
        #[prost(string, tag = "1")]
        pub name: ::prost::alloc::string::String,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListRoleParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<RoleFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListRoleReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Role>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateRoleParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<Role>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateRoleReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<Role>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateRoleParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Role>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateRoleReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum RoleNature {
    /// 初始化角色
    RoleDefault = 0,
    /// 系统角色
    RoleSystem = 1,
    /// 自定义角色
    RoleCustom = 2,
}
impl RoleNature {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            RoleNature::RoleDefault => "ROLE_DEFAULT",
            RoleNature::RoleSystem => "ROLE_SYSTEM",
            RoleNature::RoleCustom => "ROLE_CUSTOM",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "ROLE_DEFAULT" => Some(Self::RoleDefault),
            "ROLE_SYSTEM" => Some(Self::RoleSystem),
            "ROLE_CUSTOM" => Some(Self::RoleCustom),
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct User {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(string, tag = "2")]
    pub username: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub password: ::prost::alloc::string::String,
    #[prost(enumeration = "UserStatus", tag = "4")]
    pub status: i32,
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct UserFilter {
    #[prost(int64, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i64>,
    #[prost(string, repeated, tag = "2")]
    pub usernames: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(enumeration = "UserStatus", repeated, tag = "3")]
    pub states: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListUserParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<UserFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListUserReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<User>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateUserParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<User>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateUserReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<User>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateUserParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<User>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateUserReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum UserStatus {
    /// 正常
    UserNormal = 0,
    /// 禁用
    Disabled = 1,
    /// 黑名单
    Black = 2,
}
impl UserStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            UserStatus::UserNormal => "USER_NORMAL",
            UserStatus::Disabled => "USER_STATUS_DISABLED",
            UserStatus::Black => "USER_STATUS_BLACK",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "USER_NORMAL" => Some(Self::UserNormal),
            "USER_STATUS_DISABLED" => Some(Self::Disabled),
            "USER_STATUS_BLACK" => Some(Self::Black),
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Token {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(int64, tag = "2")]
    pub user_id: i64,
    #[prost(string, tag = "3")]
    pub access_token: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub refresh_token: ::prost::alloc::string::String,
    #[prost(int64, tag = "5")]
    pub issued_time: i64,
    #[prost(int64, tag = "6")]
    pub expire_time: i64,
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AuthorizationParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<Token>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AuthorizationReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
/// Generated server implementations.
pub mod auth_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with AuthServiceServer.
    #[async_trait]
    pub trait AuthService: Send + Sync + 'static {
        /// 用户
        async fn list_user(
            &self,
            request: tonic::Request<super::ListUserParams>,
        ) -> Result<tonic::Response<super::ListUserReply>, tonic::Status>;
        async fn operate_user(
            &self,
            request: tonic::Request<super::OperateUserParams>,
        ) -> Result<tonic::Response<super::OperateUserReply>, tonic::Status>;
        async fn batch_operate_user(
            &self,
            request: tonic::Request<super::BatchOperateUserParams>,
        ) -> Result<tonic::Response<super::BatchOperateUserReply>, tonic::Status>;
        /// 角色
        async fn list_role(
            &self,
            request: tonic::Request<super::ListRoleParams>,
        ) -> Result<tonic::Response<super::ListRoleReply>, tonic::Status>;
        async fn operate_role(
            &self,
            request: tonic::Request<super::OperateRoleParams>,
        ) -> Result<tonic::Response<super::OperateRoleReply>, tonic::Status>;
        async fn batch_operate_role(
            &self,
            request: tonic::Request<super::BatchOperateRoleParams>,
        ) -> Result<tonic::Response<super::BatchOperateRoleReply>, tonic::Status>;
        /// api接口
        async fn list_access(
            &self,
            request: tonic::Request<super::ListAccessParams>,
        ) -> Result<tonic::Response<super::ListAccessReply>, tonic::Status>;
        async fn operate_access(
            &self,
            request: tonic::Request<super::OperateAccessParams>,
        ) -> Result<tonic::Response<super::OperateAccessReply>, tonic::Status>;
        async fn batch_operate_access(
            &self,
            request: tonic::Request<super::BatchOperateAccessParams>,
        ) -> Result<tonic::Response<super::BatchOperateAccessReply>, tonic::Status>;
        /// api接口权限
        async fn list_permission(
            &self,
            request: tonic::Request<super::ListPermissionParams>,
        ) -> Result<tonic::Response<super::ListPermissionReply>, tonic::Status>;
        async fn operate_permission(
            &self,
            request: tonic::Request<super::OperatePermissionParams>,
        ) -> Result<tonic::Response<super::OperatePermissionReply>, tonic::Status>;
        async fn batch_operate_permission(
            &self,
            request: tonic::Request<super::BatchOperatePermissionParams>,
        ) -> Result<tonic::Response<super::BatchOperatePermissionReply>, tonic::Status>;
        /// 认证接口: 无token则创建, 如果token已过期则刷新token
        async fn authorization(
            &self,
            request: tonic::Request<super::AuthorizationParams>,
        ) -> Result<tonic::Response<super::AuthorizationReply>, tonic::Status>;
    }
    #[derive(Debug)]
    pub struct AuthServiceServer<T: AuthService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: AuthService> AuthServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for AuthServiceServer<T>
    where
        T: AuthService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/auth_service.AuthService/ListUser" => {
                    #[allow(non_camel_case_types)]
                    struct ListUserSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::ListUserParams>
                    for ListUserSvc<T> {
                        type Response = super::ListUserReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListUserParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { (*inner).list_user(request).await };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListUserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/OperateUser" => {
                    #[allow(non_camel_case_types)]
                    struct OperateUserSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::OperateUserParams>
                    for OperateUserSvc<T> {
                        type Response = super::OperateUserReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperateUserParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).operate_user(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperateUserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/BatchOperateUser" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperateUserSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::BatchOperateUserParams>
                    for BatchOperateUserSvc<T> {
                        type Response = super::BatchOperateUserReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperateUserParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).batch_operate_user(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperateUserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/ListRole" => {
                    #[allow(non_camel_case_types)]
                    struct ListRoleSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::ListRoleParams>
                    for ListRoleSvc<T> {
                        type Response = super::ListRoleReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListRoleParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { (*inner).list_role(request).await };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListRoleSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/OperateRole" => {
                    #[allow(non_camel_case_types)]
                    struct OperateRoleSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::OperateRoleParams>
                    for OperateRoleSvc<T> {
                        type Response = super::OperateRoleReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperateRoleParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).operate_role(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperateRoleSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/BatchOperateRole" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperateRoleSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::BatchOperateRoleParams>
                    for BatchOperateRoleSvc<T> {
                        type Response = super::BatchOperateRoleReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperateRoleParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).batch_operate_role(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperateRoleSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/ListAccess" => {
                    #[allow(non_camel_case_types)]
                    struct ListAccessSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::ListAccessParams>
                    for ListAccessSvc<T> {
                        type Response = super::ListAccessReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListAccessParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move { (*inner).list_access(request).await };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListAccessSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/OperateAccess" => {
                    #[allow(non_camel_case_types)]
                    struct OperateAccessSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::OperateAccessParams>
                    for OperateAccessSvc<T> {
                        type Response = super::OperateAccessReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperateAccessParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).operate_access(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperateAccessSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/BatchOperateAccess" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperateAccessSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::BatchOperateAccessParams>
                    for BatchOperateAccessSvc<T> {
                        type Response = super::BatchOperateAccessReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperateAccessParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).batch_operate_access(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperateAccessSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/ListPermission" => {
                    #[allow(non_camel_case_types)]
                    struct ListPermissionSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::ListPermissionParams>
                    for ListPermissionSvc<T> {
                        type Response = super::ListPermissionReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListPermissionParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).list_permission(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListPermissionSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/OperatePermission" => {
                    #[allow(non_camel_case_types)]
                    struct OperatePermissionSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::OperatePermissionParams>
                    for OperatePermissionSvc<T> {
                        type Response = super::OperatePermissionReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperatePermissionParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).operate_permission(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperatePermissionSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/BatchOperatePermission" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperatePermissionSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::BatchOperatePermissionParams>
                    for BatchOperatePermissionSvc<T> {
                        type Response = super::BatchOperatePermissionReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperatePermissionParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).batch_operate_permission(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperatePermissionSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/auth_service.AuthService/Authorization" => {
                    #[allow(non_camel_case_types)]
                    struct AuthorizationSvc<T: AuthService>(pub Arc<T>);
                    impl<
                        T: AuthService,
                    > tonic::server::UnaryService<super::AuthorizationParams>
                    for AuthorizationSvc<T> {
                        type Response = super::AuthorizationReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::AuthorizationParams>,
                        ) -> Self::Future {
                            let inner = self.0.clone();
                            let fut = async move {
                                (*inner).authorization(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = AuthorizationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", "12")
                                .header("content-type", "application/grpc")
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: AuthService> Clone for AuthServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
            }
        }
    }
    impl<T: AuthService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(self.0.clone())
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: AuthService> tonic::server::NamedService for AuthServiceServer<T> {
        const NAME: &'static str = "auth_service.AuthService";
    }
}
