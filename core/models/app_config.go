package models

//配置的类型
const (
	_                = iota
	ConfigTypeString //字符串配置
	ConfigTypeNumber //数值配置
)

//AppConfig 应用配置
type AppConfig struct {
	ConfigKey   string `bson:"config_key"`             //键名
	ConfigType  int    `bson:"config_type"`            //类型
	ValueString string `bson:"value_string,omitempty"` //字符串值
	ValueNumber int64  `bson:"value_number,omitempty"` //数值
}
