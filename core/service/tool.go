package service

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	//随机字符串
	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//默认的时间格式
	defaultTimeFormat = "2006-01-02 15:04:05"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//generateRandomString 生成指定长度的随机字符串
func generateRandomString(requiredLength int) string {
	b := make([]byte, requiredLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//md5String 计算字符串的MD5
func md5String(plainText string) string {
	data := []byte(plainText)
	binaryData := md5.Sum(data)
	return hex.EncodeToString(binaryData[:])
}

//FormatDateTime 使用默认的格式,把日期时间转化为字符串
func FormatDateTime(t time.Time) string {
	return t.Format(defaultTimeFormat)
}
