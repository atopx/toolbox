import { ResponseIface } from "types/api"

export interface ILoginRequestData {
    username: string
    password: string
}

export interface IRefreshRequestData {
    refresh_token: string
}

export interface Token {
    userId: number
    expireTime: number
    accessToken: string
    refreshToken: string
}

export type LoginResponseData = ResponseIface<Token>

export type RefreshResponseData = ResponseIface<Token>

export type UserInfoResponseData = ResponseIface<{
    id: number
    name: string
    username: string
    login_time: number
    create_time: number
    update_time: number
}>
