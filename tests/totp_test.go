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

//testRandomToken 测试获取随机的totp密钥
func testRandomToken(app *fiber.App, sessionID string, t *testing.T) {
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

//testTotpBind 测试绑定两步验证令牌
func testTotpBind(app *fiber.App, sessionID string, t *testing.T) {
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.LoadByID(ctx, sessionID)
	assert.NoError(t, err)
	//读取令牌信息
	tokenData, err := totp.LoadKeyDataFromSession(userSession)
	assert.NoError(t, err)
	code, err := totp2.GenerateCode(tokenData.SecretKey, time.Now())
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

//testAuth2FALogin 测试两步验证登录
func testAuth2FALogin(app *fiber.App, sessionID string, userID int64, t *testing.T) {
	captchaCode := testCaptchaShow(app, sessionID, t)
	//构造请求数据
	loginRequest := &request.LoginAccount{
		Username:    "liuguang",
		Password:    "123456",
		CaptchaCode: captchaCode,
	}
	appResponse := requestLoginAPI(app, sessionID, loginRequest, t)
	//需要身份验证
	assert.Equal(t, common.ErrorNeedAuthentication, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	totpToken := responseData.Token
	assert.NotEmpty(t, totpToken)
	//判断session 状态
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.LoadByID(ctx, sessionID)
	assert.NoError(t, err)
	assert.Empty(t, userSession.UserID)
	//根据userID获取密钥
	totpKeyData, err := totp.FindTotpKeyByUserID(ctx, userID)
	assert.NoError(t, err)
	assert.NotNil(t, totpKeyData)
	secretKey := totpKeyData.SecretKey
	code, err := totp2.GenerateCode(secretKey, time.Now())
	assert.NoError(t, err)
	verifyRequest := map[string]string{
		"token": totpToken,
		"code":  code,
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
	//判断session 状态
	userSession, err = session.LoadByID(ctx, sessionID)
	assert.NoError(t, err)
	userID = userSession.UserID
	assert.NotEqual(t, 0, userID)
}
