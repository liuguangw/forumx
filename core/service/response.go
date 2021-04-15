package service

import (
	"github.com/gofiber/fiber/v2"
)

//WriteSuccessResponse 返回成功的响应给客户端
func WriteSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

//WriteErrorResponse 返回失败的响应给客户端
func WriteErrorResponse(c *fiber.Ctx, code int, message string, httpStatus ...int) error {
	if err := c.JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    nil,
	}); err != nil {
		return err
	}
	if len(httpStatus) > 0 {
		return c.SendStatus(httpStatus[0])
	}
	return nil
}
