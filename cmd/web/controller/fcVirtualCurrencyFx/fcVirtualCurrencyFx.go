// The build tag makes sure the stub is not built in the final build.

package fcVirtualCurrencyFx

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVirtualCurrencyFx/save
func SaveFcVirtualCurrencyFxControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyFx
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
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveFcVirtualCurrencyFx(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyFx/findPage
func FindPageFcVirtualCurrencyFxControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVirtualCurrencyFx
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.CurrencyName = c.DefaultQuery("currency_name", "")
	jsonp.CurrencyChain = c.DefaultQuery("currency_chain", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.OptType = tool.Atoi(c.DefaultQuery("opt_type", ""))
	jsonp.CurrencyCode = c.DefaultQuery("currency_code", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVirtualCurrencyFx(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVirtualCurrencyFx)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVirtualCurrencyFx/findByKey
func FindByKeyFcVirtualCurrencyFxControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyFx
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
	data := modules.FindByKeyFcVirtualCurrencyFx(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyFx/update
func UpdateFcVirtualCurrencyFxControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyFx
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

	data := modules.UpdateFcVirtualCurrencyFx(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyFx/delete
func DeleteFcVirtualCurrencyFxControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyFx
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
	data := modules.DeleteFcVirtualCurrencyFx(&jsonp)
	response.SuccessJSON(c, data)
}
