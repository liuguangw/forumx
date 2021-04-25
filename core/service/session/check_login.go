package session

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/response"
)

//CheckLogin 根据session判断用户是否已经登录,如果未登录则直接返回错误码给客户端
func CheckLogin(ctx context.Context, c *fiber.Ctx) (*models.UserSession, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	userSession, err := CheckRequest(ctx, c)
	if err != nil || userSession == nil {
		return nil, err
	}
	//用户未登录
	if userSession.UserID == 0 {
		return nil, response.WriteAppError(c, &common.AppError{
			Code:    common.ErrorNotLogin,
			Message: "当前用户未登录",
		})
	}
	return userSession, nil
}
