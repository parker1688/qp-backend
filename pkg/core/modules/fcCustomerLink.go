// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcCustomerLink(vo *dos.FcCustomerLink) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcCustomerLink(page, pageSize int, vo *dos.FcCustomerLink, c *gin.Context) (ret []*dos.FcCustomerLink, total int64) {
	query := global.G_DB.Model(&dos.FcCustomerLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
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
	var dataSlice []*dos.FcCustomerLink
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcCustomerLink(vo *dos.FcCustomerLink, c *gin.Context) []*dos.FcCustomerLink {
	var data []*dos.FcCustomerLink
	query := global.G_DB.Model(&dos.FcCustomerLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
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

	query.Find(&data)
	return data
}

func FindByKeyFcCustomerLinkFirst(vo *dos.FcCustomerLink) *dos.FcCustomerLink {
	var data *dos.FcCustomerLink
	query := global.G_DB.Model(&dos.FcCustomerLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcCustomerLink(vo *dos.FcCustomerLink) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"name":          vo.Name,
		"link":          vo.Link,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"merchant_code": vo.MerchantCode,
		"merchant_name": vo.MerchantName,
		"status":        vo.Status,
	}).Error == nil
}

func DeleteFcCustomerLink(vo *dos.FcCustomerLink) bool {
	return global.G_DB.Model(&dos.FcCustomerLink{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
