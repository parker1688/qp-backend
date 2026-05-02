// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcProfitReportDay(vo *dos.FcProfitReportDay) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcProfitReportDay(page, pageSize int, vo *dos.FcProfitReportDay) (ret []*dos.FcProfitReportDay, total int64) {
	query := global.G_DB.Model(&dos.FcProfitReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.Register > 0 {
		query = query.Where("register = ?", vo.Register)
	}

	if vo.NewMemberRechargeNum > 0 {
		query = query.Where("new_member_recharge_num = ?", vo.NewMemberRechargeNum)
	}

	if vo.RelativelyNewMemberRechargeNum > 0 {
		query = query.Where("relatively_new_member_recharge_num = ?", vo.RelativelyNewMemberRechargeNum)
	}

	if vo.OldMemberRechargeNum > 0 {
		query = query.Where("old_member_recharge_num = ?", vo.OldMemberRechargeNum)
	}

	if vo.NewMemberWithdrawNum > 0 {
		query = query.Where("new_member_withdraw_num = ?", vo.NewMemberWithdrawNum)
	}

	if vo.NewMemberWithdrawCount > 0 {
		query = query.Where("new_member_withdraw_count = ?", vo.NewMemberWithdrawCount)
	}

	if vo.RelativelyNewMemberWithdrawNum > 0 {
		query = query.Where("relatively_new_member_withdraw_num = ?", vo.RelativelyNewMemberWithdrawNum)
	}

	if vo.RelativelyNewMemberWithdrawCount > 0 {
		query = query.Where("relatively_new_member_withdraw_count = ?", vo.RelativelyNewMemberWithdrawCount)
	}

	if vo.OldMemberWithdrawNum > 0 {
		query = query.Where("old_member_withdraw_num = ?", vo.OldMemberWithdrawNum)
	}

	if vo.OldMemberWithdrawCount > 0 {
		query = query.Where("old_member_withdraw_count = ?", vo.OldMemberWithdrawCount)
	}

	if vo.TotalBetCount > 0 {
		query = query.Where("total_bet_count = ?", vo.TotalBetCount)
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
	query.Order("report_date desc")
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcProfitReportDay
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcProfitReportDay(vo *dos.FcProfitReportDay) []*dos.FcProfitReportDay {
	var data []*dos.FcProfitReportDay
	query := global.G_DB.Model(&dos.FcProfitReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.Register > 0 {
		query = query.Where("register = ?", vo.Register)
	}

	if vo.NewMemberRechargeNum > 0 {
		query = query.Where("new_member_recharge_num = ?", vo.NewMemberRechargeNum)
	}

	if vo.RelativelyNewMemberRechargeNum > 0 {
		query = query.Where("relatively_new_member_recharge_num = ?", vo.RelativelyNewMemberRechargeNum)
	}

	if vo.OldMemberRechargeNum > 0 {
		query = query.Where("old_member_recharge_num = ?", vo.OldMemberRechargeNum)
	}

	if vo.NewMemberWithdrawNum > 0 {
		query = query.Where("new_member_withdraw_num = ?", vo.NewMemberWithdrawNum)
	}

	if vo.NewMemberWithdrawCount > 0 {
		query = query.Where("new_member_withdraw_count = ?", vo.NewMemberWithdrawCount)
	}

	if vo.RelativelyNewMemberWithdrawNum > 0 {
		query = query.Where("relatively_new_member_withdraw_num = ?", vo.RelativelyNewMemberWithdrawNum)
	}

	if vo.RelativelyNewMemberWithdrawCount > 0 {
		query = query.Where("relatively_new_member_withdraw_count = ?", vo.RelativelyNewMemberWithdrawCount)
	}

	if vo.OldMemberWithdrawNum > 0 {
		query = query.Where("old_member_withdraw_num = ?", vo.OldMemberWithdrawNum)
	}

	if vo.OldMemberWithdrawCount > 0 {
		query = query.Where("old_member_withdraw_count = ?", vo.OldMemberWithdrawCount)
	}

	if vo.TotalBetCount > 0 {
		query = query.Where("total_bet_count = ?", vo.TotalBetCount)
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
	query.Order("report_date desc")
	query.Find(&data)
	return data
}

func FindByKeyFcProfitReportDayFirst(vo *dos.FcProfitReportDay) *dos.FcProfitReportDay {
	var data *dos.FcProfitReportDay
	query := global.G_DB.Model(&dos.FcProfitReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if vo.Online > 0 {
		query = query.Where("online = ?", vo.Online)
	}

	if vo.Register > 0 {
		query = query.Where("register = ?", vo.Register)
	}

	if vo.NewMemberRechargeNum > 0 {
		query = query.Where("new_member_recharge_num = ?", vo.NewMemberRechargeNum)
	}

	if vo.RelativelyNewMemberRechargeNum > 0 {
		query = query.Where("relatively_new_member_recharge_num = ?", vo.RelativelyNewMemberRechargeNum)
	}

	if vo.OldMemberRechargeNum > 0 {
		query = query.Where("old_member_recharge_num = ?", vo.OldMemberRechargeNum)
	}

	if vo.NewMemberWithdrawNum > 0 {
		query = query.Where("new_member_withdraw_num = ?", vo.NewMemberWithdrawNum)
	}

	if vo.NewMemberWithdrawCount > 0 {
		query = query.Where("new_member_withdraw_count = ?", vo.NewMemberWithdrawCount)
	}

	if vo.RelativelyNewMemberWithdrawNum > 0 {
		query = query.Where("relatively_new_member_withdraw_num = ?", vo.RelativelyNewMemberWithdrawNum)
	}

	if vo.RelativelyNewMemberWithdrawCount > 0 {
		query = query.Where("relatively_new_member_withdraw_count = ?", vo.RelativelyNewMemberWithdrawCount)
	}

	if vo.OldMemberWithdrawNum > 0 {
		query = query.Where("old_member_withdraw_num = ?", vo.OldMemberWithdrawNum)
	}

	if vo.OldMemberWithdrawCount > 0 {
		query = query.Where("old_member_withdraw_count = ?", vo.OldMemberWithdrawCount)
	}

	if vo.TotalBetCount > 0 {
		query = query.Where("total_bet_count = ?", vo.TotalBetCount)
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
func UpdateFcProfitReportDay(vo *dos.FcProfitReportDay) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"report_date":                              vo.ReportDate,
		"online":                                   vo.Online,
		"profit":                                   vo.Profit,
		"register":                                 vo.Register,
		"total_recharge_amount":                    vo.TotalRechargeAmount,
		"total_withdraw_amount":                    vo.TotalWithdrawAmount,
		"new_member_recharge_probability":          vo.NewMemberRechargeProbability,
		"new_member_recharge_num":                  vo.NewMemberRechargeNum,
		"new_member_recharge_av_amount":            vo.NewMemberRechargeAvAmount,
		"new_member_recharge_amount":               vo.NewMemberRechargeAmount,
		"relatively_new_member_recharge_num":       vo.RelativelyNewMemberRechargeNum,
		"relatively_new_member_recharge_av_amount": vo.RelativelyNewMemberRechargeAvAmount,
		"relatively_new_member_recharge_amount":    vo.RelativelyNewMemberRechargeAmount,
		"old_member_recharge_num":                  vo.OldMemberRechargeNum,
		"old_member_recharge_av_amount":            vo.OldMemberRechargeAvAmount,
		"old_member_recharge_amount":               vo.OldMemberRechargeAmount,
		"new_member_withdraw_num":                  vo.NewMemberWithdrawNum,
		"new_member_withdraw_av_amount":            vo.NewMemberWithdrawAvAmount,
		"new_member_withdraw_amount":               vo.NewMemberWithdrawAmount,
		"new_member_withdraw_count":                vo.NewMemberWithdrawCount,
		"relatively_new_member_withdraw_num":       vo.RelativelyNewMemberWithdrawNum,
		"relatively_new_member_withdraw_av_amount": vo.RelativelyNewMemberWithdrawAvAmount,
		"relatively_new_member_withdraw_amount":    vo.RelativelyNewMemberWithdrawAmount,
		"relatively_new_member_withdraw_count":     vo.RelativelyNewMemberWithdrawCount,
		"old_member_withdraw_num":                  vo.OldMemberWithdrawNum,
		"old_member_withdraw_av_amount":            vo.OldMemberWithdrawAvAmount,
		"old_member_withdraw_amount":               vo.OldMemberWithdrawAmount,
		"old_member_withdraw_count":                vo.OldMemberWithdrawCount,
		"agent_invite_award_amount":                vo.AgentInviteAwardAmount,
		"agent_commission":                         vo.AgentCommission,
		"user_invite_award_amount":                 vo.UserInviteAwardAmount,
		"total_bonus":                              vo.TotalBonus,
		"total_bet_amount":                         vo.TotalBetAmount,
		"total_bet_count":                          vo.TotalBetCount,
		"total_bet_net_amount":                     vo.TotalBetNetAmount,
		"venue_cost":                               vo.VenueCost,
		"deposit_cost":                             vo.DepositCost,
		"withdraw_cost":                            vo.WithdrawCost,
		"merchant_code":                            vo.MerchantCode,
		"create_by":                                vo.CreateBy,
		"update_by":                                vo.UpdateBy,
	}).Error == nil
}

func DeleteFcProfitReportDay(vo *dos.FcProfitReportDay) bool {
	return global.G_DB.Model(&dos.FcProfitReportDay{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
