package mobile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/user"
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
	if userInfo.MobileVerified{
		return response.WriteAppError(c,common.ErrorInternalServer,"您当前已经绑定过手机了")
	}
	//todo 判断短信验证码是否正确

	return nil
}
