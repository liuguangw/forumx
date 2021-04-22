package models

import "time"

//Cache 系统的缓存
type Cache struct {
	ItemKey   string      `bson:"item_key"`   //键
	ItemValue interface{} `bson:"item_value"` //缓存值
	ExpiredAt time.Time   `bson:"expired_at"` //过期时间
}
