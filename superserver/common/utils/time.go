package utils

import "time"

const (
	DEFAULT_DATE_LAYOUT = "2006-01-02 15:04:05"
)

func TimestampDecoder(ts int64, layout string) string {
	if layout == "" {
		layout = DEFAULT_DATE_LAYOUT
	}
	return time.Unix(ts, 0).Local().Format(layout)
}
