package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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

//testMobileBindAccount 测试绑定账号
func testMobileBindAccount(t *testing.T, app *fiber.App, sessionID string) {
	//查找最后一条短信验证码
	mobile := "18566667777"
	coll, err := db.Collection(common.UserMobileCodeCollectionName)
	assert.NoError(t, err)
	codeLog := new(models.UserMobileCode)
	ctx := context.Background()
	err = coll.FindOne(ctx, bson.M{"mobile": mobile}).Decode(codeLog)
	assert.NoError(t, err)
	bindRequest := &request.BindAccount{
		Mobile: mobile,
		Code:   codeLog.Code,
	}
	requestData, err := json.Marshal(bindRequest)
	assert.NoError(t, err)
	req, err := http.NewRequest(
		"POST",
		"/api/mobile/bind-account",
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
	appResponse := new(common.AppResponse)
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
}
