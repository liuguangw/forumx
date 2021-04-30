package models

import "time"

//UserSession 表示用户会话的记录
type UserSession struct {
	ID            string            `bson:"id"`            //用户会话ID
	UserID        int64             `bson:"user_id"`       //用户ID
	Authenticated bool              `bson:"authenticated"` //是否已通过身份验证
	Data          map[string]string `bson:"data"`          //会话数据
	CreatedAt     time.Time         `bson:"created_at"`    //创建时间
	UpdatedAt     time.Time         `bson:"updated_at"`    //最后更新时间
	ExpiredAt     time.Time         `bson:"expired_at"`    //过期时间
}

//Set 设置键值对
func (userSession *UserSession) Set(key, value string) {
	if userSession.Data == nil {
		userSession.Data = make(map[string]string)
	}
	userSession.Data[key] = value
}

//Get 使用key或者值,如果key不存在则返回默认值
func (userSession *UserSession) Get(key string, defaultValue ...string) string {
	var defaultReturnValue string
	if len(defaultValue) > 0 {
		defaultReturnValue = defaultValue[0]
	}
	if userSession.Data == nil {
		return defaultReturnValue
	}
	if value, ok := userSession.Data[key]; ok {
		return value
	}
	return defaultReturnValue
}

//Delete 删除key对应的值映射
func (userSession *UserSession) Delete(key string) {
	if userSession.Data != nil {
		delete(userSession.Data, key)
	}
}
