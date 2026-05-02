// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcComplexReport(vo *dos.FcComplexReport) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcComplexReport(page, pageSize int, vo *dos.FcComplexReport, c *gin.Context) (ret []*dos.FcComplexReport, total int64) {
	query := global.G_DB.Model(&dos.FcComplexReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Day) > 0 {
		query = query.Where("day = ?", vo.Day)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if vo.RegisterNum > 0 {
		query = query.Where("register_num = ?", vo.RegisterNum)
	}

	if vo.FirstDepositNum > 0 {
		query = query.Where("first_deposit_num = ?", vo.FirstDepositNum)
	}

	if vo.DepositNum > 0 {
		query = query.Where("deposit_num = ?", vo.DepositNum)
	}

	if vo.DepositCount > 0 {
		query = query.Where("deposit_count = ?", vo.DepositCount)
	}

	if vo.NewUserDepositCount > 0 {
		query = query.Where("new_user_deposit_count = ?", vo.NewUserDepositCount)
	}

	if vo.LoginNum > 0 {
		query = query.Where("login_num = ?", vo.LoginNum)
	}

	if vo.WithdrawNum > 0 {
		query = query.Where("withdraw_num = ?", vo.WithdrawNum)
	}

	if vo.BetNum > 0 {
		query = query.Where("bet_num = ?", vo.BetNum)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcComplexReport
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcComplexReport(vo *dos.FcComplexReport, c *gin.Context) []*dos.FcComplexReport {
	var data []*dos.FcComplexReport
	query := global.G_DB.Model(&dos.FcComplexReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Day) > 0 {
		query = query.Where("day = ?", vo.Day)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if vo.RegisterNum > 0 {
		query = query.Where("register_num = ?", vo.RegisterNum)
	}

	if vo.FirstDepositNum > 0 {
		query = query.Where("first_deposit_num = ?", vo.FirstDepositNum)
	}

	if vo.DepositNum > 0 {
		query = query.Where("deposit_num = ?", vo.DepositNum)
	}

	if vo.DepositCount > 0 {
		query = query.Where("deposit_count = ?", vo.DepositCount)
	}

	if vo.NewUserDepositCount > 0 {
		query = query.Where("new_user_deposit_count = ?", vo.NewUserDepositCount)
	}

	if vo.LoginNum > 0 {
		query = query.Where("login_num = ?", vo.LoginNum)
	}

	if vo.WithdrawNum > 0 {
		query = query.Where("withdraw_num = ?", vo.WithdrawNum)
	}

	if vo.BetNum > 0 {
		query = query.Where("bet_num = ?", vo.BetNum)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcComplexReportFirst(vo *dos.FcComplexReport) *dos.FcComplexReport {
	var data *dos.FcComplexReport
	query := global.G_DB.Model(&dos.FcComplexReport{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Day) > 0 {
		query = query.Where("day = ?", vo.Day)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if vo.RegisterNum > 0 {
		query = query.Where("register_num = ?", vo.RegisterNum)
	}

	if vo.FirstDepositNum > 0 {
		query = query.Where("first_deposit_num = ?", vo.FirstDepositNum)
	}

	if vo.DepositNum > 0 {
		query = query.Where("deposit_num = ?", vo.DepositNum)
	}

	if vo.DepositCount > 0 {
		query = query.Where("deposit_count = ?", vo.DepositCount)
	}

	if vo.NewUserDepositCount > 0 {
		query = query.Where("new_user_deposit_count = ?", vo.NewUserDepositCount)
	}

	if vo.LoginNum > 0 {
		query = query.Where("login_num = ?", vo.LoginNum)
	}

	if vo.WithdrawNum > 0 {
		query = query.Where("withdraw_num = ?", vo.WithdrawNum)
	}

	if vo.BetNum > 0 {
		query = query.Where("bet_num = ?", vo.BetNum)
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
func UpdateFcComplexReport(vo *dos.FcComplexReport) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"day":                         vo.Day,
		"merchant_code":               vo.MerchantCode,
		"merchant_name":               vo.MerchantName,
		"bet_win":                     vo.BetWin,
		"bet_amount":                  vo.BetAmount,
		"game_kill_rate":              vo.GameKillRate,
		"bet_multiple":                vo.BetMultiple,
		"register_num":                vo.RegisterNum,
		"first_deposit_num":           vo.FirstDepositNum,
		"register_deposit_rate":       vo.RegisterDepositRate,
		"rebate_amount":               vo.RebateAmount,
		"deposit_num":                 vo.DepositNum,
		"deposit_count":               vo.DepositCount,
		"first_deposit_amount":        vo.FirstDepositAmount,
		"new_user_deposit_amount":     vo.NewUserDepositAmount,
		"new_user_deposit_count":      vo.NewUserDepositCount,
		"login_num":                   vo.LoginNum,
		"withdraw_num":                vo.WithdrawNum,
		"bet_num":                     vo.BetNum,
		"promotion_amount":            vo.PromotionAmount,
		"alipay_deposit_amount":       vo.AlipayDepositAmount,
		"wx_deposit_amount":           vo.WxDepositAmount,
		"bank_deposit_amount":         vo.BankDepositAmount,
		"wallet_deposit_amount":       vo.WalletDepositAmount,
		"num_cny_deposit_amount":      vo.NumCnyDepositAmount,
		"usdt_deposit_amount":         vo.UsdtDepositAmount,
		"admin_deposit_amount":        vo.AdminDepositAmount,
		"total_deposit_amount":        vo.TotalDepositAmount,
		"alipay_withdraw_amount":      vo.AlipayWithdrawAmount,
		"bank_withdraw_amount":        vo.BankWithdrawAmount,
		"wallet_withdraw_amount":      vo.WalletWithdrawAmount,
		"usdt_withdraw_amount":        vo.UsdtWithdrawAmount,
		"total_withdraw_amount":       vo.TotalWithdrawAmount,
		"deposit_withdraw_sub_amount": vo.DepositWithdrawSubAmount,
		"kill_rate":                   vo.KillRate,
		"create_by":                   vo.CreateBy,
		"update_by":                   vo.UpdateBy,
	}).Error == nil
}

func DeleteFcComplexReport(vo *dos.FcComplexReport) bool {
	return global.G_DB.Model(&dos.FcComplexReport{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
