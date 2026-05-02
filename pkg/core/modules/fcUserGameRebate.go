// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserGameRebate(vo *dos.FcUserGameRebate) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserGameRebate(page, pageSize int, vo *dos.FcUserGameRebate) (ret []*dos.FcUserGameRebate, total int64) {
	query := global.G_DB.Model(&dos.FcUserGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.Day.Timer().IsZero() {
		query = query.Where("day = ?", vo.Day)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserGameRebate
	query.Order("day desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserGameRebate(vo *dos.FcUserGameRebate) []*dos.FcUserGameRebate {
	var data []*dos.FcUserGameRebate
	query := global.G_DB.Model(&dos.FcUserGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.Day.Timer().IsZero() {
		query = query.Where("day = ?", vo.Day)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	query.Find(&data)
	return data
}

func FindByKeyFcUserGameRebateFirst(vo *dos.FcUserGameRebate) *dos.FcUserGameRebate {
	var data *dos.FcUserGameRebate
	query := global.G_DB.Model(&dos.FcUserGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.Day.Timer().IsZero() {
		query = query.Where("day = ?", vo.Day)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserGameRebate(vo *dos.FcUserGameRebate) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"day":                vo.Day,
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"game_type":          vo.GameType,
		"net_amount":         vo.NetAmount,
		"bet_amount":         vo.BetAmount,
		"valid_betamount":    vo.ValidBetamount,
		"merchant_code":      vo.MerchantCode,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
		"bonus_amount":       vo.BonusAmount,
		"bonus_amount_issue": vo.BonusAmountIssue,
		"bonus_rate":         vo.BonusRate,
	}).Error == nil
}

func DeleteFcUserGameRebate(vo *dos.FcUserGameRebate) bool {
	return global.G_DB.Model(&dos.FcUserGameRebate{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
