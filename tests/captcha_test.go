package tests

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/service/captcha"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

//testCaptchaID 测试获取验证码ID
func testCaptchaID(t *testing.T, app *fiber.App) string {
	req, err := http.NewRequest(
		"POST",
		"/api/captcha/id",
		nil,
	)
	assert.NoError(t, err)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	//
	appResponse := new(struct {
		common.AppResponse
		Data *struct {
			ID string `json:"id"` //验证码ID
		} `json:"data"` //响应数据
	})
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	assert.NotNil(t, responseData.ID)
	return responseData.ID
}

//testCaptchaShow 测试验证码接口
func testCaptchaShow(t *testing.T, app *fiber.App, captchaID string) string {
	req, err := http.NewRequest(
		"GET",
		"/api/captcha/show?id="+captchaID,
		nil,
	)
	assert.NoError(t, err)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	//获取图形验证码
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	captchaExists, captchaCode, err := captcha.LoadCaptchaCode(ctx, captchaID)
	assert.NoError(t, err)
	assert.True(t, captchaExists)
	assert.NotEmpty(t, captchaCode)
	return captchaCode
}
