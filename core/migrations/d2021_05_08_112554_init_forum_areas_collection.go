package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//InitForumAreasCollection 初始化论坛分区集合
type InitForumAreasCollection struct {
}

func (*InitForumAreasCollection) collection() (*mongo.Collection, error) {
	coll, err := db.Collection(common.ForumAreaCollectionName)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

//Name 迁移的名称
func (*InitForumAreasCollection) Name() string {
	return "d2021_05_08_112554_init_forum_areas_collection"
}

//Up 执行迁移
func (c *InitForumAreasCollection) Up() error {
	coll, err := c.collection()
	if err != nil {
		return err
	}
	timeNow := time.Now()
	itemInfo := &models.ForumArea{
		ID:          1,
		Name:        "默认分区",
		Description: "安装时自动创建的分区",
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}
	if _, err := coll.InsertOne(context.TODO(), itemInfo); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *InitForumAreasCollection) Down() error {
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
