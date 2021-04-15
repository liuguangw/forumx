package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//Error404Handle 处理HTTP404
func Error404Handle(c *fiber.Ctx) error {
	responseData := fiber.Map{
		"code":    common.ErrorInternalServer,
		"message": "Page not found",
		"data":    nil,
	}
	if err := c.JSON(responseData); err != nil {
		return err
	}
	return c.SendStatus(404)
}
