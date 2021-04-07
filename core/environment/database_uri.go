package environment

import "os"

//获取数据库连接字符串配置, 读取 `FORUM_DB_URI` 环境变量,
//若为空,则默认为 `mongodb://localhost:27017`
func DatabaseUri() string {
	envKey := "FORUM_DB_URI"
	itemValue := os.Getenv(envKey)
	if itemValue != "" {
		return itemValue
	}
	return "mongodb://localhost:27017"
}
