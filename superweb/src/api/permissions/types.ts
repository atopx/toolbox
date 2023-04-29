import { PageIface, ResponseIface } from "types/api"

export interface Access {
    id: number
    path: string
    method: string
    handler: string
    create_time: number
    update_time: number
}

export type listPermissionAccessRequest = { page: PageIface; filter: {} }
export type listPermissionAccessResponse = ResponseIface<{ pager: PageIface; search: Access[] }>

export interface Role {
    id: number
    name: string
    nature: number,
    nature_desc: string,
    create_time: string
    update_time: string
}

export type listRoleRequest = { page: PageIface; filter: { keyword: string } }
export type listRoleResponse = ResponseIface<{ pager: PageIface; list: Role[] }>