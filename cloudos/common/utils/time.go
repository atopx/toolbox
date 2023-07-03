package utils

import "time"

func Datetime(sec int64) string {
	return time.Unix(sec, 0).Format(time.DateTime)
}
