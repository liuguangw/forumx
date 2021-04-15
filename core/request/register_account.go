package request

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//RegisterAccount 注册账号的请求
type RegisterAccount struct {
	Username     string `json:"username"`      //用户名
	Nickname     string `json:"nickname"`      //昵称
	EmailAddress string `json:"email_address"` //邮箱地址
	Password     string `json:"password"`      //密码
	CaptchaID    string `json:"captcha_id"`    //验证码ID
	CaptchaCode  string `json:"captcha_code"`  //验证码
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
	if req.CaptchaID == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码ID不能为空")
	}
	if req.CaptchaCode != "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码不能为空")
	}
	return nil
}

//NewRegisterAccount 从客户端请求中初始化注册所需的参数
func NewRegisterAccount(c *fiber.Ctx) (*RegisterAccount, *common.AppError) {
	requestContentType := c.Get(requestContentTypeHeaderKey, "")
	if requestContentType == requestJSONContentType {
		requestData := new(RegisterAccount)
		if err := json.Unmarshal(c.Body(), requestData); err != nil {
			return nil, common.NewAppError(common.ErrorInputFieldInvalid, "JSON parse error: "+err.Error())
		}
		return requestData, nil
	}
	//form 表单
	return &RegisterAccount{
		Username:     c.FormValue("username", ""),
		Nickname:     c.FormValue("nickname", ""),
		EmailAddress: c.FormValue("email_address", ""),
		Password:     c.FormValue("password", ""),
		CaptchaID:    c.FormValue("captcha_id", ""),
		CaptchaCode:  c.FormValue("captcha_code", ""),
	}, nil
}
