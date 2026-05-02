// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVenueConfig(vo *dos.FcVenueConfig) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenueConfig(page, pageSize int, vo *dos.FcVenueConfig) (ret []*dos.FcVenueConfig, total int64) {
	query := global.G_DB.Model(&dos.FcVenueConfig{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVenueConfig
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueConfig(vo *dos.FcVenueConfig) []*dos.FcVenueConfig {
	var data []*dos.FcVenueConfig
	query := global.G_DB.Model(&dos.FcVenueConfig{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}
	query.Find(&data)
	return data
}

func FindByKeyFcVenueConfigFirst(vo *dos.FcVenueConfig) *dos.FcVenueConfig {
	var data *dos.FcVenueConfig
	query := global.G_DB.Model(&dos.FcVenueConfig{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Alias) > 0 {
		query = query.Where("alias = ?", vo.Alias)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
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

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}
	query.Take(&data)
	return data
}

//根据主键Update
func UpdateFcVenueConfig(vo *dos.FcVenueConfig) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"alias":      vo.Alias,
		"venue_name": vo.VenueName,
		"venue_code": vo.VenueCode,
		"remark":     vo.Remark,
		"content":    vo.Content,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
	}).Error == nil
}

func DeleteFcVenueConfig(vo *dos.FcVenueConfig) bool {
	return global.G_DB.Model(&dos.FcVenueConfig{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
