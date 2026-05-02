// The build tag makes sure the stub is not built in the final build.

package fcOrderDeposit

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcOrderDeposit/save
func SaveFcOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcOrderDeposit
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

	data, _ := modules.SaveFcOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderDeposit/findPage
func FindPageFcOrderDepositControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcOrderDeposit
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.LastStartAt = c.DefaultQuery("last_startAt", "")
	jsonp.LastEndAt = c.DefaultQuery("last_endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")

	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.EntityAccountHolder = c.DefaultQuery("entity_account_holder", "")
	jsonp.EntityAccountBankName = c.DefaultQuery("entity_account_bank_name", "")
	jsonp.EntityAccountNumber = c.DefaultQuery("entity_account_number", "")
	jsonp.RemitterAccountHolder = c.DefaultQuery("remitter_account_holder", "")
	jsonp.RemitterAccountBankName = c.DefaultQuery("remitter_account_bank_name", "")
	jsonp.RemitterAccountNumber = c.DefaultQuery("remitter_account_number", "")
	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.DepositRemark = c.DefaultQuery("deposit_remark", "")
	jsonp.Ip = c.DefaultQuery("ip", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.ChannelId = tool.Atoi(c.DefaultQuery("channel_id", ""))
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.PaymentId = tool.Atoi(c.DefaultQuery("payment_id", ""))
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total, sumAmount := modules.FindPageFcOrderDeposit(jsonp.PageNo, jsonp.PageSize, &jsonp.FcOrderDeposit, jsonp.PageTimeQuery, c)
	resp := map[string]interface{}{}
	resp["list"] = data
	resp["sum_amount"] = sumAmount
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, resp)
}

// api: api/fcOrderDeposit/findByKey
func FindByKeyFcOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcOrderDeposit
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
	data := modules.FindByKeyFcOrderDeposit(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderDeposit/update
func UpdateFcOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcOrderDeposit
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderDeposit/delete
func DeleteFcOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcOrderDeposit
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

	orderDeposit := modules.FindByKeyFcOrderDepositFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, orderDeposit.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}
