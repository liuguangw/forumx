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
	userSession, err := session.CheckRequest(ctx, c)
	if err != nil {
		return err
	}
	//使用token获取已验证过密码的用户ID
	userID, err := user.LoadTotpAuthUserID(ctx, req.Token)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "获取用户token信息失败"))
	}
	//获取用户信息
	userInfo, dbErr := user.FindUserByID(ctx, userID)
	if dbErr != nil {
		return response.WriteInternalError(c, errors.Wrap(dbErr, "find user info failed"))
	}
	//读取绑定的令牌信息
	tokenData, err := totp.FindTotpKeyByUserID(ctx, userID)
	if err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "读取绑定的令牌记录失败"))
	}
	//验证动态码是否正确
	if !totp2.Validate(req.Code, tokenData.SecretKey) {
		return response.WriteAppError(c, common.ErrorTwoFactorAuthenticationCode, "动态验证码错误")
	}
	//登录
	if err := session.LoginUser(ctx, userSession, userID); err != nil {
		return response.WriteInternalError(c, errors.Wrap(err, "用户登录失败"))
	}
	//todo 清理token
	responseData := map[string]interface{}{
		"id":       userID,
		"nickname": userInfo.Nickname,
	}
	return response.WriteSuccess(c, responseData)
}
