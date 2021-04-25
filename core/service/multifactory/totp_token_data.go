package multifactory

//TotpTokenData 临时随机随机生成的totp密钥信息
type TotpTokenData struct {
	SecretKey    string `json:"secret_key"`    //密钥
	RecoveryCode string `json:"recovery_code"` //恢复代码
}
