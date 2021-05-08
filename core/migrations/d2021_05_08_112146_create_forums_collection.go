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

//CreateForumsCollection 创建论坛集合
type CreateForumsCollection struct {
}

func (*CreateForumsCollection) collectionName() string {
	return db.CollectionFullName(common.ForumCollectionName)
}

//Name 迁移的名称
func (*CreateForumsCollection) Name() string {
	return "d2021_05_08_112146_create_forums_collection"
}

//Up 执行迁移
func (c *CreateForumsCollection) Up() error {
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
				"forum_area_id": 1,
			},
			Options: options.Index().SetName("forum_area_id_index"),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *CreateForumsCollection) Down() error {
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
