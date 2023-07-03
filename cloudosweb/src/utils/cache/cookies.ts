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
        userId: Number(localStorage.getItem(CacheKey.USER_ID)),
        accessToken: localStorage.getItem(CacheKey.ACCESS_TOKEN),
        refreshToken: localStorage.getItem(CacheKey.REFRESH_TOKEN),
        expireTime: Number(localStorage.getItem(CacheKey.TOKEN_EXPIRES))
    } as Token
}

export const setToken = (token: Token) => {
    localStorage.setItem(CacheKey.USER_ID, token.userId.toString())
    localStorage.setItem(CacheKey.ACCESS_TOKEN, token.accessToken)
    localStorage.setItem(CacheKey.REFRESH_TOKEN, token.refreshToken)
    localStorage.setItem(CacheKey.TOKEN_EXPIRES, token.expireTime.toString())
}

export const removeToken = () => {
    localStorage.removeItem(CacheKey.ACCESS_TOKEN)
    localStorage.removeItem(CacheKey.REFRESH_TOKEN)
    localStorage.removeItem(CacheKey.TOKEN_EXPIRES)
}
