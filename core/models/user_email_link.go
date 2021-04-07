package models

import "time"

const (
	_ = iota
	//绑定邮箱链接
	UserEmailLinkTypeBindAccount
	//重置密码链接
	UserEmailLinkTypeResetPassword
)

//用户绑定邮箱、重置密码时发送的邮件链接记录
type UserEmailLink struct {
	LinkType  int       `bson:"link_type"`  //链接类型
	Code      string    `bson:"code"`       //链接中的代码
	UserId    int64     `bson:"user_id"`    //用户ID
	CreatedAt time.Time `bson:"created_at"` //创建时间
}
