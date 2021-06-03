package session

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//Save 保存用户会话数据到数据库
func Save(ctx context.Context, itemInfo *models.UserSession) error {
	itemInfo.UpdatedAt = time.Now()
	if ctx == nil {
		ctx = context.Background()
	}
	if itemInfo.ID == "" {
		sessionID, err := generateUniqueID(ctx)
		if err != nil {
			return err
		}
		itemInfo.ID = sessionID
	}
	coll, err := db.Collection(common.UserSessionCollectionName)
	if err != nil {
		return err
	}
	//update
	filter := bson.M{
		"id": itemInfo.ID,
	}
	opt := options.Update().SetUpsert(true)
	if _, err := coll.UpdateOne(ctx, filter, bson.M{
		"$set": itemInfo,
	}, opt); err != nil {
		return err
	}
	return nil
}
