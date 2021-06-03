package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/app/request"
	"github.com/liuguangw/forumx/app/service/captcha"
	"github.com/liuguangw/forumx/app/service/response"
	"github.com/liuguangw/forumx/app/service/session"
	"github.com/liuguangw/forumx/app/service/tools"
	"github.com/liuguangw/forumx/app/service/user"
	"github.com/liuguangw/forumx/core/common"
	"github.com/pkg/errors"
)

//Login 处理用户登录请求
func Login(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewLoginAccount(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	//检测验证码
	captchaPassed, err := captcha.CheckCode(ctx, req.CaptchaID, req.CaptchaCode, true)
	if err != nil {
		return response.WriteAppError(c, common.ErrorInternalServer, "判断验证码出错")
	}
	if !captchaPassed {
		return response.WriteAppError(c, common.ErrorInputFieldInvalid, "验证码错误")
	}
	//判断密码是否正确
	userInfo, err := user.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "find user "+req.Username+" failed"))
	}
	if userInfo == nil {
		return response.WriteAppError(c, common.ErrorUserNotFound, "不存在此用户")
	}
	if !user.VerifyPassword(userInfo, req.Password) {
		return response.WriteAppError(c, common.ErrorPassword, "用户名或密码错误")
	}
	userID := userInfo.ID
	//没有启用双重认证的用户直接通过身份验证
	authenticated := !userInfo.Enable2FA
	userSession, expiresIn, err := session.LoginUser(ctx, userID, authenticated)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "登录出错"))
	}
	responseData := map[string]interface{}{
		"id":         userID,
		"session_id": userSession.ID,
		"expires_in": expiresIn,
	}
	if authenticated {
		return response.WriteSuccess(c, responseData)
	}
	return response.Write(c, &common.AppResponse{
		Code:    common.ErrorNeedAuthentication,
		Message: "需要二次验证",
		Data:    responseData,
	})
}
