// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcFinanceReportDay(vo *dos.FcFinanceReportDay) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcFinanceReportDay(page, pageSize int, vo *dos.FcFinanceReportDay) (ret []*dos.FcFinanceReportDay, total int64) {
	query := global.G_DB.Model(&dos.FcFinanceReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.RechargeUserCount > 0 {
		query = query.Where("recharge_user_count = ?", vo.RechargeUserCount)
	}

	if vo.BetUserCount > 0 {
		query = query.Where("bet_user_count = ?", vo.BetUserCount)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.AdminRechargeCount > 0 {
		query = query.Where("admin_recharge_count = ?", vo.AdminRechargeCount)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcFinanceReportDay
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcFinanceReportDay(vo *dos.FcFinanceReportDay) []*dos.FcFinanceReportDay {
	var data []*dos.FcFinanceReportDay
	query := global.G_DB.Model(&dos.FcFinanceReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.RechargeUserCount > 0 {
		query = query.Where("recharge_user_count = ?", vo.RechargeUserCount)
	}

	if vo.BetUserCount > 0 {
		query = query.Where("bet_user_count = ?", vo.BetUserCount)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.AdminRechargeCount > 0 {
		query = query.Where("admin_recharge_count = ?", vo.AdminRechargeCount)
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

	query.Find(&data)
	return data
}

func FindByKeyFcFinanceReportDayFirst(vo *dos.FcFinanceReportDay) *dos.FcFinanceReportDay {
	var data *dos.FcFinanceReportDay
	query := global.G_DB.Model(&dos.FcFinanceReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.RechargeUserCount > 0 {
		query = query.Where("recharge_user_count = ?", vo.RechargeUserCount)
	}

	if vo.BetUserCount > 0 {
		query = query.Where("bet_user_count = ?", vo.BetUserCount)
	}

	if vo.RechargeCount > 0 {
		query = query.Where("recharge_count = ?", vo.RechargeCount)
	}

	if vo.AdminRechargeCount > 0 {
		query = query.Where("admin_recharge_count = ?", vo.AdminRechargeCount)
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
func UpdateFcFinanceReportDay(vo *dos.FcFinanceReportDay) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"report_date":           vo.ReportDate,
		"online":                vo.Online,
		"recharge_user_count":   vo.RechargeUserCount,
		"bet_user_count":        vo.BetUserCount,
		"promotion_amount":      vo.PromotionAmount,
		"rebate_amount":         vo.RebateAmount,
		"gross_profit":          vo.GrossProfit,
		"recharge_amount":       vo.RechargeAmount,
		"recharge_count":        vo.RechargeCount,
		"admin_recharge_amount": vo.AdminRechargeAmount,
		"admin_recharge_count":  vo.AdminRechargeCount,
		"withdrawal_amount":     vo.WithdrawalAmount,
		"withdrawal_count":      vo.WithdrawalCount,
		"agent_award_amount":    vo.AgentAwardAmount,
		"merchant_code":         vo.MerchantCode,
		"create_by":             vo.CreateBy,
		"update_by":             vo.UpdateBy,
	}).Error == nil
}

func DeleteFcFinanceReportDay(vo *dos.FcFinanceReportDay) bool {
	return global.G_DB.Model(&dos.FcFinanceReportDay{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
