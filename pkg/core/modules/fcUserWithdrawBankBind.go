// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserWithdrawBankBind(vo *dos.FcUserWithdrawBankBind) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserWithdrawBankBind(page, pageSize int, vo *dos.FcUserWithdrawBankBind, c *gin.Context) (ret []*dos.FcUserWithdrawBankBind, total int64) {
	query := global.G_DB.Model(&dos.FcUserWithdrawBankBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	} else {
		return ret, total
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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
	var dataSlice []*dos.FcUserWithdrawBankBind
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserWithdrawBankBind(vo *dos.FcUserWithdrawBankBind) []*dos.FcUserWithdrawBankBind {
	var data []*dos.FcUserWithdrawBankBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBankBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

func FindByKeyFcUserWithdrawBankBind2(vo *dos.FcUserWithdrawBankBind, c *gin.Context) []*dos.FcUserWithdrawBankBind {
	var data []*dos.FcUserWithdrawBankBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBankBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserWithdrawBankBindFirst(vo *dos.FcUserWithdrawBankBind) *dos.FcUserWithdrawBankBind {
	var data *dos.FcUserWithdrawBankBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBankBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

func FindByUsersFcUserWithdrawBankBind(userId []string) []*dos.FcUserWithdrawBankBind {
	var data []*dos.FcUserWithdrawBankBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBankBind{})
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserWithdrawBankBind(vo *dos.FcUserWithdrawBankBind) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"account_bank_type": vo.AccountBankType,
		"account_bank_code": vo.AccountBankCode,
	}).Error == nil
}

func DeleteFcUserWithdrawBankBind(vo *dos.FcUserWithdrawBankBind) bool {
	return global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
