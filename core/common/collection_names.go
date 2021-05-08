package common

//定义集合的名称
const (
	CounterCollectionName        = "counters"          //计数器集合
	UserCollectionName           = "users"             //用户集合
	UserEmailLinkCollectionName  = "user_email_links"  //发给用户的验证、重置密码邮件链接集合
	UserEmailCollectionName      = "user_emails"       //用户邮箱绑定记录集合
	UserMobileCodeCollectionName = "user_mobile_codes" //发给用户的短信集合
	UserMobileCollectionName     = "user_mobiles"      //用户手机号绑定记录集合
	UserSessionCollectionName    = "user_sessions"     //用户会话数据集合
	UserTotpKeyCollectionName    = "user_totp_keys"    //用户TOTP密钥绑定集合
	CacheCollectionName          = "caches"            //缓存集合
	AppConfigCollectionName      = "app_configs"       //应用配置集合
	ForumAreaCollectionName      = "forum_areas"       //论坛分区集合
	ForumCollectionName          = "forums"            //论坛集合
)
