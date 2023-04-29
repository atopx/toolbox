package system

import (
	"superserver/common/enum"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
)

func GetErrorMessage(code ecode.ECode) string {
	return enum.Desc(enum.ECode, int32(code))
}

func NewResponse(header *common.ReplyHeader, data any) *Response {
	header.Message = GetErrorMessage(header.Code)
	return &Response{Header: header, Data: data}
}

func NewErrorResponse(header *common.ReplyHeader) *Response {
	return NewResponse(header, nil)
}
