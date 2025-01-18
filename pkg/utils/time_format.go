package utils

import "time"

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}
