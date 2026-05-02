package handler

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bytes"
	"context"
	"io"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/utils/convert"
	"github.com/kirinlabs/utils/str"
	"github.com/tidwall/sjson"
)

// 验证登录中间
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuc, user := authToken(c)
		if !isSuc {
			response.FailErrJSON(c, ecode.Unauthorized, "")
			c.Abort()
			return
		}

		if user.EnforcePwd == 1 { //需要强制修改密码
			response.FailErrJSON(c, ecode.RestPassword, "")
			c.Abort()
			return
		}
		if user.EnforcePwd == 0 { //需要绑定google验证码 MFA
			/*if modules.UseGoogleMfa() && len(user.Mfa) == 0 {
					response.FailErrJSON(c, ecode.GoogleMFABind, "")
					c.Abort()
					return
			}*/

		}
		if global.G_REDIS.Get(context.Background(), user.Token).Val() != "1" {
			response.FailErrJSON(c, ecode.GoogleMFAVerify, "")
			c.Abort()
			return
		}

		c.Set("UserInfo", user)
		//if user.UserName == "admin" {
		//	c.Next()
		//	return
		//}
		// 判断是否为超管账号
		if user.AccountType == 1 {
			c.Next()
			return
		}

		isSuc = authAPI(c, user)
		if !isSuc {
			response.FailErrJSON(c, ecode.AccessDenied, "")
			c.Abort()
			return
		}

		/*if !MerchantCodeParam(c, user) {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "无权限参数错误")
			c.Abort()
			return
		}*/

		//设置用户信息
		c.Next()
	}
}

// 验证Token
func authToken(c *gin.Context) (bool, *dos.AdminUser) {
	token := c.Request.Header.Get(enmus.LOGIN_TOKEN)
	if len(token) == 0 {
		return false, nil
	}
	tokenArray := strings.Split(token, ".")
	if len(tokenArray) != 2 {
		return false, nil
	}

	//signMd5 := tool.MD5([]byte(c.ClientIP() + tokenArray[0]))
	//if tokenArray[1] != signMd5 {
	//	return false, nil
	//}
	var err error

	//部署要打开
	sessionId, err := c.Cookie(enmus.LOGIN_COOKIE)
	if err != nil {
		return false, nil
	}
	//获取用户信息
	tokenKey := enmus.REDIS_LOGIN_TOKEN + sessionId
	userName := global.G_REDIS.Get(context.Background(), tokenKey).Val()
	if 0 == len(userName) {
		return false, nil
	}
	//单点登录验证Token
	userInfoKey := enmus.REDIS_LOGIN_USERINFO + userName
	userInfo := global.G_REDIS.Get(context.Background(), userInfoKey).Val()
	if 0 == len(userInfo) {
		return false, nil
	}
	var u *dos.AdminUser
	err = global.JSON.UnmarshalFromString(userInfo, &u)
	if err != nil {
		return false, nil
	}

	if token != u.Token {
		return false, nil
	}

	//续期Token
	seconds := global.G_REDIS.TTL(context.Background(), tokenKey).Val().Seconds()
	if seconds < 1200 {
		newSessionId := SetSessionCookie(c) //Cookie续期
		global.G_REDIS.Set(context.Background(), enmus.REDIS_LOGIN_TOKEN+newSessionId, u.UserName, 2*time.Hour)
		_ = global.G_REDIS.Set(context.Background(), u.Token, "1", 2*time.Hour) //谷歌验证码
		//_ = global.G_REDIS.Expire(tokenKey, 20*time.Minute)
		_ = global.G_REDIS.Expire(context.Background(), userInfoKey, 2*time.Hour)
	}

	return true, u
}

func authAPI(c *gin.Context, u *dos.AdminUser) bool {
	if 0 == len(u.RoleIds) {
		return false
	}
	//多个角色获取
	roles := modules.FindRoleAll()
	var userMenusIdsBuild strings.Builder
	userRolesId := "," + u.RoleIds + ","
	for i := 0; i < len(roles); i++ {
		if str.Contians(userRolesId, ","+convert.String(roles[i].Id)+",") {
			if strings.Contains(roles[i].PermsList, "*") {
				return true
			}
			if i > 0 {
				userMenusIdsBuild.WriteString(",")
			}
			userMenusIdsBuild.WriteString(roles[i].MeusIds)
		}
	}
	if userMenusIdsBuild.Len() == 0 {
		return false
	}

	//获取菜单
	menus := modules.FindMenusAll()
	userMenusIds := "," + userMenusIdsBuild.String() + ","
	for _, v := range menus {
		if !str.Contians(userMenusIds, ","+convert.String(v.Id)+",") || 0 == len(v.ApiRegular) {
			continue
		}
		ok, _ := regexp.MatchString(v.ApiRegular, c.Request.URL.Path)
		if ok {
			return true
		}
	}
	return false
}

func MerchantCodeParam(c *gin.Context, user *dos.AdminUser) bool {

	if strings.Contains(user.DepartmentId+",", "88888888888888888888") {
		return true
	}
	selectMerchantCode := user.MerchantCodes
	//selectMerchantCode := c.Request.Header.Get("MerchantCode") //提交选中的商户ID
	//if len(selectMerchantCode) == 0 || !strings.Contains(user.MerchantCodes+",", selectMerchantCode) {
	//	return false
	//}

	reqContentType := c.Request.Header.Get("Content-Type")
	isJsonRequest := strings.Contains(reqContentType, "application/json")
	isFormUrl := strings.Contains(reqContentType, "application/x-www-form-urlencoded")
	if c.Request.Method == "GET" {
		values, err := url.ParseQuery(c.Request.URL.RawQuery)
		if err != nil {
			return false
		}
		values.Set("merchant_code", selectMerchantCode)
		queryData := values.Encode()

		c.Request.URL.RawQuery = queryData
	} else if isJsonRequest {
		payload, err := c.GetRawData()
		if err != nil {
			return false
		}
		payload, err = sjson.SetBytes(payload, "merchant_code", selectMerchantCode)
		if err != nil {
			return false
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
	} else if isFormUrl {
		payload, err := c.GetRawData()
		if err != nil {
			return false
		}
		values, err := url.ParseQuery(string(payload))
		if err != nil {
			return false
		}
		values.Set("merchant_code", selectMerchantCode)
		formData := values.Encode()
		payload = []byte(formData)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
	}
	return true
}
