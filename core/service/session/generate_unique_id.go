package session

import (
	"context"
	"github.com/liuguangw/forumx/core/service/tools"
	"time"
)

//generateUniqueID 生成session ID, 并且确保此ID不存在于集合中
func generateUniqueID(ctx context.Context) (string, error) {
	var (
		sessionID      string
		sessionIDValid bool
	)
	for !sessionIDValid {
		sessionID = generateID()
		tmpSessionLog, err := LoadByID(ctx, sessionID)
		if err != nil {
			return "", err
		}
		sessionIDValid = tmpSessionLog == nil
	}
	return sessionID, nil
}

//generateID 随机生成session ID
func generateID() string {
	plainText := time.Now().Format(time.RFC3339Nano) + " / - / " + tools.GenerateRandomString(30)
	return tools.Md5String(plainText)
}
