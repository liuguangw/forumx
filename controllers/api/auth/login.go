package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service/captcha"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/user"
	"github.com/pkg/errors"
)

//Login 处理用户登录请求
func Login(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewLoginAccount(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckSession(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	//检测验证码
	if !captcha.CheckCode(ctx, userSession, req.CaptchaCode, true) {
		return response.WriteAppError(c, common.ErrorInputFieldInvalid, "验证码错误")
	}
	//判断密码是否正确
	userInfo, err := user.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "find user "+req.Username+" failed"))
	}
	if userInfo == nil {
		return response.WriteAppError(c, common.ErrorUserNotFound, "不存在此用户")
	}
	if !user.VerifyPassword(userInfo, req.Password) {
		return response.WriteAppError(c, common.ErrorPassword, "用户名或密码错误")
	}
	userID := userInfo.ID
	//未开启两步验证的直接登录成功
	if !userInfo.Enable2FA {
		if err := session.LoginUser(ctx, userSession, userID); err != nil {
			return response.WriteInternalError(c, errors.Wrap(err, "用户登录失败"))
		}
		responseData := map[string]interface{}{
			"id":       userID,
			"nickname": userInfo.Nickname,
		}
		return response.WriteSuccess(c, responseData)
	}
	//生成临时的totp token缓存
	totpAuthToken, err := user.PrepareTotpAuth(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "prepare totp auth for "+req.Username+" failed"))
	}
	totpResponse := &common.AppResponse{
		Code:    common.ErrorNeedAuthentication,
		Message: "需要二次验证",
		Data: map[string]string{
			"token": totpAuthToken,
		},
	}
	return response.Write(c, totpResponse)
}
