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
func testAuthRegister(app *fiber.App, sessionID string, t *testing.T) int64 {
	captchaCode := testCaptchaShow(app, sessionID, t)
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
	return userID
}

type loginAPIResponse = struct {
	common.AppResponse
	Data *struct {
		ID       int64  `json:"id"`       //用户ID
		Nickname string `json:"nickname"` //昵称
		Token    string `json:"token"`    //totp token
	} `json:"data"` //响应数据
}

func requestLoginAPI(app *fiber.App, sessionID string, loginRequest *request.LoginAccount, t *testing.T) *loginAPIResponse {
	requestData, err := json.Marshal(loginRequest)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		"POST",
		"/api/auth/login",
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
	appResponse := new(loginAPIResponse)
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	return appResponse
}

//testAuthLogin 测试登录
func testAuthLogin(app *fiber.App, sessionID string, t *testing.T) {
	captchaCode := testCaptchaShow(app, sessionID, t)
	//构造请求数据
	loginRequest := &request.LoginAccount{
		Username:    "liuguang1",
		Password:    "111111",
		CaptchaCode: captchaCode,
	}
	appResponse := requestLoginAPI(app, sessionID, loginRequest, t)
	//不存在此用户
	assert.Equal(t, common.ErrorUserNotFound, appResponse.Code, appResponse.Message)
	captchaCode = testCaptchaShow(app, sessionID, t)
	//构造请求数据
	loginRequest.Username = "liuguang"
	loginRequest.CaptchaCode = testCaptchaShow(app, sessionID, t)
	appResponse = requestLoginAPI(app, sessionID, loginRequest, t)
	//密码错误
	assert.Equal(t, common.ErrorPassword, appResponse.Code, appResponse.Message)
	//构造请求数据
	loginRequest.Password = "123456"
	loginRequest.CaptchaCode = testCaptchaShow(app, sessionID, t)
	appResponse = requestLoginAPI(app, sessionID, loginRequest, t)
	//登录成功
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	userID := responseData.ID
	assert.NotEmpty(t, userID)
	nickname := responseData.Nickname
	assert.NotEmpty(t, nickname)
}
