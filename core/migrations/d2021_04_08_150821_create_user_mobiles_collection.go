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

//CreateUserMobilesCollection 创建手机绑定记录集合
type CreateUserMobilesCollection struct {
}

func (*CreateUserMobilesCollection) collectionName() string {
	return db.CollectionFullName(common.UserMobileCollectionName)
}

//Name 迁移的名称
func (*CreateUserMobilesCollection) Name() string {
	return "d2021_04_08_150821_create_user_mobiles_collection"
}

//Up 执行迁移
func (c *CreateUserMobilesCollection) Up() error {
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
				"mobile": 1,
			},
			Options: options.Index().SetName("mobile_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"user_id": 1,
			},
			Options: options.Index().SetName("user_id_index"),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *CreateUserMobilesCollection) Down() error {
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
