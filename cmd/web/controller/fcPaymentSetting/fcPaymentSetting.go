// The build tag makes sure the stub is not built in the final build.

package fcPaymentSetting

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPaymentSetting/save
func SaveFcPaymentSettingControl(c *gin.Context) {
	var jsonp dos.FcPaymentSetting
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

	m := modules.FindByKeyFcPaymentSettingFirst(&dos.FcPaymentSetting{
		PaymentCode:  jsonp.PaymentCode,
		PKey:         jsonp.PKey,
		MerchantCode: jsonp.MerchantCode,
	})
	if len(m.Id) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "Key已存在")
		return
	}
	data, err := modules.SaveFcPaymentSetting(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSetting/findPage
func FindPageFcPaymentSettingControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPaymentSetting
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.PKey = c.DefaultQuery("p_key", "")
	jsonp.PValue = c.DefaultQuery("p_value", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
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
	data, total := modules.FindPageFcPaymentSetting(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPaymentSetting)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPaymentSetting/findByKey
func FindByKeyFcPaymentSettingControl(c *gin.Context) {
	var jsonp dos.FcPaymentSetting
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
	data := modules.FindByKeyFcPaymentSetting(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSetting/update
func UpdateFcPaymentSettingControl(c *gin.Context) {
	var jsonp dos.FcPaymentSetting
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

	data := modules.UpdateFcPaymentSetting(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSetting/delete
func DeleteFcPaymentSettingControl(c *gin.Context) {
	var jsonp dos.FcPaymentSetting
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
	data := modules.DeleteFcPaymentSetting(&jsonp)
	response.SuccessJSON(c, data)
}
