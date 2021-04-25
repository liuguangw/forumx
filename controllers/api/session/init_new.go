package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/pkg/errors"
	"time"
)

//InitNew 初始化Session并返回ID和过期时间给客户端
func InitNew(c *fiber.Ctx) error {
	sessionLifetime := 15 * time.Minute
	userSession := &models.UserSession{
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(sessionLifetime),
	}
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	if err := session.Save(ctx, userSession); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "init session failed"))
	}
	return response.WriteSuccess(c, fiber.Map{
		"id":         userSession.ID,
		"expired_at": tools.FormatDateTime(userSession.ExpiredAt),
	})
}
