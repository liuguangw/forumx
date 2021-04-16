package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//CreateUserSessionsCollection 创建用户会话集合
type CreateUserSessionsCollection struct {
}

func (*CreateUserSessionsCollection) collectionName() string {
	return db.CollectionFullName("user_sessions")
}

//Name 迁移的名称
func (*CreateUserSessionsCollection) Name() string {
	return "d2021_04_16_114203_create_user_sessions_collection"
}

//Up 执行迁移
func (c *CreateUserSessionsCollection) Up() error {
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
				"user_id": 1,
			},
			Options: options.Index().SetName("user_id_index"),
		},
		{
			Keys: bson.M{
				"expired_at": 1,
			},
			//ttl索引
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
func (c *CreateUserSessionsCollection) Down() error {
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
