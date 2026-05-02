// The build tag makes sure the stub is not built in the final build.

package fcOrderWithdraw

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcOrderWithdraw/save
func SaveFcOrderWithdrawControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdraw
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

	orderWithdraw := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if orderWithdraw == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, orderWithdraw.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveFcOrderWithdraw(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdraw/findPage
func FindPageFcOrderWithdrawControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcOrderWithdraw
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.PageTimeQuery.StartAt = c.DefaultQuery("startAt", "")
	jsonp.PageTimeQuery.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")

	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "-1"))
	jsonp.Province = c.DefaultQuery("province", "")
	jsonp.City = c.DefaultQuery("city", "")
	jsonp.BankAddress = c.DefaultQuery("bank_address", "")
	jsonp.AccountNumber = c.DefaultQuery("account_number", "")
	jsonp.AccountHolder = c.DefaultQuery("account_holder", "")
	jsonp.AccountBankType = c.DefaultQuery("account_bank_type", "")
	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.Ip = c.DefaultQuery("ip", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.VirtualAddress = c.DefaultQuery("virtual_address", "")
	jsonp.VirtualType = c.DefaultQuery("virtual_type", "")

	jsonp.VirtualPayNo = c.DefaultQuery("virtual_pay_no", "")
	jsonp.VirtualPayAddress = c.DefaultQuery("virtual_pay_address", "")

	jsonp.OrderType = tool.Atoi(c.DefaultQuery("order_type", ""))
	jsonp.VirtualCurrencyChain = c.DefaultQuery("virtual_currency_chain", "")
	jsonp.AnotherPayStatus = tool.Atoi(c.DefaultQuery("another_pay_status", "-1"))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	query := global.G_DB.Model(&dos.FcOrderWithdraw{})

	if len(jsonp.Id) > 0 {
		query = query.Where("id = ?", jsonp.Id)
	}

	if len(jsonp.UserId) > 0 {
		query = query.Where("user_id = ?", jsonp.UserId)
	}

	if len(jsonp.UserName) > 0 {
		query = query.Where("user_name = ?", jsonp.UserName)
	}

	if jsonp.Status > -1 {
		query = query.Where("status = ?", jsonp.Status)
	} else {
		query = query.Where("status in ?", []int{
			enmus.OrderWithdrawStats_AuditWait,
			enmus.OrderWithdrawStats_AuditReject,
		})
	}

	if len(jsonp.Province) > 0 {
		query = query.Where("province = ?", jsonp.Province)
	}

	if len(jsonp.City) > 0 {
		query = query.Where("city = ?", jsonp.City)
	}

	if len(jsonp.BankAddress) > 0 {
		query = query.Where("bank_address = ?", jsonp.BankAddress)
	}

	if len(jsonp.AccountNumber) > 0 {
		query = query.Where("account_number = ?", jsonp.AccountNumber)
	}

	if len(jsonp.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", jsonp.AccountHolder)
	}

	if len(jsonp.AccountBankType) > 0 {
		query = query.Where("account_bank_type = ?", jsonp.AccountBankType)
	}

	if len(jsonp.Remark) > 0 {
		query = query.Where("remark = ?", jsonp.Remark)
	}

	if len(jsonp.Ip) > 0 {
		query = query.Where("ip = ?", jsonp.Ip)
	}

	if !jsonp.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", jsonp.CreateTime)
	}

	if len(jsonp.CreateBy) > 0 {
		query = query.Where("create_by = ?", jsonp.CreateBy)
	}

	if !jsonp.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", jsonp.UpdateTime)
	}

	if len(jsonp.UpdateBy) > 0 {
		query = query.Where("update_by = ?", jsonp.UpdateBy)
	}

	/*if len(jsonp.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", jsonp.MerchantCode)
	}*/

	if len(jsonp.OrderSn) > 0 {
		query = query.Where("order_sn = ?", jsonp.OrderSn)
	}

	if len(jsonp.Currency) > 0 {
		query = query.Where("currency = ?", jsonp.Currency)
	}

	if len(jsonp.VirtualAddress) > 0 {
		query = query.Where("virtual_address = ?", jsonp.VirtualAddress)
	}

	if len(jsonp.VirtualType) > 0 {
		query = query.Where("virtual_type = ?", jsonp.VirtualType)
	}

	if len(jsonp.VirtualPayNo) > 0 {
		query = query.Where("virtual_pay_no = ?", jsonp.VirtualPayNo)
	}

	if len(jsonp.VirtualPayAddress) > 0 {
		query = query.Where("virtual_pay_address = ?", jsonp.VirtualPayAddress)
	}

	if jsonp.OrderType > 0 {
		query = query.Where("order_type = ?", jsonp.OrderType)
	}

	if len(jsonp.VirtualCurrencyChain) > 0 {
		query = query.Where("virtual_currency_chain = ?", jsonp.VirtualCurrencyChain)
	}

	if jsonp.AnotherPayStatus > -1 {
		query = query.Where("another_pay_status = ?", jsonp.AnotherPayStatus)
	}

	if len(jsonp.StartAt) > 0 {
		query = query.Where("create_time >= ?", jsonp.StartAt)
	}

	if len(jsonp.EndAt) > 0 {
		query = query.Where("create_time <= ?", jsonp.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = modules.QueryAdminUserMerchantCodes(c, query, jsonp.MerchantCode); !ok {
			response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, 0, []interface{}{})
			return
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcOrderWithdraw
	query.Order("create_time desc, status").Offset((jsonp.PageNo - 1) * jsonp.PageSize).Limit(jsonp.PageSize).Find(&dataSlice)
	// for _, v := range dataSlice {
	// 	v.Decrypt()
	// }

	result := []dos.FcOrderWithdrawResp{}

	for i, v := range dataSlice {
		v.Decrypt()

		if v.Status == 0 {
			dataSlice[i].UpdateTime = automaticType.Time{}
		}

		withdrawStatus := enmus.OrderWithdrawStats_No
		if v.AnotherPayStatus == enmus.OrderWithdrawAnotherPayStats_Success {
			withdrawStatus = enmus.OrderWithdrawStats_Yes
		}

		if v.OrderType == enmus.ORDER_TYPE_Virtual { // 是钱包要将钱包数据调整
			dataSlice[i].AccountBankType = v.VirtualCurrencyChain
			dataSlice[i].AccountNumber = v.VirtualAddress
		}

		if v.OrderType == enmus.ORDER_TYPE_Online {
			dataSlice[i].AccountBankType = "支付宝"
		}

		result = append(result, dos.FcOrderWithdrawResp{
			FcOrderWithdraw: *dataSlice[i],
			WithdrawStatus:  withdrawStatus,
		})
	}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, count, result)
}

// api: api/fcOrderWithdraw/findByKey
func FindByKeyFcOrderWithdrawControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdraw
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
	if statusStr, ok := c.GetQuery("status"); ok {
		jsonp.Status = tool.Atoi(statusStr)
	} else {
		jsonp.Status = -1
	}
	data := modules.FindByKeyFcOrderWithdraw(&jsonp, c)
	for i, v := range data {
		if v.Status == 0 {
			data[i].UpdateTime = automaticType.Time{}
		}
	}
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdraw/update
func UpdateFcOrderWithdrawControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdraw
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

	orderWithdraw := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if orderWithdraw == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, orderWithdraw.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcOrderWithdraw(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdraw/delete
func DeleteFcOrderWithdrawControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdraw
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

	orderWithdraw := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if orderWithdraw == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, orderWithdraw.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcOrderWithdraw(&jsonp)
	response.SuccessJSON(c, data)
}
