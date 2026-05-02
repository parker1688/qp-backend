// The build tag makes sure the stub is not built in the final build.

package fcPayChannel

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPayChannel/save
func SaveFcPayChannelControl(c *gin.Context) {
	var jsonp dos.FcPayChannel
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

	if jsonp.Currency == "" {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	if jsonp.Strategy == "" {
		jsonp.Strategy = "[]"
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, _ := modules.SaveFcPayChannel(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannel/findPage
func FindPageFcPayChannelControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPayChannel
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.Icon = c.DefaultQuery("icon", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MinLevel = tool.Atoi(c.DefaultQuery("min_level", ""))
	jsonp.MaxLevel = tool.Atoi(c.DefaultQuery("max_level", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.ChannelType = tool.Atoi(c.DefaultQuery("channel_type", ""))
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcPayChannel(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPayChannel, c)

	list := []*dos.FcPayChannelResp{}
	for _, v := range data {
		payChannelResp := dos.FcPayChannelResp{}
		tool.JsonMapper(v, &payChannelResp)
		payChannelResp.MerchantName = v.Merchant.MerchantName
		list = append(list, &payChannelResp)
	}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
}

// api: api/fcPayChannel/findByKey
func FindByKeyFcPayChannelControl(c *gin.Context) {
	var jsonp dos.FcPayChannel
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
	jsonp.MerchantCode = ""
	data := modules.FindByKeyFcPayChannel(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannel/update
func UpdateFcPayChannelControl(c *gin.Context) {
	var jsonp dos.FcPayChannel
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

	if jsonp.Currency == "" {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcPayChannel(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannel/delete
func DeleteFcPayChannelControl(c *gin.Context) {
	var jsonp dos.FcPayChannel
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

	payChannel := modules.FindByKeyFcPayChannelFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, payChannel.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcPayChannel(&jsonp)
	response.SuccessJSON(c, data)
}
