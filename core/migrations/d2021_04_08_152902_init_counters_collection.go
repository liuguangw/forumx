package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/environment"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//InitCountersCollection 初始化计数器
type InitCountersCollection struct {
}

func (*InitCountersCollection) collection() (*mongo.Collection, error) {
	counterColl, err := db.Collection("counters")
	if err != nil {
		return nil, err
	}
	return counterColl, nil
}

//Name 迁移的名称
func (*InitCountersCollection) Name() string {
	return "d2021_04_08_152902_init_counters_collection"
}

//Up 执行迁移
func (c *InitCountersCollection) Up() error {
	counterColl, err := c.collection()
	if err != nil {
		return err
	}
	itemList := []interface{}{
		&models.Counter{
			CounterKey:   "user.next_id",
			CounterValue: environment.FounderUserID(),
		},
		&models.Counter{
			CounterKey:   "forum.next_id",
			CounterValue: 1,
		},
		&models.Counter{
			CounterKey:   "topic.next_id",
			CounterValue: 1,
		},
		&models.Counter{
			CounterKey:   "topic_reply.next_id",
			CounterValue: 1,
		},
	}
	if _, err := counterColl.InsertMany(context.TODO(), itemList); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *InitCountersCollection) Down() error {
	counterColl, err := c.collection()
	if err != nil {
		return err
	}
	keyList := []string{
		"user.next_id", "forum.next_id",
		"topic.next_id", "topic_reply.next_id",
	}
	filter := bson.M{
		"counter_key": bson.M{
			"$in": keyList,
		},
	}
	if _, err := counterColl.DeleteMany(context.TODO(), filter); err != nil {
		return err
	}
	return nil
}
