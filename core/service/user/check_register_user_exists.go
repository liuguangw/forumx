package user

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckRegisterUserExists 注册用户前,判断用户名、邮箱是否已存在
func CheckRegisterUserExists(ctx context.Context, username, email string) *common.AppError {
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

//usernameExists 判断指定的用户名是否已经注册
func usernameExists(ctx context.Context, username string) (bool, error) {
	coll, err := db.Collection(common.UserCollectionName)
	if err != nil {
		return false, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	count, err := coll.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
