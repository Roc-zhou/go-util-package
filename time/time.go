package time

import "time"

// 当前时间 2006-01-02 15:04:05
func GetNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 当前时间戳 秒
func GetCurStamp() int64 {
	return time.Now().Unix()
}

// 当前时间戳 毫秒
func GetCurStamps() int64 {
	return time.Now().UnixNano() / 1000000

}
