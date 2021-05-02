package config

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//LoadConfig 读取配置
func LoadConfig(ctx context.Context, configKey string) (*models.AppConfig, error) {
	coll, err := db.Collection(common.AppConfigCollectionName)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	itemInfo := new(models.AppConfig)
	if err := coll.FindOne(ctx, bson.M{"config_key": configKey}).Decode(itemInfo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return itemInfo, nil
}
