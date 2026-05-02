package vip

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// VipUpgradation
//
//	@Description: VIP升级发奖金
//
// VipUpgradation
//
//	@Description: VIP升级发奖金
func UserVipUpgradation(userMaterial *dos.FcUserMaterial, totalRechargeAmount, totalBetAmount float64) (bool, float64) {

	var upVip *dos.FcVip
	global.G_DB.Model(&dos.FcVip{}).
		Where("level=? ", userMaterial.Level+1).First(&upVip)

	if len(upVip.Id) == 0 {
		global.G_LOG.Infof("vip upgrade! user_name:%s  未找到下一个升级VIP", userMaterial.UserName)
		return false, 0
	}

	if upVip.MinRecharegeAmount > 0 && totalRechargeAmount == 0 {
		totalRechargeAmount = getUserSumDepositAmount(userMaterial.UserId)
	}

	if upVip.MinBetAmount > 0 && totalBetAmount == 0 {
		totalBetAmount = getUserSumBetAmount(userMaterial.UserId)
	}

	if totalRechargeAmount < upVip.MinRechargeAmount || totalBetAmount < upVip.MinBetAmount {
		global.G_LOG.Infof("vip upgrade! user_name:%s  升级到%s 条件不足升级失败 目前总充值金额:%0.2f,升级所需总充值金额:%0.2f  目前总投注金额:%0.2f,升级所需总投注金额:%0.2f  ", userMaterial.UserName, upVip.VipName, totalRechargeAmount, upVip.MinRechargeAmount, totalBetAmount, upVip.MinBetAmount)
		return false, 0
	}

	//升级后的下一级
	var upVipNextVip *dos.FcVip
	upVipNextVip = modules.FindByKeyFcVipFirst(&dos.FcVip{
		Level: upVip.Level + 1,
	})

	var serVipNextLevelAmount float64 //下一级所需
	if totalBetAmount > upVipNextVip.MinBetAmount {
		serVipNextLevelAmount = 0
	} else {
		serVipNextLevelAmount = decimal.NewFromFloat(upVipNextVip.MinBetAmount).
			Sub(decimal.NewFromFloat(totalBetAmount)).Truncate(2).InexactFloat64()
		if serVipNextLevelAmount == 0 {
			serVipNextLevelAmount = 0.1
		}
	}

	global.G_REDIS.Set(context.Background(),
		fmt.Sprintf(enmus.UserVipNextLevelAmountKey,
			userMaterial.UserId),
		fmt.Sprintf("%0.2f", serVipNextLevelAmount), -1)
	/**计算出下一级所需流水 end  **/

	//下一级所需流水
	if userMaterial.Level >= upVip.Level {
		return false, 0
	}

	userVip := &dos.FcUserVipRecord{
		UserId:               userMaterial.UserId,
		UserName:             userMaterial.UserName,
		BeforLevel:           userMaterial.Level,
		Level:                upVip.Level,
		BeforVip:             userMaterial.Vip,
		Vip:                  upVip.VipName,
		TotalRecharegeAmount: totalRechargeAmount,
		TotalBetAmount:       totalBetAmount,
		MerchantCode:         userMaterial.MerchantCode,
		Bonus:                upVip.UpgradeGift,
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		rowsAffected := tx.Create(userVip).RowsAffected
		if rowsAffected != 1 {
			err := errors.New("VipUpgradation insert fc_user_vip_record  fail")
			global.G_LOG.Error(err.Error())
			return err
		}

		//待处理才进行处理
		rowsAffected = tx.Exec(`update  fc_user_material set level=?,vip=? where level=? AND  user_id=?`,
			upVip.Level, upVip.VipName, userMaterial.Level, userMaterial.UserId).RowsAffected
		if rowsAffected != 1 {
			err := errors.New("VipUpgradation update fc_user_material status fail")
			global.G_LOG.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		global.G_LOG.Error(err.Error())
		return false, 0
	}
	return true, serVipNextLevelAmount
}

func getUserSumDepositAmount(userId string) float64 {
	var amount1, amount2 float64
	global.G_DB.Model(&dos.FcOrderDeposit{}).Select("sum(amount) as amount1").Where("user_id = ? and status=3", userId).Scan(&amount1)

	global.G_DB.Model(&dos.FcOrderManageOpt{}).
		Select("sum(amount) as amount2").
		Where("user_id=? AND opt_type=?", userId, 3).
		Scan(&amount2)
	return decimal.NewFromFloat(amount1).Add(decimal.NewFromFloat(amount2)).Truncate(2).InexactFloat64()
}

func getUserSumBetAmount(userId string) float64 {
	var totalHistoryBetAmount float64
	nowTime := time.Now()
	global.G_DB.Model(&dos.FcUserReportDay{}).Select("SUM(valid_betamount) as totalHistoryBetAmount").
		Where("user_id=? AND report_date< ?", userId, nowTime.Format(tool.TimeDateLayout)).Scan(&totalHistoryBetAmount)

	var totalTodayBetAmount float64
	global.G_DB.Model(&dos.FcBetRecord{}).
		Select("sum(valid_betamount) as totalTodayBetAmount").
		Where("user_id=? AND settlement_time >=? AND settlement_time< ?", userId, nowTime.Format(tool.TimeZeroLayout), nowTime.AddDate(0, 0, 1).Format(tool.TimeZeroLayout)).
		Scan(&totalTodayBetAmount)

	return decimal.NewFromFloat(totalHistoryBetAmount).Add(decimal.NewFromFloat(totalTodayBetAmount)).Truncate(2).InexactFloat64()
}

func vipRuleAdapter(userMaterial *dos.FcUserMaterial, totalRechargeAmount, totalBetAmount float64, vipRules []*dos.FcVip) *dos.FcVip {

	for _, rule := range vipRules {
		if totalBetAmount >= rule.MinBetAmount && totalRechargeAmount >= rule.MinRecharegeAmount && rule.Level > userMaterial.Level {
			return rule
		}
	}

	return nil
}
