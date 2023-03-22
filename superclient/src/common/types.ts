export interface Page {
    index: number, // 页码
    size: number,  // 数量
    count: number,  // 总数量
    disabled: boolean,  // 禁用分页
}

export interface ResponseData<T> {
    trace_id: number,
    message: string,
    data: T,
}