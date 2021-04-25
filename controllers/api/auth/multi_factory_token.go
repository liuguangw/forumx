package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
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
	//todo 随机生成totp密钥和ID
	return nil
}
