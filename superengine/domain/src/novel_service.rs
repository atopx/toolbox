#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum NovelStatus {
    NovelNone = 0,
    /// 启用监控
    NovelEnabled = 1,
    /// 手动禁止监控
    NovelDisabled = 2,
    /// 连续N次未发生变化，停止监控
    NovelSuccess = 3,
    /// 连续N次错误, 停止监控
    NovelError = 4,
}
impl NovelStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            NovelStatus::NovelNone => "NOVEL_NONE",
            NovelStatus::NovelEnabled => "NOVEL_ENABLED",
            NovelStatus::NovelDisabled => "NOVEL_DISABLED",
            NovelStatus::NovelSuccess => "NOVEL_SUCCESS",
            NovelStatus::NovelError => "NOVEL_ERROR",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "NOVEL_NONE" => Some(Self::NovelNone),
            "NOVEL_ENABLED" => Some(Self::NovelEnabled),
            "NOVEL_DISABLED" => Some(Self::NovelDisabled),
            "NOVEL_SUCCESS" => Some(Self::NovelSuccess),
            "NOVEL_ERROR" => Some(Self::NovelError),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ScanStatus {
    ScanNone = 0,
    ScanPending = 1,
    ScanRunning = 2,
    ScanSuccess = 3,
    ScanError = 4,
}
impl ScanStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            ScanStatus::ScanNone => "SCAN_NONE",
            ScanStatus::ScanPending => "SCAN_PENDING",
            ScanStatus::ScanRunning => "SCAN_RUNNING",
            ScanStatus::ScanSuccess => "SCAN_SUCCESS",
            ScanStatus::ScanError => "SCAN_ERROR",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "SCAN_NONE" => Some(Self::ScanNone),
            "SCAN_PENDING" => Some(Self::ScanPending),
            "SCAN_RUNNING" => Some(Self::ScanRunning),
            "SCAN_SUCCESS" => Some(Self::ScanSuccess),
            "SCAN_ERROR" => Some(Self::ScanError),
            _ => None,
        }
    }
}
