package mobile

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//FindMobileBindLog 查找某个手机号的绑定记录
func FindMobileBindLog(ctx context.Context, mobile string) (*models.UserMobile, error) {
	coll, err := db.Collection(common.UserMobileCollectionName)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	itemInfo := new(models.UserMobile)
	if err := coll.FindOne(ctx, bson.M{"mobile": mobile}).Decode(itemInfo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return itemInfo, nil
}
