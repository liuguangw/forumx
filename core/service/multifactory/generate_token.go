package multifactory

import (
	"context"
	"encoding/json"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/pquerna/otp/totp"
)

//GenerateToken 随机生成totp密钥,并将其暂时存入session中
func GenerateToken(ctx context.Context, userSession *models.UserSession) (*TotpTokenData, error) {
	optKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "tmp_issuer",
		AccountName: "tmp_user",
		SecretSize:  10,
	})
	if err != nil {
		return nil, err
	}
	tokenData := TotpTokenData{
		SecretKey:    optKey.Secret(),
		RecoveryCode: tools.GenerateRandomString(8),
	}
	//序列化json
	tokenDataBytes, err := json.Marshal(tokenData)
	if err != nil {
		return nil, err
	}
	tokenDataJSONString := string(tokenDataBytes)
	//设置数据
	userSession.Set(sessionKey, tokenDataJSONString)
	//save session
	if ctx == nil {
		ctx = context.Background()
	}
	if err := session.Save(ctx, userSession); err != nil {
		return nil, err
	}
	return &tokenData, nil
}
