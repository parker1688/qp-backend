// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserLevelUpBounsRecord(vo *dos.FcUserLevelUpBounsRecord) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserLevelUpBounsRecord(page, pageSize int, vo *dos.FcUserLevelUpBounsRecord) (ret []*dos.FcUserLevelUpBounsRecord, total int64) {
	query := global.G_DB.Model(&dos.FcUserLevelUpBounsRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
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
	var dataSlice []*dos.FcUserLevelUpBounsRecord
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserLevelUpBounsRecord(vo *dos.FcUserLevelUpBounsRecord) []*dos.FcUserLevelUpBounsRecord {
	var data []*dos.FcUserLevelUpBounsRecord
	query := global.G_DB.Model(&dos.FcUserLevelUpBounsRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
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

func FindByKeyFcUserLevelUpBounsRecordFirst(vo *dos.FcUserLevelUpBounsRecord) *dos.FcUserLevelUpBounsRecord {
	var data *dos.FcUserLevelUpBounsRecord
	query := global.G_DB.Model(&dos.FcUserLevelUpBounsRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BetType > 0 {
		query = query.Where("bet_type = ?", vo.BetType)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
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
func UpdateFcUserLevelUpBounsRecord(vo *dos.FcUserLevelUpBounsRecord) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":          vo.UserId,
		"user_name":        vo.UserName,
		"bet_type":         vo.BetType,
		"level":            vo.Level,
		"bouns":            vo.Bouns,
		"valid_bet_amount": vo.ValidBetAmount,
		"create_by":        vo.CreateBy,
		"update_by":        vo.UpdateBy,
		"merchant_code":    vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserLevelUpBounsRecord(vo *dos.FcUserLevelUpBounsRecord) bool {
	return global.G_DB.Model(&dos.FcUserLevelUpBounsRecord{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
