package cache

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//PutItem 把数据写入缓存
func PutItem(itemKey string, itemValue interface{}, expiredAt time.Time) error {
	coll, err := db.Collection(collectionName)
	if err != nil {
		return err
	}
	cacheObject := &models.Cache{
		ItemKey:   itemKey,
		ItemValue: itemValue,
		ExpiredAt: expiredAt,
	}
	filter := bson.M{
		"item_key": itemKey,
	}
	opt := options.Update().SetUpsert(true)
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	if _, err := coll.UpdateOne(ctx, filter, bson.M{
		"$set": cacheObject,
	}, opt); err != nil {
		return err
	}
	return nil
}
