package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service"
	"github.com/pkg/errors"
)

//InitNew 初始化Session并返回ID和过期时间给客户端
func InitNew(c *fiber.Ctx) error {
	sessionData := new(models.UserSession)
	if err := service.SaveUserSession(sessionData); err != nil {
		return service.WriteInternalErrorResponse(c, errors.Wrap(err, "init session failed"))
	}
	return service.WriteSuccessResponse(c, fiber.Map{
		"id":         sessionData.ID,
		"expired_at": service.FormatDateTime(sessionData.ExpiredAt),
	})
}
