package migrations

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//InitForumsCollection 初始化论坛集合
type InitForumsCollection struct {
}

func (*InitForumsCollection) collection() (*mongo.Collection, error) {
	coll, err := db.Collection(common.ForumCollectionName)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

//Name 迁移的名称
func (*InitForumsCollection) Name() string {
	return "d2021_05_08_114039_init_forums_collection"
}

//Up 执行迁移
func (c *InitForumsCollection) Up() error {
	coll, err := c.collection()
	if err != nil {
		return err
	}
	timeNow := time.Now()
	itemInfo := &models.Forum{
		ID:          1,
		ForumAreaID: 1,
		Name:        "默认论坛",
		Description: "安装时自动创建的论坛",
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}
	if _, err := coll.InsertOne(context.TODO(), itemInfo); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *InitForumsCollection) Down() error {
	coll, err := c.collection()
	if err != nil {
		return err
	}
	filter := bson.M{}
	if _, err := coll.DeleteMany(context.TODO(), filter); err != nil {
		return err
	}
	return nil
}
