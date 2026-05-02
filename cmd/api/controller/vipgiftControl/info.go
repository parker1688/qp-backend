package vipgiftControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WeekGiftApply(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户不存在")
		return
	}

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	y, m, w := tool.GetDateWeek(time.Now().Format(time.DateOnly))
	week := fmt.Sprintf("%d-%d-%d", y, m, w)
	key := fmt.Sprintf("LockApplyWeekGift::%s::%s", userInfo.UserId, week)

	if !global.G_REDIS.SetNX(context.Background(), key, "1", time.Duration(3)*time.Second).Val() {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}
	var vipWeekGift *dos.FcVipWeekGift
	global.G_DB.Model(&dos.FcVipWeekGift{}).Where("user_id=? AND week=?", userInfo.UserId, week).Take(&vipWeekGift)

	if len(vipWeekGift.Id) > 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}
	vipinfo := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: userInfo.Level})
	if vipinfo.WeeklyGift <= 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "奖金未配置")
		return
	}
	vipWeekGift = &dos.FcVipWeekGift{
		UserId:       userInfo.UserId,
		UserName:     userInfo.UserName,
		MerchantCode: userInfo.MerchantCode,
		VipName:      userInfo.Vip,
		Level:        userInfo.Level,
		BonusAmount:  vipinfo.WeeklyGift,
		Week:         week,
	}
	rRow := global.G_DB.Create(&vipWeekGift)

	if rRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, rRow.Error.Error())
		return
	}

	if len(vipWeekGift.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "领取失败")
		return
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		//修改用户金额表
		err := userTransfer.UserAmountChange(tx, vipinfo.WeeklyGift, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency, week+"周礼金", userInfo.UserId, "", "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "vip_week_gift"))
		if err != nil {
			global.G_LOG.Errorf("WeekGiftApply userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s", vipinfo.WeeklyGift, userInfo.UserId, err.Error())
			return err
		}

		err = global.G_DB.Model(&dos.FcVipWeekGift{}).Where("id=?", vipWeekGift.Id).Updates(map[string]interface{}{
			"bonus_amount_issue": vipinfo.WeeklyGift,
		}).Error

		if err != nil {
			return err
		}

		fop := &dos.FcOrderPromotion{
			ApplyAmount:  vipinfo.WeeklyGift,
			AppleRate:    1,
			ApplyType:    enmus.Promotion_VIP_WeekGift,
			Amount:       vipinfo.WeeklyGift,
			Status:       enmus.ORDER_YES_STATUS,
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			TurnOver:     0,
			MerchantCode: userInfo.MerchantCode,
			Remake:       fmt.Sprintf("第%s周按照VIP等级%s领取周礼金%0.2f", week, userInfo.Vip, vipinfo.WeeklyGift),
			//Currency:     ,
		}
		fop.OrderSn = vipWeekGift.Id //关联记录ID
		rowsAffected := tx.Create(fop).RowsAffected
		if rowsAffected != 1 {
			//global.G_LOG.Error(err.Error())
			return err
		}
		return err
	})

	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	global.G_REDIS.Expire(context.Background(), key, time.Duration(24*7)*time.Hour)
	response.SuccessMsgJSON(c, true, "领取成功")
}

