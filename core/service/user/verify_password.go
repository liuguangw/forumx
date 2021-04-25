package user

import (
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/tools"
)

//hashPassword 处理密码的hash加密
func hashPassword(password, salt string) string {
	saltStr := tools.Md5String(salt) + "53f0b847-82f6-43e1-8052-d6ebb97d1e0c"
	return tools.Md5String(password + saltStr)
}

//VerifyPassword 判断用户输入的密码是否正确
func VerifyPassword(userInfo *models.User, inputPassword string) bool {
	return hashPassword(inputPassword, userInfo.Salt) == userInfo.Password
}
