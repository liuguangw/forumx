package tests

import (
	"github.com/liuguangw/forumx/cmd"
	"testing"
)

//TestAPI api单元测试入口
func TestAPI(t *testing.T) {
	app := cmd.SetupApp()
	t.Run("migrate", testMigrate)
	t.Run("index.Hello", func(t *testing.T) {
		testIndexHello(app, t)
	})
	var (
		sessionID   string
		captchaCode string
	)
	t.Run("session.InitNew", func(t *testing.T) {
		sessionID = testSessionInitNew(app, t)
	})
	t.Run("captcha.Show", func(t *testing.T) {
		captchaCode = testCaptchaShow(app, sessionID, t)
	})
	t.Run("auth.Register", func(t *testing.T) {
		testAuthRegister(app, sessionID, captchaCode, t)
	})
}
