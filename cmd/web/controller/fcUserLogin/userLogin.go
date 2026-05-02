package fcUserLogin

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kirinlabs/utils/encrypt"
)

func ResetPassword(c *gin.Context) {
	var jsonp dos.FcUserLogin
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	pwd := tool.RandString(6) + "888"
	password := encrypt.Sha256(pwd + global.CONFIG.General.ApiSHA256Salt)
	eRow := global.G_DB.Model(&dos.FcUserLogin{}).Where("id = ?", jsonp.Id).Update("password", password)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}
	//重置密码后,强制用户退出
	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.Id
	m := modules.FindByKeyFcUserLoginFirst(mm)
	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	global.G_REDIS.Del(context.Background(), memberRedisKey)
	global.G_REDIS.Del(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM).Val()
	response.SuccessJSON(c, pwd)
}

func UpdateStatus(c *gin.Context) {
	var jsonp dos.FcUserLogin
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	eRow := global.G_DB.Model(&dos.FcUserLogin{}).Where("id = ?", jsonp.Id).Update("status", jsonp.Status)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}
	if jsonp.Status == 1 { //禁用
		mm := &dos.FcUserLogin{}
		mm.Id = jsonp.Id
		m := modules.FindByKeyFcUserLoginFirst(mm)
		userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
		memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
		global.G_REDIS.Del(context.Background(), memberRedisKey)
		global.G_REDIS.Del(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM).Val()
	}
	response.SuccessJSON(c, true)
}
