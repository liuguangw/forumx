package tests

import (
	"github.com/liuguangw/forumx/cmd"
	"testing"
)

//TestAPI api单元测试入口
func TestAPI(t *testing.T) {
	app := cmd.SetupApp()
	t.Run("migration", testMigrate)
	t.Run("index.Hello", func(t *testing.T) {
		testIndexHello(app, t)
	})
}
