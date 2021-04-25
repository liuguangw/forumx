package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/multifactory"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/pkg/errors"
)

//MultiFactoryToken 返回两步验证的令牌ID、密钥给客户端
func MultiFactoryToken(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err1 := session.CheckLogin(ctx, c)
	if err1 != nil || userSession == nil {
		return err1
	}
	tokenData, err := multifactory.GenerateToken(ctx, userSession)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成totp token失败"))
	}
	return response.WriteSuccess(c, tokenData)
}
