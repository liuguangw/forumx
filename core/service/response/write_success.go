package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//WriteSuccess 成功的响应
func WriteSuccess(c *fiber.Ctx, data interface{}) error {
	return Write(c, &common.AppResponse{Data: data})
}
