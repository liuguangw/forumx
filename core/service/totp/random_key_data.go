package totp

//RandomKeyData 临时随机生成的totp密钥信息
type RandomKeyData struct {
	SecretKey    string `json:"secret_key"`    //密钥
	RecoveryCode string `json:"recovery_code"` //恢复代码
}
