package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/totp"
	totp2 "github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

//testAuthTotpRandomToken 测试获取随机的totp密钥
func testAuthTotpRandomToken(t *testing.T, app *fiber.App, sessionID string) {
	req, err := http.NewRequest(
		"GET",
		"/api/auth/totp/random-token",
		bytes.NewBuffer([]byte{}),
	)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+sessionID)
	//req.Header.Set("Content-Type", "application/json")
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
			URL          string `json:"url"`           //totp密钥URL
			RecoveryCode string `json:"recovery_code"` //恢复代码
		} `json:"data"` //响应数据
	})
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	assert.NotEmpty(t, responseData.URL)
	assert.NotEmpty(t, responseData.RecoveryCode)
}

//testAuthTotpBind 测试绑定两步验证令牌
func testAuthTotpBind(t *testing.T, app *fiber.App, sessionID string) {
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//session状态判断
	userSession, err := session.LoadByID(ctx, sessionID)
	assert.NoError(t, err)
	assert.True(t, userSession.Authenticated)
	//读取令牌信息
	secretKey, recoveryCode := totp.LoadKeyDataFromSession(userSession)
	assert.NotEmpty(t, secretKey)
	assert.NotEmpty(t, recoveryCode)
	//计算动态验证码
	code, err := totp2.GenerateCode(secretKey, time.Now())
	assert.NoError(t, err)
	requestData, err := json.Marshal(map[string]string{
		"code": code,
	})
	assert.NoError(t, err)
	//请求绑定接口
	req, err := http.NewRequest(
		"POST",
		"/api/auth/totp/bind",
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
	appResponse := new(common.AppResponse)
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
}

//testAuthTotpVerify 测试两步验证登录
func testAuthTotpVerify(t *testing.T, app *fiber.App, captchaID string) string {
	captchaCode := testCaptchaShow(t, app, captchaID)
	//构造请求数据
	loginRequest := &request.LoginAccount{
		Username:    "liuguang",
		Password:    "123456",
		CaptchaID:   captchaID,
		CaptchaCode: captchaCode,
	}
	appResponse := requestLoginAPI(t, app, loginRequest)
	//需要身份验证
	assert.Equal(t, common.ErrorNeedAuthentication, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	assert.NotNil(t, responseData.ID)
	assert.NotNil(t, responseData.SessionID)
	assert.NotNil(t, responseData.ExpiresIn)
	sessionID := responseData.SessionID
	userID := responseData.ID
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//session状态判断
	userSession, err := session.LoadByID(ctx, sessionID)
	assert.NoError(t, err)
	assert.False(t, userSession.Authenticated)
	//根据userID获取密钥
	totpKeyData, err := totp.FindTotpKeyByUserID(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, totpKeyData)
	secretKey := totpKeyData.SecretKey
	code, err := totp2.GenerateCode(secretKey, time.Now())
	assert.NoError(t, err)
	verifyRequest := map[string]string{
		"code": code,
	}
	requestData, err := json.Marshal(verifyRequest)
	assert.NoError(t, err)
	//请求验证接口
	req, err := http.NewRequest(
		"POST",
		"/api/auth/totp/verify",
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
	verifyResponse := new(common.AppResponse)
	err = json.Unmarshal(body, verifyResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, verifyResponse.Code, verifyResponse.Message)
	return sessionID
}
