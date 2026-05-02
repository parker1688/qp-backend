// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserWithdrawBlockchainBind(vo *dos.FcUserWithdrawBlockchainBind) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserWithdrawBlockchainBind(page, pageSize int, vo *dos.FcUserWithdrawBlockchainBind, c *gin.Context) (ret []*dos.FcUserWithdrawBlockchainBind, total int64) {
	query := global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.BlockchainAddress) > 0 {
		query = query.Where("blockchain_address = ?", vo.BlockchainAddress)
	}
	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}
	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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
	var dataSlice []*dos.FcUserWithdrawBlockchainBind
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserWithdrawBlockchainBind(vo *dos.FcUserWithdrawBlockchainBind) []*dos.FcUserWithdrawBlockchainBind {
	var data []*dos.FcUserWithdrawBlockchainBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}
	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}
	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.BlockchainAddress) > 0 {
		query = query.Where("blockchain_address = ?", vo.BlockchainAddress)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

func FindByKeyFcUserWithdrawBlockchainBind2(vo *dos.FcUserWithdrawBlockchainBind, c *gin.Context) []*dos.FcUserWithdrawBlockchainBind {
	var data []*dos.FcUserWithdrawBlockchainBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}
	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}
	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.BlockchainAddress) > 0 {
		query = query.Where("blockchain_address = ?", vo.BlockchainAddress)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserWithdrawBlockchainBindFirst(vo *dos.FcUserWithdrawBlockchainBind) *dos.FcUserWithdrawBlockchainBind {
	var data *dos.FcUserWithdrawBlockchainBind
	query := global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.BlockchainAddress) > 0 {
		query = query.Where("blockchain_address = ?", vo.BlockchainAddress)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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
	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserWithdrawBlockchainBind(vo *dos.FcUserWithdrawBlockchainBind) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"blockchain":         vo.Blockchain,
		"blockchain_address": vo.BlockchainAddress,
		"contract_type":      vo.ContractType,
		"is_default":         vo.IsDefault,
		"payment_code":       vo.PaymentCode,
		"payment_name":       vo.PaymentName,
		"sort":               vo.Sort,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
		"merchant_code":      vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserWithdrawBlockchainBind(vo *dos.FcUserWithdrawBlockchainBind) bool {
	return global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
