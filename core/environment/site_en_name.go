package environment

import "os"

//SiteEnName 获取站点英文名称, 读取 `SITE_EN_NAME` 环境变量,
//若为空,则默认为 `forumx`
func SiteEnName() string {
	envKey := "SITE_EN_NAME"
	itemValue := os.Getenv(envKey)
	if itemValue != "" {
		return itemValue
	}
	return "forumx"
}
