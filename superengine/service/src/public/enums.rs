/// 根据proto枚举自动生成, 构建代码查看 build.rs
use domain::public_service::{Enum, Enums, ListEnumReply};


fn get_accessstatus() -> Enums {
    let data = vec![
        Enum { code: "ACCESS_DEFAULT".to_string(), value: 0, desc: "默认状态, 需要验证权限".to_string() },
        Enum { code: "ACCESS_ANONYMOUS".to_string(), value: 1, desc: "可匿名访问，无需验证权限".to_string() },
        Enum { code: "ACCESS_DISABLED".to_string(), value: 2, desc: "已禁用，禁止访问".to_string() },
    ];
    Enums { key: "AccessStatus".to_string(), data }
}

fn get_source() -> Enums {
    let data = vec![
        Enum { code: "Public".to_string(), value: 0, desc: "公共".to_string() },
        Enum { code: "User".to_string(), value: 1, desc: "用户".to_string() },
        Enum { code: "Note".to_string(), value: 2, desc: "笔记".to_string() },
        Enum { code: "File".to_string(), value: 3, desc: "文件".to_string() },
    ];
    Enums { key: "Source".to_string(), data }
}

fn get_ecode() -> Enums {
    let data = vec![
        Enum { code: "SUCCESS".to_string(), value: 0, desc: "请求成功".to_string() },
        Enum { code: "BAD_REQUEST".to_string(), value: 400000, desc: "请求异常".to_string() },
        Enum { code: "AUTH_TOKEN_Required".to_string(), value: 400001, desc: "账户未登录".to_string() },
        Enum { code: "AUTH_TOKEN_Invalid".to_string(), value: 400002, desc: "无效Token".to_string() },
        Enum { code: "AUTH_TOKEN_NotFound".to_string(), value: 400003, desc: "token不存在".to_string() },
        Enum { code: "AUTH_TOKEN_Expired".to_string(), value: 400004, desc: "token过期".to_string() },
        Enum { code: "ACCESS_NotFound".to_string(), value: 401001, desc: "请求资源不存在".to_string() },
        Enum { code: "ACCESS_Forbidden".to_string(), value: 401002, desc: "没有访问权限".to_string() },
        Enum { code: "ACCESS_Disabled".to_string(), value: 401003, desc: "接口已禁用".to_string() },
        Enum { code: "INVALID_PARAMS".to_string(), value: 431000, desc: "参数异常".to_string() },
        Enum { code: "USER_PARAMS_ERROR_NameRequired".to_string(), value: 431001, desc: "参数缺少昵称".to_string() },
        Enum { code: "USER_PARAMS_ERROR_UsernameRequired".to_string(), value: 431002, desc: "参数缺少用户名".to_string() },
        Enum { code: "USER_PARAMS_ERROR_PasswordRequired".to_string(), value: 431003, desc: "参数缺少密码".to_string() },
        Enum { code: "USER_PARAMS_ERROR_UsernameExist".to_string(), value: 431004, desc: "用户名已存在".to_string() },
        Enum { code: "USER_PARAMS_ERROR_UserNotFound".to_string(), value: 431005, desc: "用户不存在".to_string() },
        Enum { code: "USER_PARAMS_ERROR_NameInvalid".to_string(), value: 431006, desc: "无效的用户昵称".to_string() },
        Enum { code: "USER_PARAMS_ERROR_UsernameInvalid".to_string(), value: 431007, desc: "无效的用户名".to_string() },
        Enum { code: "USER_PARAMS_ERROR_PasswordInvalid".to_string(), value: 431008, desc: "无效的密码".to_string() },
        Enum { code: "SOURCE_PARAMS_ERROR_SourceInvalid".to_string(), value: 432001, desc: "无效的实体来源".to_string() },
        Enum { code: "FOLDER_PARAMS_ERROR_NameRequired".to_string(), value: 433001, desc: "参数缺少文件夹名称".to_string() },
        Enum { code: "LABEL_PARAMS_ERROR_NameRequired".to_string(), value: 434001, desc: "参数缺少文件夹名称".to_string() },
        Enum { code: "NOTE_PARAMS_ERROR_TitleRequired".to_string(), value: 435001, desc: "参数缺少标题".to_string() },
        Enum { code: "NOTE_PARAMS_ERROR_TitleInvalid".to_string(), value: 435002, desc: "标题无效".to_string() },
        Enum { code: "ROLE_PARAMS_ERROR_NameInvalid".to_string(), value: 436001, desc: "无效的角色名".to_string() },
        Enum { code: "ROLE_PARAMS_ERROR_NameExist".to_string(), value: 436002, desc: "角色名已存在".to_string() },
        Enum { code: "SYSTEM_INTERNAL_ERROR".to_string(), value: 500000, desc: "服务异常".to_string() },
        Enum { code: "SYSTEM_ERROR_Unimplemented".to_string(), value: 500001, desc: "功能未实现".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListAccess".to_string(), value: 510001, desc: "AUTH_SERVICE_ERROR_ListAccess".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperateAccess".to_string(), value: 510002, desc: "AUTH_SERVICE_ERROR_OperateAccess".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperateAccess".to_string(), value: 510003, desc: "AUTH_SERVICE_ERROR_BatchOperateAccess".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListAuthToken".to_string(), value: 511001, desc: "AUTH_SERVICE_ERROR_ListAuthToken".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperateAuthToken".to_string(), value: 511002, desc: "AUTH_SERVICE_ERROR_OperateAuthToken".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperateAuthToken".to_string(), value: 511003, desc: "AUTH_SERVICE_ERROR_BatchOperateAuthToken".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListUser".to_string(), value: 512001, desc: "AUTH_SERVICE_ERROR_ListUser".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperateUser".to_string(), value: 512002, desc: "AUTH_SERVICE_ERROR_OperateUser".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperateUser".to_string(), value: 512003, desc: "AUTH_SERVICE_ERROR_BatchOperateUser".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListRole".to_string(), value: 513001, desc: "AUTH_SERVICE_ERROR_ListRole".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperateRole".to_string(), value: 513002, desc: "AUTH_SERVICE_ERROR_OperateRole".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperateRole".to_string(), value: 513003, desc: "AUTH_SERVICE_ERROR_BatchOperateRole".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListUserRoleRef".to_string(), value: 514001, desc: "AUTH_SERVICE_ERROR_ListUserRoleRef".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperateUserRoleRef".to_string(), value: 514002, desc: "AUTH_SERVICE_ERROR_OperateUserRoleRef".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef".to_string(), value: 514003, desc: "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_ListPermission".to_string(), value: 515001, desc: "AUTH_SERVICE_ERROR_ListPermission".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_OperatePermission".to_string(), value: 515002, desc: "AUTH_SERVICE_ERROR_OperatePermission".to_string() },
        Enum { code: "AUTH_SERVICE_ERROR_BatchOperatePermission".to_string(), value: 515003, desc: "AUTH_SERVICE_ERROR_BatchOperatePermission".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_ListLabel".to_string(), value: 520001, desc: "PUBLIC_SERVICE_ERROR_ListLabel".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_OperateLabel".to_string(), value: 520002, desc: "PUBLIC_SERVICE_ERROR_OperateLabel".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_BatchOperateLabel".to_string(), value: 520003, desc: "PUBLIC_SERVICE_ERROR_BatchOperateLabel".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_ListFolder".to_string(), value: 521001, desc: "PUBLIC_SERVICE_ERROR_ListFolder".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_OperateFolder".to_string(), value: 521002, desc: "PUBLIC_SERVICE_ERROR_OperateFolder".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_BatchOperateFolder".to_string(), value: 521003, desc: "PUBLIC_SERVICE_ERROR_BatchOperateFolder".to_string() },
        Enum { code: "PUBLIC_SERVICE_ERROR_ListEnum".to_string(), value: 522001, desc: "PUBLIC_SERVICE_ERROR_ListEnum".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_ListNote".to_string(), value: 530001, desc: "NOTE_SERVICE_ERROR_ListNote".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_OperateNote".to_string(), value: 530002, desc: "NOTE_SERVICE_ERROR_OperateNote".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_BatchOperateNote".to_string(), value: 530003, desc: "NOTE_SERVICE_ERROR_BatchOperateNote".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_ListNoteLabel".to_string(), value: 531001, desc: "NOTE_SERVICE_ERROR_ListNoteLabel".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_OperateNoteLabel".to_string(), value: 531002, desc: "NOTE_SERVICE_ERROR_OperateNoteLabel".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_BatchOperateNoteLabel".to_string(), value: 531003, desc: "NOTE_SERVICE_ERROR_BatchOperateNoteLabel".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_ListNoteTopic".to_string(), value: 532001, desc: "NOTE_SERVICE_ERROR_ListNoteTopic".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_OperateNoteTopic".to_string(), value: 532002, desc: "NOTE_SERVICE_ERROR_OperateNoteTopic".to_string() },
        Enum { code: "NOTE_SERVICE_ERROR_BatchOperateNoteTopic".to_string(), value: 532003, desc: "NOTE_SERVICE_ERROR_BatchOperateNoteTopic".to_string() },
    ];
    Enums { key: "ECode".to_string(), data }
}

fn get_userstatus() -> Enums {
    let data = vec![
        Enum { code: "USER_NORMAL".to_string(), value: 0, desc: "正常".to_string() },
        Enum { code: "USER_STATUS_DISABLED".to_string(), value: 1, desc: "禁用".to_string() },
        Enum { code: "USER_STATUS_BLACK".to_string(), value: 2, desc: "黑名单".to_string() },
    ];
    Enums { key: "UserStatus".to_string(), data }
}

fn get_rolenature() -> Enums {
    let data = vec![
        Enum { code: "ROLE_DEFAULT".to_string(), value: 0, desc: "初始化角色".to_string() },
        Enum { code: "ROLE_SYSTEM".to_string(), value: 1, desc: "系统角色".to_string() },
        Enum { code: "ROLE_CUSTOM".to_string(), value: 2, desc: "自定义角色".to_string() },
    ];
    Enums { key: "RoleNature".to_string(), data }
}

pub fn list_enum(trace_id: i64) -> ListEnumReply {
    ListEnumReply {
        header: common::header::reply(trace_id),
        data: vec![
            get_accessstatus(),
            get_source(),
            get_ecode(),
            get_userstatus(),
            get_rolenature(),
        ],
    }
}
