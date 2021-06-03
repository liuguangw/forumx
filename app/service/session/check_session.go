package session

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"strings"
)

//CheckSession 判断当前用户会话是否有效,如果无效则返回错误
func CheckSession(ctx context.Context, c *fiber.Ctx) (*models.UserSession, *common.AppError) {
	sessionID := getRequestSessionID(c)
	if sessionID == "" {
		return nil, common.NewAppError(common.ErrorSessionExpired, "未传入会话标识")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	userSession, err := LoadByID(ctx, sessionID)
	if err != nil {
		return nil, common.WrapAppError(err, "load session "+sessionID+" failed")
	}
	if userSession == nil {
		return nil, common.NewAppError(common.ErrorSessionExpired, "会话无效或者已经过期")
	}
	return userSession, nil
}

//getRequestSessionID 从客户端请求中,读取会话ID唯一标识符
func getRequestSessionID(c *fiber.Ctx) string {
	authorizationValue := c.Get("Authorization", "")
	tokenPrefix := "Bearer"
	if strings.Index(authorizationValue, tokenPrefix) == 0 {
		return authorizationValue[len(tokenPrefix)+1:]
	}
	//从URL中读取
	return c.Query("sid", "")
}
