// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcPayment(vo *dos.FcPayment) (bool, error) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func FindPageFcPayment(page, pageSize int, vo *dos.FcPayment, c *gin.Context) (ret []*dos.FcPaymentExt, total int64) {
	query := global.G_DB.Model(&dos.FcPayment{})

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

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcPaymentExt
	query.Offset((page - 1) * pageSize).Order("sort asc").Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPayment(vo *dos.FcPayment, c *gin.Context) []*dos.FcPayment {
	var data []*dos.FcPayment
	query := global.G_DB.Model(&dos.FcPayment{})

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

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

func FindByKeyFcPaymentFirst(vo *dos.FcPayment) *dos.FcPayment {
	var data *dos.FcPayment
	query := global.G_DB.Model(&dos.FcPayment{})

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

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcPayment(vo *dos.FcPayment) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"payment_name":   vo.PaymentName,
		"pay_alias_name": vo.PayAliasName,
		"payment_code":   vo.PaymentCode,
		"pay_id":         vo.PayId,
		"channel_name":   vo.ChannelName,
		"channel_code":   vo.ChannelCode,
		"status":         vo.Status,
		"min_level":      vo.MinLevel,
		"max_level":      vo.MaxLevel,
		"min_amount":     vo.MinAmount,
		"max_amount":     vo.MaxAmount,
		"day_max_amount": vo.DayMaxAmount,
		"bonus_rate":     vo.BonusRate,
		"sort":           vo.Sort,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"update_time":    vo.UpdateTime,
		"merchant_code":  vo.MerchantCode,
		"fee_rate":       vo.FeeRate,
		"amount_range":   vo.AmountRange,
		"remark":         vo.Remark,
	}).Error == nil
}

func DeleteFcPayment(vo *dos.FcPayment) bool {
	return global.G_DB.Model(&dos.FcPayment{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
