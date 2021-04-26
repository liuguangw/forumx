package user

import (
	"context"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//Register 处理用户注册
func Register(ctx context.Context, username, nickname, email, password, clientIP string) (*models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	//构造用户数据
	timeNow := time.Now()
	salt := tools.GenerateRandomString(8)
	encodedPassword := hashPassword(password, salt)
	userInfo := &models.User{
		Username:   username,
		Nickname:   nickname,
		Password:   encodedPassword,
		Salt:       salt,
		RegisterIP: clientIP,
		CreatedAt:  timeNow,
		UpdatedAt:  timeNow,
	}
	//使用事务
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		return registerUserCallback(sessionContext, userInfo)
	})
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

//registerUserCallback 注册用户时,在MongoDB事务内调用的函数
func registerUserCallback(sessionContext mongo.SessionContext, userInfo *models.User) error {
	//计数器 用户ID+1
	counterColl, err := db.Collection(counterCollectionName)
	if err != nil {
		return err
	}
	counterFilter := bson.M{
		"counter_key": models.CounterKeyUserNextID,
	}
	updateData := bson.M{
		"$inc": bson.M{
			"counter_value": 1,
		},
	}
	updatedDocument := &models.Counter{}
	if err := counterColl.FindOneAndUpdate(sessionContext, counterFilter, updateData).
		Decode(updatedDocument); err != nil {
		return err
	}
	//设置用户ID
	userInfo.ID = updatedDocument.CounterValue
	//插入用户数据
	userColl, err := db.Collection(collectionName)
	if err != nil {
		return err
	}
	if _, err := userColl.InsertOne(sessionContext, userInfo); err != nil {
		return err
	}
	return nil
}
