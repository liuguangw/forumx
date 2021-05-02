package migrations

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//InitAppConfigsCollection 初始化应用配置集合
type InitAppConfigsCollection struct {
}

func (*InitAppConfigsCollection) collection() (*mongo.Collection, error) {
	coll, err := db.Collection(common.AppConfigCollectionName)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

//Name 迁移的名称
func (*InitAppConfigsCollection) Name() string {
	return "d2021_05_02_175826_init_app_configs_collection"
}

//Up 执行迁移
func (c *InitAppConfigsCollection) Up() error {
	coll, err := c.collection()
	if err != nil {
		return err
	}
	itemList := []interface{}{
		&models.AppConfig{
			ConfigKey:   "sms.driver.type",
			ConfigType:  models.ConfigTypeNumber,
			ValueNumber: 1,
		},
		&models.AppConfig{
			ConfigKey:   "tencent.secret_id",
			ConfigType:  models.ConfigTypeString,
			ValueString: "111111111",
		},
		&models.AppConfig{
			ConfigKey:   "tencent.secret_key",
			ConfigType:  models.ConfigTypeString,
			ValueString: "KeyKeyKey",
		},
		&models.AppConfig{
			ConfigKey:   "tencent.sms_sdk.app_id",
			ConfigType:  models.ConfigTypeString,
			ValueString: "11223344",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.sign_text",
			ConfigType:  models.ConfigTypeString,
			ValueString: "流光网",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.bind_mobile.template_id",
			ConfigType:  models.ConfigTypeString,
			ValueString: "1111111",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.reset_password.template_id",
			ConfigType:  models.ConfigTypeString,
			ValueString: "1111111",
		},
	}
	if _, err := coll.InsertMany(context.TODO(), itemList); err != nil {
		return err
	}
	return nil
}

//Down 回滚迁移
func (c *InitAppConfigsCollection) Down() error {
	coll, err := c.collection()
	if err != nil {
		return err
	}
	keyList := []string{
		"sms.tencent.app_id", "sms.tencent.app_key", "sms.tencent.sign_text",
		"sms.tencent.bind_mobile.template_id", "sms.tencent.reset_password.template_id",
	}
	filter := bson.M{
		"config_key": bson.M{
			"$in": keyList,
		},
	}
	if _, err := coll.DeleteMany(context.TODO(), filter); err != nil {
		return err
	}
	return nil
}
