package user

import (
	"context"
	"github.com/liuguangw/forumx/app/service/cache"
	"github.com/liuguangw/forumx/app/service/tools"
	"time"
)

//TotpAuthData 当用户登录成功后,临时保存的用户信息,
//用于在验证动态码时读取
type TotpAuthData struct {
	UserID     int64 `bson:"user_id"`     //用户ID
	ErrorCount int   `bson:"error_count"` //动态码验证失败的次数
}

func totpAuthUserIDCacheKey(totpAuthToken string) string {
	return "totp_tokens." + totpAuthToken + ".auth_data"
}

//PrepareTotpAuth 为totp认证做准备
// 此函数会随机生成一个token,根据token生成key,并将已验证过密码的用户信息作为value写入缓存,
// 返回随机生成的token
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
		cachedAuthData, err := LoadTotpAuthData(ctx, totpAuthToken)
		if err != nil {
			return "", err
		}
		tokenValid = cachedAuthData == nil
	}
	//缓存
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	expiredAt := time.Now().Add(15 * time.Minute)
	authData := TotpAuthData{UserID: userID}
	if err := cache.PutItem(ctx, cacheKey, &authData, expiredAt); err != nil {
		return "", err
	}
	return totpAuthToken, nil
}

//LoadTotpAuthData 根据token获取已验证密码的用户信息
func LoadTotpAuthData(ctx context.Context, totpAuthToken string) (*TotpAuthData, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	var cachedItem struct {
		ItemValue *TotpAuthData `bson:"item_value"`
	}
	cacheExists, err := cache.GetItem(ctx, cacheKey, &cachedItem)
	//fmt.Println(totpAuthToken, cacheKey, cacheExists, cachedItem, err)
	if err != nil {
		return nil, err
	}
	if !cacheExists {
		return nil, nil
	}
	return cachedItem.ItemValue, nil
}

//ClearTotpAuthData 当用户验证成功后,清理缓存的验证信息
func ClearTotpAuthData(ctx context.Context, totpAuthToken string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	return cache.DeleteItem(ctx, cacheKey)
}

//IncrTotpAuthFailedCount 在两步验证失败时,增加失败次数的计数
func IncrTotpAuthFailedCount(ctx context.Context, totpAuthToken string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	authData, err := LoadTotpAuthData(ctx, totpAuthToken)
	if err != nil {
		return err
	}
	//不存在记录
	if authData == nil {
		return nil
	}
	//错误次数+1
	authData.ErrorCount++
	cacheKey := totpAuthUserIDCacheKey(totpAuthToken)
	expiredAt := time.Now().Add(15 * time.Minute)
	return cache.PutItem(ctx, cacheKey, &authData, expiredAt)
}
