package models

import "time"

//ForumTopic 论坛帖子信息
type ForumTopic struct {
	ID          int64     `bson:"id"`            //帖子ID
	ForumAreaID int64     `bson:"forum_area_id"` //分区ID
	ForumID     int64     `bson:"forum_id"`      //论坛ID
	TopicType   int       `bson:"topic_type"`    //帖子类型
	Title       string    `bson:"title"`         //标题
	Content     string    `bson:"content"`       //帖子内容
	UserID      int64     `bson:"user_id"`       //发帖人用户ID
	ViewCount   int64     `bson:"view_count"`    //浏览次数
	ReplyCount  int64     `bson:"reply_count"`   //回复次数
	LikeCount   int64     `bson:"like_count"`    //点赞次数
	Locked      bool      `bson:"locked"`        //是否已锁定
	Blocked     bool      `bson:"blocked"`       //是否已经屏蔽
	Deleted     bool      `bson:"deleted"`       //是否已经删除
	Published   bool      `bson:"published"`     //是否已经发布
	PublishedAt time.Time `bson:"published_at"`  //发布时间
	OrderID     int       `bson:"order_id"`      //排序值
	CreatedAt   time.Time `bson:"created_at"`    //创建时间
	UpdatedAt   time.Time `bson:"updated_at"`    //更新时间
}
