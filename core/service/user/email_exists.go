package user

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
)

//EmailExists 判断指定的邮箱是否已存在绑定记录
func EmailExists(ctx context.Context, emailAddress string) (bool, error) {
	coll, err := db.Collection(userEmailCollectionName)
	if err != nil {
		return false, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	count, err := coll.CountDocuments(ctx, bson.M{"email_address": emailAddress})
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
