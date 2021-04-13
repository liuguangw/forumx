package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"os"
	"runtime"
)

//RecoverHandle 处理应用的panic
func RecoverHandle() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				//记录错误信息
				defaultStackTraceHandler(r)
				//返回给客户端的响应
				responseData := map[string]interface{}{
					"code":    common.ErrorInternalServer,
					"message": "系统内部错误",
					"data":    nil,
				}
				c.JSON(responseData)
				c.SendStatus(500)
			}
		}()
		return c.Next()
	}
}

func defaultStackTraceHandler(e interface{}) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}
