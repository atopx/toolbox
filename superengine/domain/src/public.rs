#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ECode {
    /// 默认无错误
    Success = 0,
    /// 客户端错误
    ///
    /// 请求异常
    BadRequest = 40000,
    /// 无效的请求参数
    InvalidParams = 40001,
    /// 查询的记录不存在
    RecodeNotFound = 40002,
    /// 资源冲突
    RequestConflict = 40003,
    /// 账户未登录
    AuthNotin = 40100,
    /// 无效Token
    AuthInvalid = 40102,
    /// 账户已过期
    AuthExpired = 40103,
    /// 限制登录
    AuthLimit = 40104,
    /// 不是access-token
    AuthNotAccessToken = 40105,
    /// 禁止访问
    AccessForbidden = 40300,
    /// 服务端错误
    SystemError = 50000,
    DbError = 50100,
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
            ECode::InvalidParams => "INVALID_PARAMS",
            ECode::RecodeNotFound => "RECODE_NOT_FOUND",
            ECode::RequestConflict => "REQUEST_CONFLICT",
            ECode::AuthNotin => "AUTH_NOTIN",
            ECode::AuthInvalid => "AUTH_INVALID",
            ECode::AuthExpired => "AUTH_EXPIRED",
            ECode::AuthLimit => "AUTH_LIMIT",
            ECode::AuthNotAccessToken => "AUTH_NOT_ACCESS_TOKEN",
            ECode::AccessForbidden => "ACCESS_FORBIDDEN",
            ECode::SystemError => "SYSTEM_ERROR",
            ECode::DbError => "DB_ERROR",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "SUCCESS" => Some(Self::Success),
            "BAD_REQUEST" => Some(Self::BadRequest),
            "INVALID_PARAMS" => Some(Self::InvalidParams),
            "RECODE_NOT_FOUND" => Some(Self::RecodeNotFound),
            "REQUEST_CONFLICT" => Some(Self::RequestConflict),
            "AUTH_NOTIN" => Some(Self::AuthNotin),
            "AUTH_INVALID" => Some(Self::AuthInvalid),
            "AUTH_EXPIRED" => Some(Self::AuthExpired),
            "AUTH_LIMIT" => Some(Self::AuthLimit),
            "AUTH_NOT_ACCESS_TOKEN" => Some(Self::AuthNotAccessToken),
            "ACCESS_FORBIDDEN" => Some(Self::AccessForbidden),
            "SYSTEM_ERROR" => Some(Self::SystemError),
            "DB_ERROR" => Some(Self::DbError),
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
    #[prost(enumeration = "ECode", tag = "1")]
    pub code: i32,
    #[prost(string, tag = "2")]
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
