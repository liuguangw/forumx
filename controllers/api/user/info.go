package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
)

//Info 获取当前登录的用户信息
func Info(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckLogin(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	//todo 处理用户信息的展示
	return nil
}
