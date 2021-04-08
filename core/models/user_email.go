package models

import "time"

//UserEmail 用户绑定的邮箱记录
type UserEmail struct {
	UserID       int64     `bson:"user_id"`       //用户ID
	EmailAddress string    `bson:"email_address"` //邮箱地址
	CreatedAt    time.Time `bson:"created_at"`    //绑定时间
}
