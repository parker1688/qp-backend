// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVip(vo *dos.FcVip) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcVip(page, pageSize int, vo *dos.FcVip) (ret []*dos.FcVip, total int64) {
	query := global.G_DB.Model(&dos.FcVip{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.DailyWithdrawalTimes > 0 {
		query = query.Where("daily_withdrawal_times = ?", vo.DailyWithdrawalTimes)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

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

	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}

	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVip
	query.Offset((page - 1) * pageSize).Order("level asc").Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVip(vo *dos.FcVip) []*dos.FcVip {
	var data []*dos.FcVip
	query := global.G_DB.Model(&dos.FcVip{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.DailyWithdrawalTimes > 0 {
		query = query.Where("daily_withdrawal_times = ?", vo.DailyWithdrawalTimes)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

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

	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}

	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}

	query.Order("level asc").Find(&data)
	return data
}

func FindByKeyFcVipFirst(vo *dos.FcVip) *dos.FcVip {
	var data *dos.FcVip
	query := global.G_DB.Model(&dos.FcVip{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	} else {
		query = query.Where("level = ?", vo.Level)
	}

	if vo.DailyWithdrawalTimes > 0 {
		query = query.Where("daily_withdrawal_times = ?", vo.DailyWithdrawalTimes)
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

	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}

	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVip(vo *dos.FcVip) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"vip_name":                vo.VipName,
		"level":                   vo.Level,
		"min_recharege_amount":    vo.MinRecharegeAmount,
		"min_bet_amount":          vo.MinBetAmount,
		"min_withdraw_amount":     vo.MinWithdrawAmount,
		"relegation_bet_amount":   vo.RelegationBetAmount,
		"daily_withdrawal_times":  vo.DailyWithdrawalTimes,
		"daily_withdrawal_amount": vo.DailyWithdrawalAmount,
		"withdrawal_fee":          vo.WithdrawalFee,
		"upgrade_gift":            vo.UpgradeGift,
		"birthday_gift":           vo.BirthdayGift,
		"weekly_gift":             vo.WeeklyGift,
		"monthly_gift":            vo.MonthlyGift,
		"merchant_code":           vo.MerchantCode,
		"create_by":               vo.CreateBy,
		"update_by":               vo.UpdateBy,
		"min_recharge_amount":     vo.MinRechargeAmount,
		"rank":                    vo.Rank,
		"rank_flag":               vo.RankFlag,
	}).Error == nil
}

func DeleteFcVip(vo *dos.FcVip) bool {
	return global.G_DB.Model(&dos.FcVip{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
