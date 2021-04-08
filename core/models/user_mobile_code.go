package models

import "time"

const (
	_ = iota
	//绑定手机
	MobileCodeTypeBindAccount
	//重置密码
	MobileCodeTypeResetPassword
)

//用户绑定手机号、重置密码时发送的短信验证码记录
type UserMobileCode struct {
	Mobile    string    `bson:"mobile"`     //手机号
	CodeType  int       `bson:"code_type"`  //验证码类型
	Code      string    `bson:"code"`       //验证码
	UserId    int64     `bson:"user_id"`    //用户ID
	LinkUsed  bool      `bson:"link_used"`  //是否已使用
	CreatedAt time.Time `bson:"created_at"` //创建时间
	UpdatedAt time.Time `bson:"updated_at"` //最后更新时间
}
