// The build tag makes sure the stub is not built in the final build.

package fcPayment

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPayment/save
func SaveFcPaymentControl(c *gin.Context) {
	var jsonp dos.FcPayment
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
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, err := modules.SaveFcPayment(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcPayment/findPage
func FindPageFcPaymentControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPayment
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.PaymentName = c.DefaultQuery("payment_name", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.PayId = c.DefaultQuery("pay_id", "")
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
	data, total := modules.FindPageFcPayment(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPayment, c)

	list := []*dos.FcPaymentResp{}
	for _, v := range data {
		paymentResp := dos.FcPaymentResp{}
		tool.JsonMapper(v, &paymentResp)
		paymentResp.MerchantName = v.Merchant.MerchantName
		list = append(list, &paymentResp)
	}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
}

// api: api/fcPayment/findByKey
func FindByKeyFcPaymentControl(c *gin.Context) {
	var jsonp dos.FcPayment
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
	data := modules.FindByKeyFcPayment(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcPayment/update
func UpdateFcPaymentControl(c *gin.Context) {
	var jsonp dos.FcPayment
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
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcPayment(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayment/delete
func DeleteFcPaymentControl(c *gin.Context) {
	var jsonp dos.FcPayment
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

	payment := modules.FindByKeyFcPaymentFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, payment.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcPayment(&jsonp)
	response.SuccessJSON(c, data)
}
