//Package routes 定义了系统所有的路由
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/controllers/api/auth"
	"github.com/liuguangw/forumx/app/controllers/api/captcha"
	"github.com/liuguangw/forumx/app/controllers/api/index"
	"github.com/liuguangw/forumx/app/controllers/api/mobile"
	"github.com/liuguangw/forumx/app/controllers/api/user"
)

//LoadAPIRoutes 加载系统路由和中间件配置
func LoadAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/", index.Hello)
	apiGroup.Get("/panic", index.PanicDemo)

	apiGroup.Post("/captcha/id", captcha.GenerateID)
	apiGroup.Get("/captcha/show", captcha.Show)

	apiGroup.Post("/auth/register", auth.Register)
	apiGroup.Post("/auth/login", auth.Login)
	apiGroup.Get("/auth/totp/random-token", auth.TotpRandomToken)
	apiGroup.Post("/auth/totp/bind", auth.TotpBind)
	apiGroup.Post("/auth/totp/verify", auth.TotpVerify)
	apiGroup.Post("/mobile/send-code", mobile.SendCode)
	apiGroup.Post("/mobile/bind-account", mobile.BindAccount)

	apiGroup.Get("/user/info", user.Info)
}
