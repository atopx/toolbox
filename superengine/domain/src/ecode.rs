#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ECode {
    /// 请求成功
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
    ///
    /// 请求资源不存在
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
    ///
    /// 参数缺少昵称
    UserParamsErrorNameRequired = 431001,
    /// 参数缺少用户名
    UserParamsErrorUsernameRequired = 431002,
    /// 参数缺少密码
    UserParamsErrorPasswordRequired = 431003,
    /// 用户名已存在
    UserParamsErrorUsernameExist = 431004,
    /// 用户不存在
    UserParamsErrorUserNotFound = 431005,
    /// user login
    ///
    /// 无效的用户昵称
    UserParamsErrorNameInvalid = 431006,
    /// 无效的用户名
    UserParamsErrorUsernameInvalid = 431007,
    /// 无效的密码
    UserParamsErrorPasswordInvalid = 431008,
    /// public source
    ///
    /// 无效的实体来源
    SourceParamsErrorSourceInvalid = 432001,
    /// public folder
    ///
    /// 参数缺少文件夹名称
    FolderParamsErrorNameRequired = 433001,
    /// public label
    ///
    /// 参数缺少文件夹名称
    LabelParamsErrorNameRequired = 434001,
    /// note
    ///
    /// 参数缺少标题
    NoteParamsErrorTitleRequired = 435001,
    /// 标题无效
    NoteParamsErrorTitleInvalid = 435002,
    /// role
    ///
    /// 无效的角色名
    RoleParamsErrorNameInvalid = 436001,
    /// 角色名已存在
    RoleParamsErrorNameExist = 436002,
    /// ######### 服务端错误 #########
    ///
    /// 服务异常
    SystemInternalError = 500000,
    /// 功能未实现
    SystemErrorUnimplemented = 500001,
    /// auth-service access error
    ///
    /// AUTH_SERVICE_ERROR_ListAccess
    AuthServiceErrorListAccess = 510001,
    /// AUTH_SERVICE_ERROR_OperateAccess
    AuthServiceErrorOperateAccess = 510002,
    /// AUTH_SERVICE_ERROR_BatchOperateAccess
    AuthServiceErrorBatchOperateAccess = 510003,
    /// auth-service token error
    ///
    /// AUTH_SERVICE_ERROR_ListAuthToken
    AuthServiceErrorListAuthToken = 511001,
    /// AUTH_SERVICE_ERROR_OperateAuthToken
    AuthServiceErrorOperateAuthToken = 511002,
    /// AUTH_SERVICE_ERROR_BatchOperateAuthToken
    AuthServiceErrorBatchOperateAuthToken = 511003,
    /// auth-service user error
    ///
    /// AUTH_SERVICE_ERROR_ListUser
    AuthServiceErrorListUser = 512001,
    /// AUTH_SERVICE_ERROR_OperateUser
    AuthServiceErrorOperateUser = 512002,
    /// AUTH_SERVICE_ERROR_BatchOperateUser
    AuthServiceErrorBatchOperateUser = 512003,
    /// auth-service role error
    ///
    /// AUTH_SERVICE_ERROR_ListRole
    AuthServiceErrorListRole = 513001,
    /// AUTH_SERVICE_ERROR_OperateRole
    AuthServiceErrorOperateRole = 513002,
    /// AUTH_SERVICE_ERROR_BatchOperateRole
    AuthServiceErrorBatchOperateRole = 513003,
    /// auth-service permission error
    ///
    /// AUTH_SERVICE_ERROR_ListUserRoleRef
    AuthServiceErrorListUserRoleRef = 514001,
    /// AUTH_SERVICE_ERROR_OperateUserRoleRef
    AuthServiceErrorOperateUserRoleRef = 514002,
    /// AUTH_SERVICE_ERROR_BatchOperateUserRoleRef
    AuthServiceErrorBatchOperateUserRoleRef = 514003,
    /// auth-service permission error
    ///
    /// AUTH_SERVICE_ERROR_ListPermission
    AuthServiceErrorListPermission = 515001,
    /// AUTH_SERVICE_ERROR_OperatePermission
    AuthServiceErrorOperatePermission = 515002,
    /// AUTH_SERVICE_ERROR_BatchOperatePermission
    AuthServiceErrorBatchOperatePermission = 515003,
    /// public-service label error
    ///
    /// PUBLIC_SERVICE_ERROR_ListLabel
    PublicServiceErrorListLabel = 520001,
    /// PUBLIC_SERVICE_ERROR_OperateLabel
    PublicServiceErrorOperateLabel = 520002,
    /// PUBLIC_SERVICE_ERROR_BatchOperateLabel
    PublicServiceErrorBatchOperateLabel = 520003,
    /// public-service list enum error
    ///
    /// PUBLIC_SERVICE_ERROR_ListEnum
    PublicServiceErrorListEnum = 521001,
    /// note-service note error
    ///
    /// NOTE_SERVICE_ERROR_ListNote
    NoteServiceErrorListNote = 530001,
    /// NOTE_SERVICE_ERROR_OperateNote
    NoteServiceErrorOperateNote = 530002,
    /// NOTE_SERVICE_ERROR_BatchOperateNote
    NoteServiceErrorBatchOperateNote = 530003,
    /// note-service note-label error
    ///
    /// NOTE_SERVICE_ERROR_ListNoteLabel
    NoteServiceErrorListNoteLabel = 531001,
    /// NOTE_SERVICE_ERROR_OperateNoteLabel
    NoteServiceErrorOperateNoteLabel = 531002,
    /// NOTE_SERVICE_ERROR_BatchOperateNoteLabel
    NoteServiceErrorBatchOperateNoteLabel = 531003,
    /// novel-service book error
    ///
    /// NOVEL_SERVICE_ERROR_ListBook
    NovelServiceErrorListBook = 540001,
    /// NOVEL_SERVICE_ERROR_OperateBook
    NovelServiceErrorOperateBook = 540002,
    /// NOVEL_SERVICE_ERROR_BatchOperateBook
    NovelServiceErrorBatchOperateBook = 540003,
    /// novel-service chapter error
    ///
    /// NOVEL_SERVICE_ERROR_ListChapter
    NovelServiceErrorListChapter = 541001,
    /// NOVEL_SERVICE_ERROR_OperateChapter
    NovelServiceErrorOperateChapter = 541002,
    /// NOVEL_SERVICE_ERROR_BatchOperateChapter
    NovelServiceErrorBatchOperateChapter = 541003,
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
            ECode::SourceParamsErrorSourceInvalid => "SOURCE_PARAMS_ERROR_SourceInvalid",
            ECode::FolderParamsErrorNameRequired => "FOLDER_PARAMS_ERROR_NameRequired",
            ECode::LabelParamsErrorNameRequired => "LABEL_PARAMS_ERROR_NameRequired",
            ECode::NoteParamsErrorTitleRequired => "NOTE_PARAMS_ERROR_TitleRequired",
            ECode::NoteParamsErrorTitleInvalid => "NOTE_PARAMS_ERROR_TitleInvalid",
            ECode::RoleParamsErrorNameInvalid => "ROLE_PARAMS_ERROR_NameInvalid",
            ECode::RoleParamsErrorNameExist => "ROLE_PARAMS_ERROR_NameExist",
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
            ECode::PublicServiceErrorListLabel => "PUBLIC_SERVICE_ERROR_ListLabel",
            ECode::PublicServiceErrorOperateLabel => "PUBLIC_SERVICE_ERROR_OperateLabel",
            ECode::PublicServiceErrorBatchOperateLabel => {
                "PUBLIC_SERVICE_ERROR_BatchOperateLabel"
            }
            ECode::PublicServiceErrorListEnum => "PUBLIC_SERVICE_ERROR_ListEnum",
            ECode::NoteServiceErrorListNote => "NOTE_SERVICE_ERROR_ListNote",
            ECode::NoteServiceErrorOperateNote => "NOTE_SERVICE_ERROR_OperateNote",
            ECode::NoteServiceErrorBatchOperateNote => {
                "NOTE_SERVICE_ERROR_BatchOperateNote"
            }
            ECode::NoteServiceErrorListNoteLabel => "NOTE_SERVICE_ERROR_ListNoteLabel",
            ECode::NoteServiceErrorOperateNoteLabel => {
                "NOTE_SERVICE_ERROR_OperateNoteLabel"
            }
            ECode::NoteServiceErrorBatchOperateNoteLabel => {
                "NOTE_SERVICE_ERROR_BatchOperateNoteLabel"
            }
            ECode::NovelServiceErrorListBook => "NOVEL_SERVICE_ERROR_ListBook",
            ECode::NovelServiceErrorOperateBook => "NOVEL_SERVICE_ERROR_OperateBook",
            ECode::NovelServiceErrorBatchOperateBook => {
                "NOVEL_SERVICE_ERROR_BatchOperateBook"
            }
            ECode::NovelServiceErrorListChapter => "NOVEL_SERVICE_ERROR_ListChapter",
            ECode::NovelServiceErrorOperateChapter => {
                "NOVEL_SERVICE_ERROR_OperateChapter"
            }
            ECode::NovelServiceErrorBatchOperateChapter => {
                "NOVEL_SERVICE_ERROR_BatchOperateChapter"
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
            "SOURCE_PARAMS_ERROR_SourceInvalid" => {
                Some(Self::SourceParamsErrorSourceInvalid)
            }
            "FOLDER_PARAMS_ERROR_NameRequired" => {
                Some(Self::FolderParamsErrorNameRequired)
            }
            "LABEL_PARAMS_ERROR_NameRequired" => Some(Self::LabelParamsErrorNameRequired),
            "NOTE_PARAMS_ERROR_TitleRequired" => Some(Self::NoteParamsErrorTitleRequired),
            "NOTE_PARAMS_ERROR_TitleInvalid" => Some(Self::NoteParamsErrorTitleInvalid),
            "ROLE_PARAMS_ERROR_NameInvalid" => Some(Self::RoleParamsErrorNameInvalid),
            "ROLE_PARAMS_ERROR_NameExist" => Some(Self::RoleParamsErrorNameExist),
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
            "PUBLIC_SERVICE_ERROR_ListLabel" => Some(Self::PublicServiceErrorListLabel),
            "PUBLIC_SERVICE_ERROR_OperateLabel" => {
                Some(Self::PublicServiceErrorOperateLabel)
            }
            "PUBLIC_SERVICE_ERROR_BatchOperateLabel" => {
                Some(Self::PublicServiceErrorBatchOperateLabel)
            }
            "PUBLIC_SERVICE_ERROR_ListEnum" => Some(Self::PublicServiceErrorListEnum),
            "NOTE_SERVICE_ERROR_ListNote" => Some(Self::NoteServiceErrorListNote),
            "NOTE_SERVICE_ERROR_OperateNote" => Some(Self::NoteServiceErrorOperateNote),
            "NOTE_SERVICE_ERROR_BatchOperateNote" => {
                Some(Self::NoteServiceErrorBatchOperateNote)
            }
            "NOTE_SERVICE_ERROR_ListNoteLabel" => {
                Some(Self::NoteServiceErrorListNoteLabel)
            }
            "NOTE_SERVICE_ERROR_OperateNoteLabel" => {
                Some(Self::NoteServiceErrorOperateNoteLabel)
            }
            "NOTE_SERVICE_ERROR_BatchOperateNoteLabel" => {
                Some(Self::NoteServiceErrorBatchOperateNoteLabel)
            }
            "NOVEL_SERVICE_ERROR_ListBook" => Some(Self::NovelServiceErrorListBook),
            "NOVEL_SERVICE_ERROR_OperateBook" => Some(Self::NovelServiceErrorOperateBook),
            "NOVEL_SERVICE_ERROR_BatchOperateBook" => {
                Some(Self::NovelServiceErrorBatchOperateBook)
            }
            "NOVEL_SERVICE_ERROR_ListChapter" => Some(Self::NovelServiceErrorListChapter),
            "NOVEL_SERVICE_ERROR_OperateChapter" => {
                Some(Self::NovelServiceErrorOperateChapter)
            }
            "NOVEL_SERVICE_ERROR_BatchOperateChapter" => {
                Some(Self::NovelServiceErrorBatchOperateChapter)
            }
            _ => None,
        }
    }
}
