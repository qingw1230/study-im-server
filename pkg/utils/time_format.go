package utils

import "time"

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

func UnixNanoSecondToTime(nanoSecond int64) time.Time {
	return time.Unix(0, nanoSecond)
}

func UnixMillSecondToTime(millSecond int64) time.Time {
	return time.Unix(0, millSecond*1e6)
}

func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}
