package session

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/pkg/errors"
	"strings"
)

//CheckSession 判断用户的session ID是否存在切有效,如果无效则直接返回错误码给客户端
func CheckSession(ctx context.Context, c *fiber.Ctx) (*models.UserSession, error) {
	sessionID := getRequestSessionID(c)
	if sessionID == "" {
		return nil, response.WriteAppError(c, common.ErrorSessionExpired, "未传入会话ID")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	userSession, err := LoadByID(ctx, sessionID)
	if err != nil {
		return nil, response.WriteInternalError(c, errors.Wrap(err, "load session "+sessionID+" failed"))
	}
	if userSession == nil {
		return nil, response.WriteAppError(c, common.ErrorSessionExpired, "会话无效或者已经过期")
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
