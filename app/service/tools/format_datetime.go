package tools

import "time"

//默认的时间格式
const defaultTimeFormat = "2006-01-02 15:04:05"

//FormatDateTime 使用默认的格式,把日期时间转化为字符串
func FormatDateTime(t time.Time) string {
	return t.Format(defaultTimeFormat)
}
