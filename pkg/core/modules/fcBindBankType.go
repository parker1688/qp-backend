// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcBindBankType(vo *dos.FcBindBankType) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcBindBankType(page, pageSize int, vo *dos.FcBindBankType, c *gin.Context) (ret []*dos.FcBindBankType, total int64) {
	query := global.G_DB.Model(&dos.FcBindBankType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BankImg) > 0 {
		query = query.Where("bank_img = ?", vo.BankImg)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}
	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBindBankType
	query.Order("sort desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBindBankType(vo *dos.FcBindBankType, c *gin.Context) []*dos.FcBindBankType {
	var data []*dos.FcBindBankType
	query := global.G_DB.Model(&dos.FcBindBankType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BankImg) > 0 {
		query = query.Where("bank_img = ?", vo.BankImg)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}
	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcBindBankTypeFirst(vo *dos.FcBindBankType) *dos.FcBindBankType {
	var data *dos.FcBindBankType
	query := global.G_DB.Model(&dos.FcBindBankType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BankName) > 0 {
		query = query.Where("bank_name = ?", vo.BankName)
	}

	if len(vo.BankCode) > 0 {
		query = query.Where("bank_code = ?", vo.BankCode)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BankImg) > 0 {
		query = query.Where("bank_img = ?", vo.BankImg)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcBindBankType(vo *dos.FcBindBankType) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"bank_name": vo.BankName,
		"bank_code": vo.BankCode,
		"sort":      vo.Sort,
		"bank_img":  vo.BankImg,
		"status":    vo.Status,
		"currency":  vo.Currency,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcBindBankType(vo *dos.FcBindBankType) bool {
	return global.G_DB.Model(&dos.FcBindBankType{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
