package sms

import (
	"context"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/service/config"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/db"
	"github.com/pkg/errors"
	"time"
)

//SendSms 发送短信验证码,并将其保存到集合中
func SendSms(ctx context.Context, codeLog *models.UserMobileCode) error {
	codeLog.Code = generateRandomCode(5)
	codeLog.CreatedAt = time.Now()
	codeLog.UpdatedAt = codeLog.CreatedAt
	//短信验证码的保存天数
	storeDaysConfig, err := config.LoadConfig(ctx, "mobile_code.store.days")
	if err != nil {
		return err
	}
	storeDaysDuration := time.Duration(storeDaysConfig.ValueNumber) * 24 * time.Hour
	codeLog.ExpiredAt = time.Now().Add(storeDaysDuration)
	//获取当前短信验证码驱动类型
	if ctx == nil {
		ctx = context.Background()
	}
	driverConfig, err := config.LoadConfig(ctx, "sms.driver.type")
	if err != nil {
		return err
	}
	//发送短信
	if err := processSendSms(ctx, int(driverConfig.ValueNumber), codeLog); err != nil {
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
func processSendSms(ctx context.Context, smsDriverType int, codeLog *models.UserMobileCode) error {
	if smsDriverType == models.SmsNullDriver {
		return nil
	}
	//获取短信验证码的有效分钟数
	codeDurationMinutes := 15
	codeLifetimeMinutesConfig, err := config.LoadConfig(ctx, "mobile_code.lifetime.minutes")
	if err != nil {
		return err
	}
	codeDurationMinutes = int(codeLifetimeMinutesConfig.ValueNumber)
	if smsDriverType == models.SmsTencentDriver {
		return sendTencentSms(ctx, codeDurationMinutes, codeLog)
	}
	return errors.New("无效的短信发送驱动类型")
}
