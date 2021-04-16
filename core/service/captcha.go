package service

import (
	"github.com/liuguangw/forumx/core/models"
	"github.com/mojocn/base64Captcha"
	"strings"
)

const captchaCodeSessionKey = "captcha_code" //验证码存储的Session key

//CreateCaptcha 生成图形验证码所需的二进制字节
func CreateCaptcha(userSession *models.UserSession) ([]byte, error) {
	captchaDriver := &base64Captcha.DriverString{
		Height:          60,
		Width:           165,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          4,
		Source:          letterBytes,
		Fonts:           []string{"actionj.ttf"},
	}
	//generate code string
	_, _, captchaCode := captchaDriver.GenerateIdQuestionAnswer()
	userSession.Data = map[string]interface{}{
		captchaCodeSessionKey: strings.ToLower(captchaCode),
	}
	//save session
	if err := SaveUserSession(userSession); err != nil {
		return nil, err
	}
	//draw
	item, err := captchaDriver.DrawCaptcha(captchaCode)
	if err != nil {
		return nil, err
	}
	var binData []byte
	if itemChar, ok := item.(*base64Captcha.ItemChar); ok {
		binData = itemChar.BinaryEncoding()
	}
	return binData, nil
}

//CheckCaptchaCode 检测用户输入的验证码是否正确
func CheckCaptchaCode(userSession *models.UserSession, inputCaptchaCode string, clear bool) bool {
	sessionData := userSession.Data
	if captchaCode, ok := sessionData[captchaCodeSessionKey]; ok {
		captchaCodeStr := strings.ToLower(captchaCode.(string))
		result := captchaCodeStr == inputCaptchaCode
		if clear {
			delete(sessionData, captchaCodeSessionKey)
			_ = SaveUserSession(userSession)
		}
		return result
	}
	return false
}
