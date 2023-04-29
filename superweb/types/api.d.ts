/** 所有 api 接口的响应数据都应该准守该格式 */
export interface Header {
  trace_id: number
  message: string
  code: number
}

interface ResponseIface<T> {
  header: Header
  data: T
}

export interface PageIface {
  count: number
  disabled: boolean
  index: number
  size: number
}
