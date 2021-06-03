package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
)

//SendSms 发送短信验证码的请求
type SendSms struct {
	CaptchaID   string `json:"captcha_id" form:"captcha_id"`     //图形验证码ID
	CaptchaCode string `json:"captcha_code" form:"captcha_code"` //图形验证码
	CodeType    int    `json:"code_type" form:"code_type"`       //验证码类型
	Mobile      string `json:"mobile" form:"mobile"`             //手机号
}

//CheckRequest 检测用户输入
func (req *SendSms) CheckRequest() *common.AppError {
	if req.CaptchaID == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码ID不能为空")
	}
	if req.CaptchaCode == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码不能为空")
	}
	if !(req.CodeType == models.MobileCodeTypeBindAccount || req.CodeType == models.MobileCodeTypeResetPassword) {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "无效的验证码类型")
	}
	if req.Mobile == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "手机号不能为空")
	}
	return nil
}

//NewSendSms 从客户端请求中初始化发送短信所需的参数
func NewSendSms(c *fiber.Ctx) (*SendSms, *common.AppError) {
	req := new(SendSms)
	if err := c.BodyParser(req); err != nil {
		return nil, common.NewAppError(common.ErrorBadRequest, err.Error())
	}
	return req, nil
}
