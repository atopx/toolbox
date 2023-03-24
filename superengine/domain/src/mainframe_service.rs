#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum MainframePowerStatus {
    /// 未知状态
    MainframePowerUnknown = 0,
    /// 关机状态
    MainframePowerOff = 1,
    /// 开机状态
    MainframePowerOn = 2,
}
impl MainframePowerStatus {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            MainframePowerStatus::MainframePowerUnknown => "MAINFRAME_POWER_UNKNOWN",
            MainframePowerStatus::MainframePowerOff => "MAINFRAME_POWER_OFF",
            MainframePowerStatus::MainframePowerOn => "MAINFRAME_POWER_ON",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "MAINFRAME_POWER_UNKNOWN" => Some(Self::MainframePowerUnknown),
            "MAINFRAME_POWER_OFF" => Some(Self::MainframePowerOff),
            "MAINFRAME_POWER_ON" => Some(Self::MainframePowerOn),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum MainframeOperate {
    /// 关机
    MainframeClose = 0,
    /// 开机
    MainframeOpen = 1,
    /// 探测
    MainframeDetect = 2,
}
impl MainframeOperate {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            MainframeOperate::MainframeClose => "MAINFRAME_CLOSE",
            MainframeOperate::MainframeOpen => "MAINFRAME_OPEN",
            MainframeOperate::MainframeDetect => "MAINFRAME_DETECT",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "MAINFRAME_CLOSE" => Some(Self::MainframeClose),
            "MAINFRAME_OPEN" => Some(Self::MainframeOpen),
            "MAINFRAME_DETECT" => Some(Self::MainframeDetect),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum MainframePortProtocol {
    Unknown = 0,
    Ftpdata = 20,
    Ftpctl = 21,
    Ssh = 22,
    Telnet = 23,
    Smtp = 25,
    Dns = 53,
    Dhcp = 67,
    Http = 80,
    Pop3 = 110,
    Ntp = 123,
    Imap = 143,
    Snmp = 161,
    Https = 443,
    Ipsec = 500,
    Radius = 1645,
    L2tp = 1701,
    Pptp = 1723,
    Rdp = 3389,
}
impl MainframePortProtocol {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            MainframePortProtocol::Unknown => "MAINFRAME_PORT_PROTOCOL_UNKNOWN",
            MainframePortProtocol::Ftpdata => "MAINFRAME_PORT_PROTOCOL_FTPDATA",
            MainframePortProtocol::Ftpctl => "MAINFRAME_PORT_PROTOCOL_FTPCTL",
            MainframePortProtocol::Ssh => "MAINFRAME_PORT_PROTOCOL_SSH",
            MainframePortProtocol::Telnet => "MAINFRAME_PORT_PROTOCOL_TELNET",
            MainframePortProtocol::Smtp => "MAINFRAME_PORT_PROTOCOL_SMTP",
            MainframePortProtocol::Dns => "MAINFRAME_PORT_PROTOCOL_DNS",
            MainframePortProtocol::Dhcp => "MAINFRAME_PORT_PROTOCOL_DHCP",
            MainframePortProtocol::Http => "MAINFRAME_PORT_PROTOCOL_HTTP",
            MainframePortProtocol::Pop3 => "MAINFRAME_PORT_PROTOCOL_POP3",
            MainframePortProtocol::Ntp => "MAINFRAME_PORT_PROTOCOL_NTP",
            MainframePortProtocol::Imap => "MAINFRAME_PORT_PROTOCOL_IMAP",
            MainframePortProtocol::Snmp => "MAINFRAME_PORT_PROTOCOL_SNMP",
            MainframePortProtocol::Https => "MAINFRAME_PORT_PROTOCOL_HTTPS",
            MainframePortProtocol::Ipsec => "MAINFRAME_PORT_PROTOCOL_IPSEC",
            MainframePortProtocol::Radius => "MAINFRAME_PORT_PROTOCOL_RADIUS",
            MainframePortProtocol::L2tp => "MAINFRAME_PORT_PROTOCOL_L2TP",
            MainframePortProtocol::Pptp => "MAINFRAME_PORT_PROTOCOL_PPTP",
            MainframePortProtocol::Rdp => "MAINFRAME_PORT_PROTOCOL_RDP",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "MAINFRAME_PORT_PROTOCOL_UNKNOWN" => Some(Self::Unknown),
            "MAINFRAME_PORT_PROTOCOL_FTPDATA" => Some(Self::Ftpdata),
            "MAINFRAME_PORT_PROTOCOL_FTPCTL" => Some(Self::Ftpctl),
            "MAINFRAME_PORT_PROTOCOL_SSH" => Some(Self::Ssh),
            "MAINFRAME_PORT_PROTOCOL_TELNET" => Some(Self::Telnet),
            "MAINFRAME_PORT_PROTOCOL_SMTP" => Some(Self::Smtp),
            "MAINFRAME_PORT_PROTOCOL_DNS" => Some(Self::Dns),
            "MAINFRAME_PORT_PROTOCOL_DHCP" => Some(Self::Dhcp),
            "MAINFRAME_PORT_PROTOCOL_HTTP" => Some(Self::Http),
            "MAINFRAME_PORT_PROTOCOL_POP3" => Some(Self::Pop3),
            "MAINFRAME_PORT_PROTOCOL_NTP" => Some(Self::Ntp),
            "MAINFRAME_PORT_PROTOCOL_IMAP" => Some(Self::Imap),
            "MAINFRAME_PORT_PROTOCOL_SNMP" => Some(Self::Snmp),
            "MAINFRAME_PORT_PROTOCOL_HTTPS" => Some(Self::Https),
            "MAINFRAME_PORT_PROTOCOL_IPSEC" => Some(Self::Ipsec),
            "MAINFRAME_PORT_PROTOCOL_RADIUS" => Some(Self::Radius),
            "MAINFRAME_PORT_PROTOCOL_L2TP" => Some(Self::L2tp),
            "MAINFRAME_PORT_PROTOCOL_PPTP" => Some(Self::Pptp),
            "MAINFRAME_PORT_PROTOCOL_RDP" => Some(Self::Rdp),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum MainframePortTransport {
    MainframeTransportUnknown = 0,
    MainframeTransportTcp = 1,
    MainframeTransportUdp = 2,
}
impl MainframePortTransport {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            MainframePortTransport::MainframeTransportUnknown => {
                "MAINFRAME_TRANSPORT_UNKNOWN"
            }
            MainframePortTransport::MainframeTransportTcp => "MAINFRAME_TRANSPORT_TCP",
            MainframePortTransport::MainframeTransportUdp => "MAINFRAME_TRANSPORT_UDP",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "MAINFRAME_TRANSPORT_UNKNOWN" => Some(Self::MainframeTransportUnknown),
            "MAINFRAME_TRANSPORT_TCP" => Some(Self::MainframeTransportTcp),
            "MAINFRAME_TRANSPORT_UDP" => Some(Self::MainframeTransportUdp),
            _ => None,
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum MainframePortUseType {
    /// 空
    MainframePortUsetypeNone = 0,
    /// 网络探测
    MainframePortUsetypeDetect = 1,
}
impl MainframePortUseType {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            MainframePortUseType::MainframePortUsetypeNone => {
                "MAINFRAME_PORT_USETYPE_NONE"
            }
            MainframePortUseType::MainframePortUsetypeDetect => {
                "MAINFRAME_PORT_USETYPE_DETECT"
            }
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "MAINFRAME_PORT_USETYPE_NONE" => Some(Self::MainframePortUsetypeNone),
            "MAINFRAME_PORT_USETYPE_DETECT" => Some(Self::MainframePortUsetypeDetect),
            _ => None,
        }
    }
}
