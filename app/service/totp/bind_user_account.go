package totp

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//BindUserAccount 处理用户绑定两步验证令牌
func BindUserAccount(ctx context.Context, userInfo *models.User, secretKey, recoveryCode string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	//使用事务
	client, err := db.Client()
	if err != nil {
		return err
	}
	return client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		return bindUserAccountCallback(sessionContext, userInfo, secretKey, recoveryCode)
	})
}

//bindUserAccountCallback 绑定两步验证令牌时,在MongoDB事务内调用的函数
func bindUserAccountCallback(sessionContext mongo.SessionContext, userInfo *models.User, secretKey, recoveryCode string) error {
	coll, err := db.Collection(common.UserTotpKeyCollectionName)
	if err != nil {
		return err
	}
	//insert row
	timeNow := time.Now()
	itemInfo := &models.UserTotpKey{
		UserID:       userInfo.ID,
		SecretKey:    secretKey,
		RecoveryCode: recoveryCode,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}
	if _, err := coll.InsertOne(sessionContext, itemInfo); err != nil {
		return err
	}
	//update
	userColl, err := db.Collection(common.UserCollectionName)
	if err != nil {
		return err
	}
	filter := bson.M{
		"id": userInfo.ID,
	}
	if _, err := userColl.UpdateOne(sessionContext, filter, bson.M{
		"$set": bson.M{
			"enable_2fa": true,
		},
	}); err != nil {
		return err
	}
	return nil
}
