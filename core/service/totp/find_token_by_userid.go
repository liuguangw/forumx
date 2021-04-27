package totp

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//FindTokenByUserID 根据用户ID查找用户绑定的令牌记录,如果数据不存在则返回nil
func FindTokenByUserID(ctx context.Context, userID int64) (*models.UserTotpKey, error) {
	coll, err := db.Collection(common.UserTotpKeyCollectionName)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	tokenInfo := new(models.UserTotpKey)
	if err := coll.FindOne(ctx, bson.M{"user_id": userID}).Decode(tokenInfo); err != nil {
		//不存在记录
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		//其它错误
		return nil, err
	}
	return tokenInfo, nil
}
