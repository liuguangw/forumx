package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/environment"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/totp"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/pkg/errors"
)

//TotpRandomToken 返回随机的两步验证密钥、恢复码给客户端
func TotpRandomToken(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, sessionErr := session.CheckLogin(ctx, c)
	if sessionErr != nil {
		return sessionErr.WriteResponse(c)
	}
	secretKey, recoveryCode, err := totp.GenerateRandomKeyData(ctx, userSession)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成totp密钥失败"))
	}
	//获取用户信息
	userInfo, err := user.FindUserByID(ctx, userSession.UserID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	//生成totp URL
	username := userInfo.Username
	siteEnName := environment.SiteEnName()
	totpURL := "otpauth://totp/" + siteEnName + ":" + username +
		"?secret=" + secretKey + "&issuer=" + siteEnName
	responseData := map[string]string{
		"url":           totpURL,
		"recovery_code": recoveryCode,
	}
	return response.WriteSuccess(c, responseData)
}
