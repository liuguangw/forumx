package mobile

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/service/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//CheckCode 判断用户输入的短信验证码是否正确
func CheckCode(ctx context.Context, mobile string, codeType int, inputCode string) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	//获取短信验证码记录
	lastCodeLog, err := findLastCodeLog(ctx, mobile, codeType)
	if err != nil {
		return false, err
	}
	if lastCodeLog == nil {
		return false, errors.New("短信验证码记录不存在")
	}
	//获取短信验证码的有效分钟数
	codeDurationMinutes := 15
	codeLifetimeMinutesConfig, err := config.LoadConfig(ctx, "mobile_code.lifetime.minutes")
	if err != nil {
		return false, err
	}
	codeDurationMinutes = int(codeLifetimeMinutesConfig.ValueNumber)
	codeDuration := time.Duration(codeDurationMinutes) * time.Minute
	//判断验证码是否有效
	if lastCodeLog.CreatedAt.Add(codeDuration).Before(time.Now()) {
		return false, errors.New("短信验证码已失效,请重新获取")
	}
	//判断短信验证码是否正确
	isValid := lastCodeLog.Code == inputCode
	//如果验证成功,则标记为已使用
	if isValid {
		filter := bson.M{
			"mobile":    mobile,
			"code_type": codeType,
			"code_used": false,
		}
		updateData := bson.M{
			"$set": bson.M{"code_used": true,
				"updated_at": time.Now(),
			},
		}
		coll, err := db.Collection(common.UserMobileCodeCollectionName)
		if err != nil {
			return false, err
		}
		if _, err := coll.UpdateOne(ctx, filter, updateData); err != nil {
			return false, errors.Wrap(err, "更新验证码数据出错")
		}
	}
	return isValid, nil
}
