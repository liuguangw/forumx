package tools

import "time"

//GenerateHashID 生成32位 hash ID
func GenerateHashID() string {
	plainText := time.Now().Format(defaultTimeFormat) + "||" + GenerateRandomString(30)
	return Md5String(plainText)
}
