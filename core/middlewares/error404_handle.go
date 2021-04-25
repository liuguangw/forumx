package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/service/response"
)

//Error404Handle 处理HTTP404
func Error404Handle(c *fiber.Ctx) error {
	resp := &common.AppResponse{
		Code:    common.ErrorInternalServer,
		Message: "Page not found",
	}
	return response.Write(c, resp, 404)
}
