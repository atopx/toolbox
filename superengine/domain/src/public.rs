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
    #[prost(enumeration = "super::ecode::ECode", tag = "2")]
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
    /// 默认
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
pub enum BooleanScope {
    /// 默认
    BoolAll = 0,
    /// FALSE
    BoolFalse = 1,
    /// TRUE
    BoolTrue = 2,
}
impl BooleanScope {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            BooleanScope::BoolAll => "BOOL_ALL",
            BooleanScope::BoolFalse => "BOOL_FALSE",
            BooleanScope::BoolTrue => "BOOL_TRUE",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "BOOL_ALL" => Some(Self::BoolAll),
            "BOOL_FALSE" => Some(Self::BoolFalse),
            "BOOL_TRUE" => Some(Self::BoolTrue),
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
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum TaskStatus {
    Pending = 0,
    Running = 1,
    Success = 2,
    Failure = 4,
    Finally = 5,
}
impl TaskStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            TaskStatus::Pending => "PENDING",
            TaskStatus::Running => "RUNNING",
            TaskStatus::Success => "SUCCESS",
            TaskStatus::Failure => "FAILURE",
            TaskStatus::Finally => "FINALLY",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "PENDING" => Some(Self::Pending),
            "RUNNING" => Some(Self::Running),
            "SUCCESS" => Some(Self::Success),
            "FAILURE" => Some(Self::Failure),
            "FINALLY" => Some(Self::Finally),
            _ => None,
        }
    }
}
