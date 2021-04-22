package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service"
	"github.com/pkg/errors"
)

//Login 处理用户登录请求
func Login(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewLoginAccount(c)
	if err != nil {
		return service.WriteAppErrorResponse(c, err)
	}
	if err := req.CheckRequest(); err != nil {
		return service.WriteAppErrorResponse(c, err)
	}
	//加载session
	sessionID := service.GetRequestSessionID(c)
	userSession, err1 := service.GetUserSessionByID(sessionID)
	if err1 != nil {
		return service.WriteInternalErrorResponse(c, errors.Wrap(err1, "load session "+sessionID+" failed"))
	}
	if userSession == nil {
		return service.WriteAppErrorResponse(c, &common.AppError{
			Code:    common.ErrorSessionExpired,
			Message: "会话已失效",
		})
	}
	//检测验证码
	if !service.CheckCaptchaCode(userSession, req.CaptchaCode, true) {
		return service.WriteAppErrorResponse(c, &common.AppError{
			Code:    common.ErrorInputFieldInvalid,
			Message: "验证码错误",
		})
	}
	//判断密码是否正确
	userInfo, dbErr := service.FindUserByUsername(req.Username)
	if dbErr != nil {
		return service.WriteInternalErrorResponse(c, errors.Wrap(err1, "find user "+req.Username+" failed"))
	}
	if userInfo == nil {
		return service.WriteAppErrorResponse(c, &common.AppError{
			Code:    common.ErrorUserNotFound,
			Message: "不存在此用户",
		})
	}
	if !service.VerifyPassword(userInfo, req.Password) {
		return service.WriteAppErrorResponse(c, &common.AppError{
			Code:    common.ErrorPassword,
			Message: "用户名或密码错误",
		})
	}
	userSession.UserID = userInfo.ID
	userSession.Authed = true
	if userInfo.Enable2FA {
		userSession.Authed = false
	}
	//保存会话数据
	if err := service.SaveUserSession(userSession); err != nil {
		return service.WriteInternalErrorResponse(c, errors.Wrap(err1, "save session "+sessionID+" failed"))
	}
	if userInfo.Enable2FA {
		return service.WriteAppErrorResponse(c, &common.AppError{
			Code:    common.ErrorNeedAuthentication,
			Message: "需要身份验证",
		})
	}
	responseData := map[string]interface{}{
		"id":       userInfo.ID,
		"nickname": userInfo.Nickname,
	}
	return service.WriteSuccessResponse(c, responseData)
}
