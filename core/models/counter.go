package models

//计数器
type Counter struct {
	CounterKey   string `bson:"counter_key"`   //键名
	CounterValue int64  `bson:"counter_value"` //值
}
