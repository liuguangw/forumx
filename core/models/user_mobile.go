package models

import "time"

//用户绑定手机的记录
type UserMobile struct {
	UserId    int64     `bson:"user_id"`    //用户ID
	Mobile    string    `bson:"mobile"`     //手机号
	CreatedAt time.Time `bson:"created_at"` //绑定时间
}
