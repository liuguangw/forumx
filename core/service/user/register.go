package user

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//Register 处理用户注册
func Register(ctx context.Context, username, nickname, email, password, clientIP string) (*models.User, *common.AppError) {
	if ctx == nil {
		ctx = context.Background()
	}
	//判断存在性,防止重复
	if err := checkRegisterUserExists(ctx, username, email); err != nil {
		return nil, err
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
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		return registerUserCallback(userInfo, sessionContext)
	})
	if err != nil {
		return nil, common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	return userInfo, nil
}

//checkRegisterUserExists 注册用户前,判断用户名、邮箱是否已存在
func checkRegisterUserExists(ctx context.Context, username, email string) *common.AppError {
	//判断用户名是否已存在
	userExists, err := usernameExists(ctx, username)
	if err != nil {
		return common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	if userExists {
		return common.NewAppError(common.ErrorUsernameExists, "此用户名已存在")
	}
	//判断邮箱是否已存在绑定
	emailExists, err := EmailExists(ctx, email)
	if err != nil {
		return common.NewAppError(common.ErrorInternalServer, "数据库服务异常")
	}
	if emailExists {
		return common.NewAppError(common.ErrorEmailExists, "此邮箱已存在")
	}
	return nil
}

//registerUserCallback 注册用户时,在MongoDB事务内调用的函数
func registerUserCallback(userInfo *models.User, sessionContext mongo.SessionContext) error {
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
