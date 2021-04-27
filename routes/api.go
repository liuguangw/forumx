//Package routes 定义了系统所有的路由
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/controllers/api/auth"
	"github.com/liuguangw/forumx/controllers/api/captcha"
	"github.com/liuguangw/forumx/controllers/api/index"
	"github.com/liuguangw/forumx/controllers/api/session"
)

//LoadAPIRoutes 加载系统路由和中间件配置
func LoadAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/", index.Hello)
	apiGroup.Get("/panic", index.PanicDemo)

	apiGroup.Post("/session/init", session.InitNew)
	apiGroup.Get("/captcha/show", captcha.Show)

	apiGroup.Post("/auth/register", auth.Register)
	apiGroup.Post("/auth/login", auth.Login)
	apiGroup.Get("/auth/totp/random-token", auth.TotpRandomToken)
	apiGroup.Post("/auth/totp/bind", auth.TotpBind)
	apiGroup.Post("/auth/totp/verify", auth.TotpVerify)
}
