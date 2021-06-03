package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
)

//Info 获取当前登录的用户信息
func Info(c *fiber.Ctx) error {
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, sessionError := session.CheckLogin(ctx, c)
	if sessionError != nil {
		return sessionError.WriteResponse(c)
	}
	//处理用户信息的展示
	userID := userSession.UserID
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户信息失败"))
	}
	if userInfo == nil {
		return response.WriteAppError(c, common.ErrorCommonMessage, "不存在此用户")
	}
	//格式化返回给用户的数据
	responseData := map[string]interface{}{
		"id":              userID,
		"username":        userInfo.Username,
		"nickname":        userInfo.Nickname,
		"avatar_url":      userInfo.AvatarURL,
		"coin_amount":     userInfo.CoinAmount,
		"exp_amount":      userInfo.ExpAmount,
		"email_verified":  userInfo.EmailVerified,
		"mobile_verified": userInfo.MobileVerified,
		"enable_2fa":      userInfo.Enable2FA,
		"register_ip":     userInfo.RegisterIP,
		"created_at":      tools.FormatDateTime(userInfo.CreatedAt),
		"updated_at":      tools.FormatDateTime(userInfo.UpdatedAt),
	}
	return response.WriteSuccess(c, responseData)
}
