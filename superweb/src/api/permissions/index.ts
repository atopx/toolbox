import { request } from "@/utils/service"
import * as types from "./types"

// /api/permission/access
export function listPermissionAccess(data: types.listPermissionAccessRequest) {
    return request<types.listPermissionAccessResponse>({
        url: "/permission/access",
        method: "post",
        data
    })
}

export function listRoleApi(data: types.listRoleRequest) {
    return request<types.listRoleResponse>({
        url: "/role/list",
        method: "post",
        data
    })
}

export function createRoleApi(name: string) {
    return request<null>({
        url: "/role/create",
        method: "post",
        data: { name: name }
    })
}

export function updateRoleApi(id: number, name: string) {
    return request<null>({
        url: "/role/update",
        method: "post",
        data: { id: id, name: name }
    })
}


export function deleteRoleApi(id: number) {
    return request<null>({
        url: "/role/delete",
        method: "delete",
        data: { id: id }
    })
}