package totp

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/pquerna/otp/totp"
)

//GenerateRandomKeyData 随机生成totp密钥,并将其暂时存入session中，返回密钥和恢复代码
func GenerateRandomKeyData(ctx context.Context, userSession *models.UserSession) (string, string, error) {
	optKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "tmp_issuer",
		AccountName: "tmp_user",
		SecretSize:  10,
	})
	if err != nil {
		return "", "", err
	}
	secretKey := optKey.Secret()
	recoveryCode := tools.GenerateRandomString(8)
	userSession.Set(secretKeySessionKey, secretKey)
	userSession.Set(recoveryCodeSessionKey, recoveryCode)
	//save session
	if ctx == nil {
		ctx = context.Background()
	}
	if err := session.Save(ctx, userSession); err != nil {
		return "", "", err
	}
	return secretKey, recoveryCode, nil
}
