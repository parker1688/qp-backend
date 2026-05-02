// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcBanksDetails(vo *dos.FcBanksDetails) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcBanksDetails(page, pageSize int, vo *dos.FcBanksDetails) (ret []*dos.FcBanksDetails, total int64) {
	query := global.G_DB.Model(&dos.FcBanksDetails{})

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

	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBanksDetails
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBanksDetails(vo *dos.FcBanksDetails) []*dos.FcBanksDetails {
	var data []*dos.FcBanksDetails
	query := global.G_DB.Model(&dos.FcBanksDetails{})

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

	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcBanksDetailsFirst(vo *dos.FcBanksDetails) *dos.FcBanksDetails {
	var data *dos.FcBanksDetails
	query := global.G_DB.Model(&dos.FcBanksDetails{})

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

	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
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
func UpdateFcBanksDetails(vo *dos.FcBanksDetails) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"bank_name":                vo.BankName,
		"bank_code":                vo.BankCode,
		"min_level":                vo.MinLevel,
		"max_level":                vo.MaxLevel,
		"sort":                     vo.Sort,
		"update_by":                vo.UpdateBy,
		"entity_account_holder":    vo.EntityAccountHolder,
		"entity_account_bank_name": vo.EntityAccountBankName,
		"entity_account_number":    vo.EntityAccountNumber,
		"status":                   vo.Status,
		"day_max_amount":           vo.DayMaxAmount,
		"min_amount":               vo.MinAmount,
		"max_amount":               vo.MaxAmount,
		"merchant_code":            vo.MerchantCode,
	}).Error == nil
}

func DeleteFcBanksDetails(vo *dos.FcBanksDetails) bool {
	return global.G_DB.Model(&dos.FcBanksDetails{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
