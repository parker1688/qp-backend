// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/shopspring/decimal"
)

func SaveFcUserReport(vo *dos.FcUserReport) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserReport(page, pageSize int, vo *dos.FcUserReport, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcUserReport, total int64) {
	query := global.G_DB.Model(&dos.FcUserReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.BetCount > 0 {
		query = query.Where("bet_count = ?", vo.BetCount)
	}

	if vo.WithdrawalCount > 0 {
		query = query.Where("withdrawal_count = ?", vo.WithdrawalCount)
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

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", vo.CreateTime)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", vo.CreateTime)
	}

	query.Order("update_time desc")
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserReport
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserReport(vo *dos.FcUserReport) []*dos.FcUserReport {
	var data []*dos.FcUserReport
	query := global.G_DB.Model(&dos.FcUserReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.BetCount > 0 {
		query = query.Where("bet_count = ?", vo.BetCount)
	}

	if vo.WithdrawalCount > 0 {
		query = query.Where("withdrawal_count = ?", vo.WithdrawalCount)
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
	query.Order("update_time desc")
	query.Find(&data)
	return data
}

func FindByKeyFcUserReportFirst(vo *dos.FcUserReport) *dos.FcUserReport {
	var data *dos.FcUserReport
	query := global.G_DB.Model(&dos.FcUserReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.BetCount > 0 {
		query = query.Where("bet_count = ?", vo.BetCount)
	}

	if vo.WithdrawalCount > 0 {
		query = query.Where("withdrawal_count = ?", vo.WithdrawalCount)
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
func UpdateFcUserReport(vo *dos.FcUserReport) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":               vo.UserId,
		"user_name":             vo.UserName,
		"recharge_amount":       vo.RechargeAmount,
		"recharge_count":        vo.RechargeCount,
		"bet_amount":            vo.BetAmount,
		"bet_count":             vo.BetCount,
		"valid_betamount":       vo.ValidBetamount,
		"withdrawal_amount":     vo.WithdrawalAmount,
		"withdrawal_count":      vo.WithdrawalCount,
		"promotion_amount":      vo.PromotionAmount,
		"rebate_amount":         vo.RebateAmount,
		"win_amount":            vo.WinAmount,
		"absolutely_win_amount": vo.AbsolutelyWinAmount,
		"friends_bonus_amount":  vo.FriendsBonusAmount,
		"merchant_code":         vo.MerchantCode,
		"create_by":             vo.CreateBy,
		"update_by":             vo.UpdateBy,
	}).Error == nil
}

func DeleteFcUserReport(vo *dos.FcUserReport) bool {
	return global.G_DB.Model(&dos.FcUserReport{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetUserReportTotalPromotionExVal - 获取额外福利累计
// @param {string} userId
// @returns float64
func GetUserReportTotalPromotionExVal(userId string) float64 {
	if userId == "" || userId == "0" {
		return 0.00
	}

	// += 签到累计 + 任务累计 + 活动累计
	return GetDailyBonusPromotionTotalVal(userId) +
		GetUserTaskPromotionTotalVal(userId) +
		GetUserActivityPromotionTotalVal(userId)
}

// GetDepositWithdrawSubAmount - 获取用户充提差
// @param {string} userId
// @returns float64
func GetDepositWithdrawSubAmount(userId string) float64 {
	var rechargeAmount float64
	global.G_DB.Model(&dos.FcOrderDeposit{}).
		Select("IFNULL(sum(amount), 0) as rechargeAmount").
		Where("user_id = ? and status = 3", userId).Scan(&rechargeAmount)

	var withdrawalAmount float64
	global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).
		Select("IFNULL(sum(amount), 0) as withdrawalAmount").
		Where("user_id = ? and status = 3", userId).Scan(&withdrawalAmount)

	return decimal.NewFromFloat(rechargeAmount).
		Sub(decimal.NewFromFloat(withdrawalAmount)).
		Truncate(2).InexactFloat64()
}
