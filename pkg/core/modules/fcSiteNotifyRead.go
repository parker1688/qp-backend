// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSiteNotifyRead(vo *dos.FcSiteNotifyRead) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcSiteNotifyRead(page, pageSize int, vo *dos.FcSiteNotifyRead) (ret []*dos.FcSiteNotifyRead, total int64) {
	query := global.G_DB.Model(&dos.FcSiteNotifyRead{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SiteNotifyId) > 0 {
		query = query.Where("site_notify_id = ?", vo.SiteNotifyId)
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

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcSiteNotifyRead
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSiteNotifyRead(vo *dos.FcSiteNotifyRead) []*dos.FcSiteNotifyRead {
	var data []*dos.FcSiteNotifyRead
	query := global.G_DB.Model(&dos.FcSiteNotifyRead{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SiteNotifyId) > 0 {
		query = query.Where("site_notify_id = ?", vo.SiteNotifyId)
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

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcSiteNotifyReadFirst(vo *dos.FcSiteNotifyRead) *dos.FcSiteNotifyRead {
	var data *dos.FcSiteNotifyRead
	query := global.G_DB.Model(&dos.FcSiteNotifyRead{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SiteNotifyId) > 0 {
		query = query.Where("site_notify_id = ?", vo.SiteNotifyId)
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

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcSiteNotifyRead(vo *dos.FcSiteNotifyRead) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"site_notify_id": vo.SiteNotifyId,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"user_id":        vo.UserId,
		"status":         vo.Status,
	}).Error == nil
}

func DeleteFcSiteNotifyRead(vo *dos.FcSiteNotifyRead) bool {
	return global.G_DB.Model(&dos.FcSiteNotifyRead{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
