import { ResponseIface } from "types/api"

export interface ILoginRequestData {
    username: string
    password: string
}

export interface IRefreshRequestData {
    refresh_token: string
}

export type LoginResponseData = ResponseIface<{
    user_id: number
    name: string
    username: string
    access_token: string
    refresh_token: string
    expires: number
}>

export interface Token {
    access_token: string
    refresh_token: string
    expires: number
}

export type RefreshResponseData = ResponseIface<Token>

export type UserInfoResponseData = ResponseIface<{
    id: number
    name: string
    username: string
    login_time: number
    create_time: number
    update_time: number
}>
