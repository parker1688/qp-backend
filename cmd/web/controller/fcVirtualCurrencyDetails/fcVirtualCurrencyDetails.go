// The build tag makes sure the stub is not built in the final build.

package fcVirtualCurrencyDetails

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVirtualCurrencyDetails/save
func SaveFcVirtualCurrencyDetailsControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyDetails
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

	data, _ := modules.SaveFcVirtualCurrencyDetails(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyDetails/findPage
func FindPageFcVirtualCurrencyDetailsControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVirtualCurrencyDetails
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.CurrencyName = c.DefaultQuery("currency_name", "")
	jsonp.CurrencyChain = c.DefaultQuery("currency_chain", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.ToAddr = c.DefaultQuery("to_addr", "")
	jsonp.ToAddrQrPre = c.DefaultQuery("to_addr_qr_pre", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVirtualCurrencyDetails(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVirtualCurrencyDetails)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVirtualCurrencyDetails/findByKey
func FindByKeyFcVirtualCurrencyDetailsControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyDetails
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
	data := modules.FindByKeyFcVirtualCurrencyDetails(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyDetails/update
func UpdateFcVirtualCurrencyDetailsControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyDetails
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

	data := modules.UpdateFcVirtualCurrencyDetails(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVirtualCurrencyDetails/delete
func DeleteFcVirtualCurrencyDetailsControl(c *gin.Context) {
	var jsonp dos.FcVirtualCurrencyDetails
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
	data := modules.DeleteFcVirtualCurrencyDetails(&jsonp)
	response.SuccessJSON(c, data)
}
