package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//CreateUserMultiFactoryTokensCollection 创建用户两步验证token集合
type CreateUserMultiFactoryTokensCollection struct {
}

func (*CreateUserMultiFactoryTokensCollection) collectionName() string {
	return db.CollectionFullName("user_multi_factor_tokens")
}

//Name 迁移的名称
func (*CreateUserMultiFactoryTokensCollection) Name() string {
	return "d2021_04_21_163324_create_user_multi_factor_tokens_collection"
}

//Up 执行迁移
func (c *CreateUserMultiFactoryTokensCollection) Up() error {
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
				"user_id": 1,
			},
			Options: options.Index().SetName("user_id_uni").SetUnique(true),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *CreateUserMultiFactoryTokensCollection) Down() error {
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
