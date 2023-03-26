package system

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
)

var CodeMap = map[ecode.ECode]string{
	ecode.ECode_AUTH_SERVICE_ERROR_ListAccess: "auth_service.ListAccess error",
	ecode.ECode_ACCESS_NotFound:               "访问的资源不存在",
}

func GetErrorMessage(code ecode.ECode) string {
	return CodeMap[code]
}

func NewResponse(header *common.ReplyHeader, data any) *Response {
	return &Response{Header: header, Data: data}
}

func NewErrorResponse(header *common.ReplyHeader) *Response {
	return &Response{Header: header}
}
