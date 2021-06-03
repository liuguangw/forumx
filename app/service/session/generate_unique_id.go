package session

import (
	"context"
	"github.com/liuguangw/forumx/app/service/tools"
)

//generateUniqueID 生成session ID, 并且确保此ID不存在于集合中
func generateUniqueID(ctx context.Context) (string, error) {
	var (
		sessionID      string
		sessionIDValid bool
	)
	for !sessionIDValid {
		sessionID = tools.GenerateHashID()
		tmpSessionLog, err := LoadByID(ctx, sessionID)
		if err != nil {
			return "", err
		}
		sessionIDValid = tmpSessionLog == nil
	}
	return sessionID, nil
}
