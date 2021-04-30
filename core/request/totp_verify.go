package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//TotpVerify 两步验证请求
type TotpVerify struct {
	Code string `json:"code" form:"code"` //动态码
}

//CheckRequest 检测用户输入
func (req *TotpVerify) CheckRequest() *common.AppError {
	if req.Code == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "动态验证码不能为空")
	}
	return nil
}

//NewTotpVerify 从客户端请求中初始化,两步验证登录所需的参数
func NewTotpVerify(c *fiber.Ctx) (*TotpVerify, *common.AppError) {
	req := new(TotpVerify)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
