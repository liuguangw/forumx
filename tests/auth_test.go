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
func testAuthRegister(t *testing.T, app *fiber.App, captchaID string) {
	captchaCode := testCaptchaShow(t, app, captchaID)
	//构造请求数据
	registerRequest := &request.RegisterAccount{
		Username:     "liuguang",
		Nickname:     "流光",
		EmailAddress: "admin@liuguang.vip",
		Password:     "123456",
		CaptchaID:    captchaID,
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

type loginAPIResponse = struct {
	common.AppResponse
	Data *struct {
		ID        int64  `json:"id"`         //用户ID
		SessionID string `json:"session_id"` //本次登录的会话ID
		ExpiresIn int64  `json:"expires_in"` //会话的有效期,单位秒
	} `json:"data"` //响应数据
}

func requestLoginAPI(t *testing.T, app *fiber.App, loginRequest *request.LoginAccount) *loginAPIResponse {
	requestData, err := json.Marshal(loginRequest)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		"POST",
		"/api/auth/login",
		bytes.NewBuffer(requestData),
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	appResponse := new(loginAPIResponse)
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	return appResponse
}

//testAuthLogin 测试登录
func testAuthLogin(t *testing.T, app *fiber.App, captchaID string) string {
	captchaCode := testCaptchaShow(t, app, captchaID)
	//构造请求数据
	loginRequest := &request.LoginAccount{
		Username:    "liuguang1",
		Password:    "111111",
		CaptchaID:   captchaID,
		CaptchaCode: captchaCode,
	}
	appResponse := requestLoginAPI(t, app, loginRequest)
	//不存在此用户
	assert.Equal(t, common.ErrorUserNotFound, appResponse.Code, appResponse.Message)
	//构造请求数据
	loginRequest.Username = "liuguang"
	loginRequest.CaptchaCode = testCaptchaShow(t, app, captchaID)
	appResponse = requestLoginAPI(t, app, loginRequest)
	//密码错误
	assert.Equal(t, common.ErrorPassword, appResponse.Code, appResponse.Message)
	//构造请求数据
	loginRequest.Password = "123456"
	loginRequest.CaptchaCode = testCaptchaShow(t, app, captchaID)
	appResponse = requestLoginAPI(t, app, loginRequest)
	//登录成功
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	assert.NotNil(t, responseData.ID)
	assert.NotNil(t, responseData.SessionID)
	assert.NotNil(t, responseData.ExpiresIn)
	return responseData.SessionID
}
