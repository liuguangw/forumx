package captcha

import (
	"context"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/mojocn/base64Captcha"
	"strings"
)

//随机字符串
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//CreateImage 生成图形验证码所需的二进制字节
func CreateImage(ctx context.Context, userSession *models.UserSession) ([]byte, error) {
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
		sessionKey: strings.ToLower(captchaCode),
	}
	//save session
	if ctx == nil {
		ctx = context.Background()
	}
	if err := session.Save(ctx, userSession); err != nil {
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
