package migrations

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
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
			Description: "短信验证码驱动类型",
			ValueNumber: 1,
		},
		&models.AppConfig{
			ConfigKey:   "tencent.secret_id",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云API密钥 SecretId",
			ValueString: "111111111",
		},
		&models.AppConfig{
			ConfigKey:   "tencent.secret_key",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云API密钥 SecretKey",
			ValueString: "KeyKeyKey",
		},
		&models.AppConfig{
			ConfigKey:   "tencent.sms_sdk.app_id",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云短信应用SDK AppID",
			ValueString: "11223344",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.sign_text",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云短信签名",
			ValueString: "流光网",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.bind_mobile.template_id",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云短信模板(绑定手机号)ID",
			ValueString: "1111111",
		},
		&models.AppConfig{
			ConfigKey:   "sms.tencent.reset_password.template_id",
			ConfigType:  models.ConfigTypeString,
			Description: "腾讯云短信模板(重置密码验证码)ID",
			ValueString: "1111111",
		},
		&models.AppConfig{
			ConfigKey:   "mobile_code.lifetime.minutes",
			ConfigType:  models.ConfigTypeNumber,
			Description: "短信验证码多少分钟内有效",
			ValueNumber: 15,
		},
		&models.AppConfig{
			ConfigKey:   "mobile_code.store.days",
			ConfigType:  models.ConfigTypeNumber,
			Description: "短信验证码记录保存在数据库的天数",
			ValueNumber: 30,
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
		"mobile_code.lifetime.minutes", "mobile_code.store.days",
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
