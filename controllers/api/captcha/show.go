package captcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/captcha"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/pkg/errors"
)

//Show 显示图形验证码
func Show(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckSession(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	binData, err := captcha.CreateImage(ctx, userSession)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成验证码失败"))
	}
	if _, err := c.Write(binData); err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return nil
}
