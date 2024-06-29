package helpers

import "time"

func GetTimeNow() int64 {
	return time.Now().Unix()
}

func GetDateFromTimestamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}
