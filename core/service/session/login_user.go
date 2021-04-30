package session

import (
	"context"
	"github.com/liuguangw/forumx/core/models"
	"time"
)

//LoginUser 使用当前的会话登录用户账户
func LoginUser(ctx context.Context, userID int64, authenticated bool) (*models.UserSession, int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	timeNow := time.Now()
	sessionLifeTime := 5 * 24 * time.Hour
	expiredAt := time.Now().Add(sessionLifeTime)
	userSession := &models.UserSession{
		UserID:        userID,
		Authenticated: authenticated,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		ExpiredAt:     expiredAt,
	}
	//保存会话数据
	if err := Save(ctx, userSession); err != nil {
		return nil, 0, err
	}
	return userSession, int64(sessionLifeTime.Seconds()), nil
}
