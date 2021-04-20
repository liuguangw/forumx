package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
)

//WriteSuccessResponse 返回成功的响应给客户端
func WriteSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return (&common.AppResponse{Data: data}).Send(c)
}

//WriteErrorResponse 返回失败的响应给客户端
func WriteErrorResponse(c *fiber.Ctx, code int, message string, httpStatus ...int) error {
	if err := (&common.AppResponse{
		Code:    code,
		Message: message,
	}).Send(c); err != nil {
		return err
	}
	if len(httpStatus) > 0 {
		return c.SendStatus(httpStatus[0])
	}
	return nil
}
