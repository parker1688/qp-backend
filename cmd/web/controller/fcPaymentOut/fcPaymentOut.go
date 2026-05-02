// The build tag makes sure the stub is not built in the final build.

package fcPaymentOut

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPaymentOut/save
func SaveFcPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcPaymentOut
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

	data, _ := modules.SaveFcPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentOut/findPage
func FindPageFcPaymentOutControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPaymentOut
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.PaymentName = c.DefaultQuery("payment_name", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MinLevel = tool.Atoi(c.DefaultQuery("min_level", ""))
	jsonp.MaxLevel = tool.Atoi(c.DefaultQuery("max_level", ""))

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
	data, total := modules.FindPageFcPaymentOut(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPaymentOut, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPaymentOut/findByKey
func FindByKeyFcPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcPaymentOut
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
	data := modules.FindByKeyFcPaymentOut(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentOut/update
func UpdateFcPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcPaymentOut
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
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentOut/delete
func DeleteFcPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcPaymentOut
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

	/*paymentOut := modules.FindByKeyFcPaymentOutFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, paymentOut.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.DeleteFcPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}
