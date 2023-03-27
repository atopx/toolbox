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
        access_token: localStorage.getItem(CacheKey.ACCESS_TOKEN),
        refresh_token: localStorage.getItem(CacheKey.REFRESH_TOKEN),
        expires: Number(localStorage.getItem(CacheKey.TOKEN_EXPIRES))
    } as Token
}

export const setToken = (token: Token) => {
    localStorage.setItem(CacheKey.ACCESS_TOKEN, token.access_token)
    localStorage.setItem(CacheKey.REFRESH_TOKEN, token.refresh_token)
    localStorage.setItem(CacheKey.TOKEN_EXPIRES, token.expires.toString())
}

export const removeToken = () => {
    localStorage.removeItem(CacheKey.ACCESS_TOKEN)
    localStorage.removeItem(CacheKey.REFRESH_TOKEN)
    localStorage.removeItem(CacheKey.TOKEN_EXPIRES)
}
