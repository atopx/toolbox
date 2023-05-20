import { PageIface, ResponseIface } from "types/api"



export interface Note {
value: any
    id: number // 主键
    title: string // 标题
    content: string // 笔记内容
    labels: number[] // 标签
    topics: number[] // 主题
    create_time: string // 创建时间 时间戳：秒
    update_time: string // 最后更新时间 时间戳：秒
}

export type noteListRequest = { pager: PageIface, filter: { keyword: string } }
export type noteListResponse = ResponseIface<{ pager: PageIface, list: Note[] }>

export type noteSaveRequest = {
    id: number,
    title: string,
    content: string,
    topics: number[],
    labels: number[],
}
