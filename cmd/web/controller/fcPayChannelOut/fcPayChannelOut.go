// The build tag makes sure the stub is not built in the final build.

package fcPayChannelOut

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPayChannelOut/save
func SaveFcPayChannelOutControl(c *gin.Context) {
	var jsonp dos.FcPayChannelOut
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

	/*if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveFcPayChannelOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelOut/findPage
func FindPageFcPayChannelOutControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPayChannelOut
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
	data, total := modules.FindPageFcPayChannelOut(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPayChannelOut, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPayChannelOut/findByKey
func FindByKeyFcPayChannelOutControl(c *gin.Context) {
	var jsonp dos.FcPayChannelOut
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
	data := modules.FindByKeyFcPayChannelOut(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelOut/update
func UpdateFcPayChannelOutControl(c *gin.Context) {
	var jsonp dos.FcPayChannelOut
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

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	/*payChannelOut := modules.FindByKeyFcPayChannelOutFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, payChannelOut.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.UpdateFcPayChannelOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelOut/delete
func DeleteFcPayChannelOutControl(c *gin.Context) {
	var jsonp dos.FcPayChannelOut
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

	/*payChannelOut := modules.FindByKeyFcPayChannelOutFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, payChannelOut.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.DeleteFcPayChannelOut(&jsonp)
	response.SuccessJSON(c, data)
}
