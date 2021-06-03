package tools

import "math/rand"

//随机字符串
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(requiredLength int) string {
	b := make([]byte, requiredLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
