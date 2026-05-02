// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSiteLink(vo *dos.FcSiteLink) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcSiteLink(page, pageSize int, vo *dos.FcSiteLink) (ret []*dos.FcSiteLink, total int64) {
	query := global.G_DB.Model(&dos.FcSiteLink{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AppKey) > 0 {
		query = query.Where("app_key = ?", vo.AppKey)
	}

	if len(vo.AppLink) > 0 {
		query = query.Where("app_link = ?", vo.AppLink)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	//if len(vo.Domain) > 0 {
	//	query = query.Where("domain = ?", vo.Domain)
	//}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcSiteLink
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSiteLink(vo *dos.FcSiteLink) []*dos.FcSiteLink {
	var data []*dos.FcSiteLink
	query := global.G_DB.Model(&dos.FcSiteLink{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AppKey) > 0 {
		query = query.Where("app_key = ?", vo.AppKey)
	}

	if len(vo.AppLink) > 0 {
		query = query.Where("app_link = ?", vo.AppLink)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	//if len(vo.Domain) > 0 {
	//	query = query.Where("domain = ?", vo.Domain)
	//}

	query.Find(&data)
	return data
}

func FindByKeyFcSiteLinkFirst(vo *dos.FcSiteLink) *dos.FcSiteLink {
	var data *dos.FcSiteLink
	query := global.G_DB.Model(&dos.FcSiteLink{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AppKey) > 0 {
		query = query.Where("app_key = ?", vo.AppKey)
	}

	if len(vo.AppLink) > 0 {
		query = query.Where("app_link = ?", vo.AppLink)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	//if len(vo.Domain) > 0 {
	//	query = query.Where("domain = ?", vo.Domain)
	//}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcSiteLink(vo *dos.FcSiteLink) bool {
	u := map[string]interface{}{
		"app_key":        vo.AppKey,
		"app_link":       vo.AppLink,
		"content":        vo.Content,
		"update_by":      vo.UpdateBy,
		"img_link":       vo.ImgLink,
		"content_detail": vo.ContentDetail,
		"download_link":  vo.DownloadLink,
		//"domain":    vo.Domain,
	}
	if len(vo.MerchantCode) > 0 {
		u["merchant_code"] = vo.MerchantCode
	}
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(u).Error == nil
}

func DeleteFcSiteLink(vo *dos.FcSiteLink) bool {
	return global.G_DB.Model(&dos.FcSiteLink{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
