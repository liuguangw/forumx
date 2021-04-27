package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//CreateCachesCollection 创建缓存集合
type CreateCachesCollection struct {
}

func (*CreateCachesCollection) collectionName() string {
	return db.CollectionFullName(common.CacheCollectionName)
}

//Name 迁移的名称
func (*CreateCachesCollection) Name() string {
	return "d2021_04_22_144239_create_caches_collection"
}

//Up 执行迁移
func (c *CreateCachesCollection) Up() error {
	database, err := db.Database()
	if err != nil {
		return err
	}
	//创建集合
	collectionFullName := c.collectionName()
	opts := options.CreateCollection()
	if err := database.CreateCollection(context.TODO(), collectionFullName, opts); err != nil {
		return err
	}
	//创建索引
	coll := database.Collection(collectionFullName)
	indexView := coll.Indexes()
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{
				"item_key": 1,
			},
			Options: options.Index().SetName("item_key_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"expired_at": 1,
			},
			Options: options.Index().SetName("expired_at_ttl").SetExpireAfterSeconds(1),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *CreateCachesCollection) Down() error {
	database, err := db.Database()
	if err != nil {
		return err
	}
	coll := database.Collection(c.collectionName())
	if err := coll.Drop(context.TODO()); err != nil {
		return err
	}
	return nil
}
