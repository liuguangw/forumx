package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/totp"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
	totp2 "github.com/pquerna/otp/totp"
)

//TotpVerify 登录后,使用动态码进行两步验证
func TotpVerify(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewTotpVerify(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, sessionErr := session.CheckSession(ctx, c)
	if sessionErr != nil {
		return sessionErr.WriteResponse(c)
	}
	userID := userSession.UserID
	//读取绑定的令牌信息
	userTotpKeyData, err := totp.FindTotpKeyByUserID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "读取绑定的令牌记录失败"))
	}
	if userTotpKeyData == nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取绑定的密钥信息失败"))
	}
	//验证动态码是否正确
	if !totp2.Validate(req.Code, userTotpKeyData.SecretKey) {
		//todo 验证错误次数+1
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "动态验证码错误")
	}
	userSession.Authenticated = true
	if err := session.Save(ctx, userSession); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "保存信息失败"))
	}
	return response.WriteSuccess(c, nil)
}
