// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserLevelWeekSetting(vo *dos.FcUserLevelWeekSetting) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserLevelWeekSetting(page, pageSize int, vo *dos.FcUserLevelWeekSetting) (ret []*dos.FcUserLevelWeekSetting, total int64) {
	query := global.G_DB.Model(&dos.FcUserLevelWeekSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserLevelWeekSetting
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserLevelWeekSetting(vo *dos.FcUserLevelWeekSetting) []*dos.FcUserLevelWeekSetting {
	var data []*dos.FcUserLevelWeekSetting
	query := global.G_DB.Model(&dos.FcUserLevelWeekSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
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

func FindByKeyFcUserLevelWeekSettingFirst(vo *dos.FcUserLevelWeekSetting) *dos.FcUserLevelWeekSetting {
	var data *dos.FcUserLevelWeekSetting
	query := global.G_DB.Model(&dos.FcUserLevelWeekSetting{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
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

	query.First(&data)
	return data
}

// 根据主键Update
func UpdateFcUserLevelWeekSetting(vo *dos.FcUserLevelWeekSetting) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"level":            vo.Level,
		"bet_type":         vo.BetType,
		"valid_bet_amount": vo.ValidBetAmount,
		"level_bouns":      vo.LevelBouns,
		"week_bouns":       vo.WeekBouns,
		"create_by":        vo.CreateBy,
		"update_by":        vo.UpdateBy,
		"merchant_code":    vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserLevelWeekSetting(vo *dos.FcUserLevelWeekSetting) bool {
	return global.G_DB.Model(&dos.FcUserLevelWeekSetting{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
