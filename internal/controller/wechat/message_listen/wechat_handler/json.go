// Package wechat_handler 微信消息处理器
package wechat_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// TODO 微信消息匹配事件, 基于事件驱动
type JsonHandler struct {
	Message string
}

func (h *JsonHandler) Deal() string {
	jsonStr, err := strconv.Unquote(fmt.Sprintf("\"%s\"", h.Message))
	if err != nil {
		return "Error: " + err.Error()
	}
	var out bytes.Buffer
	if err := json.Indent(&out, []byte(jsonStr), "", "    "); err != nil {
		return err.Error()
	}
	return out.String()
}
