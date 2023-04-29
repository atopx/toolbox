import { ResponseIface } from "types/api"

export interface ILoginRequestData {
    username: string
    password: string
}

export interface IRefreshRequestData {
    refresh_token: string
}

export interface Token {
    id: number
    user_id: number
    access_token: string
    refresh_token: string
    issued_time: number
    expire_time: number
    create_time: number
    update_time: number
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
