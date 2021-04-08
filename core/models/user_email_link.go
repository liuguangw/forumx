package models

import "time"

const (
	_ = iota
	//EmailLinkTypeBindAccount 用于绑定邮箱链接
	EmailLinkTypeBindAccount
	//EmailLinkTypeResetPassword 用于重置密码链接
	EmailLinkTypeResetPassword
)

//UserEmailLink 用户绑定邮箱、重置密码时发送的邮件链接记录
type UserEmailLink struct {
	LinkType  int       `bson:"link_type"`  //链接类型
	Code      string    `bson:"code"`       //链接中的代码
	UserID    int64     `bson:"user_id"`    //用户ID
	Email     string    `bson:"email"`      //email地址
	LinkUsed  bool      `bson:"link_used"`  //是否已使用
	CreatedAt time.Time `bson:"created_at"` //创建时间
	UpdatedAt time.Time `bson:"updated_at"` //最后更新时间
}
