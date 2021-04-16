package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
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
			appError := common.NewAppError(common.ErrorInternalServer, "系统内部错误")
			_ = appError.WriteResponse(c, 500)
		}
	}()
	return c.Next()
}

func defaultStackTraceHandler(e interface{}) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}