func MonthGiftApply(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户不存在")
		return
	}

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	month := time.Now().Format("2006-01")

	key := fmt.Sprintf("LockApplyMonthGift::%s::%s", userInfo.UserId, month)

	if !global.G_REDIS.SetNX(context.Background(), key, "1", time.Duration(3)*time.Second).Val() {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}
	var vipMonthGift *dos.FcVipMonthGift
	global.G_DB.Model(&dos.FcVipMonthGift{}).Where("user_id=? AND month=?", userInfo.UserId, month).Take(&vipMonthGift)

	if len(vipMonthGift.Id) > 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}
	vipinfo := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: userInfo.Level})
	if vipinfo.MonthlyGift <= 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "奖金未配置")
		return
	}
	vipMonthGift = &dos.FcVipMonthGift{
		UserId:       userInfo.UserId,
		UserName:     userInfo.UserName,
		MerchantCode: userInfo.MerchantCode,
		VipName:      userInfo.Vip,
		Level:        userInfo.Level,
		BonusAmount:  vipinfo.MonthlyGift,
		Month:        month,
	}
	rRow := global.G_DB.Create(&vipMonthGift)

	if rRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, rRow.Error.Error())
		return
	}

	if len(vipMonthGift.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "领取失败")
		return
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		//修改用户金额表
		err := userTransfer.UserAmountChange(tx, vipinfo.MonthlyGift, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency, month+"月礼金", userInfo.UserId, "", "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "vip_month_gift"))
		if err != nil {
			global.G_LOG.Errorf("MonthGiftApply userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s", vipinfo.WeeklyGift, userInfo.UserId, err.Error())
			return err
		}

		err = global.G_DB.Model(&dos.FcVipMonthGift{}).Where("id=?", vipMonthGift.Id).Updates(map[string]interface{}{
			"bonus_amount_issue": vipinfo.MonthlyGift,
		}).Error

		if err != nil {
			return err
		}

		fop := &dos.FcOrderPromotion{
			ApplyAmount:  vipinfo.MonthlyGift,
			AppleRate:    1,
			ApplyType:    enmus.Promotion_VIP_MonthGift,
			Amount:       vipinfo.MonthlyGift,
			Status:       enmus.ORDER_YES_STATUS,
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			TurnOver:     0,
			MerchantCode: userInfo.MerchantCode,
			Remake:       fmt.Sprintf("第%s月按照VIP等级%s领取月礼金%0.2f", month, userInfo.Vip, vipinfo.MonthlyGift),
			//Currency:     ,
		}
		fop.OrderSn = vipMonthGift.Id //关联记录ID
		rowsAffected := tx.Create(fop).RowsAffected
		if rowsAffected != 1 {
			//global.G_LOG.Error(err.Error())
			return err
		}
		return err
	})

	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	global.G_REDIS.Expire(context.Background(), key, time.Duration(24*30)*time.Hour)
	response.SuccessMsgJSON(c, true, "领取成功")
}

func VipUpGiftApply(c *gin.Context) {
	level := tool.Atoi(c.DefaultQuery("level", "0"))
	if level < 1 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "请求参数错误")
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户不存在")
		return
	}

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	key := fmt.Sprintf("LockApplyVipUpGift::%s::%d", userInfo.UserId, level)

	if !global.G_REDIS.SetNX(context.Background(), key, "1", time.Duration(30)*time.Second).Val() {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}

	upLog := modules.FindByKeyFcUserVipRecordFirst(&dos.FcUserVipRecord{UserId: userInfo.UserId, Level: level})

	if len(upLog.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "领取失败还未晋级")
		return
	}

	if upLog.IssueBonus >= upLog.Bonus {
		response.FailErrJSON(c, response.ERROR_SERVER, "已经领取")
		return
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		//修改用户金额表
		err := userTransfer.UserAmountChange(tx, upLog.Bonus, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency, upLog.Vip+"晋级礼金", userInfo.UserId, "", "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "vip_up_gift"))
		if err != nil {
			global.G_LOG.Errorf("VipUpGiftApply userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s", upLog.Bonus, userInfo.UserId, err.Error())
			return err
		}

		err = global.G_DB.Model(&dos.FcUserVipRecord{}).Where("id=?", upLog.Id).Updates(map[string]interface{}{
			"issue_bonus": upLog.Bonus,
		}).Error

		if err != nil {
			return err
		}

		fop := &dos.FcOrderPromotion{
			ApplyAmount:  upLog.Bonus,
			AppleRate:    1,
			ApplyType:    enmus.Promotion_Vip_Upgradation,
			Amount:       upLog.Bonus,
			Status:       enmus.ORDER_YES_STATUS,
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			TurnOver:     0,
			MerchantCode: userInfo.MerchantCode,
			Remake:       fmt.Sprintf("晋级VIP等级%s领取礼金%0.2f", upLog.Vip, upLog.Bonus),
			//Currency:     ,
		}
		fop.OrderSn = upLog.Id //关联记录ID
		rowsAffected := tx.Create(fop).RowsAffected
		if rowsAffected != 1 {
			//global.G_LOG.Error(err.Error())
			return err
		}
		return err
	})

	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessMsgJSON(c, true, "领取成功")

}
