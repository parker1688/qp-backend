// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserRankReportDay(vo *dos.FcUserRankReportDay) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserRankReportDay(page, pageSize int, vo *dos.FcUserRankReportDay) (ret []*dos.FcUserRankReportDay, total int64) {
	query := global.G_DB.Model(&dos.FcUserRankReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.UserType > 0 {
		query = query.Where("user_type = ?", vo.UserType)
	}

	if len(vo.Business) > 0 {
		query = query.Where("business = ?", vo.Business)
	}

	if vo.GradeOneNum > 0 {
		query = query.Where("grade_one_num = ?", vo.GradeOneNum)
	}

	if vo.GradeOneRechargeNum > 0 {
		query = query.Where("grade_one_recharge_num = ?", vo.GradeOneRechargeNum)
	}

	if vo.GradeTwoNum > 0 {
		query = query.Where("grade_two_num = ?", vo.GradeTwoNum)
	}

	if vo.GradeTwoRechargeNum > 0 {
		query = query.Where("grade_two_recharge_num = ?", vo.GradeTwoRechargeNum)
	}

	if vo.GradeThreeNum > 0 {
		query = query.Where("grade_three_num = ?", vo.GradeThreeNum)
	}

	if vo.GradeThreeRechargeNum > 0 {
		query = query.Where("grade_three_recharge_num = ?", vo.GradeThreeRechargeNum)
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

	if vo.TotalNum > 0 {
		query = query.Where("total_num = ?", vo.TotalNum)
	}

	if vo.TotalRechargeNum > 0 {
		query = query.Where("total_recharge_num = ?", vo.TotalRechargeNum)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserRankReportDay
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserRankReportDay(vo *dos.FcUserRankReportDay) []*dos.FcUserRankReportDay {
	var data []*dos.FcUserRankReportDay
	query := global.G_DB.Model(&dos.FcUserRankReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.UserType > 0 {
		query = query.Where("user_type = ?", vo.UserType)
	}

	if len(vo.Business) > 0 {
		query = query.Where("business = ?", vo.Business)
	}

	if vo.GradeOneNum > 0 {
		query = query.Where("grade_one_num = ?", vo.GradeOneNum)
	}

	if vo.GradeOneRechargeNum > 0 {
		query = query.Where("grade_one_recharge_num = ?", vo.GradeOneRechargeNum)
	}

	if vo.GradeTwoNum > 0 {
		query = query.Where("grade_two_num = ?", vo.GradeTwoNum)
	}

	if vo.GradeTwoRechargeNum > 0 {
		query = query.Where("grade_two_recharge_num = ?", vo.GradeTwoRechargeNum)
	}

	if vo.GradeThreeNum > 0 {
		query = query.Where("grade_three_num = ?", vo.GradeThreeNum)
	}

	if vo.GradeThreeRechargeNum > 0 {
		query = query.Where("grade_three_recharge_num = ?", vo.GradeThreeRechargeNum)
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

	if vo.TotalNum > 0 {
		query = query.Where("total_num = ?", vo.TotalNum)
	}

	if vo.TotalRechargeNum > 0 {
		query = query.Where("total_recharge_num = ?", vo.TotalRechargeNum)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserRankReportDayFirst(vo *dos.FcUserRankReportDay) *dos.FcUserRankReportDay {
	var data *dos.FcUserRankReportDay
	query := global.G_DB.Model(&dos.FcUserRankReportDay{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if !vo.ReportDate.Timer().IsZero() {
		query = query.Where("report_date = ?", vo.ReportDate)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.UserType > 0 {
		query = query.Where("user_type = ?", vo.UserType)
	}

	if len(vo.Business) > 0 {
		query = query.Where("business = ?", vo.Business)
	}

	if vo.GradeOneNum > 0 {
		query = query.Where("grade_one_num = ?", vo.GradeOneNum)
	}

	if vo.GradeOneRechargeNum > 0 {
		query = query.Where("grade_one_recharge_num = ?", vo.GradeOneRechargeNum)
	}

	if vo.GradeTwoNum > 0 {
		query = query.Where("grade_two_num = ?", vo.GradeTwoNum)
	}

	if vo.GradeTwoRechargeNum > 0 {
		query = query.Where("grade_two_recharge_num = ?", vo.GradeTwoRechargeNum)
	}

	if vo.GradeThreeNum > 0 {
		query = query.Where("grade_three_num = ?", vo.GradeThreeNum)
	}

	if vo.GradeThreeRechargeNum > 0 {
		query = query.Where("grade_three_recharge_num = ?", vo.GradeThreeRechargeNum)
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

	if vo.TotalNum > 0 {
		query = query.Where("total_num = ?", vo.TotalNum)
	}

	if vo.TotalRechargeNum > 0 {
		query = query.Where("total_recharge_num = ?", vo.TotalRechargeNum)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserRankReportDay(vo *dos.FcUserRankReportDay) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"report_date":                    vo.ReportDate,
		"user_id":                        vo.UserId,
		"user_name":                      vo.UserName,
		"user_type":                      vo.UserType,
		"business":                       vo.Business,
		"support_amount":                 vo.SupportAmount,
		"grade_one_num":                  vo.GradeOneNum,
		"grade_one_recharge_num":         vo.GradeOneRechargeNum,
		"grade_one_recharge_fee":         vo.GradeOneRechargeFee,
		"grade_one_recharge_amount":      vo.GradeOneRechargeAmount,
		"grade_one_recharge_av_amount":   vo.GradeOneRechargeAvAmount,
		"grade_one_recharge_perce":       vo.GradeOneRechargePerce,
		"grade_two_num":                  vo.GradeTwoNum,
		"grade_two_recharge_num":         vo.GradeTwoRechargeNum,
		"grade_two_recharge_amount":      vo.GradeTwoRechargeAmount,
		"grade_two_recharge_av_amount":   vo.GradeTwoRechargeAvAmount,
		"grade_two_recharge_perce":       vo.GradeTwoRechargePerce,
		"grade_three_num":                vo.GradeThreeNum,
		"grade_three_recharge_num":       vo.GradeThreeRechargeNum,
		"grade_three_recharge_amount":    vo.GradeThreeRechargeAmount,
		"grade_three_recharge_av_amount": vo.GradeThreeRechargeAvAmount,
		"grade_three_recharge_perce":     vo.GradeThreeRechargePerce,
		"user_invite_bonus_amount":       vo.UserInviteBonusAmount,
		"merchant_code":                  vo.MerchantCode,
		"create_by":                      vo.CreateBy,
		"update_by":                      vo.UpdateBy,
		"total_num":                      vo.TotalNum,
		"total_recharge_num":             vo.TotalRechargeNum,
		"total_recharge_amount":          vo.TotalRechargeAmount,
		"total_recharge_av_amount":       vo.TotalRechargeAvAmount,
		"total_recharge_fee":             vo.TotalRechargeFee,
		"total_recharge_perce":           vo.TotalRechargePerce,
		"remark":                         vo.Remark,
	}).Error == nil
}

func DeleteFcUserRankReportDay(vo *dos.FcUserRankReportDay) bool {
	return global.G_DB.Model(&dos.FcUserRankReportDay{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
