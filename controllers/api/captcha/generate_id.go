package captcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/captcha"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/pkg/errors"
)

//GenerateID 随机生成一个验证码ID
func GenerateID(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	captchaID, err := captcha.CreateCaptchaID(ctx)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成验证码ID出错"))
	}
	return response.WriteSuccess(c, fiber.Map{
		"id": captchaID,
	})
}
