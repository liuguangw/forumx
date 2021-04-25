package cache

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//GetItem 从缓存中加载数据,如果缓存存在则为true
func GetItem(ctx context.Context, itemKey string, cachedItem interface{}) (bool, error) {
	coll, err := db.Collection(collectionName)
	if err != nil {
		return false, err
	}
	filter := bson.M{
		"item_key": itemKey,
	}
	if ctx == nil {
		ctx = context.Background()
	}
	findResult := coll.FindOne(ctx, filter)
	if err := findResult.Decode(cachedItem); err != nil {
		//不存在的key
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		//其它错误
		return false, err
	}
	//判断过期时间
	tmpItem := new(models.Cache)
	if err := findResult.Decode(tmpItem); err != nil {
		return false, err
	}
	//缓存已经过期了
	if tmpItem.ExpiredAt.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}
