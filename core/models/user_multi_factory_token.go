package models

import "time"

//UserMultiFactoryToken 用户两步验证密钥
type UserMultiFactoryToken struct {
	UserID    int64     `bson:"user_id"`    //用户ID
	SecretKey string    `bson:"secret_key"` //密钥
	CreatedAt time.Time `bson:"created_at"` //创建时间
	UpdatedAt time.Time `bson:"updated_at"` //最后更新时间
}
