package models

import "time"

//User 表示注册的用户信息
type User struct {
	ID             int64     `bson:"id"`              //用户ID
	Username       string    `bson:"username"`        //用户名
	Nickname       string    `bson:"nickname"`        //昵称
	AvatarURL      string    `bson:"avatar_url"`      //头像
	CoinAmount     int64     `bson:"coin_amount"`     //金币余额
	ExpAmount      int64     `bson:"exp_amount"`      //经验值
	EmailVerified  bool      `bson:"email_verified"`  //email是否已验证
	MobileVerified bool      `bson:"mobile_verified"` //手机号是否已验证
	Enable2FA      bool      `bson:"enable_2fa"`      //是否启用了双重验证
	Password       string    `bson:"password"`        //密码
	Salt           string    `bson:"salt"`            //密码salt
	RegisterIP     string    `bson:"register_ip"`     //注册时的IP地址
	CreatedAt      time.Time `bson:"created_at"`      //注册时间
	UpdatedAt      time.Time `bson:"updated_at"`      //最后更新时间
}
