package cache

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//DeleteItem 根据键名删除缓存
func DeleteItem(itemKey string) error {
	coll, err := db.Collection(collectionName)
	if err != nil {
		return err
	}
	filter := bson.M{
		"item_key": itemKey,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
