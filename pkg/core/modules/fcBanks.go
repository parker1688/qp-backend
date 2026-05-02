// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcBanks(vo *dos.FcBanks) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcBanks(page, pageSize int, vo *dos.FcBanks) (ret []*dos.FcBanks, total int64) {
	query := global.G_DB.Model(&dos.FcBanks{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBanks
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBanks(vo *dos.FcBanks) []*dos.FcBanks {
	var data []*dos.FcBanks
	query := global.G_DB.Model(&dos.FcBanks{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcBanksFirst(vo *dos.FcBanks) *dos.FcBanks {
	var data *dos.FcBanks
	query := global.G_DB.Model(&dos.FcBanks{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
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
func UpdateFcBanks(vo *dos.FcBanks) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"bank_name": vo.BankName,
		"bank_code": vo.BankCode,
		"min_level": vo.MinLevel,
		"max_level": vo.MaxLevel,
		"sort":      vo.Sort,
		"update_by": vo.UpdateBy,
		"status":    vo.Status,
		"currency":  vo.Currency,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcBanks(vo *dos.FcBanks) bool {
	return global.G_DB.Model(&dos.FcBanks{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
