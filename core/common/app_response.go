package common

import "github.com/gofiber/fiber/v2"

//AppResponse API接口返回给客户端的响应
type AppResponse struct {
	Code    int         `json:"code"`    //状态码
	Message string      `json:"message"` //错误消息
	Data    interface{} `json:"data"`    //负载数据
}

//Send 发送response给客户端
func (res *AppResponse) Send(c *fiber.Ctx) error {
	return c.JSON(res)
}
