package system

import "go.uber.org/zap/zapcore"

type ReplyError struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 异常消息
}

func (err ReplyError) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddInt("response_code", err.Code)
	encoder.AddString("response_message", err.Message)
	return nil
}
