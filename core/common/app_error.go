package common

import "github.com/gofiber/fiber/v2"

//AppError 表示应用出现的错误
type AppError struct {
	Code    int    //错误码
	Message string //错误消息
}

//WriteResponse 将错误消息返回给客户端的快捷方法
func (err *AppError) WriteResponse(c *fiber.Ctx, httpStatus ...int) error {
	responseHTTPStatus := 200
	if err.Code == ErrorInternalServer {
		responseHTTPStatus = 500
	}
	if len(httpStatus) > 0 {
		responseHTTPStatus = httpStatus[0]
	}
	return c.Status(responseHTTPStatus).JSON(fiber.Map{
		"code":    err.Code,
		"message": err.Message,
		"data":    nil,
	})
}

//NewAppError 构造一个错误对象
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
