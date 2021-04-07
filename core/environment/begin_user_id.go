package environment

import (
	"os"
	"strconv"
)

//除创始人外,注册时用户ID的开始值, 读取 `FORUM_BEGIN_USER_ID` 环境变量,
//若为空,则默认为创始人用户ID + 1
//
//安装后此值会写入计数器集合
func BeginUserId() int64 {
	envKey := "FORUM_BEGIN_USER_ID"
	itemValue := os.Getenv(envKey)
	if itemValue == "" {
		return FounderUserId() + 1
	}
	userId, err := strconv.ParseInt(itemValue, 10, 0)
	if err != nil {
		panic("invalid " + envKey + " value: " + err.Error())
	}
	return userId
}
