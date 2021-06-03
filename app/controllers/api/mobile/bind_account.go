package mobile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/models"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/app/service/mobile"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
)

//BindAccount 处理手机号绑定账号
func BindAccount(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewBindAccount(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//必须登录
	userSession, sessionErr := session.CheckLogin(ctx, c)
	if sessionErr != nil {
		return sessionErr.WriteResponse(c)
	}
	//获取用户的基本信息
	userID := userSession.UserID
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	if userInfo == nil {
		return response.WriteAppError(c, common.ErrorCommonMessage, "不存在此用户")
	}
	if userInfo.MobileVerified {
		return response.WriteAppError(c, common.ErrorInternalServer, "您当前已经绑定过手机了")
	}
	//判断验证码是否正确
	isValidCode, err := mobile.CheckCode(ctx, req.Mobile, models.MobileCodeTypeBindAccount, req.Code)
	if err != nil {
		return response.WriteInternalError(c, err)
	}
	if !isValidCode {
		return response.WriteAppError(c, common.ErrorInputFieldInvalid, "短信验证码错误")
	}
	//插入或者更新绑定记录
	if err := mobile.SaveUserBindLog(ctx, userID, req.Mobile); err != nil {
		return response.WriteAppError(c, common.ErrorCommonMessage, err.Error())
	}
	return response.WriteSuccess(c, nil)
}
