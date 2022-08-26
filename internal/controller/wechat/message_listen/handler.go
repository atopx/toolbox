package message_listen

import (
	"strings"
	"toolbox/internal/controller/wechat/message_listen/wechat_handler"
)

type handler interface {
	Deal() string
}

func NewHandler(content string) handler {
	if strings.HasPrefix(content, "atopx") {
		return &wechat_handler.MessageHandler{}
	}
	if strings.HasPrefix(content, "json") {
		return &wechat_handler.JsonHandler{Message: content[5:]}
	}
	if strings.HasPrefix(content, "ts") {
		return &wechat_handler.TimestampHandler{Message: content[3:]}
	}
	if strings.HasPrefix(content, "compress") {
		return &wechat_handler.CompcressHandler{Message: content[9:]}
	}

	return &wechat_handler.BasicHandler{Message: content}
}
