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

//CreateUserMobileCodesCollection 创建短信验证码集合
type CreateUserMobileCodesCollection struct {
}

func (*CreateUserMobileCodesCollection) collectionName() string {
	return db.CollectionFullName(common.UserMobileCodeCollectionName)
}

//Name 迁移的名称
func (*CreateUserMobileCodesCollection) Name() string {
	return "d2021_04_08_145948_create_user_mobile_codes_collection"
}

//Up 执行迁移
func (c *CreateUserMobileCodesCollection) Up() error {
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
				{Key: "mobile", Value: 1},
				{Key: "code_type", Value: 1},
				{Key: "code", Value: 1},
			},
			Options: options.Index().SetName("mobile_code_index"),
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
func (c *CreateUserMobileCodesCollection) Down() error {
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
