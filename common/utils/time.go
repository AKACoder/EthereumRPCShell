package utils

import "time"

const (
	defTimeFormat = "2006-01-02 15:04:05"
)

func CurrentDatetime() string {
	timestamp := time.Now().Unix()
	utcTime := time.Unix(timestamp, 0).UTC().Format(defTimeFormat)
	return utcTime
}

func TimeBefore(duration time.Duration) string {
	now := time.Now()
	timestamp := now.Add(0 - duration).Unix()
	return time.Unix(timestamp, 0).UTC().Format(defTimeFormat)
}

func TimeStringOf(secStamp int64) string {
	return time.Unix(secStamp, 0).UTC().Format(defTimeFormat)
}
