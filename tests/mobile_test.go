package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/request"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

//testMobileSendCode 测试获取绑定手机号的短信验证码
func testMobileSendCode(t *testing.T, app *fiber.App, captchaID, sessionID string) {
	captchaCode := testCaptchaShow(t, app, captchaID)
	smsRequest := &request.SendSms{
		CaptchaID:   captchaID,
		CaptchaCode: captchaCode,
		CodeType:    models.MobileCodeTypeBindAccount,
		Mobile:      "18566667777",
	}
	smsRequestData, err := json.Marshal(smsRequest)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		"POST",
		"/api/mobile/send-code",
		bytes.NewBuffer(smsRequestData),
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
	appResponse := new(common.AppResponse)
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
}
