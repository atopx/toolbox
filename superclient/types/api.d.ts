/** 所有 api 接口的响应数据都应该准守该格式 */
interface ResponseIface<T> {
  success: boolean
  trace_id: number,
  message: string
  data: T
}


export interface PageIface {
  count: number
  disabled: boolean
  index: number
  size: number
}
