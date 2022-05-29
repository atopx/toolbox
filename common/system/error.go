package system

import (
	"go.uber.org/zap/zapcore"
	"net/http"
)

const REPLY_ERROR_KEY = "error_reply"

type ReplyError struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 异常消息
	error   string // 内部字段，不导出
}

func NewReplyError(code int, err string) ReplyError {
	return ReplyError{
		Code:    code,
		Message: http.StatusText(code),
		error:   err,
	}
}

func (reply ReplyError) GetErrorMessage() string {
	return reply.error
}

func (reply ReplyError) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddInt("code", reply.Code)
	encoder.AddString("message", reply.Message)
	encoder.AddString("error", reply.error)
	return nil
}
