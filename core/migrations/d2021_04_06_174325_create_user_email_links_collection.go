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

//CreateUserEmailLinksCollection 创建绑定邮箱、重置密码的邮件链接记录集合
type CreateUserEmailLinksCollection struct {
}

func (*CreateUserEmailLinksCollection) collectionName() string {
	return db.CollectionFullName(common.UserEmailLinkCollectionName)
}

//Name 迁移的名称
func (*CreateUserEmailLinksCollection) Name() string {
	return "d2021_04_06_174325_create_user_email_links_collection"
}

//Up 执行迁移
func (c *CreateUserEmailLinksCollection) Up() error {
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
			Keys: bson.D{
				{Key: "link_type", Value: 1},
				{Key: "code", Value: 1},
			},
			Options: options.Index().SetName("link_code_uni").SetUnique(true),
		},
		{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetName("email_index"),
		},
		{
			Keys: bson.M{
				"created_at": 1,
			},
			//15分钟后自动失效
			Options: options.Index().SetName("created_at_ttl").SetExpireAfterSeconds(900),
		},
	}
	indexOpts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	if _, err := indexView.CreateMany(context.TODO(), indexModels, indexOpts); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *CreateUserEmailLinksCollection) Down() error {
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
