package mobile

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//findLastCodeLog 获取指定手机号的最后一条未使用的短信验证码记录
func findLastCodeLog(ctx context.Context, mobile string, codeType int) (*models.UserMobileCode, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	coll, err := db.Collection(common.UserMobileCodeCollectionName)
	if err != nil {
		return nil, err
	}
	itemInfo := new(models.UserMobileCode)
	filter := bson.M{
		"mobile":    mobile,
		"code_type": codeType,
		"code_used": false,
	}
	//排序方式
	opt := options.FindOne().SetSort(bson.M{"created_at": -1})
	if err := coll.FindOne(ctx, filter, opt).Decode(itemInfo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return itemInfo, nil
}
