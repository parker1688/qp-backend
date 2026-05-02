// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcPaymentSetting(vo *dos.FcPaymentSetting) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcPaymentSetting(page, pageSize int, vo *dos.FcPaymentSetting) (ret []*dos.FcPaymentSetting, total int64) {
	query := global.G_DB.Model(&dos.FcPaymentSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PKey) > 0 {
		query = query.Where("p_key = ?", vo.PKey)
	}

	if len(vo.PValue) > 0 {
		query = query.Where("p_value = ?", vo.PValue)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcPaymentSetting
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPaymentSetting(vo *dos.FcPaymentSetting) []*dos.FcPaymentSetting {
	var data []*dos.FcPaymentSetting
	query := global.G_DB.Model(&dos.FcPaymentSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PKey) > 0 {
		query = query.Where("p_key = ?", vo.PKey)
	}

	if len(vo.PValue) > 0 {
		query = query.Where("p_value = ?", vo.PValue)
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

	query.Find(&data)
	return data
}

func FindByKeyFcPaymentSettingFirst(vo *dos.FcPaymentSetting) *dos.FcPaymentSetting {
	var data *dos.FcPaymentSetting
	query := global.G_DB.Model(&dos.FcPaymentSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.PKey) > 0 {
		query = query.Where("p_key = ?", vo.PKey)
	}

	if len(vo.PValue) > 0 {
		query = query.Where("p_value = ?", vo.PValue)
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
func UpdateFcPaymentSetting(vo *dos.FcPaymentSetting) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"payment_code": vo.PaymentCode,
		"p_key":        vo.PKey,
		"p_value":      vo.PValue,
		"sort":         vo.Sort,
		//"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"merchant_code": vo.MerchantCode,
		"remark":        vo.Remark,
	}).Error == nil
}

func DeleteFcPaymentSetting(vo *dos.FcPaymentSetting) bool {
	return global.G_DB.Model(&dos.FcPaymentSetting{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
