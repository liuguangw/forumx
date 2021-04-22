package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
	"os"
)

//WriteResponse 返回响应给客户端
func WriteResponse(c *fiber.Ctx, resp *common.AppResponse, httpStatus ...int) error {
	responseHTTPStatus := 200
	if resp.Code == common.ErrorInternalServer {
		responseHTTPStatus = 500
	}
	if len(httpStatus) > 0 {
		responseHTTPStatus = httpStatus[0]
	}
	return c.Status(responseHTTPStatus).JSON(resp)
}

//WriteSuccessResponse 成功的响应
func WriteSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return WriteResponse(c, &common.AppResponse{Data: data})
}

//WriteAppErrorResponse 应用错误的response
func WriteAppErrorResponse(c *fiber.Ctx, err *common.AppError) error {
	resp := common.AppResponse{
		Code:    err.Code,
		Message: err.Message,
	}
	return WriteResponse(c, &resp)
}

//WriteInternalErrorResponse 系统内部出现错误时的响应
func WriteInternalErrorResponse(c *fiber.Ctx, systemError error) error {
	resp := common.AppResponse{
		Code:    common.ErrorInternalServer,
		Message: "系统内部错误",
	}
	err := fmt.Errorf("%+v", errors.WithStack(systemError))
	_, _ = os.Stderr.WriteString(err.Error() + "\n")
	return WriteResponse(c, &resp)
}
