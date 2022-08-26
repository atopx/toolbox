package wechat_handler

import (
	"strings"
)

type CompcressHandler struct {
	Message string
}

var replacer = strings.NewReplacer(" ", "", "\n", "")

func (h *CompcressHandler) Deal() string {
	return replacer.Replace(h.Message)
}
