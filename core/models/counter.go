package models

//计数器
type Counter struct {
	CounterKey   string `json:"counter_key"`   //键名
	CounterValue int64  `json:"counter_value"` //值
}
