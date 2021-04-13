package index

import "github.com/gofiber/fiber/v2"

//PanicDemo panic的demo
func PanicDemo(c *fiber.Ctx) error {
	panic("panic demo ...")
}
