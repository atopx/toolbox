package wechat_handler

import (
	"strconv"
	"time"
)

type TimestampHandler struct {
	Message string
}

func (h *TimestampHandler) Deal() string {
	// try convert message to int64
	ts, err := strconv.ParseInt(h.Message, 10, 64)
	if err != nil {
		return "Error: " + err.Error()
	}
	return time.Unix(ts, 0).Local().Format("2006-01-02 15:04:05")
}
