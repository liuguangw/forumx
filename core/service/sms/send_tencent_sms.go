package sms

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/liuguangw/forumx/core/models"
	"github.com/liuguangw/forumx/core/service/config"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

//sendTencentSms 调用腾讯云短信
func sendTencentSms(ctx context.Context, codeDurationMinutes int, codeLog *models.UserMobileCode) error {
	if ctx == nil {
		ctx = context.Background()
	}
	secretIDConfig, err := config.LoadConfig(ctx, "tencent.secret_id")
	if err != nil {
		return err
	}
	secretKeyConfig, err := config.LoadConfig(ctx, "tencent.secret_key")
	if err != nil {
		return err
	}
	smsSdkAppIDConfig, err := config.LoadConfig(ctx, "tencent.sms_sdk.app_id")
	if err != nil {
		return err
	}
	signTextConfig, err := config.LoadConfig(ctx, "sms.tencent.sign_text")
	if err != nil {
		return err
	}
	var templateID string
	if codeLog.CodeType == models.MobileCodeTypeBindAccount {
		templateIDConfig, err := config.LoadConfig(ctx, "sms.tencent.bind_mobile.template_id")
		if err != nil {
			return err
		}
		templateID = templateIDConfig.ValueString
	} else if codeLog.CodeType == models.MobileCodeTypeResetPassword {
		templateIDConfig, err := config.LoadConfig(ctx, "sms.tencent.reset_password.template_id")
		if err != nil {
			return err
		}
		templateID = templateIDConfig.ValueString
	}
	templateParamSet := []string{
		codeLog.Code,
		strconv.Itoa(codeDurationMinutes),
	}
	return requestTencentSmsAPI(ctx, secretIDConfig.ValueString, secretKeyConfig.ValueString, smsSdkAppIDConfig.ValueString,
		signTextConfig.ValueString, templateID, codeLog.Mobile, templateParamSet)
}

//requestTencentSmsApi 请求腾讯云短信api
func requestTencentSmsAPI(ctx context.Context, secretID, secretKey, smsSdkAppID,
	signText, templateID, mobileNumber string, templateParamSet []string) error {
	requestData := map[string]interface{}{
		"PhoneNumberSet":   []string{"+86" + mobileNumber},
		"TemplateID":       templateID,
		"Sign":             signText,
		"TemplateParamSet": templateParamSet,
		"SmsSdkAppid":      smsSdkAppID,
	}
	payloadData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestBody := string(payloadData)
	apiHost := "sms.tencentcloudapi.com"
	requestTime := time.Now()
	requestTimestamp := requestTime.Unix()
	requestDate := requestTime.UTC().Format("2006-01-02")
	//fmt.Println("time="+requestTime.UTC().Format("2006-01-02 15:04:05"))
	authorization := getSmsAPIAuthorization(apiHost, requestBody, requestTimestamp, requestDate,
		secretID, secretKey)
	client := resty.New().
		SetHeaders(map[string]string{
			"Authorization":  authorization,
			"Content-Type":   "application/json",
			"X-TC-Action":    "SendSms",
			"X-TC-Timestamp": strconv.FormatInt(requestTimestamp, 10),
			"X-TC-Version":   "2019-07-11",
		})
	resp, err := client.R().
		SetBody(requestBody).
		Post("https://" + apiHost)
	if err != nil {
		return err
	}
	statusCode := resp.StatusCode()
	if statusCode != 200 {
		return errors.New(resp.String())
	}
	var responseData struct {
		Response *struct {
			RequestID string `json:"RequestId"`
			Error     *struct {
				Code    string
				Message string
			}
			SendStatusSet []*struct {
				SerialNo       string
				PhoneNumber    string
				Fee            int
				SessionContext string
				Code           string
				Message        string
				IsoCode        string
			}
		}
	}
	if err := json.Unmarshal(resp.Body(), &responseData); err != nil {
		return err
	}
	responseError := responseData.Response.Error
	if responseError != nil {
		return errors.New("[" + responseError.Code + "]" + responseError.Message)
	}
	sendStatus := responseData.Response.SendStatusSet[0]
	if sendStatus.Code != "Ok" {
		return errors.New("[" + sendStatus.Code + "]" + sendStatus.Message)
	}
	return nil
}

//getSmsAPIAuthorization 计算签名
func getSmsAPIAuthorization(apiHost, requestBody string, requestTimestamp int64, requestDate, secretID, secretKey string) string {
	algorithm := "TC3-HMAC-SHA256"
	canonicalHeaders := "content-type:application/json\n" + "host:" + apiHost + "\n"
	signedHeaders := "content-type;host"
	hashedRequestPayload := sha256Hex(requestBody)
	canonicalRequest := "POST\n/\n\n" +
		canonicalHeaders + "\n" +
		signedHeaders + "\n" + hashedRequestPayload
	//fmt.Println("canonicalRequest=" + canonicalRequest)

	serviceName := "sms"
	credentialScope := requestDate + "/" + serviceName + "/tc3_request"
	hashedCanonicalRequest := sha256Hex(canonicalRequest)
	stringToSign := algorithm + "\n" +
		strconv.FormatInt(requestTimestamp, 10) + "\n" +
		credentialScope + "\n" +
		hashedCanonicalRequest
	//fmt.Println("stringToSign=" + stringToSign)

	secretDate := hmacSha256(requestDate, "TC3"+secretKey)
	secretService := hmacSha256(serviceName, secretDate)
	secretSigning := hmacSha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacSha256(stringToSign, secretSigning)))
	return algorithm + " Credential=" + secretID + "/" + credentialScope +
		", SignedHeaders=" + signedHeaders +
		", Signature=" + signature
}
