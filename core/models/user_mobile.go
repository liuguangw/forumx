package models

import "time"

//UserMobile 用户绑定的手机记录
type UserMobile struct {
	UserID    int64     `bson:"user_id"`    //用户ID
	Mobile    string    `bson:"mobile"`     //手机号
	CreatedAt time.Time `bson:"created_at"` //绑定时间
}
