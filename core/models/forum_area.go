package models

import "time"

//ForumArea 论坛分区
type ForumArea struct {
	ID          int64     `bson:"id"`          //论坛分区ID
	Name        string    `bson:"name"`        //名称
	Description string    `bson:"description"` //描述信息
	CreatedAt   time.Time `bson:"created_at"`  //创建时间
	UpdatedAt   time.Time `bson:"updated_at"`  //更新时间
}
