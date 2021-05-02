package sms

import (
	"context"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/config"
	"github.com/pkg/errors"
	"time"
)

//SendSms 发送短信验证码,并将其保存到集合中
func SendSms(ctx context.Context, codeLog *models.UserMobileCode) error {
	codeLog.Code = generateRandomCode(5)
	codeLog.CreatedAt = time.Now()
	codeLog.UpdatedAt = codeLog.CreatedAt
	//短信验证码的有效时间(分钟)
	codeDurationMinutes := 15
	codeLog.ExpiredAt = time.Now().Add(time.Duration(codeDurationMinutes) * time.Minute)
	//获取当前短信验证码驱动类型
	if ctx == nil {
		ctx = context.Background()
	}
	driverConfig, err := config.LoadConfig(ctx, "sms.driver.type")
	if err != nil {
		return err
	}
	//发送短信
	if err := processSendSms(ctx, int(driverConfig.ValueNumber), codeDurationMinutes, codeLog); err != nil {
		return err
	}
	//保存记录
	coll, err := db.Collection(common.UserMobileCodeCollectionName)
	if err != nil {
		return err
	}
	if _, err := coll.InsertOne(ctx, codeLog); err != nil {
		return err
	}
	return nil
}

//processSendSms 处理发送短信
func processSendSms(ctx context.Context, smsDriverType, codeDurationMinutes int, codeLog *models.UserMobileCode) error {
	if smsDriverType == models.SmsNullDriver {
		return nil
	}
	if smsDriverType == models.SmsTencentDriver {
		return sendTencentSms(ctx, codeDurationMinutes, codeLog)
	}
	return errors.New("无效的短信发送驱动类型")
}
