package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//BindAccount 手机号绑定账号的请求
type BindAccount struct {
	Mobile string `json:"mobile" form:"mobile"` //手机号
	Code   string `json:"code" form:"code"`     //短信验证码
}

//CheckRequest 检测用户输入
func (req *BindAccount) CheckRequest() *common.AppError {
	if req.Mobile == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "手机号不能为空")
	}
	if req.Code == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "短信验证码不能为空")
	}
	return nil
}

//NewBindAccount 从客户端请求中初始化发送短信所需的参数
func NewBindAccount(c *fiber.Ctx) (*BindAccount, *common.AppError) {
	req := new(BindAccount)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
