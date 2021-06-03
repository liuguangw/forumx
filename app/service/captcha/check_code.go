package captcha

import (
	"context"
	"strings"
)

//CheckCode 检测用户输入的验证码是否正确
func CheckCode(ctx context.Context, captchaID, inputCaptchaCode string, clear bool) (bool, error) {
	cacheExists, captchaCode, err := LoadCaptchaCode(ctx, captchaID)
	if err != nil {
		return false, err
	}
	if !cacheExists {
		return false, nil
	}
	if clear {
		if err := cacheCaptchaCode(ctx, captchaID, ""); err != nil {
			return false, err
		}
	}
	return captchaCode == strings.ToLower(inputCaptchaCode), nil
}
