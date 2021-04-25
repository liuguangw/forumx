package captcha

import (
	"context"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/session"
	"strings"
)

//CheckCode 检测用户输入的验证码是否正确
func CheckCode(ctx context.Context, userSession *models.UserSession, inputCaptchaCode string, clear bool) bool {
	captchaCode := userSession.Get(sessionKey)
	//存储的验证码为空,表示客户端未请求图片
	if captchaCode == "" {
		return false
	}
	if clear {
		userSession.Delete(sessionKey)
		_ = session.Save(ctx, userSession)
	}
	return captchaCode == strings.ToLower(inputCaptchaCode)
}
