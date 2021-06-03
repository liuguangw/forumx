package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//LoginAccount 登录账号的请求
type LoginAccount struct {
	Username    string `json:"username" form:"username"`         //用户名
	Password    string `json:"password" form:"password"`         //密码
	CaptchaID   string `json:"captcha_id" form:"captcha_id"`     //验证码ID
	CaptchaCode string `json:"captcha_code" form:"captcha_code"` //验证码
}

//CheckRequest 检测用户输入
func (req *LoginAccount) CheckRequest() *common.AppError {
	if req.Username == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "用户名不能为空")
	}
	if req.Password == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "密码不能为空")
	}
	if req.CaptchaID == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码ID不能为空")
	}
	if req.CaptchaCode == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码不能为空")
	}
	return nil
}

//NewLoginAccount 从客户端请求中初始化登录所需的参数
func NewLoginAccount(c *fiber.Ctx) (*LoginAccount, *common.AppError) {
	req := new(LoginAccount)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
