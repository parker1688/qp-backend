// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcFirstOrderDeposit(vo *dos.FcFirstOrderDeposit) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcFirstOrderDeposit(page, pageSize int, vo *dos.FcFirstOrderDeposit) (ret []*dos.FcFirstOrderDeposit, total int64) {
	query := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if !vo.PayTime.Timer().IsZero() {
		query = query.Where("pay_time = ?", vo.PayTime)
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

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayAliasName) > 0 {
		query = query.Where("pay_alias_name = ?", vo.PayAliasName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if vo.OrderSecondType > 0 {
		query = query.Where("order_second_type = ?", vo.OrderSecondType)
	}

	if len(vo.AuthBy) > 0 {
		query = query.Where("auth_by = ?", vo.AuthBy)
	}

	if !vo.AuthTime.Timer().IsZero() {
		query = query.Where("auth_time = ?", vo.AuthTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcFirstOrderDeposit
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcFirstOrderDeposit(vo *dos.FcFirstOrderDeposit) []*dos.FcFirstOrderDeposit {
	var data []*dos.FcFirstOrderDeposit
	query := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if !vo.PayTime.Timer().IsZero() {
		query = query.Where("pay_time = ?", vo.PayTime)
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

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayAliasName) > 0 {
		query = query.Where("pay_alias_name = ?", vo.PayAliasName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if vo.OrderSecondType > 0 {
		query = query.Where("order_second_type = ?", vo.OrderSecondType)
	}

	if len(vo.AuthBy) > 0 {
		query = query.Where("auth_by = ?", vo.AuthBy)
	}

	if !vo.AuthTime.Timer().IsZero() {
		query = query.Where("auth_time = ?", vo.AuthTime)
	}
	query.Find(&data)
	return data
}

func FindByKeyFcFirstOrderDepositFirst(vo *dos.FcFirstOrderDeposit) *dos.FcFirstOrderDeposit {
	var data *dos.FcFirstOrderDeposit
	query := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if !vo.PayTime.Timer().IsZero() {
		query = query.Where("pay_time = ?", vo.PayTime)
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

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayAliasName) > 0 {
		query = query.Where("pay_alias_name = ?", vo.PayAliasName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if vo.OrderSecondType > 0 {
		query = query.Where("order_second_type = ?", vo.OrderSecondType)
	}

	if len(vo.AuthBy) > 0 {
		query = query.Where("auth_by = ?", vo.AuthBy)
	}

	if !vo.AuthTime.Timer().IsZero() {
		query = query.Where("auth_time = ?", vo.AuthTime)
	}
	query.Take(&data)
	return data
}

//根据主键Update
func UpdateFcFirstOrderDeposit(vo *dos.FcFirstOrderDeposit) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":           vo.UserId,
		"user_name":         vo.UserName,
		"order_sn":          vo.OrderSn,
		"amount":            vo.Amount,
		"fact_amount":       vo.FactAmount,
		"bonus_amount":      vo.BonusAmount,
		"bonus_rate":        vo.BonusRate,
		"remark":            vo.Remark,
		"deposit_remark":    vo.DepositRemark,
		"ip":                vo.Ip,
		"level":             vo.Level,
		"pay_time":          vo.PayTime,
		"create_by":         vo.CreateBy,
		"update_by":         vo.UpdateBy,
		"merchant_code":     vo.MerchantCode,
		"channel_id":        vo.ChannelId,
		"channel_code":      vo.ChannelCode,
		"payment_id":        vo.PaymentId,
		"payment_code":      vo.PaymentCode,
		"pay_alias_name":    vo.PayAliasName,
		"payment_name":      vo.PaymentName,
		"currency":          vo.Currency,
		"order_type":        vo.OrderType,
		"order_second_type": vo.OrderSecondType,
		"fee_rate":          vo.FeeRate,
		"fee":               vo.Fee,
		"auth_by":           vo.AuthBy,
		"auth_time":         vo.AuthTime,
	}).Error == nil
}

func DeleteFcFirstOrderDeposit(vo *dos.FcFirstOrderDeposit) bool {
	return global.G_DB.Model(&dos.FcFirstOrderDeposit{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
