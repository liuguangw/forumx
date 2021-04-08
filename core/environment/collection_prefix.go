package environment

import "os"

//CollectionPrefix 获取集合前缀, 读取 `FORUM_COLLECTION_PREFIX` 环境变量,
//若为空,则默认为空字符串
func CollectionPrefix() string {
	envKey := "FORUM_COLLECTION_PREFIX"
	return os.Getenv(envKey)
}
