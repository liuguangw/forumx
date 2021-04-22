package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service"
	"github.com/pkg/errors"
)

//Register 处理用户注册请求
func Register(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewRegisterAccount(c)
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
	//注册账号
	clientIP := c.IP()
	user, registerError := service.RegisterUser(req.Username, req.Nickname,
		req.EmailAddress, req.Password, clientIP)
	if registerError != nil {
		return service.WriteAppErrorResponse(c, registerError)
	}
	responseData := fiber.Map{
		"id":       user.ID,
		"nickname": user.Nickname,
	}
	return service.WriteSuccessResponse(c, responseData)
}
