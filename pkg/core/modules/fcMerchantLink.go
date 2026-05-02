// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcMerchantLink(vo *dos.FcMerchantLink) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcMerchantLink(page, pageSize int, vo *dos.FcMerchantLink, c *gin.Context) (ret []*dos.FcMerchantLink, total int64) {
	query := global.G_DB.Model(&dos.FcMerchantLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcMerchantLink
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcMerchantLink(vo *dos.FcMerchantLink, c *gin.Context) []*dos.FcMerchantLink {
	var data []*dos.FcMerchantLink
	query := global.G_DB.Model(&dos.FcMerchantLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcMerchantLinkFirst(vo *dos.FcMerchantLink) *dos.FcMerchantLink {
	var data *dos.FcMerchantLink
	query := global.G_DB.Model(&dos.FcMerchantLink{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
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

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcMerchantLink(vo *dos.FcMerchantLink) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"link":          vo.Link,
		"alias":         vo.Alias,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"merchant_code": vo.MerchantCode,
		"merchant_name": vo.MerchantName,
		"logo_img":      vo.LogoImg,
		"banner_img":    vo.BannerImg,
	}).Error == nil
}

func DeleteFcMerchantLink(vo *dos.FcMerchantLink) bool {
	return global.G_DB.Model(&dos.FcMerchantLink{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func GetFcMerchantLinkData(merchantCode string) dos.FcMerchantLink {
	data := dos.FcMerchantLink{}
	err := global.G_DB.Model(&dos.FcMerchantLink{}).Where("merchant_code = ?", merchantCode).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcMerchantLinkData] Find merchant link failed: merchant_code=%s err=%v",
			merchantCode, err.Error())
		return dos.FcMerchantLink{}
	}

	return data
}
