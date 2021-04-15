package common

//定义应用的错误码
const (
	//40系列(用户账号相关)

	ErrorNotLogin                    = 4001 //未登录
	ErrorPassword                    = 4002 //密码错误
	ErrorAccountLocked               = 4003 //账号已锁定
	ErrorUserNotFound                = 4004 //用户不存在
	ErrorUsernameExists              = 4005 //用户名已存在
	ErrorEmailExists                 = 4006 //邮箱已存在
	ErrorNeedAuthentication          = 4007 //需要进行身份验证
	ErrorTwoFactorAuthenticationCode = 4008 //通过两步验证动态码验证身份失败
	ErrorEmailAuthenticationCode     = 4009 //通过邮箱验证码验证身份失败
	ErrorMobileAuthenticationCode    = 4010 //通过短信验证码验证身份失败

	//50系列

	ErrorInternalServer    = 5000 //内部错误
	ErrorInputFieldIsEmpty = 5001 //所需字段为空
	ErrorInputFieldInvalid = 5001 //输入的字段值无效
)
