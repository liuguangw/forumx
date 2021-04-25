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
	resp := common.AppResponse{
		Code:    common.ErrorInternalServer,
		Message: "系统内部错误",
	}
	err := fmt.Errorf("%+v", errors.WithStack(systemError))
	_, _ = os.Stderr.WriteString(err.Error() + "\n")
	return Write(c, &resp)
}
