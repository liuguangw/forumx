package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//创建users集合
type CreateUsersCollection struct {
}

func (*CreateUsersCollection) collectionName() string {
	return db.CollectionFullName("users")
}

//迁移的名称
func (*CreateUsersCollection) Name() string {
	return "d2021_04_06_172714_create_users_collection"
}

//执行迁移
func (c *CreateUsersCollection) Up() error {
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
				"id": 1,
			},
			Options: options.Index().SetName("id_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"username": 1,
			},
			Options: options.Index().SetName("username_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"coin_amount": -1,
			},
			Options: options.Index().SetName("coin_amount_index"),
		},
		{
			Keys: bson.M{
				"exp_amount": -1,
			},
			Options: options.Index().SetName("exp_amount_index"),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//回滚迁移
func (c *CreateUsersCollection) Down() error {
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
