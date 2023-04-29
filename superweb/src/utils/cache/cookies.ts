// 本地存储管理Token信息
import { Token } from "@/api/login/types/login"
import CacheKey from "@/constants/cacheKey"

export const getAccessToken = () => {
    return localStorage.getItem(CacheKey.ACCESS_TOKEN)
}

export const getRefreshToken = () => {
    return localStorage.getItem(CacheKey.ACCESS_TOKEN)
}

export const getTokenExpires = () => {
    return Number(localStorage.getItem(CacheKey.TOKEN_EXPIRES))
}

export const getToken = () => {
    return {
        id: 0,
        user_id: Number(localStorage.getItem(CacheKey.USER_ID)),
        access_token: localStorage.getItem(CacheKey.ACCESS_TOKEN),
        refresh_token: localStorage.getItem(CacheKey.REFRESH_TOKEN),
        issued_time: 0,
        expire_time: Number(localStorage.getItem(CacheKey.TOKEN_EXPIRES)),
        create_time: 0,
        update_time: 0
    } as Token
}

export const setToken = (token: Token) => {
    localStorage.setItem(CacheKey.USER_ID, token.user_id.toString())
    localStorage.setItem(CacheKey.ACCESS_TOKEN, token.access_token)
    localStorage.setItem(CacheKey.REFRESH_TOKEN, token.refresh_token)
    localStorage.setItem(CacheKey.TOKEN_EXPIRES, token.expire_time.toString())
}

export const removeToken = () => {
    localStorage.removeItem(CacheKey.ACCESS_TOKEN)
    localStorage.removeItem(CacheKey.REFRESH_TOKEN)
    localStorage.removeItem(CacheKey.TOKEN_EXPIRES)
}
