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
)

//Register 处理用户注册请求
func Register(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewRegisterAccount(c)
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
	//注册账号
	clientIP := c.IP()
	userInfo, registerError := user.Register(ctx, req.Username, req.Nickname,
		req.EmailAddress, req.Password, clientIP)
	if registerError != nil {
		return response.WriteAppError(c, registerError)
	}
	responseData := fiber.Map{
		"id":       userInfo.ID,
		"nickname": userInfo.Nickname,
	}
	return response.WriteSuccess(c, responseData)
}
