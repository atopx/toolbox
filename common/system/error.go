package system

import (
	"net/http"

	"go.uber.org/zap/zapcore"
)

const REPLY_ERROR_KEY = "error_reply"

type ReplyError struct {
	TraceId int64  `json:"trace_id"`
	Message string `json:"message"` // 异常消息
	error   string // 内部字段，不导出
}

func NewReplyError(traceId int64, code int, err string) ReplyError {
	return ReplyError{
		TraceId: traceId,
		Message: http.StatusText(code),
		error:   err,
	}
}

func (reply ReplyError) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("message", reply.Message)
	encoder.AddString("error", reply.error)
	return nil
}
