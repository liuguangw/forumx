package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//Write 返回响应给客户端
func Write(c *fiber.Ctx, resp *common.AppResponse, httpStatus ...int) error {
	responseHTTPStatus := 200
	if resp.Code == common.ErrorInternalServer {
		responseHTTPStatus = 500
	}
	if len(httpStatus) > 0 {
		responseHTTPStatus = httpStatus[0]
	}
	return c.Status(responseHTTPStatus).JSON(resp)
}
