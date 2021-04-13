package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/middlewares"
	"github.com/liuguangw/forumx/routes"
	"github.com/urfave/cli/v2"
	"strconv"
)

//处理HTTP服务
func serveCommand() *cli.Command {
	serveCmd := &cli.Command{
		Name:  "serve",
		Usage: "Run application API server",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "port", Usage: "server port", Value: 3000},
		},
		Action: func(c *cli.Context) error {
			app := fiber.New()
			app.Use(middlewares.RecoverHandle())
			//加载api路由
			routes.LoadAPIRoutes(app)
			//端口
			port := c.Int("port")
			return app.Listen(":" + strconv.Itoa(port))
		},
	}
	return serveCmd
}
