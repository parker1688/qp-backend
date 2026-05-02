// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserSiteNotify(vo *dos.FcUserSiteNotify) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserSiteNotify(page, pageSize int, vo *dos.FcUserSiteNotify) (ret []*dos.FcUserSiteNotify, total int64) {
	query := global.G_DB.Model(&dos.FcUserSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.NotifyId) > 0 {
		query = query.Where("notify_id = ?", vo.NotifyId)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserSiteNotify
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserSiteNotify(vo *dos.FcUserSiteNotify) []*dos.FcUserSiteNotify {
	var data []*dos.FcUserSiteNotify
	query := global.G_DB.Model(&dos.FcUserSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.NotifyId) > 0 {
		query = query.Where("notify_id = ?", vo.NotifyId)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserSiteNotifyFirst(vo *dos.FcUserSiteNotify) *dos.FcUserSiteNotify {
	var data *dos.FcUserSiteNotify
	query := global.G_DB.Model(&dos.FcUserSiteNotify{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.NotifyId) > 0 {
		query = query.Where("notify_id = ?", vo.NotifyId)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserSiteNotify(vo *dos.FcUserSiteNotify) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":     vo.UserId,
		"notify_id":   vo.NotifyId,
		"title":       vo.Title,
		"content":     vo.Content,
		"language":    vo.Language,
		"sort":        vo.Sort,
		"notify_type": vo.NotifyType,
		//"create_by":     vo.CreateBy,
		"update_by": vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
		"class_type": vo.ClassType,
		"status":     vo.Status,
	}).Error == nil
}

func DeleteFcUserSiteNotify(vo *dos.FcUserSiteNotify) bool {
	return global.G_DB.Model(&dos.FcUserSiteNotify{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
