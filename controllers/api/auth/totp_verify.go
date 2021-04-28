package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/core/common"
	"github.com/liuguangw/forumx/core/request"
	"github.com/liuguangw/forumx/core/service/response"
	"github.com/liuguangw/forumx/core/service/session"
	"github.com/liuguangw/forumx/core/service/tools"
	"github.com/liuguangw/forumx/core/service/totp"
	"github.com/liuguangw/forumx/core/service/user"
	"github.com/pkg/errors"
	totp2 "github.com/pquerna/otp/totp"
)

//TotpVerify 登录后,使用动态码进行两步验证
func TotpVerify(c *fiber.Ctx) error {
	//获取所需参数
	req, requestErr := request.NewTotpVerify(c)
	if requestErr != nil {
		return requestErr.WriteResponse(c)
	}
	if err := req.CheckRequest(); err != nil {
		return err.WriteResponse(c)
	}
	//加载session
	ctx, cancel := tools.DefaultExecContext()
	defer cancel()
	userSession, err := session.CheckSession(ctx, c)
	if err != nil || userSession == nil {
		return err
	}
	//使用token获取已验证过密码的用户数据
	totpAuthToken := req.Token
	authData, err := user.LoadTotpAuthData(ctx, totpAuthToken)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户token信息失败"))
	}
	if authData == nil {
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "登录已失效,请返回重新输入密码")
	}
	//fmt.Println(totpAuthToken, authData)
	//输入动态码的错误次数限制
	if authData.ErrorCount >= 5 {
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "您的失败次数过多,请返回重新输入密码")
	}
	userID := authData.UserID
	//获取用户信息
	userInfo, err := user.FindUserByID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "find user info failed"))
	}
	//读取绑定的令牌信息
	userTotpKeyData, err := totp.FindTotpKeyByUserID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "读取绑定的令牌记录失败"))
	}
	if userTotpKeyData == nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取绑定的密钥信息失败"))
	}
	//验证动态码是否正确
	if !totp2.Validate(req.Code, userTotpKeyData.SecretKey) {
		//验证错误次数+1
		if err := user.IncrTotpAuthFailedCount(ctx, totpAuthToken); err != nil {
			return response.WriteInternalError(c, errors.Wrap(err, "增加动态码验证失败次数出错"))
		}
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "动态验证码错误")
	}
	//登录
	if err := session.LoginUser(ctx, userSession, userID); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "用户登录失败"))
	}
	//清理token
	if err := user.ClearTotpAuthData(ctx, totpAuthToken); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "清理缓存的totp验证信息失败"))
	}
	responseData := map[string]interface{}{
		"id":       userID,
		"nickname": userInfo.Nickname,
	}
	return response.WriteSuccess(c, responseData)
}
