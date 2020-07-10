package time

import "time"

var lay string = "2006-01-02 15:04:05"

// 当前时间 2006-01-02 15:04:05
func GetNow() string {
	return time.Now().Format(lay)
}

// 当前时间戳 秒
func GetCurStamp() int64 {
	return time.Now().Unix()
}

// 当前时间戳 毫秒
func GetCurStamps() int64 {
	return time.Now().UnixNano() / 1000000
}

// 时间戳转时间字符串 毫秒
func StampToString(timeUnix int64) string {
	return time.Unix(timeUnix/1000, 0).Format(lay)
}

// 时间字符串转时间戳 毫秒
func StringToStamp(formatTimeStr string) int64 {
	stamp, err := time.Parse(lay, formatTimeStr)
	if err != nil {
		return 0
	}
	return stamp.Unix() * 1000
}
