package multifactory

import (
	"encoding/json"
	"github.com/liuguangw/forumx/core/models"
)

//LoadTokenFromSession 从session中加载临时生成的两步验证令牌信息,此函数会在数据不存在时返回nil
func LoadTokenFromSession(userSession *models.UserSession) (*TotpTokenData, error) {
	tokenDataJSONString := userSession.Get(sessionKey)
	if tokenDataJSONString == "" {
		return nil, nil
	}
	tokenData := new(TotpTokenData)
	//JSON 解码
	if err := json.Unmarshal([]byte(tokenDataJSONString), tokenData); err != nil {
		return nil, err
	}
	return tokenData, nil
}
