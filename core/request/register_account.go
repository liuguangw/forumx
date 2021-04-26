package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//RegisterAccount 注册账号的请求
type RegisterAccount struct {
	Username     string `json:"username" form:"username"`           //用户名
	Nickname     string `json:"nickname" form:"nickname"`           //昵称
	EmailAddress string `json:"email_address" form:"email_address"` //邮箱地址
	Password     string `json:"password" form:"password"`           //密码
	CaptchaCode  string `json:"captcha_code" form:"captcha_code"`   //验证码
}

//CheckRequest 检测用户输入
func (req *RegisterAccount) CheckRequest() *common.AppError {
	if req.Username == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "用户名不能为空")
	}
	if req.Nickname == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "昵称不能为空")
	}
	if req.EmailAddress == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "邮箱地址不能为空")
	}
	if req.Password == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "密码不能为空")
	}
	if req.CaptchaCode == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码不能为空")
	}
	return nil
}

//NewRegisterAccount 从客户端请求中初始化注册所需的参数
func NewRegisterAccount(c *fiber.Ctx) (*RegisterAccount, *common.AppError) {
	req := new(RegisterAccount)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
