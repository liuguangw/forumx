package totp

import (
	"github.com/liuguangw/forumx/core/models"
)

//LoadKeyDataFromSession 从session中加载临时生成的两步验证令牌信息,此函数会在数据不存在时返回空字符串
func LoadKeyDataFromSession(userSession *models.UserSession) (secretKey string, recoveryCode string) {
	secretKey = userSession.Get(secretKeySessionKey)
	recoveryCode = userSession.Get(recoveryCodeSessionKey)
	return secretKey, recoveryCode
}
