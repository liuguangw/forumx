package captcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/service"
)

//Show 显示图形验证码
func Show(c *fiber.Ctx) error {
	sessionID := service.GetRequestSessionID(c)
	userSession, err := service.GetUserSessionByID(sessionID)
	if err != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "读取会话失败")
	}
	if userSession == nil {
		return service.WriteErrorResponse(c, common.ErrorSessionExpired, "会话已失效")
	}
	binData, err := service.CreateCaptcha(userSession)
	if err != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "生成验证码失败")
	}
	if _, err := c.Write(binData); err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return nil
}
