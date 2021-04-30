package session

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/models"
)

//CheckLogin 根据session判断用户是否已经登录,如果未登录则返回错误
func CheckLogin(ctx context.Context, c *fiber.Ctx) (*models.UserSession, *common.AppError) {
	if ctx == nil {
		ctx = context.Background()
	}
	userSession, err := CheckSession(ctx, c)
	if err != nil {
		return nil, err
	}
	//用户未登录
	if userSession.UserID == 0 {
		return nil, common.NewAppError(common.ErrorNotLogin, "当前用户未登录")
	}
	if !userSession.Authenticated {
		return nil, common.NewAppError(common.ErrorNeedAuthentication, "需要进行身份验证")
	}
	return userSession, nil
}
