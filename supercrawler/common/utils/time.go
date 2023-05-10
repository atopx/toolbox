package utils

import "time"

func TimeLoad(date, layout string) int64 {
	t, _ := time.Parse(layout, date)
	return t.Unix()
}
