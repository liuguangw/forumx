package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

//testAuthRegister 测试注册
func testAuthRegister(app *fiber.App, sessionID, captchaCode string, t *testing.T) {
	//构造请求数据
	registerRequest := &request.RegisterAccount{
		Username:     "liuguang",
		Nickname:     "流光",
		EmailAddress: "admin@liuguang.vip",
		Password:     "123456",
		CaptchaCode:  captchaCode,
	}
	requestData, err := json.Marshal(registerRequest)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		"POST",
		"/api/auth/register",
		bytes.NewBuffer(requestData),
	)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+sessionID)
	req.Header.Set("Content-Type", "application/json")
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
			ID       int64  `json:"id"`       //用户ID
			Nickname string `json:"nickname"` //昵称
		} `json:"data"` //响应数据
	})
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	userID := responseData.ID
	assert.NotEmpty(t, userID)
	nickname := responseData.Nickname
	assert.Equal(t, registerRequest.Nickname, nickname)
}
