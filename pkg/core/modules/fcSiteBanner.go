// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSiteBanner(vo *dos.FcSiteBanner) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcSiteBanner(page, pageSize int, vo *dos.FcSiteBanner) (ret []*dos.FcSiteBanner, total int64) {
	query := global.G_DB.Model(&dos.FcSiteBanner{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BannerLink) > 0 {
		query = query.Where("banner_link = ?", vo.BannerLink)
	}

	if len(vo.BannerHref) > 0 {
		query = query.Where("banner_href = ?", vo.BannerHref)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BannerType) > 0 {
		query = query.Where("banner_type = ?", vo.BannerType)
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

	if vo.BannerOtherType > 0 {
		query = query.Where("banner_other_type = ?", vo.BannerOtherType)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcSiteBanner
	query.Order("sort desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSiteBanner(vo *dos.FcSiteBanner) []*dos.FcSiteBanner {
	var data []*dos.FcSiteBanner
	query := global.G_DB.Model(&dos.FcSiteBanner{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BannerLink) > 0 {
		query = query.Where("banner_link = ?", vo.BannerLink)
	}

	if len(vo.BannerHref) > 0 {
		query = query.Where("banner_href = ?", vo.BannerHref)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BannerType) > 0 {
		query = query.Where("banner_type = ?", vo.BannerType)
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

	if vo.BannerOtherType > 0 {
		query = query.Where("banner_other_type = ?", vo.BannerOtherType)
	}

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcSiteBannerFirst(vo *dos.FcSiteBanner) *dos.FcSiteBanner {
	var data *dos.FcSiteBanner
	query := global.G_DB.Model(&dos.FcSiteBanner{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.BannerLink) > 0 {
		query = query.Where("banner_link = ?", vo.BannerLink)
	}

	if len(vo.BannerHref) > 0 {
		query = query.Where("banner_href = ?", vo.BannerHref)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BannerType) > 0 {
		query = query.Where("banner_type = ?", vo.BannerType)
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

	if vo.BannerOtherType > 0 {
		query = query.Where("banner_other_type = ?", vo.BannerOtherType)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcSiteBanner(vo *dos.FcSiteBanner) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"banner_link":       vo.BannerLink,
		"banner_href":       vo.BannerHref,
		"language":          vo.Language,
		"sort":              vo.Sort,
		"banner_type":       vo.BannerType,
		"update_by":         vo.UpdateBy,
		"merchant_code":     vo.MerchantCode,
		"banner_other_type": vo.BannerOtherType,
	}).Error == nil
}

func DeleteFcSiteBanner(vo *dos.FcSiteBanner) bool {
	return global.G_DB.Model(&dos.FcSiteBanner{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
