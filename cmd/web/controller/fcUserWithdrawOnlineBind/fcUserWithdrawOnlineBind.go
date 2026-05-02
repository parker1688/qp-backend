// The build tag makes sure the stub is not built in the final build.

package fcUserWithdrawOnlineBind

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/vo"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserWithdrawOnlineBind/save
func SaveFcUserWithdrawOnlineBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawOnlineBind
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcUserWithdrawOnlineBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawOnlineBind/findPage
func FindPageFcUserWithdrawOnlineBindControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserWithdrawOnlineBind
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.AccountNumber = c.DefaultQuery("account_number", "")
	jsonp.AccountHolder = c.DefaultQuery("account_holder", "")
	jsonp.IsDefault = tool.Atoi(c.DefaultQuery("is_default", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcUserWithdrawOnlineBind(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserWithdrawOnlineBind, c)

	for _, v := range data {
		v.Decrypt()
		p := &vo.FcPrivate{}
		p.AccountHolder = v.AccountHolder
		p.AccountNumber = v.AccountNumber
		tool.PrivateDataHandler(p)
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserWithdrawOnlineBind/findByKey
func FindByKeyFcUserWithdrawOnlineBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawOnlineBind
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
	data := modules.FindByKeyFcUserWithdrawOnlineBind2(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawOnlineBind/update
func UpdateFcUserWithdrawOnlineBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawOnlineBind
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcUserWithdrawOnlineBind(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawOnlineBind/delete
func DeleteFcUserWithdrawOnlineBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawOnlineBind
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

	userWithdrawOnlineBind := modules.FindByKeyFcUserWithdrawOnlineBindFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, userWithdrawOnlineBind.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcUserWithdrawOnlineBind(&jsonp)
	response.SuccessJSON(c, data)
}
