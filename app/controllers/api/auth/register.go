package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/app/service/captcha"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/liuguangw/forumx/core/common"
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
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//检测验证码
	captchaPassed, err := captcha.CheckCode(ctx, req.CaptchaID, req.CaptchaCode, true)
	if err != nil {
		return response.WriteAppError(c, common.ErrorInternalServer, "判断验证码出错")
	}
	if !captchaPassed {
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
