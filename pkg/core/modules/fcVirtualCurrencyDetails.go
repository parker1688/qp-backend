// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVirtualCurrencyDetails(vo *dos.FcVirtualCurrencyDetails) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVirtualCurrencyDetails(page, pageSize int, vo *dos.FcVirtualCurrencyDetails) (ret []*dos.FcVirtualCurrencyDetails, total int64) {
	query := global.G_DB.Model(&dos.FcVirtualCurrencyDetails{})

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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ToAddr) > 0 {
		query = query.Where("to_addr = ?", vo.ToAddr)
	}

	if len(vo.ToAddrQrPre) > 0 {
		query = query.Where("to_addr_qr_pre = ?", vo.ToAddrQrPre)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVirtualCurrencyDetails
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVirtualCurrencyDetails(vo *dos.FcVirtualCurrencyDetails) []*dos.FcVirtualCurrencyDetails {
	var data []*dos.FcVirtualCurrencyDetails
	query := global.G_DB.Model(&dos.FcVirtualCurrencyDetails{})

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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ToAddr) > 0 {
		query = query.Where("to_addr = ?", vo.ToAddr)
	}

	if len(vo.ToAddrQrPre) > 0 {
		query = query.Where("to_addr_qr_pre = ?", vo.ToAddrQrPre)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVirtualCurrencyDetailsFirst(vo *dos.FcVirtualCurrencyDetails) *dos.FcVirtualCurrencyDetails {
	var data *dos.FcVirtualCurrencyDetails
	query := global.G_DB.Model(&dos.FcVirtualCurrencyDetails{})

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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ToAddr) > 0 {
		query = query.Where("to_addr = ?", vo.ToAddr)
	}

	if len(vo.ToAddrQrPre) > 0 {
		query = query.Where("to_addr_qr_pre = ?", vo.ToAddrQrPre)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVirtualCurrencyDetails(vo *dos.FcVirtualCurrencyDetails) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"currency_name":  vo.CurrencyName,
		"currency_chain": vo.CurrencyChain,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"status":         vo.Status,
		"to_addr":        vo.ToAddr,
		"to_addr_qr_pre": vo.ToAddrQrPre,
		"merchant_code":  vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVirtualCurrencyDetails(vo *dos.FcVirtualCurrencyDetails) bool {
	return global.G_DB.Model(&dos.FcVirtualCurrencyDetails{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
