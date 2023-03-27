const SYSTEM_NAME = "super"

/** 缓存数据时用到的 Key */
class CacheKey {
    static ACCESS_TOKEN = `${SYSTEM_NAME}-access_token`
    static REFRESH_TOKEN = `${SYSTEM_NAME}-refresh_token`
    static TOKEN_EXPIRES = `${SYSTEM_NAME}-token_expires`
    static SIDEBAR_STATUS = `${SYSTEM_NAME}-sidebar-status-key`
    static ACTIVE_THEME_NAME = `${SYSTEM_NAME}-active-theme-name-key`
}

export default CacheKey
