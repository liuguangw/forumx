package captcha

import (
	"context"
	"github.com/liuguangw/forumx/app/service/cache"
	"github.com/liuguangw/forumx/app/service/tools"
	"time"
)

const captchaCacheLifetime = 20 * time.Minute

//CreateCaptchaID 生成一个唯一的验证码ID, 生成验证码时可以使用此ID来缓存验证码
func CreateCaptchaID(ctx context.Context) (string, error) {
	var (
		captchaID      string //验证码ID
		captchaIDValid bool   //验证码ID是否有效
	)
	if ctx == nil {
		ctx = context.Background()
	}
	//随机生成验证码ID
	for !captchaIDValid {
		captchaID = tools.GenerateHashID()
		cacheExists, _, err := LoadCaptchaCode(ctx, captchaID)
		if err != nil {
			return "", err
		}
		//验证码ID缓存必须不存在
		captchaIDValid = !cacheExists
	}
	//初始化缓存
	if err := cacheCaptchaCode(ctx, captchaID, ""); err != nil {
		return "", err
	}
	return captchaID, nil
}

//cacheCaptchaCode 缓存验证码
func cacheCaptchaCode(ctx context.Context, captchaID, code string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	cacheKey := getCaptchaCacheKey(captchaID)
	expiredAt := time.Now().Add(captchaCacheLifetime)
	return cache.PutItem(ctx, cacheKey, code, expiredAt)
}

//LoadCaptchaCode 根据验证码ID读取验证码,
//此函数只在当前包内部、单元测试时调用,其它地方不应该调用此函数
func LoadCaptchaCode(ctx context.Context, captchaID string) (bool, string, error) {
	var cachedItem struct {
		ItemValue string `bson:"item_value"`
	}
	if ctx == nil {
		ctx = context.Background()
	}
	cacheKey := getCaptchaCacheKey(captchaID)
	cacheExists, err := cache.GetItem(ctx, cacheKey, &cachedItem)
	if err != nil {
		return false, "", err
	}
	return cacheExists, cachedItem.ItemValue, nil
}

//getCaptchaCacheKey 计算验证码ID对应的缓存key
func getCaptchaCacheKey(captchaID string) string {
	return "captcha_ids." + captchaID + ".code"
}
