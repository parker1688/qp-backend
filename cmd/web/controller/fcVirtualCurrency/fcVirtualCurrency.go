// The build tag makes sure the stub is not built in the final build.

package fcVirtualCurrency

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVirtualCurrency/save
func SaveFcVirtualCurrencyControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrency
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

	data, _ := modules.SaveFcVirtualCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrency/findPage
func FindPageFcVirtualCurrencyControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVirtualCurrency
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.CurrencyName = c.DefaultQuery("currency_name", "")
	jsonp.CurrencyNameImg = c.DefaultQuery("currency_name_img", "")
	jsonp.CurrencyChain = c.DefaultQuery("currency_chain", "")
	jsonp.CurrencyProtocol = c.DefaultQuery("currency_protocol", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVirtualCurrency(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVirtualCurrency)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVirtualCurrency/findByKey
func FindByKeyFcVirtualCurrencyControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrency
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
	data := modules.FindByKeyFcVirtualCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrency/update
func UpdateFcVirtualCurrencyControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrency
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

	data := modules.UpdateFcVirtualCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrency/delete
func DeleteFcVirtualCurrencyControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrency
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
	data := modules.DeleteFcVirtualCurrency(&jsonp)
	response.SuccessJSON(c, data)
}
