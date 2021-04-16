package service

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//usernameExists 判断指定的用户名是否已经注册
func usernameExists(username string) (bool, error) {
	coll, err := db.Collection(userCollectionName)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	count, err := coll.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

//RegisterUser 处理用户注册
func RegisterUser(username, nickname, email, password, clientIP string) (*models.User, *common.AppError) {
	//判断用户名是否已存在
	userExists, err := usernameExists(username)
	if err != nil {
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	if userExists {
		return nil, common.NewAppError(common.ErrorUsernameExists, "此用户名已存在")
	}
	//判断邮箱是否已存在绑定
	emailExists, err := userEmailExists(email)
	if err != nil {
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	if emailExists {
		return nil, common.NewAppError(common.ErrorEmailExists, "此邮箱已存在")
	}
	//构造用户数据
	timeNow := time.Now()
	salt := generateRandomString(8)
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
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	//定义超时时间
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		//更新计数器
		counterColl, err1 := db.Collection(counterCollectionName)
		if err1 != nil {
			return err1
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
		if err1 := counterColl.FindOneAndUpdate(sessionContext, counterFilter, updateData).
			Decode(updatedDocument); err1 != nil {
			return err1
		}
		//设置用户ID
		userInfo.ID = updatedDocument.CounterValue
		//插入用户数据
		userColl, err1 := db.Collection(userCollectionName)
		if err1 != nil {
			return err1
		}
		if _, err1 := userColl.InsertOne(sessionContext, userInfo); err1 != nil {
			return err1
		}
		return nil
	})
	if err != nil {
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	return userInfo, nil
}

//hashPassword 处理密码的hash加密
func hashPassword(password, salt string) string {
	saltStr := md5String(salt) + "53f0b847-82f6-43e1-8052-d6ebb97d1e0c"
	return md5String(password + saltStr)
}
