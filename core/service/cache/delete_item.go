package cache

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteItem 根据键名删除缓存
func DeleteItem(ctx context.Context, itemKey string) error {
	coll, err := db.Collection(collectionName)
	if err != nil {
		return err
	}
	filter := bson.M{
		"item_key": itemKey,
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
