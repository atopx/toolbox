#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ECode {
    /// 默认无错误
    Success = 0,
    /// ######### 客户端错误 #########
    ///
    /// 请求异常
    BadRequest = 400000,
    /// auth token error
    ///
    /// 账户未登录
    AuthTokenRequired = 400001,
    /// 无效Token
    AuthTokenInvalid = 400002,
    /// token不存在
    AuthTokenNotFound = 400003,
    /// token过期
    AuthTokenExpired = 400004,
    /// access error
    AccessNotFound = 401001,
    /// 没有访问权限
    AccessForbidden = 401002,
    /// 接口已禁用
    AccessDisabled = 401003,
    /// ######### 参数错误 #########
    ///
    /// 参数异常
    InvalidParams = 431000,
    /// user params error
    UserParamsErrorNameRequired = 431001,
    UserParamsErrorUsernameRequired = 431002,
    UserParamsErrorPasswordRequired = 431003,
    UserParamsErrorUsernameExist = 431004,
    UserParamsErrorUserNotFound = 431005,
    /// user login
    UserParamsErrorNameInvalid = 431006,
    UserParamsErrorUsernameInvalid = 431007,
    UserParamsErrorPasswordInvalid = 431008,
    /// ######### 服务端错误 #########
    ///
    /// 服务异常
    SystemInternalError = 500000,
    /// 功能未实现
    SystemErrorUnimplemented = 500001,
    /// auth-service access error
    AuthServiceErrorListAccess = 510001,
    AuthServiceErrorOperateAccess = 510002,
    AuthServiceErrorBatchOperateAccess = 510003,
    /// auth-service token error
    AuthServiceErrorListAuthToken = 511001,
    AuthServiceErrorOperateAuthToken = 511002,
    AuthServiceErrorBatchOperateAuthToken = 511003,
    /// auth-service user error
    AuthServiceErrorListUser = 512001,
    AuthServiceErrorOperateUser = 512002,
    AuthServiceErrorBatchOperateUser = 512003,
    /// auth-service role error
    AuthServiceErrorListRole = 513001,
    AuthServiceErrorOperateRole = 513002,
    AuthServiceErrorBatchOperateRole = 513003,
    /// auth-service permission error
    AuthServiceErrorListUserRoleRef = 514001,
    AuthServiceErrorOperateUserRoleRef = 514002,
    AuthServiceErrorBatchOperateUserRoleRef = 514003,
    /// auth-service permission error
    AuthServiceErrorListPermission = 515001,
    AuthServiceErrorOperatePermission = 515002,
    AuthServiceErrorBatchOperatePermission = 515003,
}
impl ECode {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            ECode::Success => "SUCCESS",
            ECode::BadRequest => "BAD_REQUEST",
            ECode::AuthTokenRequired => "AUTH_TOKEN_Required",
            ECode::AuthTokenInvalid => "AUTH_TOKEN_Invalid",
            ECode::AuthTokenNotFound => "AUTH_TOKEN_NotFound",
            ECode::AuthTokenExpired => "AUTH_TOKEN_Expired",
            ECode::AccessNotFound => "ACCESS_NotFound",
            ECode::AccessForbidden => "ACCESS_Forbidden",
            ECode::AccessDisabled => "ACCESS_Disabled",
            ECode::InvalidParams => "INVALID_PARAMS",
            ECode::UserParamsErrorNameRequired => "USER_PARAMS_ERROR_NameRequired",
            ECode::UserParamsErrorUsernameRequired => {
                "USER_PARAMS_ERROR_UsernameRequired"
            }
            ECode::UserParamsErrorPasswordRequired => {
                "USER_PARAMS_ERROR_PasswordRequired"
            }
            ECode::UserParamsErrorUsernameExist => "USER_PARAMS_ERROR_UsernameExist",
            ECode::UserParamsErrorUserNotFound => "USER_PARAMS_ERROR_UserNotFound",
            ECode::UserParamsErrorNameInvalid => "USER_PARAMS_ERROR_NameInvalid",
            ECode::UserParamsErrorUsernameInvalid => "USER_PARAMS_ERROR_UsernameInvalid",
            ECode::UserParamsErrorPasswordInvalid => "USER_PARAMS_ERROR_PasswordInvalid",
            ECode::SystemInternalError => "SYSTEM_INTERNAL_ERROR",
            ECode::SystemErrorUnimplemented => "SYSTEM_ERROR_Unimplemented",
            ECode::AuthServiceErrorListAccess => "AUTH_SERVICE_ERROR_ListAccess",
            ECode::AuthServiceErrorOperateAccess => "AUTH_SERVICE_ERROR_OperateAccess",
            ECode::AuthServiceErrorBatchOperateAccess => {
                "AUTH_SERVICE_ERROR_BatchOperateAccess"
            }
            ECode::AuthServiceErrorListAuthToken => "AUTH_SERVICE_ERROR_ListAuthToken",
            ECode::AuthServiceErrorOperateAuthToken => {
                "AUTH_SERVICE_ERROR_OperateAuthToken"
            }
            ECode::AuthServiceErrorBatchOperateAuthToken => {
                "AUTH_SERVICE_ERROR_BatchOperateAuthToken"
            }
            ECode::AuthServiceErrorListUser => "AUTH_SERVICE_ERROR_ListUser",
            ECode::AuthServiceErrorOperateUser => "AUTH_SERVICE_ERROR_OperateUser",
            ECode::AuthServiceErrorBatchOperateUser => {
                "AUTH_SERVICE_ERROR_BatchOperateUser"
            }
            ECode::AuthServiceErrorListRole => "AUTH_SERVICE_ERROR_ListRole",
            ECode::AuthServiceErrorOperateRole => "AUTH_SERVICE_ERROR_OperateRole",
            ECode::AuthServiceErrorBatchOperateRole => {
                "AUTH_SERVICE_ERROR_BatchOperateRole"
            }
            ECode::AuthServiceErrorListUserRoleRef => {
                "AUTH_SERVICE_ERROR_ListUserRoleRef"
            }
            ECode::AuthServiceErrorOperateUserRoleRef => {
                "AUTH_SERVICE_ERROR_OperateUserRoleRef"
            }
            ECode::AuthServiceErrorBatchOperateUserRoleRef => {
                "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef"
            }
            ECode::AuthServiceErrorListPermission => "AUTH_SERVICE_ERROR_ListPermission",
            ECode::AuthServiceErrorOperatePermission => {
                "AUTH_SERVICE_ERROR_OperatePermission"
            }
            ECode::AuthServiceErrorBatchOperatePermission => {
                "AUTH_SERVICE_ERROR_BatchOperatePermission"
            }
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "SUCCESS" => Some(Self::Success),
            "BAD_REQUEST" => Some(Self::BadRequest),
            "AUTH_TOKEN_Required" => Some(Self::AuthTokenRequired),
            "AUTH_TOKEN_Invalid" => Some(Self::AuthTokenInvalid),
            "AUTH_TOKEN_NotFound" => Some(Self::AuthTokenNotFound),
            "AUTH_TOKEN_Expired" => Some(Self::AuthTokenExpired),
            "ACCESS_NotFound" => Some(Self::AccessNotFound),
            "ACCESS_Forbidden" => Some(Self::AccessForbidden),
            "ACCESS_Disabled" => Some(Self::AccessDisabled),
            "INVALID_PARAMS" => Some(Self::InvalidParams),
            "USER_PARAMS_ERROR_NameRequired" => Some(Self::UserParamsErrorNameRequired),
            "USER_PARAMS_ERROR_UsernameRequired" => {
                Some(Self::UserParamsErrorUsernameRequired)
            }
            "USER_PARAMS_ERROR_PasswordRequired" => {
                Some(Self::UserParamsErrorPasswordRequired)
            }
            "USER_PARAMS_ERROR_UsernameExist" => Some(Self::UserParamsErrorUsernameExist),
            "USER_PARAMS_ERROR_UserNotFound" => Some(Self::UserParamsErrorUserNotFound),
            "USER_PARAMS_ERROR_NameInvalid" => Some(Self::UserParamsErrorNameInvalid),
            "USER_PARAMS_ERROR_UsernameInvalid" => {
                Some(Self::UserParamsErrorUsernameInvalid)
            }
            "USER_PARAMS_ERROR_PasswordInvalid" => {
                Some(Self::UserParamsErrorPasswordInvalid)
            }
            "SYSTEM_INTERNAL_ERROR" => Some(Self::SystemInternalError),
            "SYSTEM_ERROR_Unimplemented" => Some(Self::SystemErrorUnimplemented),
            "AUTH_SERVICE_ERROR_ListAccess" => Some(Self::AuthServiceErrorListAccess),
            "AUTH_SERVICE_ERROR_OperateAccess" => {
                Some(Self::AuthServiceErrorOperateAccess)
            }
            "AUTH_SERVICE_ERROR_BatchOperateAccess" => {
                Some(Self::AuthServiceErrorBatchOperateAccess)
            }
            "AUTH_SERVICE_ERROR_ListAuthToken" => {
                Some(Self::AuthServiceErrorListAuthToken)
            }
            "AUTH_SERVICE_ERROR_OperateAuthToken" => {
                Some(Self::AuthServiceErrorOperateAuthToken)
            }
            "AUTH_SERVICE_ERROR_BatchOperateAuthToken" => {
                Some(Self::AuthServiceErrorBatchOperateAuthToken)
            }
            "AUTH_SERVICE_ERROR_ListUser" => Some(Self::AuthServiceErrorListUser),
            "AUTH_SERVICE_ERROR_OperateUser" => Some(Self::AuthServiceErrorOperateUser),
            "AUTH_SERVICE_ERROR_BatchOperateUser" => {
                Some(Self::AuthServiceErrorBatchOperateUser)
            }
            "AUTH_SERVICE_ERROR_ListRole" => Some(Self::AuthServiceErrorListRole),
            "AUTH_SERVICE_ERROR_OperateRole" => Some(Self::AuthServiceErrorOperateRole),
            "AUTH_SERVICE_ERROR_BatchOperateRole" => {
                Some(Self::AuthServiceErrorBatchOperateRole)
            }
            "AUTH_SERVICE_ERROR_ListUserRoleRef" => {
                Some(Self::AuthServiceErrorListUserRoleRef)
            }
            "AUTH_SERVICE_ERROR_OperateUserRoleRef" => {
                Some(Self::AuthServiceErrorOperateUserRoleRef)
            }
            "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef" => {
                Some(Self::AuthServiceErrorBatchOperateUserRoleRef)
            }
            "AUTH_SERVICE_ERROR_ListPermission" => {
                Some(Self::AuthServiceErrorListPermission)
            }
            "AUTH_SERVICE_ERROR_OperatePermission" => {
                Some(Self::AuthServiceErrorOperatePermission)
            }
            "AUTH_SERVICE_ERROR_BatchOperatePermission" => {
                Some(Self::AuthServiceErrorBatchOperatePermission)
            }
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Pager {
    /// 页码
    #[prost(uint64, tag = "1")]
    pub index: u64,
    /// 数量
    #[prost(uint64, tag = "2")]
    pub size: u64,
    /// 总数量
    #[prost(uint64, tag = "3")]
    pub count: u64,
    /// 禁用分页
    #[prost(bool, tag = "5")]
    pub disabled: bool,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RangeI64 {
    /// 区间左
    #[prost(int64, tag = "1")]
    pub left: i64,
    /// 区间右
    #[prost(int64, tag = "2")]
    pub right: i64,
    /// 应用范围
    #[prost(enumeration = "RangeScope", tag = "3")]
    pub scope: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Header {
    #[prost(int64, tag = "1")]
    pub trace_id: i64,
    #[prost(string, tag = "2")]
    pub source: ::prost::alloc::string::String,
    #[prost(int32, tag = "3")]
    pub operator: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ReplyHeader {
    #[prost(int64, tag = "1")]
    pub trace_id: i64,
    #[prost(enumeration = "ECode", tag = "2")]
    pub code: i32,
    #[prost(string, tag = "3")]
    pub message: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BetweenInt64 {
    #[prost(enumeration = "RangeScope", tag = "1")]
    pub scope: i32,
    #[prost(int64, tag = "2")]
    pub left: i64,
    #[prost(int64, tag = "3")]
    pub right: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Sort {
    #[prost(string, tag = "1")]
    pub field: ::prost::alloc::string::String,
    #[prost(enumeration = "SortDirection", tag = "2")]
    pub direction: i32,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum RangeScope {
    /// 含左右
    RangeAll = 0,
    /// 含左
    RangeLeft = 1,
    /// 含右
    RangeRight = 2,
}
impl RangeScope {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            RangeScope::RangeAll => "RANGE_ALL",
            RangeScope::RangeLeft => "RANGE_LEFT",
            RangeScope::RangeRight => "RANGE_RIGHT",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "RANGE_ALL" => Some(Self::RangeAll),
            "RANGE_LEFT" => Some(Self::RangeLeft),
            "RANGE_RIGHT" => Some(Self::RangeRight),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum SortDirection {
    SortAsc = 0,
    SortDesc = 1,
}
impl SortDirection {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            SortDirection::SortAsc => "SORT_ASC",
            SortDirection::SortDesc => "SORT_DESC",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "SORT_ASC" => Some(Self::SortAsc),
            "SORT_DESC" => Some(Self::SortDesc),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Operation {
    /// 基础操作类型
    Create = 0,
    Update = 1,
    Delete = 2,
    RealDelete = 3,
    /// 高阶操作类型
    Upsert = 101,
}
impl Operation {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Operation::Create => "OPERATION_CREATE",
            Operation::Update => "OPERATION_UPDATE",
            Operation::Delete => "OPERATION_DELETE",
            Operation::RealDelete => "OPERATION_REAL_DELETE",
            Operation::Upsert => "OPERATION_UPSERT",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "OPERATION_CREATE" => Some(Self::Create),
            "OPERATION_UPDATE" => Some(Self::Update),
            "OPERATION_DELETE" => Some(Self::Delete),
            "OPERATION_REAL_DELETE" => Some(Self::RealDelete),
            "OPERATION_UPSERT" => Some(Self::Upsert),
            _ => None,
        }
    }
}
