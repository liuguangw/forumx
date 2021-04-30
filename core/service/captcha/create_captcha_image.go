package captcha

import (
	"context"
	"errors"
	"github.com/mojocn/base64Captcha"
	"strings"
)

//随机字符串
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//CreateCaptchaImage 生成图形验证码所需的二进制字节
func CreateCaptchaImage(ctx context.Context, captchaID string) ([]byte, error) {
	//判断验证码ID是否已存在
	if ctx == nil {
		ctx = context.Background()
	}
	cacheExists, _, err := LoadCaptchaCode(ctx, captchaID)
	if !cacheExists {
		return nil, errors.New("验证码ID已失效")
	}
	//生成验证码
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
	//save cache
	if err := cacheCaptchaCode(ctx, captchaID, strings.ToLower(captchaCode)); err != nil {
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
