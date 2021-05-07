package tests

import (
	"github.com/liuguangw/forumx/cmd"
	"testing"
)

//TestAPI api单元测试入口
func TestAPI(t *testing.T) {
	app := cmd.SetupApp()
	t.Run("migrate", testMigrate)
	t.Run("cache", testCache)
	t.Run("index.Hello", func(t *testing.T) {
		testIndexHello(t, app)
	})
	var (
		captchaID string
		sessionID string
	)
	t.Run("captcha.GenerateID", func(t *testing.T) {
		captchaID = testCaptchaID(t, app)
	})
	t.Run("captcha.Show", func(t *testing.T) {
		testCaptchaShow(t, app, captchaID)
	})
	t.Run("auth.Register", func(t *testing.T) {
		testAuthRegister(t, app, captchaID)
	})
	t.Run("auth.Login", func(t *testing.T) {
		sessionID = testAuthLogin(t, app, captchaID)
	})
	t.Run("auth.TotpRandomToken", func(t *testing.T) {
		testAuthTotpRandomToken(t, app, sessionID)
	})
	t.Run("auth.TotpBind", func(t *testing.T) {
		testAuthTotpBind(t, app, sessionID)
	})
	//初始化新的session ID
	t.Run("auth.TotpVerify", func(t *testing.T) {
		sessionID = testAuthTotpVerify(t, app, captchaID)
	})
	t.Run("mobile.SendCode", func(t *testing.T) {
		testMobileSendCode(t, app, captchaID, sessionID)
	})
	t.Run("mobile.BindAccount", func(t *testing.T) {
		testMobileBindAccount(t, app, sessionID)
	})
}
