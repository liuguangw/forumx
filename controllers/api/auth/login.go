package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service"
)

//Login 处理用户登录请求
func Login(c *fiber.Ctx) error {
	//获取所需参数
	req, err := request.NewLoginAccount(c)
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
	//判断密码是否正确
	userInfo, dbErr := service.FindUserByUsername(req.Username)
	if dbErr != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "后台服务异常", 500)
	}
	if userInfo == nil {
		return service.WriteErrorResponse(c, common.ErrorUserNotFound, "不存在此用户")
	}
	if !service.VerifyPassword(userInfo, req.Password) {
		return service.WriteErrorResponse(c, common.ErrorPassword, "用户名或密码错误")
	}
	userSession.UserID = userInfo.ID
	userSession.Authed = true
	if userInfo.Enable2FA {
		userSession.Authed = false
	}
	//保存会话数据
	if err := service.SaveUserSession(userSession); err != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "后台服务异常", 500)
	}
	if userInfo.Enable2FA {
		return service.WriteErrorResponse(c, common.ErrorNeedAuthentication, "需要身份验证")
	}
	responseData := map[string]interface{}{
		"id":       userInfo.ID,
		"nickname": userInfo.Nickname,
	}
	return service.WriteSuccessResponse(c, responseData)
}
