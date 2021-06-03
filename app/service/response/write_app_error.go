package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//WriteAppError 应用错误的response
func WriteAppError(c *fiber.Ctx, code int, message string) error {
	return Write(c, &common.AppResponse{
		Code:    code,
		Message: message,
	})
}
