package utils

import "time"

func TimestampDecoder(ts int64, layout string) string {
	return time.Unix(ts, 0).Local().Format(layout)
}

func TimeLoad(date, layout string) int64 {
	t, _ := time.Parse(layout, date)
	return t.Unix()
}

func TodayZeroTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func TodayZeroDate() string {
	return time.Now().Format(time.DateOnly)
}
