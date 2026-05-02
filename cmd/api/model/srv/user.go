package srv

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserMaterial(userNameM string) (material *dos.FcUserMaterial) {
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	v := global.G_REDIS.Get(context.Background(), memberRedisKey).Val()
	err := tool.JsonUnmarshalFromString(v, &material)
	if err != nil {
		global.G_LOG.Errorf("redis JsonUnmarshalFromString GetUserMaterial err:%v", err)
	}
	return
}

func SetUserMaterial(material *dos.FcUserMaterial) {
	userNameM := fmt.Sprintf("%s:%s", material.MerchantCode, material.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err := global.G_REDIS.Set(context.Background(), memberRedisKey, tool.String(material), 30*24*time.Hour).Err()
	if err != nil {
		global.G_LOG.Errorf("redis set err:%v", err)
	}
}

// GetLoginUserMaterial
//
//	@Description: 获取提交的用户信息
//	@param c HTTP上下文
//	@return bool true: 有用户 false 无用户
//	@return material 用户信息
func GetLoginUserMaterial(c *gin.Context) (bool, *dos.FcUserMaterial) {
	tokenStr := c.Request.Header.Get(enmus.LOGIN_TOKEN)
	isLogin, userNameM := VerifyJWTToken(tokenStr)
	if !isLogin {
		return false, nil
	}
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	userVal := global.G_REDIS.Get(context.Background(), memberRedisKey).Val()
	var user *dos.FcUserMaterial
	err := tool.JsonUnmarshalFromString(userVal, &user)
	if err != nil {
		return false, nil
	}
	return true, user
}

// GetLoginUserMaterialByTokenStr
//
//	@Description: 获取提交的用户信息
//	@param tokenStr token字符串
//	@return bool true: 有用户 false 无用户
//	@return material 用户信息
func GetLoginUserMaterialByTokenStr(tokenStr string) (bool, *dos.FcUserMaterial) {
	isLogin, userNameM := VerifyJWTToken(tokenStr)
	if !isLogin {
		return false, nil
	}
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	userVal := global.G_REDIS.Get(context.Background(), memberRedisKey).Val()
	var user *dos.FcUserMaterial
	err := tool.JsonUnmarshalFromString(userVal, &user)
	if err != nil {
		return false, nil
	}
	return true, user
}

// GetUserInfo - 获取用户信息
// @param {*gin.Context} c
// @returns *dos.FcUserMaterial
func GetUserInfo(c *gin.Context) (*dos.FcUserMaterial, error) {
	userInfoF, ret := c.Get(vo.USER_NAME_INFO_G)
	if !ret {
		return nil, errors.New("获取用户信息失败")
	}
	return userInfoF.(*dos.FcUserMaterial), nil
}
