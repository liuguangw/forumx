package session

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ExampleCheckLogin() {
	//用户请求的Ctx
	var c *fiber.Ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userSession, err := CheckLogin(ctx, c)
	if err != nil {
		_ = err.WriteResponse(c)
		return
	}
	fmt.Println("continue other things")
	fmt.Println(userSession.UserID)
}
