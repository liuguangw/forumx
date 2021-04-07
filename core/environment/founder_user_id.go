package environment

import (
	"os"
	"strconv"
)

//安装时创始人的用户ID, 读取 `FORUM_FOUNDER_USER_ID` 环境变量,
//若为空,则默认为1
func FounderUserId() int64 {
	envKey := "FORUM_FOUNDER_USER_ID"
	itemValue := os.Getenv(envKey)
	if itemValue == "" {
		return 1
	}
	userId, err := strconv.ParseInt(itemValue, 10, 0)
	if err != nil {
		panic("invalid " + envKey + " value: " + err.Error())
	}
	return userId
}
