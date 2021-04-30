package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/totp"
	"github.com/liuguangw/forumx/core/service/user"
	"github.com/pkg/errors"
	totp2 "github.com/pquerna/otp/totp"
)

//TotpBind 绑定两步验证令牌
func TotpBind(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewTotpBind(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, sessionErr := session.CheckLogin(ctx, c)
	if sessionErr != nil {
		return sessionErr.WriteResponse(c)
	}
	//判断用户是否已经绑定过令牌了
	userInfo, err := user.FindUserByID(ctx, userSession.UserID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	if userInfo.Enable2FA {
		return response.WriteAppError(c, common.ErrorCommonMessage, "您的账户已经启用过两步验证了")
	}
	//读取令牌信息
	secretKey, recoveryCode := totp.LoadKeyDataFromSession(userSession)
	if secretKey == "" || recoveryCode == "" {
		return response.WriteInternalError(c, errors.Wrap(err, "解码令牌数据失败"))
	}
	//验证动态码是否正确
	if !totp2.Validate(req.Code, secretKey) {
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "动态验证码错误")
	}
	if err := totp.BindUserAccount(ctx, userInfo, secretKey, recoveryCode); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "绑定两步验证令牌失败"))
	}
	return response.WriteSuccess(c, nil)
}
