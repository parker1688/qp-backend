package dailyBonusControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 签到信息
func DailyBonusInfo(c *gin.Context) {
	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	// 签到触发及获取用户签到数据
	modules.DoDailyBonusAction(userInfo.UserId, merchantCode)

	// 获取签到信息
	result, err1 := modules.GetDailyBonusInfo(userInfo.UserId, merchantCode)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	response.SuccessJSON(c, result)
}

// 签到领取
func DailyBonusReward(c *gin.Context) {
	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	dailyBonusData, err1 := modules.GetUserDailyBonusData(userInfo.UserId, true)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err1.Error())
		return
	}

	// 判断签到状态
	switch dailyBonusData.Status {
	case enmus.DailyBonusStats_None:
		response.FailErrJSON(c, response.ERROR_SERVER,
			fmt.Sprintf("当日打码量不足%.2f, 不可签到",
				modules.GetDailyBonusDiffAmount(merchantCode, dailyBonusData)))
		return
	case enmus.DailyBonusStats_Rewarded:
		response.FailErrJSON(c, response.ERROR_SERVER, "签到奖励已领取")
		return
	}

	var dailyBonusConfig *dos.DailyBonusField
	var dailyBonusExtConfig *dos.DailyBonusField

	// 处理可领取
	dailyBonusConfigs, err2 := modules.GetDailyBonusConfig(merchantCode)
	if err2 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err2.Error())
		return
	}
	for _, config := range dailyBonusConfigs {
		// 基础奖励
		if config.Days == dailyBonusData.Day &&
			config.Reward > 0 {
			dailyBonusConfig = &config
		}

		// 额外奖励
		if config.Days == dailyBonusData.ExtDay &&
			config.ExtraReward > 0 {
			dailyBonusExtConfig = &config
		}
	}

	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		var err3 error
		// 领取基础奖励
		if dailyBonusConfig != nil {
			err3 = userTransfer.UserAmountChange(tx, dailyBonusConfig.Reward, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency,
				fmt.Sprintf("第%d天", dailyBonusConfig.Days), userInfo.UserId, "", "",
				modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus"))
			if err3 != nil {
				global.G_LOG.Errorf("[DailyBonusReward][Reward] userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s",
					dailyBonusConfig.Reward, userInfo.UserId, err3.Error())
				return err3
			}
		}

		// 领取额外奖励
		if dailyBonusExtConfig != nil {
			err3 = userTransfer.UserAmountChange(tx, dailyBonusExtConfig.ExtraReward, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency,
				fmt.Sprintf("第%d天", dailyBonusExtConfig.Days), userInfo.UserId, "", "",
				modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus_extra"))
			if err3 != nil {
				global.G_LOG.Errorf("[DailyBonusReward][ExtraReward] userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s",
					dailyBonusExtConfig.ExtraReward, userInfo.UserId, err3.Error())
				return err3
			}
		}

		return err3
	})

	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	dailyBonusData.Status = enmus.DailyBonusStats_Rewarded // 更新状态为已领
	modules.SetUserDailyBonusData(userInfo.UserId, dailyBonusData)

	response.SuccessMsgJSON(c, nil, "领取成功")
}
