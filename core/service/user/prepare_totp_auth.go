package user

import (
	"context"
	"github.com/liuguangw/forumx/core/service/cache"
	"github.com/liuguangw/forumx/core/service/tools"
	"time"
)

func totpAuthUserIDCacheKey(totpAuthToken string) string {
	return "totp_auth_tokens." + totpAuthToken + ".user_id"
}

//PrepareTotpAuth 为totp认证做准备
// 此函数会随机生成一个key,并将用户ID作为值写入缓存
func PrepareTotpAuth(ctx context.Context, userID int64) (string, error) {
	var (
		totpAuthToken string //随机token
		tokenValid    bool   //token是否有效
	)
	if ctx == nil {
		ctx = context.Background()
	}
	//生成一个不存在的token
	for !tokenValid {
		totpAuthToken = tools.GenerateHashID()
		cachedUserID, err := LoadTotpAuthUserID(ctx, totpAuthToken)
		if err != nil {
			return "", err
		}
		tokenValid = cachedUserID == 0
	}
	//缓存
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	expiredAt := time.Now().Add(15 * time.Minute)
	if err := cache.PutItem(ctx, cacheKey, &userID, expiredAt); err != nil {
		return "", err
	}
	return totpAuthToken, nil
}

//LoadTotpAuthUserID 根据token获取已验证密码的用户ID,如果不存在则返回0
func LoadTotpAuthUserID(ctx context.Context, totpAuthToken string) (int64, error) {
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	if ctx == nil {
		ctx = context.Background()
	}
	var cachedUserID int64
	_, err := cache.GetItem(ctx, cacheKey, &cachedUserID)
	if err != nil {
		return 0, err
	}
	return cachedUserID, nil
}
