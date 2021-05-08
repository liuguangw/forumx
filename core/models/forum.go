package models

import "time"

//Forum 论坛信息
type Forum struct {
	ID          int64     `bson:"id"`            //论坛ID
	ForumAreaID int64     `bson:"forum_area_id"` //所属分区ID
	Name        string    `bson:"name"`          //名称
	Description string    `bson:"description"`   //描述信息
	TopicCount  int64     `bson:"topic_count"`   //累计帖子总数
	ReplyCount  int64     `bson:"reply_count"`   //累计回复数量
	CreatedAt   time.Time `bson:"created_at"`    //创建时间
	UpdatedAt   time.Time `bson:"updated_at"`    //更新时间
}
