// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserReportDay(vo *dos.FcUserReportDay) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserReportDay(page, pageSize int, vo *dos.FcUserReportDay) (ret []*dos.FcUserReportDay, total int64) {
	query := global.G_DB.Model(&dos.FcUserReportDay{})
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

	//if !vo.ReportDate.Timer().IsZero() {
	//	query = query.Where("report_date = ?", vo.ReportDate)
	//}

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
	var dataSlice []*dos.FcUserReportDay
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindPageFcUserReportDayBills(page, pageSize int, vo *dos.FcUserReportDayBills) (ret []*dos.FcUserReportDayBills, total int64) {
	query := global.G_DB.Model(&dos.FcUserReportDay{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	//if !vo.ReportDate.Timer().IsZero() {
	//	query = query.Where("report_date = ?", vo.ReportDate)
	//}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	query.Order("update_time desc")
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserReportDayBills
	query.Table("fc_user_report_day").Offset((page - 1) * pageSize).Limit(pageSize).Scan(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserReportDay(vo *dos.FcUserReportDay) []*dos.FcUserReportDay {
	var data []*dos.FcUserReportDay
	query := global.G_DB.Model(&dos.FcUserReportDay{})
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

	//if !vo.ReportDate.Timer().IsZero() {
	//	query = query.Where("report_date = ?", vo.ReportDate)
	//}

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

func FindByKeyFcUserReportDayFirst(vo *dos.FcUserReportDay) *dos.FcUserReportDay {
	var data *dos.FcUserReportDay
	query := global.G_DB.Model(&dos.FcUserReportDay{})
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

	//if !vo.ReportDate.Timer().IsZero() {
	//	query = query.Where("report_date = ?", vo.ReportDate)
	//}

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
func UpdateFcUserReportDay(vo *dos.FcUserReportDay) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":               vo.UserId,
		"user_name":             vo.UserName,
		"recharge_amount":       vo.RechargeAmount,
		"recharge_count":        vo.RechargeCount,
		"recharge_fee":          vo.RechargeFee,
		"bet_amount":            vo.BetAmount,
		"bet_count":             vo.BetCount,
		"valid_betamount":       vo.ValidBetamount,
		"withdrawal_amount":     vo.WithdrawalAmount,
		"withdrawal_count":      vo.WithdrawalCount,
		"withdrawal_fee":        vo.WithdrawalFee,
		"promotion_amount":      vo.PromotionAmount,
		"rebate_amount":         vo.RebateAmount,
		"win_amount":            vo.WinAmount,
		"absolutely_win_amount": vo.AbsolutelyWinAmount,
		"report_date":           vo.ReportDate,
		"merchant_code":         vo.MerchantCode,
		"create_by":             vo.CreateBy,
		"update_by":             vo.UpdateBy,
	}).Error == nil
}

func DeleteFcUserReportDay(vo *dos.FcUserReportDay) bool {
	return global.G_DB.Model(&dos.FcUserReportDay{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
