import { request } from "@/utils/service"
import type * as types from "./types"

// /api/permission/access
export function listPermissionAccess(data: types.listPermissionAccessRequest) {
    return request<types.listPermissionAccessResponse>({
        url: "/permission/access",
        method: "post",
        data
    })
}
