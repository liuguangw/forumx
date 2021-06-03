package mobile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/app/service/captcha"
	"github.com/liuguangw/forumx/app/service/mobile"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/sms"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
)

//SendCode 处理发送短信验证码(绑定手机需要登录状态,重置密码无需登录)
func SendCode(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewSendSms(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//检测图形验证码
	captchaPassed, err := captcha.CheckCode(ctx, req.CaptchaID, req.CaptchaCode, true)
	if err != nil {
		return response.WriteAppError(c, common.ErrorInternalServer, "判断图形验证码出错")
	}
	if !captchaPassed {
		return response.WriteAppError(c, common.ErrorInputFieldInvalid, "图形验证码错误")
	}
	var userID int64
	if req.CodeType == models.MobileCodeTypeBindAccount {
		//判断用户是否已经登录
		userSession, sessionErr := session.CheckLogin(ctx, c)
		if sessionErr != nil {
			return sessionErr.WriteResponse(c)
		}
		userID = userSession.UserID
	} else {
		//重置密码,获取绑定记录
		bindLog, err := mobile.FindMobileBindLog(ctx, req.Mobile)
		if err != nil {
			return response.WriteInternalError(c, errors.Wrap(err, "获取绑定记录失败"))
		}
		if bindLog == nil {
			return response.WriteAppError(c, common.ErrorCommonMessage, "此手机号未绑定任何账号")
		}
		userID = bindLog.UserID
	}
	//获取用户信息
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	codeLog := &models.UserMobileCode{
		Mobile:   req.Mobile,
		CodeType: req.CodeType,
		UserID:   userID,
		CodeUsed: false,
	}
	if req.CodeType == models.MobileCodeTypeBindAccount {
		if userInfo.MobileVerified {
			return response.WriteAppError(c, common.ErrorCommonMessage, "你的账号已经绑定过手机了")
		}
	}
	//发送短信验证码
	if err := sms.SendSms(ctx, codeLog); err != nil {
		return response.WriteAppError(c, common.ErrorCommonMessage, err.Error())
	}
	return response.WriteSuccess(c, nil)
}
