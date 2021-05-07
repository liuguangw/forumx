package mobile

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/service/user"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//SaveUserBindLog 保存用户的手机绑定记录
func SaveUserBindLog(ctx context.Context, userID int64, mobile string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if userInfo == nil {
		return errors.New("不存在此用户ID")
	}
	//判断此手机号是否已存在其他绑定记录
	existsLog, err := FindMobileBindLog(ctx, mobile)
	if err != nil {
		return err
	}
	if existsLog != nil {
		if existsLog.UserID != userID {
			return errors.New("此手机号已绑定了其它账号")
		}
	}
	//事务处理
	client, err := db.Client()
	if err != nil {
		return err
	}
	return client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		return saveUserBindLogCallback(sessionContext, userID, mobile)
	})
}

//saveUserBindLogCallback 保存用户绑定记录时,在MongoDB事务内调用的函数
func saveUserBindLogCallback(sessionContext mongo.SessionContext, userID int64, mobile string) error {
	updatedAt := time.Now()
	//更新用户绑定状态
	userColl, err := db.Collection(common.UserCollectionName)
	if err != nil {
		return err
	}
	if _, err := userColl.UpdateOne(sessionContext, bson.M{
		"mobile_verified": true,
		"updated_at":      updatedAt,
	}, bson.M{
		"$set": bson.M{
			"id": userID,
		},
	}); err != nil {
		return err
	}
	//插入或者更新绑定记录
	userMobileColl, err := db.Collection(common.UserMobileCollectionName)
	if err != nil {
		return err
	}
	filter := bson.M{"user_id": userID}
	updateData := bson.M{
		"$set": bson.M{
			"user_id":    userID,
			"mobile":     mobile,
			"created_at": updatedAt,
		},
	}
	opt := options.Update().SetUpsert(true)
	if _, err := userMobileColl.UpdateOne(sessionContext, filter, updateData, opt); err != nil {
		return err
	}
	return nil
}
