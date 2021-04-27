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

//CreateUsersEmailsCollection 创建用户绑定的邮箱记录集合
type CreateUsersEmailsCollection struct {
}

func (*CreateUsersEmailsCollection) collectionName() string {
	return db.CollectionFullName(common.UserEmailCollectionName)
}

//Name 迁移的名称
func (*CreateUsersEmailsCollection) Name() string {
	return "d2021_04_08_145219_create_user_emails_collection"
}

//Up 执行迁移
func (c *CreateUsersEmailsCollection) Up() error {
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
				"email_address": 1,
			},
			Options: options.Index().SetName("email_address_uni").SetUnique(true),
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
func (c *CreateUsersEmailsCollection) Down() error {
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
