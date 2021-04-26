package response

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
	"os"
)

//WriteInternalError 系统内部出现错误时返回统一的错误提示给客户端,并向stderr输出错误信息堆栈
func WriteInternalError(c *fiber.Ctx, systemError error) error {
	//打印错误堆栈信息
	err := fmt.Errorf("%+v", errors.WithStack(systemError))
	_, _ = os.Stderr.WriteString(err.Error() + "\n")
	//返回给客户端的统一响应
	return Write(c, &common.AppResponse{
		Code:    common.ErrorInternalServer,
		Message: "系统内部错误",
	})
}
