package sms

import "math/rand"

//随机字符串
const letterBytes = "0123456789"

//generateRandomCode 随机生成指定长度的验证码
func generateRandomCode(codeLength int) string {
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
