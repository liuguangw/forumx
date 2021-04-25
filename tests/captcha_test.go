package tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//testCaptchaShow 测试验证码接口
func testCaptchaShow(app *fiber.App, sessionID string, t *testing.T) string {
	req, err := http.NewRequest(
		"GET",
		"/api/captcha/show",
		nil,
	)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+sessionID)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	//获取图形验证码
	userSession, err := session.LoadByID(nil, sessionID)
	assert.NoError(t, err)
	return userSession.Data["captcha_code"].(string)
}
