package tests

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

//testSessionInitNew 测试初始化会话
func testSessionInitNew(app *fiber.App, t *testing.T) string {
	var sessionID string
	req, err := http.NewRequest(
		"POST",
		"/api/session/init",
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
			ID        string `json:"id"`         //session ID
			ExpiredAt string `json:"expired_at"` //过期时间
		} `json:"data"` //响应数据
	})
	err = json.Unmarshal(body, appResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, appResponse.Code, appResponse.Message)
	responseData := appResponse.Data
	assert.NotNil(t, responseData)
	sessionID = responseData.ID
	assert.NotEmpty(t, sessionID)
	expiredAt := responseData.ExpiredAt
	assert.NotEmpty(t, expiredAt)
	return sessionID
}
