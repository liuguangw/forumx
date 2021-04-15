package common

import "github.com/gofiber/fiber/v2"

//AppError 表示应用出现的错误
type AppError struct {
	Code    int    //错误码
	Message string //错误消息
}

//Error 实现error接口
func (err *AppError) Error() string {
	return err.Message
}

//WriteResponse 将错误消息返回给客户端的快捷方法
func (err *AppError) WriteResponse(c *fiber.Ctx, httpStatus ...int) error {
	if err := c.JSON(fiber.Map{
		"code":    err.Code,
		"message": err.Message,
		"data":    nil,
	}); err != nil {
		return err
	}
	if len(httpStatus) > 0 {
		return c.SendStatus(httpStatus[0])
	}
	return nil
}

//NewAppError 构造一个错误对象
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
