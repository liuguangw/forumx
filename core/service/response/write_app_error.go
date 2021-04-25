package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//WriteAppError 应用错误的response
func WriteAppError(c *fiber.Ctx, err *common.AppError) error {
	resp := common.AppResponse{
		Code:    err.Code,
		Message: err.Message,
	}
	return Write(c, &resp)
}
