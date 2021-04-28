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

//Register 处理用户注册请求
func Register(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewRegisterAccount(c)
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
	//判断存在性,防止重复
	username := req.Username
	email := req.EmailAddress
	if err := user.CheckRegisterUserExists(ctx, username, email); err != nil {
		return err.WriteResponse(c)
	}
	//注册账号
	clientIP := c.IP()
	userInfo, err := user.Register(ctx, username, req.Nickname,
		email, req.Password, clientIP)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "数据库服务异常"))
	}
	responseData := fiber.Map{
		"id":       userInfo.ID,
		"nickname": userInfo.Nickname,
	}
	return response.WriteSuccess(c, responseData)
}
