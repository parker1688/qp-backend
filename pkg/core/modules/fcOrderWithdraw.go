// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"

	"github.com/gin-gonic/gin"
)

func SaveFcOrderWithdraw(vo *dos.FcOrderWithdraw) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcOrderWithdraw(page, pageSize int, vo *dos.FcOrderWithdraw, pageQuery response.PageTimeQuery) (ret []*dos.FcOrderWithdraw, total int64) {
	query := global.G_DB.Model(&dos.FcOrderWithdraw{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Province) > 0 {
		query = query.Where("province = ?", vo.Province)
	}

	if len(vo.City) > 0 {
		query = query.Where("city = ?", vo.City)
	}

	if len(vo.BankAddress) > 0 {
		query = query.Where("bank_address = ?", vo.BankAddress)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if len(vo.AccountBankType) > 0 {
		query = query.Where("account_bank_type = ?", vo.AccountBankType)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.VirtualAddress) > 0 {
		query = query.Where("virtual_address = ?", vo.VirtualAddress)
	}

	if len(vo.VirtualType) > 0 {
		query = query.Where("virtual_type = ?", vo.VirtualType)
	}

	if len(vo.VirtualPayNo) > 0 {
		query = query.Where("virtual_pay_no = ?", vo.VirtualPayNo)
	}

	if len(vo.VirtualPayAddress) > 0 {
		query = query.Where("virtual_pay_address = ?", vo.VirtualPayAddress)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.VirtualCurrencyChain) > 0 {
		query = query.Where("virtual_currency_chain = ?", vo.VirtualCurrencyChain)
	}

	if vo.AnotherPayStatus > -1 {
		query = query.Where("another_pay_status = ?", vo.AnotherPayStatus)
	}

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
	}

	if len(pageQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcOrderWithdraw
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcOrderWithdraw(vo *dos.FcOrderWithdraw, c *gin.Context) []*dos.FcOrderWithdraw {
	var data []*dos.FcOrderWithdraw
	query := global.G_DB.Model(&dos.FcOrderWithdraw{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Province) > 0 {
		query = query.Where("province = ?", vo.Province)
	}

	if len(vo.City) > 0 {
		query = query.Where("city = ?", vo.City)
	}

	if len(vo.BankAddress) > 0 {
		query = query.Where("bank_address = ?", vo.BankAddress)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if len(vo.AccountBankType) > 0 {
		query = query.Where("account_bank_type = ?", vo.AccountBankType)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.VirtualAddress) > 0 {
		query = query.Where("virtual_address = ?", vo.VirtualAddress)
	}

	if len(vo.VirtualType) > 0 {
		query = query.Where("virtual_type = ?", vo.VirtualType)
	}

	if len(vo.VirtualPayNo) > 0 {
		query = query.Where("virtual_pay_no = ?", vo.VirtualPayNo)
	}

	if len(vo.VirtualPayAddress) > 0 {
		query = query.Where("virtual_pay_address = ?", vo.VirtualPayAddress)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.VirtualCurrencyChain) > 0 {
		query = query.Where("virtual_currency_chain = ?", vo.VirtualCurrencyChain)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcOrderWithdrawFirst(vo *dos.FcOrderWithdraw) *dos.FcOrderWithdraw {
	data := &dos.FcOrderWithdraw{}
	query := global.G_DB.Model(&dos.FcOrderWithdraw{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Province) > 0 {
		query = query.Where("province = ?", vo.Province)
	}

	if len(vo.City) > 0 {
		query = query.Where("city = ?", vo.City)
	}

	if len(vo.BankAddress) > 0 {
		query = query.Where("bank_address = ?", vo.BankAddress)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if len(vo.AccountBankType) > 0 {
		query = query.Where("account_bank_type = ?", vo.AccountBankType)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.VirtualAddress) > 0 {
		query = query.Where("virtual_address = ?", vo.VirtualAddress)
	}

	if len(vo.VirtualType) > 0 {
		query = query.Where("virtual_type = ?", vo.VirtualType)
	}

	if len(vo.VirtualPayNo) > 0 {
		query = query.Where("virtual_pay_no = ?", vo.VirtualPayNo)
	}

	if len(vo.VirtualPayAddress) > 0 {
		query = query.Where("virtual_pay_address = ?", vo.VirtualPayAddress)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.VirtualCurrencyChain) > 0 {
		query = query.Where("virtual_currency_chain = ?", vo.VirtualCurrencyChain)
	}

	err := query.Take(data).Error
	if err != nil {
		return nil
	}
	return data
}

// 根据主键Update
func UpdateFcOrderWithdraw(vo *dos.FcOrderWithdraw) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":                vo.UserId,
		"user_name":              vo.UserName,
		"amount":                 vo.Amount,
		"status":                 vo.Status,
		"province":               vo.Province,
		"city":                   vo.City,
		"bank_address":           vo.BankAddress,
		"account_number":         vo.AccountNumber,
		"account_holder":         vo.AccountHolder,
		"account_bank_type":      vo.AccountBankType,
		"remark":                 vo.Remark,
		"ip":                     vo.Ip,
		"create_by":              vo.CreateBy,
		"update_by":              vo.UpdateBy,
		"merchant_code":          vo.MerchantCode,
		"order_sn":               vo.OrderSn,
		"currency":               vo.Currency,
		"virtual_address":        vo.VirtualAddress,
		"virtual_type":           vo.VirtualType,
		"virtual_num":            vo.VirtualNum,
		"virtual_fx":             vo.VirtualFx,
		"virtual_pay_no":         vo.VirtualPayNo,
		"virtual_pay_address":    vo.VirtualPayAddress,
		"virtual_pay_amount":     vo.VirtualPayAmount,
		"order_type":             vo.OrderType,
		"virtual_currency_chain": vo.VirtualCurrencyChain,
	}).Error == nil
}

func DeleteFcOrderWithdraw(vo *dos.FcOrderWithdraw) bool {
	return global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// IsWithdrawOrderTodo - 是否提款订单处理中
// @param {string} userId
// @returns bool
func IsWithdrawOrderTodo(userId string) bool {
	var count int64
	err := global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("user_id = ? AND status = ?",
		userId, enmus.OrderWithdrawStats_AuditWait).Count(&count).Error
	if err != nil {
		global.G_LOG.Errorf("[IsWithdrawOrderTodo] Find order withdraw count failed: userId=%s, err=%s",
			userId, err.Error())
		return false
	}

	if count > 0 {
		return true // 存在待审核单子
	}

	count = 0

	err = global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where("user_id = ? AND status in ? AND withdraw_status = ?",
		userId,
		[]int{
			enmus.OrderWithdrawPaymentOutStats_Prepare,
			enmus.OrderWithdrawPaymentOutStats_Progress,
		},
		enmus.OrderWithdrawStats_No,
	).Count(&count).Error
	if err != nil {
		global.G_LOG.Errorf("[IsWithdrawOrderTodo] Find order withdraw payment out count failed: userId=%s, err=%s",
			userId, err.Error())
		return false
	}

	if count > 0 {
		return true // 存在未处理完的单子
	}

	return false
}
