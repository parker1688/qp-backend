// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcPaymentSum(vo *dos.FcPaymentSum) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcPaymentSum(page, pageSize int, vo *dos.FcPaymentSum) (ret []*dos.FcPaymentSum, total int64) {
	query := global.G_DB.Model(&dos.FcPaymentSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayId) > 0 {
		query = query.Where("pay_id = ?", vo.PayId)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcPaymentSum
	query.Offset((page - 1) * pageSize).Order("sort asc").Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPaymentSum(vo *dos.FcPaymentSum) []*dos.FcPaymentSum {
	var data []*dos.FcPaymentSum
	query := global.G_DB.Model(&dos.FcPaymentSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayId) > 0 {
		query = query.Where("pay_id = ?", vo.PayId)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcPaymentSumFirst(vo *dos.FcPaymentSum) *dos.FcPaymentSum {
	var data *dos.FcPaymentSum
	query := global.G_DB.Model(&dos.FcPaymentSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PayId) > 0 {
		query = query.Where("pay_id = ?", vo.PayId)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcPaymentSum(vo *dos.FcPaymentSum) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"payment_name":   vo.PaymentName,
		"pay_alias_name": vo.PayAliasName,
		"payment_code":   vo.PaymentCode,
		"pay_id":         vo.PayId,
		"channel_name":   vo.ChannelName,
		"channel_code":   vo.ChannelCode,
		"status":         vo.Status,
		"min_amount":     vo.MinAmount,
		"max_amount":     vo.MaxAmount,
		"day_max_amount": vo.DayMaxAmount,
		"bonus_rate":     vo.BonusRate,
		"sort":           vo.Sort,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"update_time":    vo.UpdateTime,
		"fee_rate":       vo.FeeRate,
		"amount_range":   vo.AmountRange,
		"min_level":      vo.MinLevel,
		"max_level":      vo.MaxLevel,
		"remark":         vo.Remark,
	}).Error == nil
}

func DeleteFcPaymentSum(vo *dos.FcPaymentSum) bool {
	return global.G_DB.Model(&dos.FcPaymentSum{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
