package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service"
)

//Register 处理用户注册请求
func Register(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewRegisterAccount(c)
	if err != nil {
		return err.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	sessionID := service.GetRequestSessionID(c)
	userSession, err1 := service.GetUserSessionByID(sessionID)
	if err1 != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "读取会话失败")
	}
	if userSession == nil {
		return service.WriteErrorResponse(c, common.ErrorSessionExpired, "会话已失效")
	}
	//检测验证码
	if !service.CheckCaptchaCode(userSession, req.CaptchaCode, true) {
		return service.WriteErrorResponse(c, common.ErrorInputFieldInvalid, "验证码错误")
	}
	//注册账号
	clientIP := c.IP()
	user, registerError := service.RegisterUser(req.Username, req.Nickname,
		req.EmailAddress, req.Password, clientIP)
	if registerError != nil {
		return registerError.WriteResponse(c)
	}
	responseData := fiber.Map{
		"id":       user.ID,
		"nickname": user.Nickname,
	}
	return service.WriteSuccessResponse(c, responseData)
}
