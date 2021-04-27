package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/environment"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/totp"
	"github.com/liuguangw/forumx/core/service/user"
	"github.com/pkg/errors"
)

//TotpRandomToken 返回随机的两步验证密钥、恢复码给客户端
func TotpRandomToken(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckLogin(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	tokenData, err := totp.GenerateToken(ctx, userSession)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成totp token失败"))
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
		"?secret=" + tokenData.SecretKey + "&issuer=" + siteEnName
	responseData := map[string]string{
		"url":           totpURL,
		"recovery_code": tokenData.RecoveryCode,
	}
	return response.WriteSuccess(c, responseData)
}
