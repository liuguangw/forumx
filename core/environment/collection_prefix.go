package environment

import "os"

//获取数据库名称配置, 读取 `FORUM_COLLECTION_PREFIX` 环境变量,
//若为空,则默认为空字符串
func CollectionPrefix() string {
	envKey := "FORUM_COLLECTION_PREFIX"
	return os.Getenv(envKey)
}
