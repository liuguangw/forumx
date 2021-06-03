package captcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/service/captcha"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
)

//Show 显示图形验证码
func Show(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	captchaID := c.Query("id")
	if captchaID == "" {
		return response.WriteAppError(c, common.ErrorCommonMessage, "缺少id参数")
	}
	binData, err := captcha.CreateCaptchaImage(ctx, captchaID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "生成验证码失败"))
	}
	if _, err := c.Write(binData); err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return nil
}
