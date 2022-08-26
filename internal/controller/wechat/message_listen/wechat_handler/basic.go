package wechat_handler

import "fmt"

type BasicHandler struct {
	Message string
}

func (h *BasicHandler) Deal() string {
	return fmt.Sprintf("Sorry, I don't understand: 「%s」", h.Message)
}
