package request

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//LoginAccount 登录账号的请求
type LoginAccount struct {
	Username    string `json:"username"`     //用户名
	Password    string `json:"password"`     //密码
	CaptchaCode string `json:"captcha_code"` //验证码
}

//CheckRequest 检测用户输入
func (req *LoginAccount) CheckRequest() *common.AppError {
	if req.Username == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "用户名不能为空")
	}
	if req.Password == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "密码不能为空")
	}
	if req.CaptchaCode == "" {
		return common.NewAppError(common.ErrorInputFieldIsEmpty, "验证码不能为空")
	}
	return nil
}

//NewLoginAccount 从客户端请求中初始化登录所需的参数
func NewLoginAccount(c *fiber.Ctx) (*LoginAccount, *common.AppError) {
	requestContentType := c.Get(requestContentTypeHeaderKey, "")
	if requestContentType == requestJSONContentType {
		requestData := new(LoginAccount)
		if err := json.Unmarshal(c.Body(), requestData); err != nil {
			return nil, common.NewAppError(common.ErrorInputFieldInvalid, "JSON parse error: "+err.Error())
		}
		return requestData, nil
	}
	//form 表单
	return &LoginAccount{
		Username:    c.FormValue("username", ""),
		Password:    c.FormValue("password", ""),
		CaptchaCode: c.FormValue("captcha_code", ""),
	}, nil
}
