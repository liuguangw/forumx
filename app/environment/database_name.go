package environment

import "os"

//DatabaseName 获取数据库名称配置, 读取 `FORUM_DB_NAME` 环境变量,
//若为空,则默认为 `forumx`
func DatabaseName() string {
	envKey := "FORUM_DB_NAME"
	itemValue := os.Getenv(envKey)
	if itemValue != "" {
		return itemValue
	}
	return "forumx"
}
