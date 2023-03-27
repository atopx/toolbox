import { request } from "@/utils/service"
import type * as Login from "./types/login"

/** 登录并返回 Token */
export function loginApi(data: Login.ILoginRequestData) {
    return request<Login.LoginResponseData>({
        url: "/user/login",
        method: "post",
        data
    })
}

export function refreshTokenApi(data: Login.IRefreshRequestData) {
    return request<Login.RefreshResponseData>({
        url: "/user/refresh",
        method: "post",
        data
    })
}

/** 获取用户详情 */
export function getUserInfoApi() {
    return request<Login.UserInfoResponseData>({
        url: "/user/info",
        method: "get"
    })
}
