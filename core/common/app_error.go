package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"os"
)

//AppError 表示应用出现的错误
type AppError struct {
	Code       int    //错误码
	Message    string //错误消息
	InnerError error  //内部错误信息
}

//WriteResponse 将错误消息返回给客户端的快捷方法
func (err *AppError) WriteResponse(c *fiber.Ctx, httpStatus ...int) error {
	responseHTTPStatus := 200
	if err.Code == ErrorInternalServer {
		responseHTTPStatus = 500
		//打印错误堆栈信息
		if err.InnerError != nil {
			//打印错误堆栈信息
			err1 := fmt.Errorf("%+v", errors.WithStack(err.InnerError))
			_, _ = os.Stderr.WriteString(err1.Error() + "\n")
		}
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

//Error 实现error接口
func (err *AppError) Error() string {
	return err.Message
}

//NewAppError 构造一个错误对象
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

//WrapAppError 返回一个包含内部错误信息的AppError
func WrapAppError(err error, message string) *AppError {
	return &AppError{
		Code:       ErrorInternalServer,
		Message:    "系统内部错误",
		InnerError: errors.Wrap(err, message),
	}
}
