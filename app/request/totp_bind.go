package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//TotpBind 绑定二步验证令牌的请求
type TotpBind struct {
	Code string `json:"code" form:"code"` //动态码
}

//CheckRequest 检测用户输入
func (req *TotpBind) CheckRequest() *common.AppError {
	if req.Code == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "动态验证码不能为空")
	}
	return nil
}

//NewTotpBind 从客户端请求中初始化,绑定两步验证令牌所需的参数
func NewTotpBind(c *fiber.Ctx) (*TotpBind, *common.AppError) {
	req := new(TotpBind)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
