package system

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
)

func GetErrorMessage(code ecode.ECode) string {
	value, ok := CodeMap[code]
	if !ok {
		value = code.String()
	}
	return value
}

func NewResponse(header *common.ReplyHeader, data any) *Response {
	header.Message = GetErrorMessage(header.Code)
	return &Response{Header: header, Data: data}
}

func NewErrorResponse(header *common.ReplyHeader) *Response {
	return NewResponse(header, nil)
}

var CodeMap = map[ecode.ECode]string{
	ecode.ECode_AUTH_SERVICE_ERROR_ListAccess: "auth_service.ListAccess error",
	ecode.ECode_ACCESS_NotFound:               "访问的资源不存在",
}
