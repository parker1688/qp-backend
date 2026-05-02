// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSiteNotify(vo *dos.FcSiteNotify) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcSiteNotify(page, pageSize int, vo *dos.FcSiteNotify) (ret []*dos.FcSiteNotify, total int64) {
	query := global.G_DB.Model(&dos.FcSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.NotifyType) > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
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

	if vo.ClassType > 0 {
		query = query.Where("class_type = ?", vo.ClassType)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcSiteNotify
	query.Order("sort,create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSiteNotify(vo *dos.FcSiteNotify) []*dos.FcSiteNotify {
	var data []*dos.FcSiteNotify
	query := global.G_DB.Model(&dos.FcSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.NotifyType) > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
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

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcSiteNotifyFirst(vo *dos.FcSiteNotify) *dos.FcSiteNotify {
	var data *dos.FcSiteNotify
	query := global.G_DB.Model(&dos.FcSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.NotifyType) > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
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

	query.Order("sort,create_time desc").Take(&data)
	return data
}

// 根据主键Update
func UpdateFcSiteNotify(vo *dos.FcSiteNotify) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"title":       vo.Title,
		"language":    vo.Language,
		"sort":        vo.Sort,
		"notify_type": vo.NotifyType,
		"update_by":   vo.UpdateBy,
		"title_img":   vo.TitleImg,
		// "merchant_code": vo.MerchantCode,
		"class_type": vo.ClassType,
	}).Error == nil
}

func DeleteFcSiteNotify(vo *dos.FcSiteNotify) bool {
	return global.G_DB.Model(&dos.FcSiteNotify{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func UpdateFcSiteNotifyContent(vo *dos.FcSiteNotify) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"content":   vo.Content,
		"update_by": vo.UpdateBy,
	}).Error == nil
}
