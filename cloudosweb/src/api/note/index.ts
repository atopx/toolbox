import { request } from "@/utils/service"
import * as types from "./types"
import { NullResponse, ResponseIface } from "types/api"

export function noteListApi(data: types.noteListRequest) {
    return request<types.noteListResponse>({
        url: "/note/list",
        method: "post",
        data
    })
}

export function noteSaveApi(data: types.noteSaveRequest) {
    return request<NullResponse>({
        url: "/note/save",
        method: "post",
        data
    })
}

export function noteDeleteApi(id: number) {
    return request<NullResponse>({
        url: "/note/delete",
        method: "post",
        data: { id: id }
    })
}

export function noteInfoApi(id: number) {
    return request<ResponseIface<types.Note>>({
        url: "/note/info",
        method: "get",
        params: { id: id }
    })
}

export function noteTopicsApi() {
    return request<ResponseIface<string[]>>({
        url: "/note/topics",
        method: "get"
    })
}
