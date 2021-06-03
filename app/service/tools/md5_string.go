package tools

import (
	"crypto/md5"
	"encoding/hex"
)

//Md5String 计算字符串的MD5
func Md5String(plainText string) string {
	data := []byte(plainText)
	binaryData := md5.Sum(data)
	return hex.EncodeToString(binaryData[:])
}
