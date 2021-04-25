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
	"time"
)

//Login 处理用户登录请求
func Login(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewLoginAccount(c)
	if err != nil {
		return response.WriteAppError(c, err)
	}
	if err := req.CheckRequest(); err != nil {
		return response.WriteAppError(c, err)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err1 := session.CheckRequest(ctx, c)
	if err1 != nil || userSession == nil {
		return err1
	}
	//检测验证码
	if !captcha.CheckCode(ctx, userSession, req.CaptchaCode, true) {
		return response.WriteAppError(c, &common.AppError{
			Code:    common.ErrorInputFieldInvalid,
			Message: "验证码错误",
		})
	}
	//判断密码是否正确
	userInfo, dbErr := user.FindUserByUsername(ctx, req.Username)
	if dbErr != nil {
		return response.WriteInternalError(c, errors.Wrap(err1, "find user "+req.Username+" failed"))
	}
	if userInfo == nil {
		return response.WriteAppError(c, &common.AppError{
			Code:    common.ErrorUserNotFound,
			Message: "不存在此用户",
		})
	}
	if !user.VerifyPassword(userInfo, req.Password) {
		return response.WriteAppError(c, &common.AppError{
			Code:    common.ErrorPassword,
			Message: "用户名或密码错误",
		})
	}
	userSession.UserID = userInfo.ID
	userSession.Authed = true
	if userInfo.Enable2FA {
		userSession.Authed = false
	}
	//session生命周期重新设置
	userSession.ExpiredAt = time.Now().Add(5 * 24 * time.Hour)
	//保存会话数据
	if err := session.Save(ctx, userSession); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err1, "save session "+userSession.ID+" failed"))
	}
	if userInfo.Enable2FA {
		return response.WriteAppError(c, &common.AppError{
			Code:    common.ErrorNeedAuthentication,
			Message: "需要身份验证",
		})
	}
	responseData := map[string]interface{}{
		"id":       userInfo.ID,
		"nickname": userInfo.Nickname,
	}
	return response.WriteSuccess(c, responseData)
}
