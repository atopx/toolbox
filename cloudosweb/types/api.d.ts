/** 所有 api 接口的响应数据都应该准守该格式 */
export interface ResponseIface<Data> {
  code: number;
  message: string;
  traceId: number;
  data: Data;
}

export type NullResponse = ResponseIface<null>

export interface PageIface {
  index: number;
  size: number;
  count: number;
  disable: boolean;
}

declare interface TimeRangeIface {
  left: number;
  right: number;
}

declare interface OptionIface {
  label: string;
  value: string;
}
