package tests

import (
	"github.com/liuguangw/forumx/app/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

//testMigrate 测试执行数据迁移
func testMigrate(t *testing.T) {
	args := []string{"forumx", "migrate", "refresh"}
	err := cmd.Execute(args)
	assert.NoError(t, err)
}
