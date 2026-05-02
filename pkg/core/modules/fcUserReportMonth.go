// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserReportMonth(vo *dos.FcUserReportMonth) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserReportMonth(page, pageSize int, vo *dos.FcUserReportMonth) (ret []*dos.FcUserReportMonth, total int64) {
	query := global.G_DB.Model(&dos.FcUserReportMonth{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserReportMonth
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserReportMonth(vo *dos.FcUserReportMonth) []*dos.FcUserReportMonth {
	var data []*dos.FcUserReportMonth
	query := global.G_DB.Model(&dos.FcUserReportMonth{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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

func FindByKeyFcUserReportMonthFirst(vo *dos.FcUserReportMonth) *dos.FcUserReportMonth {
	var data *dos.FcUserReportMonth
	query := global.G_DB.Model(&dos.FcUserReportMonth{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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
func UpdateFcUserReportMonth(vo *dos.FcUserReportMonth) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":               vo.UserId,
		"user_name":             vo.UserName,
		"month":                 vo.Month,
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
		"merchant_code":         vo.MerchantCode,
		"create_by":             vo.CreateBy,
		"update_by":             vo.UpdateBy,
	}).Error == nil
}

func DeleteFcUserReportMonth(vo *dos.FcUserReportMonth) bool {
	return global.G_DB.Model(&dos.FcUserReportMonth{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
