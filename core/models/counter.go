package models

//Counter 系统计数器
type Counter struct {
	CounterKey   string `bson:"counter_key"`   //键名
	CounterValue int64  `bson:"counter_value"` //值
}

//计数器key定义
const (
	CounterKeyUserNextID       = "user.next_id"        //下一个用户ID的key
	CounterKeyForumNextID      = "forum.next_id"       //下一个论坛ID的key
	CounterKeyTopicNextID      = "topic.next_id"       //下一个帖子ID的key
	CounterKeyTopicReplyNextID = "topic_reply.next_id" //下一个帖子回复ID的key
)
