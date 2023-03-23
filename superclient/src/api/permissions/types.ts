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
export type listPermissionAccessResponse = ResponseIface<{ page: PageIface; search: Access[] }>
