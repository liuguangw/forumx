package models

import "time"

//ForumReply 帖子回复信息
type ForumReply struct {
	ID           int64     `bson:"id"`             //回复ID
	ForumTopicID int64     `bson:"forum_topic_id"` //帖子ID
	Content      string    `bson:"content"`        //回复内容
	UserID       int64     `bson:"user_id"`        //回复人用户ID
	LikeCount    int64     `bson:"like_count"`     //点赞次数
	Blocked      bool      `bson:"blocked"`        //是否已经屏蔽
	Deleted      bool      `bson:"deleted"`        //是否已经删除
	CreatedAt    time.Time `bson:"created_at"`     //创建时间
	UpdatedAt    time.Time `bson:"updated_at"`     //更新时间
}
