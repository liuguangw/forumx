package models

import "time"

//UserSession 表示用户会话的记录
type UserSession struct {
	ID        string                 `bson:"id"`         //用户会话ID
	UserID    int64                  `bson:"user_id"`    //用户ID
	Authed    bool                   `bson:"authed"`     //是否通过了身份验证
	Data      map[string]interface{} `bson:"data"`       //会话数据
	CreatedAt time.Time              `bson:"created_at"` //创建时间
	UpdatedAt time.Time              `bson:"updated_at"` //最后更新时间
	ExpiredAt time.Time              `bson:"expired_at"` //过期时间
}
