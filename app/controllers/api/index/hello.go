package index

import "github.com/gofiber/fiber/v2"

//Hello 简单的hello world Controller
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
