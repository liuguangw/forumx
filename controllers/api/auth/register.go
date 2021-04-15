package auth

import (
	"github.com/gofiber/fiber/v2"
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
	//todo 检测验证码是否正确
	//注册账号
	user, registerError := service.RegisterUser(req.Username, req.Nickname,
		req.EmailAddress, req.Password)
	if registerError != nil {
		return registerError.WriteResponse(c)
	}
	responseData := fiber.Map{
		"id":       user.ID,
		"nickname": user.Nickname,
	}
	return service.WriteSuccessResponse(c, responseData)
}
