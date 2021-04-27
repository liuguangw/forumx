package totp

import (
	"encoding/json"
	"github.com/liuguangw/forumx/core/models"
)

//LoadKeyDataFromSession 从session中加载临时生成的两步验证令牌信息,此函数会在数据不存在时返回nil
func LoadKeyDataFromSession(userSession *models.UserSession) (*RandomKeyData, error) {
	keyDataJSONString := userSession.Get(sessionKey)
	if keyDataJSONString == "" {
		return nil, nil
	}
	keyData := new(RandomKeyData)
	//JSON 解码
	if err := json.Unmarshal([]byte(keyDataJSONString), keyData); err != nil {
		return nil, err
	}
	return keyData, nil
}
