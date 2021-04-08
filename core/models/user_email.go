package models

import "time"

//用户绑定邮箱的记录
type UserEmail struct {
	UserId       int64     `bson:"user_id"`       //用户ID
	EmailAddress string    `bson:"email_address"` //邮箱地址
	CreatedAt    time.Time `bson:"created_at"`    //绑定时间
}
