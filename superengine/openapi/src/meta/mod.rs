/// 根据proto枚举自动生成, 构建代码查看 build.rs
use axum::{Json, http::StatusCode};
use axum::extract::Path;
use serde::Serialize;

#[derive(Serialize)]
pub struct EnumMate {
    code: &'static str,
    desc: &'static str,
    value: i32,
}

pub async fn ecode_data() -> (StatusCode, Json<Vec<EnumMate>>) {
    let data = vec![
        EnumMate { code: "SUCCESS", value: 0, desc: "请求成功" },
        EnumMate { code: "BAD_REQUEST", value: 400000, desc: "请求异常" },
        EnumMate { code: "AUTH_TOKEN_Required", value: 400001, desc: "账户未登录" },
        EnumMate { code: "AUTH_TOKEN_Invalid", value: 400002, desc: "无效Token" },
        EnumMate { code: "AUTH_TOKEN_NotFound", value: 400003, desc: "token不存在" },
        EnumMate { code: "AUTH_TOKEN_Expired", value: 400004, desc: "token过期" },
        EnumMate { code: "ACCESS_NotFound", value: 401001, desc: "请求资源不存在" },
        EnumMate { code: "ACCESS_Forbidden", value: 401002, desc: "没有访问权限" },
        EnumMate { code: "ACCESS_Disabled", value: 401003, desc: "接口已禁用" },
        EnumMate { code: "INVALID_PARAMS", value: 431000, desc: "参数异常" },
        EnumMate { code: "USER_PARAMS_ERROR_NameRequired", value: 431001, desc: "参数缺少昵称" },
        EnumMate { code: "USER_PARAMS_ERROR_UsernameRequired", value: 431002, desc: "参数缺少用户名" },
        EnumMate { code: "USER_PARAMS_ERROR_PasswordRequired", value: 431003, desc: "参数缺少密码" },
        EnumMate { code: "USER_PARAMS_ERROR_UsernameExist", value: 431004, desc: "用户名已存在" },
        EnumMate { code: "USER_PARAMS_ERROR_UserNotFound", value: 431005, desc: "用户不存在" },
        EnumMate { code: "USER_PARAMS_ERROR_NameInvalid", value: 431006, desc: "无效的用户昵称" },
        EnumMate { code: "USER_PARAMS_ERROR_UsernameInvalid", value: 431007, desc: "无效的用户名" },
        EnumMate { code: "USER_PARAMS_ERROR_PasswordInvalid", value: 431008, desc: "无效的密码" },
        EnumMate { code: "SOURCE_PARAMS_ERROR_SourceInvalid", value: 432001, desc: "无效的实体来源" },
        EnumMate { code: "FOLDER_PARAMS_ERROR_NameRequired", value: 433001, desc: "参数缺少文件夹名称" },
        EnumMate { code: "LABEL_PARAMS_ERROR_NameRequired", value: 434001, desc: "参数缺少文件夹名称" },
        EnumMate { code: "NOTE_PARAMS_ERROR_TitleRequired", value: 435001, desc: "参数缺少标题" },
        EnumMate { code: "NOTE_PARAMS_ERROR_TitleInvalid", value: 435002, desc: "标题无效" },
        EnumMate { code: "SYSTEM_INTERNAL_ERROR", value: 500000, desc: "服务异常" },
        EnumMate { code: "SYSTEM_ERROR_Unimplemented", value: 500001, desc: "功能未实现" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListAccess", value: 510001, desc: "AUTH_SERVICE_ERROR_ListAccess" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperateAccess", value: 510002, desc: "AUTH_SERVICE_ERROR_OperateAccess" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperateAccess", value: 510003, desc: "AUTH_SERVICE_ERROR_BatchOperateAccess" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListAuthToken", value: 511001, desc: "AUTH_SERVICE_ERROR_ListAuthToken" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperateAuthToken", value: 511002, desc: "AUTH_SERVICE_ERROR_OperateAuthToken" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperateAuthToken", value: 511003, desc: "AUTH_SERVICE_ERROR_BatchOperateAuthToken" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListUser", value: 512001, desc: "AUTH_SERVICE_ERROR_ListUser" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperateUser", value: 512002, desc: "AUTH_SERVICE_ERROR_OperateUser" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperateUser", value: 512003, desc: "AUTH_SERVICE_ERROR_BatchOperateUser" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListRole", value: 513001, desc: "AUTH_SERVICE_ERROR_ListRole" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperateRole", value: 513002, desc: "AUTH_SERVICE_ERROR_OperateRole" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperateRole", value: 513003, desc: "AUTH_SERVICE_ERROR_BatchOperateRole" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListUserRoleRef", value: 514001, desc: "AUTH_SERVICE_ERROR_ListUserRoleRef" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperateUserRoleRef", value: 514002, desc: "AUTH_SERVICE_ERROR_OperateUserRoleRef" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef", value: 514003, desc: "AUTH_SERVICE_ERROR_BatchOperateUserRoleRef" },
        EnumMate { code: "AUTH_SERVICE_ERROR_ListPermission", value: 515001, desc: "AUTH_SERVICE_ERROR_ListPermission" },
        EnumMate { code: "AUTH_SERVICE_ERROR_OperatePermission", value: 515002, desc: "AUTH_SERVICE_ERROR_OperatePermission" },
        EnumMate { code: "AUTH_SERVICE_ERROR_BatchOperatePermission", value: 515003, desc: "AUTH_SERVICE_ERROR_BatchOperatePermission" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_ListLabel", value: 520001, desc: "PUBLIC_SERVICE_ERROR_ListLabel" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_OperateLabel", value: 520002, desc: "PUBLIC_SERVICE_ERROR_OperateLabel" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_BatchOperateLabel", value: 520003, desc: "PUBLIC_SERVICE_ERROR_BatchOperateLabel" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_ListFolder", value: 521001, desc: "PUBLIC_SERVICE_ERROR_ListFolder" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_OperateFolder", value: 521002, desc: "PUBLIC_SERVICE_ERROR_OperateFolder" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_BatchOperateFolder", value: 521003, desc: "PUBLIC_SERVICE_ERROR_BatchOperateFolder" },
        EnumMate { code: "PUBLIC_SERVICE_ERROR_TransCurl", value: 522001, desc: "PUBLIC_SERVICE_ERROR_TransCurl" },
        EnumMate { code: "NOTE_SERVICE_ERROR_ListNote", value: 530001, desc: "NOTE_SERVICE_ERROR_ListNote" },
        EnumMate { code: "NOTE_SERVICE_ERROR_OperateNote", value: 530002, desc: "NOTE_SERVICE_ERROR_OperateNote" },
        EnumMate { code: "NOTE_SERVICE_ERROR_BatchOperateNote", value: 530003, desc: "NOTE_SERVICE_ERROR_BatchOperateNote" },
        EnumMate { code: "NOTE_SERVICE_ERROR_ListNoteLabel", value: 531001, desc: "NOTE_SERVICE_ERROR_ListNoteLabel" },
        EnumMate { code: "NOTE_SERVICE_ERROR_OperateNoteLabel", value: 531002, desc: "NOTE_SERVICE_ERROR_OperateNoteLabel" },
        EnumMate { code: "NOTE_SERVICE_ERROR_BatchOperateNoteLabel", value: 531003, desc: "NOTE_SERVICE_ERROR_BatchOperateNoteLabel" },
        EnumMate { code: "NOTE_SERVICE_ERROR_ListNoteTopic", value: 532001, desc: "NOTE_SERVICE_ERROR_ListNoteTopic" },
        EnumMate { code: "NOTE_SERVICE_ERROR_OperateNoteTopic", value: 532002, desc: "NOTE_SERVICE_ERROR_OperateNoteTopic" },
        EnumMate { code: "NOTE_SERVICE_ERROR_BatchOperateNoteTopic", value: 532003, desc: "NOTE_SERVICE_ERROR_BatchOperateNoteTopic" },
    ];
    (StatusCode::OK, Json(data))
}

pub async fn source_data() -> (StatusCode, Json<Vec<EnumMate>>) {
    let data = vec![
        EnumMate { code: "Public", value: 0, desc: "公共" },
        EnumMate { code: "User", value: 1, desc: "用户" },
        EnumMate { code: "Note", value: 2, desc: "笔记" },
        EnumMate { code: "File", value: 3, desc: "文件" },
    ];
    (StatusCode::OK, Json(data))
}

pub async fn booleanscope_data() -> (StatusCode, Json<Vec<EnumMate>>) {
    let data = vec![
        EnumMate { code: "BOOL_ALL", value: 0, desc: "默认" },
        EnumMate { code: "BOOL_FALSE", value: 1, desc: "FALSE" },
        EnumMate { code: "BOOL_TRUE", value: 2, desc: "TRUE" },
    ];
    (StatusCode::OK, Json(data))
}

pub async fn rangescope_data() -> (StatusCode, Json<Vec<EnumMate>>) {
    let data = vec![
        EnumMate { code: "RANGE_ALL", value: 0, desc: "默认" },
        EnumMate { code: "RANGE_LEFT", value: 1, desc: "含左" },
        EnumMate { code: "RANGE_RIGHT", value: 2, desc: "含右" },
    ];
    (StatusCode::OK, Json(data))
}

pub(crate) async fn meta_handler(Path(name): Path<String>) -> (StatusCode, Json<Vec<EnumMate>>) {
    match name.as_str() {
        "ecode" => ecode_data().await,
        "source" => source_data().await,
        "booleanscope" => booleanscope_data().await,
        "rangescope" => rangescope_data().await,
        _=> { (StatusCode::OK, Json(vec![])) }
    }
}

