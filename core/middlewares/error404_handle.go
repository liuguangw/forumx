package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/service"
)

//Error404Handle 处理HTTP404
func Error404Handle(c *fiber.Ctx) error {
	response := common.AppResponse{
		Code:    common.ErrorInternalServer,
		Message: "Page not found",
	}
	return service.WriteResponse(c, &response, 404)
}
