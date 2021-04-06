package db

//环境变量定义
const (
	//数据库连接字符串环境变量key
	//
	//可用值例如 mongodb://localhost:27017
	dbUriEnvKey = "FORUM_DB_URI"

	//数据库名称环境变量key
	//
	//可用值例如 forumx
	dbNameEnvKey = "FORUM_DB_NAME"

	//集合名称前缀
	//
	//可用值例如 pre_
	collectionPrefixEnvKey = "FORUM_COLLECTION_PREFIX"
)
