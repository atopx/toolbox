import { PageIface, ResponseIface, TimeRangeIface } from "types/api"

export interface Note {
    id: number // 主键
    title: string // 标题
    content: string // 笔记内容
    labels: string[] // 标签
    topic: string // 主题
    createTime: string // 创建时间 时间戳：秒
    updateTime: string // 最后更新时间 时间戳：秒
}

export type noteListRequest = { pager: PageIface; keyword: string; topic: string; updateTimeRange?: TimeRangeIface }
export type noteListResponse = ResponseIface<{ pager: PageIface; list: Note[] }>

export type noteSaveRequest = {
    id: number
    title: string
    topic: string
    content: string
    labels: string[]
}
