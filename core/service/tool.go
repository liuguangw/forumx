package service

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//随机字符串
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
