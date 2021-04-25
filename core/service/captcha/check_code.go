package captcha

import (
	"context"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/session"
	"strings"
)

//CheckCode 检测用户输入的验证码是否正确
func CheckCode(ctx context.Context, userSession *models.UserSession, inputCaptchaCode string, clear bool) bool {
	sessionData := userSession.Data
	if captchaCode, ok := sessionData[sessionKey]; ok {
		captchaCodeStr := strings.ToLower(captchaCode.(string))
		result := captchaCodeStr == inputCaptchaCode
		if clear {
			delete(sessionData, sessionKey)
			_ = session.Save(ctx, userSession)
		}
		return result
	}
	return false
}
