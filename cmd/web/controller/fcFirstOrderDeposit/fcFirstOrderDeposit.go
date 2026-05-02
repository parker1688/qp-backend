// The build tag makes sure the stub is not built in the final build.

package fcFirstOrderDeposit

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcFirstOrderDeposit/save
func SaveFcFirstOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcFirstOrderDeposit
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

	data, err := modules.SaveFcFirstOrderDeposit(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcFirstOrderDeposit/findPage
func FindPageFcFirstOrderDepositControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcFirstOrderDeposit
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")

	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.DepositRemark = c.DefaultQuery("deposit_remark", "")
	jsonp.Ip = c.DefaultQuery("ip", "")
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.PayTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("pay_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.ChannelId = tool.Atoi(c.DefaultQuery("channel_id", ""))
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.PaymentId = tool.Atoi(c.DefaultQuery("payment_id", ""))
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.PayAliasName = c.DefaultQuery("pay_alias_name", "")
	jsonp.PaymentName = c.DefaultQuery("payment_name", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.OrderType = tool.Atoi(c.DefaultQuery("order_type", ""))
	jsonp.OrderSecondType = tool.Atoi(c.DefaultQuery("order_second_type", ""))

	jsonp.AuthBy = c.DefaultQuery("auth_by", "")
	jsonp.AuthTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("auth_time", "")))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcFirstOrderDeposit(jsonp.PageNo, jsonp.PageSize, &jsonp.FcFirstOrderDeposit)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcFirstOrderDeposit/findByKey
func FindByKeyFcFirstOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcFirstOrderDeposit
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
	data := modules.FindByKeyFcFirstOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcFirstOrderDeposit/update
func UpdateFcFirstOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcFirstOrderDeposit
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

	data := modules.UpdateFcFirstOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcFirstOrderDeposit/delete
func DeleteFcFirstOrderDepositControl(c *gin.Context) {
	var jsonp dos.FcFirstOrderDeposit
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
	data := modules.DeleteFcFirstOrderDeposit(&jsonp)
	response.SuccessJSON(c, data)
}
