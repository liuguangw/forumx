package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/service/response"
	"os"
	"runtime"
)

//RecoverHandle 处理应用的panic
func RecoverHandle(c *fiber.Ctx) error {
	// Catch panics
	defer func() {
		if r := recover(); r != nil {
			//记录错误信息
			defaultStackTraceHandler(r)
			//返回给客户端的响应
			_ = response.WriteAppError(c, &common.AppError{
				Code:    common.ErrorInternalServer,
				Message: "服务器异常",
			})
		}
	}()
	return c.Next()
}

func defaultStackTraceHandler(e interface{}) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}
