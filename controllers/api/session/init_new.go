package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service"
)

//InitNew 初始化Session并返回ID和过期时间给客户端
func InitNew(c *fiber.Ctx) error {
	sessionData := new(models.UserSession)
	if err := service.SaveUserSession(sessionData); err != nil {
		return service.WriteErrorResponse(c, common.ErrorInternalServer, "初始化会话失败")
	}
	return service.WriteSuccessResponse(c, fiber.Map{
		"id":         sessionData.ID,
		"expired_at": service.FormatDateTime(sessionData.ExpiredAt),
	})
}
