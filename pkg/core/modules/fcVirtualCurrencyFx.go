// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVirtualCurrencyFx(vo *dos.FcVirtualCurrencyFx) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVirtualCurrencyFx(page, pageSize int, vo *dos.FcVirtualCurrencyFx) (ret []*dos.FcVirtualCurrencyFx, total int64) {
	query := global.G_DB.Model(&dos.FcVirtualCurrencyFx{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
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

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.CurrencyCode) > 0 {
		query = query.Where("currency_code = ?", vo.CurrencyCode)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVirtualCurrencyFx
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVirtualCurrencyFx(vo *dos.FcVirtualCurrencyFx) []*dos.FcVirtualCurrencyFx {
	var data []*dos.FcVirtualCurrencyFx
	query := global.G_DB.Model(&dos.FcVirtualCurrencyFx{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
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

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.CurrencyCode) > 0 {
		query = query.Where("currency_code = ?", vo.CurrencyCode)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVirtualCurrencyFxFirst(vo *dos.FcVirtualCurrencyFx) *dos.FcVirtualCurrencyFx {
	var data *dos.FcVirtualCurrencyFx
	query := global.G_DB.Model(&dos.FcVirtualCurrencyFx{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
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

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.CurrencyCode) > 0 {
		query = query.Where("currency_code = ?", vo.CurrencyCode)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVirtualCurrencyFx(vo *dos.FcVirtualCurrencyFx) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"currency_name":  vo.CurrencyName,
		"currency_chain": vo.CurrencyChain,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"fx_amount":      vo.FxAmount,
		"opt_type":       vo.OptType,
		"currency_code":  vo.CurrencyCode,
		"merchant_code":  vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVirtualCurrencyFx(vo *dos.FcVirtualCurrencyFx) bool {
	return global.G_DB.Model(&dos.FcVirtualCurrencyFx{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
