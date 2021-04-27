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

//TotpVerify 登录后,使用动态码进行两步验证
func TotpVerify(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewMultiFactoryBind(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckLogin(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	//判断用户是否已经绑定过令牌了
	userID := userSession.UserID
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	if !userInfo.Enable2FA {
		return response.WriteAppError(c, common.ErrorCommonMessage, "当前账户还未启用两步验证")
	}
	//读取绑定的令牌信息
	tokenData, err := totp.FindTokenByUserID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "读取绑定的令牌记录失败"))
	}
	//验证动态码是否正确
	if !totp2.Validate(req.Code, tokenData.SecretKey) {
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "动态验证码错误")
	}
	//保存会话
	userSession.Authed = true
	if err := session.Save(ctx, userSession); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "保存会话数据失败"))
	}
	return response.WriteSuccess(c, nil)
}
