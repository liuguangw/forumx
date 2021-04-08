package cmd

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/urfave/cli/v2"
	"strconv"
)

//展示版本信息的命令
func serveCommand() *cli.Command {
	serveCmd := &cli.Command{
		Name:  "serve",
		Usage: "Run application API server",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "port", Usage: "server port", Value: 3000},
		},
		Action: func(c *cli.Context) error {
			app := fiber.New()
			app.Use(recover2.New(recover2.Config{
				EnableStackTrace: true,
			}))
			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World 👋!")
			})
			app.Get("/panic", func(c *fiber.Ctx) error {
				panic("normally this would crash your app")
			})
			port := c.Int("port")
			return app.Listen(":" + strconv.Itoa(port))
		},
	}
	return serveCmd
}
