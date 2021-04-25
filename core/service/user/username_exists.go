package user

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
)

//usernameExists 判断指定的用户名是否已经注册
func usernameExists(ctx context.Context, username string) (bool, error) {
	coll, err := db.Collection(collectionName)
	if err != nil {
		return false, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	count, err := coll.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
