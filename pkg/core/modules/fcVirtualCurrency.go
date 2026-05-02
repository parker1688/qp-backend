// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVirtualCurrency(vo *dos.FcVirtualCurrency) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVirtualCurrency(page, pageSize int, vo *dos.FcVirtualCurrency) (ret []*dos.FcVirtualCurrency, total int64) {
	query := global.G_DB.Model(&dos.FcVirtualCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyNameImg) > 0 {
		query = query.Where("currency_name_img = ?", vo.CurrencyNameImg)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
	}

	if len(vo.CurrencyProtocol) > 0 {
		query = query.Where("currency_protocol = ?", vo.CurrencyProtocol)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVirtualCurrency
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVirtualCurrency(vo *dos.FcVirtualCurrency) []*dos.FcVirtualCurrency {
	var data []*dos.FcVirtualCurrency
	query := global.G_DB.Model(&dos.FcVirtualCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyNameImg) > 0 {
		query = query.Where("currency_name_img = ?", vo.CurrencyNameImg)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
	}

	if len(vo.CurrencyProtocol) > 0 {
		query = query.Where("currency_protocol = ?", vo.CurrencyProtocol)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVirtualCurrencyFirst(vo *dos.FcVirtualCurrency) *dos.FcVirtualCurrency {
	var data *dos.FcVirtualCurrency
	query := global.G_DB.Model(&dos.FcVirtualCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.CurrencyName) > 0 {
		query = query.Where("currency_name = ?", vo.CurrencyName)
	}

	if len(vo.CurrencyNameImg) > 0 {
		query = query.Where("currency_name_img = ?", vo.CurrencyNameImg)
	}

	if len(vo.CurrencyChain) > 0 {
		query = query.Where("currency_chain = ?", vo.CurrencyChain)
	}

	if len(vo.CurrencyProtocol) > 0 {
		query = query.Where("currency_protocol = ?", vo.CurrencyProtocol)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVirtualCurrency(vo *dos.FcVirtualCurrency) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"currency_name":     vo.CurrencyName,
		"currency_name_img": vo.CurrencyNameImg,
		"currency_chain":    vo.CurrencyChain,
		"currency_protocol": vo.CurrencyProtocol,
		"create_by":         vo.CreateBy,
		"update_by":         vo.UpdateBy,
		"status":            vo.Status,
		"merchant_code":     vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVirtualCurrency(vo *dos.FcVirtualCurrency) bool {
	return global.G_DB.Model(&dos.FcVirtualCurrency{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
