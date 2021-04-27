package session

import (
	"context"
	"github.com/liuguangw/forumx/core/models"
	"time"
)

//LoginUser 使用当前的会话登录用户账户
func LoginUser(ctx context.Context, userSession *models.UserSession, userID int64) error {
	if ctx == nil {
		ctx = context.Background()
	}
	userSession.UserID = userID
	//重新设置过期时间
	userSession.ExpiredAt = time.Now().Add(5 * 24 * time.Hour)
	return Save(ctx, userSession)
}
