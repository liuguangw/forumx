package user

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//FindUserByUsername 使用用户名查找用户信息
func FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	coll, err := db.Collection(common.UserCollectionName)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	userInfo := new(models.User)
	if err := coll.FindOne(ctx, bson.M{"username": username}).Decode(userInfo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return userInfo, nil
}
