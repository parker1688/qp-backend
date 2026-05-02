package gooleMFA

import (
	"bootpkg/cmd/web/srv"
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(c *gin.Context) {
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	if len(userInfo.Mfa) > 0 {
		response.FailErrJSON(c, ecode.GoogleMFABindOK, "")
		return
	}
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      c.Request.Host + userInfo.UserName,
		AccountName: userInfo.UserName,
	})
	secret := key.Secret()
	img, _ := key.Image(200, 200)
	var buf strings.Builder
	png.Encode(&buf, img)

	base64Image := base64.StdEncoding.EncodeToString([]byte(buf.String()))
	response.SuccessJSON(c, struct {
		Secret string `json:"secret"`
		Qr     string `json:"qr"`
	}{
		Secret: secret,
		Qr:     "data:image/png;base64," + base64Image,
	})

	// 更新验证时间
	userInfo.UpdateTime = automaticType.Time(time.Now())
	global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", userInfo.Id).
		Update("update_time", userInfo.UpdateTime)
	c.Set("UserInfo", userInfo)
}

func ValidateTOTP(c *gin.Context) {
	var jsonp struct {
		Secret       string `form:"secret" json:"secret"`
		ValidateCode string `form:"validate_code" json:"validate_code"`
		//Date         int64  `json:"date"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	if len(userInfo.Mfa) > 0 {
		response.FailErrJSON(c, ecode.GoogleMFABindOK, "")
		return
	}
	adminUser := modules.FindByKeyAdmindUserFirst(&dos.AdminUser{
		BaseDos: dos.BaseDos{Id: userInfo.Id},
	})
	if adminUser != nil && automaticType.Now().Timer().Unix()-adminUser.UpdateTime.Timer().Unix() > 7200 {
		global.G_LOG.Errorf("[ValidateTOTP] MFA secret key is expires: now=%d, updateTime=%d, date=%v",
			automaticType.Now().Timer().Unix(), userInfo.UpdateTime.Timer().Unix(), userInfo.UpdateTime)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "待绑定密钥已失效，请重新获取")
		return
	}
	localTime := time.Now() //time.Unix(jsonp.Date, 0)
	subTime := time.Since(localTime)
	isSuc, err := totp.ValidateCustom(jsonp.ValidateCode, jsonp.Secret, localTime, totp.ValidateOpts{
		Period:    30,
		Skew:      2,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if !isSuc || err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "6位效验码错误,请重试")
		return
	}
	secretEncrypt := TOTPEncryptAES(jsonp.Secret, userInfo.UserName)
	if len(secretEncrypt) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "处理失败请重试")
		return
	}
	m := &dos.AdminUser{
		UpdateBy: userInfo.UserName,
		Mfa:      secretEncrypt,
		MfaHour:  int(math.Floor(subTime.Seconds())),
	}
	m.Id = userInfo.Id
	saveSuc := modules.UpdateAdminUserMfa(m)
	if !saveSuc {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "数据保存失败,请重试")
		return
	}
	if !srv.ExitLogin(c) {
	}
	response.SuccessJSON(c, saveSuc)
}

func ValidateUserTOTP(c *gin.Context) {
	var jsonp struct {
		ValidateCode string `form:"validate_code" json:"validate_code"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	if len(userInfo.Mfa) == 0 {
		response.FailErrJSON(c, ecode.GoogleMFABind, "")
		return
	}
	mfaErrCount := global.G_REDIS.Get(context.Background(), enmus.REDIS_LOGIN_MFA+userInfo.UserName).Val()
	if tool.Int(mfaErrCount) >= 7 {
		f := global.G_REDIS.TTL(context.Background(), enmus.REDIS_LOGIN_MFA+userInfo.UserName).Val().Minutes()
		response.FailErrJSON(c, response.ERROR_PARAMETER, "虚拟6位验证码错误次数过多，账号将被临时锁定 "+tool.String(int(f))+" 分钟")
		return
	}
	secret := TOTPDecryptAES(userInfo.Mfa, userInfo.UserName)
	//subTime := time.Duration(userInfo.MfaHour) * time.Second
	isSuc, err := totp.ValidateCustom(jsonp.ValidateCode, secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      2,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	/*if global.CONFIG.General.ENV == enmus.Debug {
		t := time.Now().Format("2006010215")
		if jsonp.ValidateCode == t {
			isSuc = true
			err = nil
		}
	}*/

	if err != nil {
		global.G_LOG.Errorf("[ValidateUserTOTP] validate custom failed: %v", err.Error())
	}

	if !isSuc || err != nil {
		count := global.G_REDIS.Incr(context.Background(), enmus.REDIS_LOGIN_MFA+userInfo.UserName).Val()
		global.G_REDIS.Expire(context.Background(), enmus.REDIS_LOGIN_MFA+userInfo.UserName, 30*time.Minute)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "虚拟6位验证码错误，您还可以尝试 "+tool.String(7-count)+" 次（账号将被临时锁定 30 分钟）")
		return
	}
	global.G_REDIS.Del(context.Background(), enmus.REDIS_LOGIN_MFA+userInfo.UserName)
	sessionId, _ := c.Cookie(enmus.LOGIN_COOKIE)
	//获取用户信息
	tokenKey := enmus.REDIS_LOGIN_TOKEN + sessionId
	seconds := global.G_REDIS.TTL(context.Background(), tokenKey).Val().Seconds()
	_ = global.G_REDIS.Set(context.Background(), userInfo.Token, "1", time.Duration(seconds)*time.Second) //谷歌验证码
	response.SuccessJSON(c, isSuc)
}

func TOTPEncryptAES(secret, username string) string {
	aesKey := fmt.Sprintf("%016s", username)
	B, err := tool.EncryptAES([]byte(secret), []byte(aesKey))
	if err != nil {
		return ""
	}
	randStr := strings.ToUpper(tool.RandString(len(username)))
	return randStr + base64.StdEncoding.EncodeToString(B)
}

func TOTPDecryptAES(data, username string) string {
	if len(data) < len(username) {
		return ""
	}
	secret := data[len(username):]
	secretByte, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return ""
	}
	aesKey := fmt.Sprintf("%016s", username)
	B, err := tool.DecryptAES(secretByte, []byte(aesKey))
	if err != nil {
		return ""
	}
	return string(B)
}
