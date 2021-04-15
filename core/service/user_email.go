package service

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//userEmailExists 判断指定的邮箱是否已存在绑定记录
func userEmailExists(emailAddress string) (bool, error) {
	coll, err := db.Collection(userEmailCollectionName)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	count, err := coll.CountDocuments(ctx, bson.M{"email_address": emailAddress})
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
