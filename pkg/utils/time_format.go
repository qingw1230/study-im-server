package utils

import "time"

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}
