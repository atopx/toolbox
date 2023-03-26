package utils

import (
	"time"
)

const (
	DefaultDateLayout = "2006-01-02 15:04:05"
)

func TimestampDecoder(ts int64, layout string) string {
	if layout == "" {
		layout = DefaultDateLayout
	}
	return time.Unix(ts, 0).Local().Format(layout)
}

func NewTraceId() int64 {
	return time.Now().Local().UnixNano()
}
