package activityControl

import (
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func GetPromotionInfo(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcPromotionInfo
	}{}
	jsonp.PageTimeQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageTimeQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.PromotionType = tool.Atoi(c.DefaultQuery("promotion_type", ""))
	jsonp.GameType = tool.Atoi(c.DefaultQuery("game_type", ""))

	jsonp.Status = 1 // 查询开启的活动

	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	//jsonp.ClientType = clientType
	//jsonp.Language = language
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	jsonp.MerchantCode = merchantCode

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcPromotionInfo(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPromotionInfo, &jsonp.PageTimeQuery, nil)
	newData := make([]*vo.PromotionInfoResp, len(data))
	tool.JsonMapper(data, &newData)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, newData)
}

func GetPromotionInfoDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	m := &dos.FcPromotionInfo{
		MerchantCode: merchantCode,
		//Language:     language,
	}
	m.Id = id
	data := modules.FindByKeyFcPromotionInfoFirst(m)
	var newData *vo.PromotionInfoResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}

// api: /api/user/activity/info - 活动信息
func ActivityInfo(c *gin.Context) {
	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// ============= 活动（回血包）以下
	remaingAmount, ratio := modules.GetUserRemainingAmountAndRechargeRatio(userInfo.UserId)
	modules.DoUserActivityAction(enmus.ActivityTypes_HealthPack,
		userInfo.UserId, []modules.ActivityActionParam{
			{
				RechargeBalanceRatio: ratio,
				Balance:              remaingAmount,
				RegTime:              userInfo.CreateTime.String(),
				FirstRechargeAmount:  modules.GetUserFirstRechargeAmount(userInfo.UserId),
			},
		})
	// ============= 活动（回血包）以上

	response.SuccessJSON(c, struct {
		ActivityHealthPackInfo  []modules.UserActivityHealthPackResult     `json:"activity_health_pack_info"`
		ActivityRedEnvelopeRain *modules.UserActivityRedEnvelopeRainResult `json:"activity_red_envelope_rain_info"`
	}{
		ActivityHealthPackInfo:  modules.GetUserActivityHealthPackInfo(userInfo.UserId),
		ActivityRedEnvelopeRain: modules.GetUserActivityRedEnvelopeRainInfo(userInfo.UserId, 0),
	})
}

// api: /api/user/activity/reward - 活动领取
func ActivityReward(c *gin.Context) {
	var jsonp struct {
		Id   string `json:"id"`
		Type int    `json:"type"`
	}

	c.ShouldBindJSON(&jsonp)

	if len(jsonp.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	activityConfig := modules.GetActivityConfig(jsonp.Id)
	if activityConfig == nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "活动配置错误")
		return
	}

	switch jsonp.Type {
	case enmus.ActivityTypes_HealthPack: // 回血包
		userActivityMap := modules.GetUserActivityDataToMap(userInfo.UserId)
		if v, ok := userActivityMap[jsonp.Id]; ok {
			switch v.Status {
			case enmus.UserActivityStats_None:
				response.FailErrJSON(c, response.ERROR_PARAMETER, "活动奖励不可领")
				return
			case enmus.UserActivityStats_Rewarded:
				response.FailErrJSON(c, response.ERROR_PARAMETER, "活动奖励已领取")
				return
			}

			err = global.G_DB.Transaction(func(tx *gorm.DB) error {
				err1 := userTransfer.UserAmountChange(tx, activityConfig.BonusAmount, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency,
					activityConfig.Title, userInfo.UserId, "", "",
					modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "activity"))
				if err1 != nil {
					global.G_LOG.Errorf("[ActivityReward][HealthPack.Reward] userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s",
						activityConfig.BonusAmount, userInfo.UserId, err1.Error())
					return err1
				}
				return err1
			})

			if err != nil {
				response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
				return
			}

			// 已领取
			v.Status = enmus.UserActivityStats_Rewarded
			userActivityMap[jsonp.Id] = v
			modules.SetUserActivityDataByMap(userInfo.UserId, userActivityMap)

			response.SuccessMsgJSON(c, struct {
				BonusAmount float64 `json:"bonus_amount"`
			}{
				BonusAmount: activityConfig.BonusAmount,
			}, "领取成功")
			return
		}
	case enmus.ActivityTypes_RedEnvelopeRain: // 红包雨
		res := modules.GetUserActivityRedEnvelopeRainInfo(userInfo.UserId, 10) // 缓冲结束10s
		if res == nil || res.DateRange == nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "活动已结束无法领取")
			return
		}

		if res.Cd > 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "活动未开始请稍后再领取")
			return
		}

		if len(res.DateRange.Amounts) == 0 { // 金额金区间必须存在
			response.FailErrJSON(c, response.ERROR_SERVER, "活动配置错误(无金额区间)")
			return
		}

		userActivityMap := modules.GetUserActivityDataToMap(userInfo.UserId)
		if v, ok := userActivityMap[jsonp.Id]; ok {
			var bonusAmount float64

			// 金额区间随机
			var totalWeight float64
			for _, v := range res.DateRange.Amounts {
				totalWeight += v.Prob
			}

			/*sort.Slice(res.DateRange.Amounts, func(i, j int) bool {
				return res.DateRange.Amounts[i].Prob < res.DateRange.Amounts[j].Prob
			})*/

			weight := tool.RandInt(0, int(totalWeight))
			fWeight := float64(weight)
			var currWeight float64
			for _, v := range res.DateRange.Amounts {
				currWeight += v.Prob
				if fWeight < currWeight {
					if len(v.AmountScope) == 2 {
						randAmount := tool.RandInt(int(v.AmountScope[0]*100), int(v.AmountScope[1]*100))
						bonusAmount = tool.TruncateFloat(float64(randAmount)/100, 2)
					}
					break
				}
			}

			if bonusAmount == 0 {
				global.G_LOG.Errorf("[ActivityReward][RedEnvelopeRain] bonusAmount is wrong: bonusAmount=0, userId=%s", userInfo.UserId)
				response.FailErrJSON(c, response.ERROR_SERVER, "活动配置错误(金额区间)")
				return
			}

			err = global.G_DB.Transaction(func(tx *gorm.DB) error {
				err1 := userTransfer.UserAmountChange(tx, bonusAmount, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency,
					activityConfig.Title, userInfo.UserId, "", "",
					modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "activity"))
				if err1 != nil {
					global.G_LOG.Errorf("[ActivityReward][RedEnvelopeRain.Reward] userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s",
						bonusAmount, userInfo.UserId, err1.Error())
					return err1
				}
				return err1
			})

			if err != nil {
				response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
				return
			}

			v.Status = enmus.UserActivityStats_Rewarded
			userActivityMap[jsonp.Id] = v
			modules.SetUserActivityDataByMap(userInfo.UserId, userActivityMap)

			response.SuccessMsgJSON(c, struct {
				BonusAmount float64 `json:"bonus_amount"`
			}{
				BonusAmount: bonusAmount,
			}, "领取成功")
			return
		}
	}

	response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
}
