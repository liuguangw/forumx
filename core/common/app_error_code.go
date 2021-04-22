package common

//40系列(用户账号相关)
const (
	ErrorNotLogin                    = 4001 + iota //未登录
	ErrorPassword                                  //密码错误
	ErrorAccountLocked                             //账号已锁定
	ErrorUserNotFound                              //用户不存在
	ErrorUsernameExists                            //用户名已存在
	ErrorEmailExists                               //邮箱已存在
	ErrorNeedAuthentication                        //需要进行身份验证
	ErrorTwoFactorAuthenticationCode               //通过两步验证动态码验证身份失败
	ErrorEmailAuthenticationCode                   //通过邮箱验证码验证身份失败
	ErrorMobileAuthenticationCode                  //通过短信验证码验证身份失败
)

//50系列
const (
	ErrorInternalServer    = 5000 + iota //内部错误
	ErrorBadRequest                      //输入无效(表单或者json解析失败)
	ErrorInputFieldIsEmpty               //所需字段为空
	ErrorInputFieldInvalid               //输入的字段值无效
	ErrorSessionExpired                  //sessionID无效或者已经过期
)
