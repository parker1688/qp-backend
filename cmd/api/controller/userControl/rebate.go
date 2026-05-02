package userControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	commonResp "bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/userTransfer"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func RebateRecordList(c *gin.Context) {
	jsonp := struct {
		commonResp.PageTimeQuery
		dos.FcUserRebateRecords
	}{}
	jsonp.PageTimeQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageTimeQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.TimeType = tool.Atoi(c.DefaultQuery("time_type", "-1"))

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	jsonp.UserId = userInfo.UserId

	data, total, totalBonusAmount := modules.FcUserRebateRecordsList(&jsonp.FcUserRebateRecords, &jsonp.PageTimeQuery)
	//newData := make([]*vo.FcUserRebateRecordsVO, len(data))
	//tool.JsonMapper(&data, &newData)

	respData := vo.FcUserRebateRecordsListResp{}
	respData.TotalBonusAmount = totalBonusAmount
	respData.RebateData = data

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, respData)
}

// 返水详情记录列表接口
func RebateDetailRecordList(c *gin.Context) {
	jsonp := struct {
		Date string `json:"date"`
	}{}

	jsonp.Date = c.DefaultQuery("date", "")

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	data := modules.GetFcUserRebateDetailRecordList(userInfo.UserId, jsonp.Date)
	response.SuccessJSON(c, data)
}

// 领取返水，洗码
func RebateApply(c *gin.Context) {
	var jsonp vo.FcUserRebateApplyReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	// 判断是否达到可领取返水
	/*needRebateBonusVal := GetDictConfigRebateMinBonusAmount(c.GetHeader(vo.MerchantCode_KEY_G))
	if modules.GetCacheUserRebateBonusValue(userInfo.UserId) < needRebateBonusVal {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("返水金额未达到领取要求(%.f)", needRebateBonusVal))
		return
	}*/

	flowKey := fmt.Sprintf(enmus.UserRebateFlow, userInfo.UserId)
	if jsonp.UserId != userInfo.UserId {
		global.G_LOG.Errorf("jsonp.UserId: %s not match userInfo.UserId: %s", jsonp.UserId, userInfo.UserId)
		response.FailErrJSON(c, response.ERROR_SERVER, "userId not match")
		return
	}
	if jsonp.UserName != userInfo.UserName {
		global.G_LOG.Errorf("jsonp.UserName: %s not match userInfo.UserName: %s", jsonp.UserName, userInfo.UserName)
		response.FailErrJSON(c, response.ERROR_SERVER, "userName not match")
		return
	}
	lockKey := fmt.Sprintf(enmus.ApplyUserRebateFlowLock, userInfo.UserId)
	isWait := global.G_REDIS.SetNX(context.Background(), lockKey, "1", 8*time.Second).Val()
	if !isWait { //操作频繁
		response.SuccessMsgJSON(c, nil, langs.GetWithLocaleGin(c, "opt_frequent"))
		return
	}

	// 获取用户的洗码
	flowMap, err := global.G_REDIS.HGetAll(context.Background(), flowKey).Result()
	if err != nil {
		global.G_LOG.Errorf("Get RebateApply username=%s key=%s err: %s", userInfo.UserName, flowKey, err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if len(flowMap) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "no rebate data")
		return
	}

	rebateVenueCode := ""

	gameTypeArr := []string{}            // 获取所有不重复的 gameType
	betAmountMap := map[string]float64{} // 洗码map, key 为 gameType
	rebateDatas := []*dos.FcUserRebateRecords{}
	for k, v := range flowMap {
		kArr := strings.Split(k, ":")
		if len(kArr) < 2 {
			global.G_LOG.Errorf("RebateApply username=%s key=%s invalid", userInfo.UserName, k)
			continue
		}
		amount, err := strconv.ParseFloat(v, 64)
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s amount=%s invalid", userInfo.UserName, v)
			continue
		}

		//amount = amount * (1 - GetDictRebateCostRateValue())

		venueCode := kArr[0]
		gameType := kArr[1]

		rebateVenueCode = venueCode

		tmp := &dos.FcUserRebateRecords{}
		tmp.UserId = userInfo.UserId
		tmp.UserName = userInfo.UserName
		tmp.MerchantCode = userInfo.MerchantCode
		tmp.Level = userInfo.Level
		tmp.VenueCode = venueCode
		tmp.GameType = gameType
		tmp.BetAmount = amount
		tmp.Status = 1
		tmp.RebateType = 3
		tmp.CreateBy = "system"
		tmp.UpdateBy = "system"
		rebateDatas = append(rebateDatas, tmp)

		tmpBet, ok := betAmountMap[gameType]
		if !ok {
			betAmountMap[gameType] = amount
			gameTypeArr = append(gameTypeArr, gameType)
		} else {
			betAmountMap[gameType] = amount + tmpBet
		}
	}

	// 获取用户游戏累积打码量, 如果没有查到则就直接使用打码过来的数据
	queryCtx, queryCancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer queryCancel()
	userGameReportRows := []*dos.FcUserGameReport{}
	err = global.G_DB.WithContext(queryCtx).Model(&dos.FcUserGameReport{}).Where("user_id = ? AND game_type IN ?", userInfo.UserId, gameTypeArr).Find(&userGameReportRows).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			global.G_LOG.Errorf("RebateApply query FcUserGameReport gameType: %v err: %v", gameTypeArr, err)
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
	}
	// 用户历史累积游戏类型数据分类, key 为 gameType
	userGameReporMap := map[string]*dos.FcUserGameReport{}
	for i, v := range userGameReportRows {
		tmp := userGameReportRows[i]
		userGameReporMap[v.GameType] = tmp
	}

	// 获取所有类型的返水比例数据根据 gameType
	gameRebateRows := []*dos.FcGameRebate{}
	err = global.G_DB.WithContext(queryCtx).Model(&dos.FcGameRebate{}).Where("game_type IN ?", gameTypeArr).Order("bonus_rate asc").Find(&gameRebateRows).Error
	if err != nil {
		global.G_LOG.Errorf("RebateApply query FcGameRebate gameType: %v err: %v", gameTypeArr, err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if len(gameRebateRows) == 0 {
		response.FailErrJSON(c, response.ERROR_SERVER, "no rebate config data")
		return
	}
	gameMaxRebateMap := map[string]*dos.FcGameRebate{} // 存储每个游戏类型最大的返水比例和返水上限
	gameMinRebateMap := map[string]*dos.FcGameRebate{} // 存储每个游戏类型最小的返水比例和返水上限
	for i := range gameRebateRows {
		row := gameRebateRows[i]
		tmpMax, ok := gameMaxRebateMap[row.GameType]
		if !ok {
			gameMaxRebateMap[row.GameType] = row
		} else {
			if tmpMax.MaxBetAmount < row.MaxBetAmount {
				gameMaxRebateMap[row.GameType] = row
			}
		}

		tmpMin, ok := gameMinRebateMap[row.GameType]
		if !ok {
			gameMinRebateMap[row.GameType] = row
		} else {
			if tmpMin.MaxBetAmount > row.MaxBetAmount {
				gameMinRebateMap[row.GameType] = row
			}
		}
	}

	// 获取上次用户各个返水区间的打码量
	flowHisKey := fmt.Sprintf(enmus.UserHisRebateFlow, userInfo.UserId)
	flowHisMap, err := global.G_REDIS.HGetAll(context.Background(), flowHisKey).Result()
	if err != nil {
		global.G_LOG.Errorf("Get RebateApply username=%s flowHisKey=%s err: %s", userInfo.UserName, flowHisKey, err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if len(flowHisMap) == 0 {
		flowHisMap = map[string]string{}
	}
	gameTypeRebateRateMap := map[string]float64{} // 游戏类型返水比例流水  key 为 gameType:rebate比例，值为流水
	for k, v := range flowHisMap {
		// key 为 gameType:rebate比例, 值为流水
		amount, err := strconv.ParseFloat(v, 64)
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s flowHisMapKey=%s amount=%s invalid", userInfo.UserName, v)
			continue
		}

		gameTypeRebateRateMap[k] = amount
	}

	// 获取返水比例以及返水记录
	userRebateRecords := []*dos.FcUserRebateRecords{}
	limitRecords := []*dos.FcUserRebateRecords{} // 超过最大返水的数据
	for k1, v1 := range rebateDatas {
		tmpRebateData := rebateDatas[k1]
		tmpGameType := v1.GameType
		//fmt.Printf("=======================username: %v venueCode: %v gameType: %v betAmount: %v\n", v1.UserName, v1.VenueCode, v1.GameType, v1.BetAmount)

		batchFlow, ok := betAmountMap[tmpGameType]
		if !ok {
			continue
		}
		venueFlow := v1.BetAmount // 场馆流水
		if venueFlow == 0 {
			continue
		}

		// 如果历史累积打码量不存在，则用刚获取的数据量进行打码计算
		ugrRow, ok := userGameReporMap[tmpGameType]
		if !ok {
			ugrRow = &dos.FcUserGameReport{}
		}

		// 历史累积的游戏类型流水
		hisFlow := ugrRow.RebateFlow
		totalFlow := hisFlow + batchFlow
		//fmt.Printf("=====================hisFlow: %v totalFlow: %v\n", hisFlow, totalFlow)

		tmpRebateData.HisBetAmount = hisFlow
		tmpRebateData.BatchBetAmount = batchFlow

		// 如果超过最大返水配置的打码量则直接进行计算
		tmpGameMaxRebate, ok := gameMaxRebateMap[tmpGameType]
		if ok {
			// 历史流水超过最大值
			if hisFlow >= tmpGameMaxRebate.MaxBetAmount {
				tmpRebateData.BonusRate = tmpGameMaxRebate.BonusRate
				tmpRebateData.BonusAmount = venueFlow * tmpGameMaxRebate.BonusRate
				userRebateRecords = append(userRebateRecords, tmpRebateData)
				continue
			}
		} else {
			global.G_LOG.Errorf("RebateApply username: %v venueCode: %v gameType: %v flow: %v not support", userInfo.UserName, v1.VenueCode, tmpGameType, v1.BetAmount)
			continue
		}

		// 计算阶梯式返水比例
		thSumBet := 0.00
		thMaxRate := 0.00
		for _, v2 := range gameRebateRows {
			if v2.GameType != tmpGameType {
				continue
			}

			// 最小流水不满足
			if v2.MinBetAmount > totalFlow {
				continue
			}
			tmpRate := v2.BonusRate
			if tmpRate > thMaxRate {
				thMaxRate = tmpRate
			}

			// 如果历史流水比该层次最大流水大则直接跳过找再高一层次流水配置
			if hisFlow > v2.MaxBetAmount {
				continue
			}
			venueFlow2 := tmpRebateData.BetAmount
			if venueFlow2 == 0.00 {
				continue
			}

			// 返水区间
			rebateInterval := v2.MaxBetAmount - v2.MinBetAmount
			// 如果是最低的返水区间，则区间值为最低返水区间的最大值
			tmpGameMinRebate, ok := gameMinRebateMap[tmpGameType]
			if ok {
				if v2.BonusRate == tmpGameMinRebate.BonusRate {
					rebateInterval = tmpGameMinRebate.MaxBetAmount
				}
			}

			key1 := tmpGameType + ":" + strconv.FormatFloat(v2.BonusRate, 'f', -1, 64)
			// 游戏类型返水比例对应的累计流水
			gameTypeRebateRateFlow, ok := gameTypeRebateRateMap[key1]
			if !ok {
				gameTypeRebateRateFlow = 0.0
			}
			//fmt.Printf("gameTypeRebateRateFlow: %v key1: %v venueFlow2: %v rebateInterval: %v==============\n", gameTypeRebateRateFlow, key1, venueFlow2, rebateInterval)
			// 该类型区间流水已返水
			if gameTypeRebateRateFlow >= rebateInterval {
				continue
			}

			// 如果场馆流水+游戏类型返水比例对应的累计流水大于返水区间值
			if venueFlow2+gameTypeRebateRateFlow > rebateInterval {
				// 剩余的区间返水值
				venueAvailFlow := rebateInterval - gameTypeRebateRateFlow

				// 更改数据,剩余的数据接着下一次再算
				tmp2 := &dos.FcUserRebateRecords{}
				tool.JsonMapper(tmpRebateData, tmp2)
				tmpRebateData.BetAmount = venueFlow2 - venueAvailFlow

				tmp2.BetAmount = venueAvailFlow
				tmp2.BonusRate = v2.BonusRate
				tmp2.BonusAmount = tmp2.BetAmount * v2.BonusRate
				userRebateRecords = append(userRebateRecords, tmp2)
				thSumBet += tmp2.BetAmount

				//fmt.Printf("gameTypeRebateRateMap111 key=%v value=%v\n", key1, rebateInterval)
				gameTypeRebateRateMap[key1] = rebateInterval
				continue
			} else { // 如果场馆流水+游戏类型返水比例对应的累计流水小于返水区间值, 直接 continue
				tmpRebateData.BonusRate = v2.BonusRate
				tmpRebateData.BonusAmount = tmpRebateData.BetAmount * v2.BonusRate
				userRebateRecords = append(userRebateRecords, tmpRebateData)

				thSumBet += tmpRebateData.BetAmount

				//fmt.Printf("gameTypeRebateRateMap222 key=%v value=%v\n", key1, gameTypeRebateRateFlow+venueFlow2)
				gameTypeRebateRateMap[key1] = gameTypeRebateRateFlow + venueFlow2
				break
			}
		}

		//fmt.Printf("thMaxRate: %v maxRate: %v thSumBet: %v venueFlow: %v\n", thMaxRate, tmpGameMaxRebate.BonusRate, thSumBet, venueFlow)
		// 如果到达了最大返水率 且 还有剩余的流水量
		if thMaxRate == tmpGameMaxRebate.BonusRate && thSumBet < venueFlow {
			limitRecords = append(limitRecords, tmpRebateData)
		}
	}

	// 处理超过最大返水的数据
	for i, v := range limitRecords {
		row := limitRecords[i]
		tmpGameType := v.GameType

		// 如果超过最大返水配置的打码量则直接进行计算
		tmpGameMaxRebate, ok := gameMaxRebateMap[tmpGameType]
		if ok {
			row.BonusRate = tmpGameMaxRebate.BonusRate
			row.BonusAmount = row.BetAmount * tmpGameMaxRebate.BonusRate
			userRebateRecords = append(userRebateRecords, row)
		}
	}

	nowTime := automaticType.Now()
	totalBonusAmount := 0.00

	hisFlowMap := map[string]float64{}        // 游戏类型历史缓存流水, key 为 gameType
	useUpBetAmountMap := map[string]float64{} // 用户消耗的场馆流水, key 为 venueCode:gameType
	for i := range userRebateRecords {
		row := userRebateRecords[i]
		row.CreateTime = nowTime
		row.UpdateTime = nowTime

		if row.BetAmount <= 0.00 {
			continue
		}
		totalBonusAmount += row.BonusAmount

		tmpFlow, ok := hisFlowMap[row.GameType]
		if !ok {
			hisFlowMap[row.GameType] = row.BetAmount
		} else {
			hisFlowMap[row.GameType] = tmpFlow + row.BetAmount
		}

		// 用户消耗的场馆流水, key 为 venueCode:gameType
		useUpBetKey := row.VenueCode + ":" + row.GameType
		tmpUseUpBetAmount, ok := useUpBetAmountMap[useUpBetKey]
		if !ok {
			useUpBetAmountMap[useUpBetKey] = row.BetAmount
		} else {
			useUpBetAmountMap[useUpBetKey] = tmpUseUpBetAmount + row.BetAmount
		}
	}

	if totalBonusAmount == 0.00 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "not enough game data")
		return
	}

	totalBonusAmount = tool.TruncateFloat(totalBonusAmount, 2)

	// 给用户返水加钱
	optType := userTransfer.TranRebate
	currency := global.CONFIG.General.DefaultCurrency
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		err = userTransfer.UserAmountChange(tx, totalBonusAmount, optType, currency, "", userInfo.UserId, "system", "",
			rebateVenueCode+modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "rebate"))
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s updateUserMoney err: %s", userInfo.UserName, err.Error())
			return err
		}
		return err
	})
	if err != nil {
		global.G_LOG.Errorf("RebateApply username=%s userAmountChange err: %v", userInfo.UserName, err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	// 处理返水报表以及个人返水的一些数据
	go RebateDataReportHandle(userInfo, userRebateRecords, totalBonusAmount, userGameReportRows, hisFlowMap, useUpBetAmountMap, gameTypeRebateRateMap)

	resp := &vo.FcUserRebateApplyResp{}
	resp.BonusAmount = totalBonusAmount

	response.SuccessMsgJSON(c, resp, "领取成功")

	// modules.DelCacheUserRebateBonusValue(userInfo.UserId)
}

// 获取洗码数据
func GetRebateData(c *gin.Context) {
	var jsonp vo.FcUserRebateApplyReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	flowKey := fmt.Sprintf(enmus.UserRebateFlow, userInfo.UserId)
	if jsonp.UserId != userInfo.UserId {
		global.G_LOG.Errorf("jsonp.UserId: %s not match userInfo.UserId: %s", jsonp.UserId, userInfo.UserId)
		response.FailErrJSON(c, response.ERROR_SERVER, "userId not match")
		return
	}
	if jsonp.UserName != userInfo.UserName {
		global.G_LOG.Errorf("jsonp.UserName: %s not match userInfo.UserName: %s", jsonp.UserName, userInfo.UserName)
		response.FailErrJSON(c, response.ERROR_SERVER, "userName not match")
		return
	}

	//lockKey := fmt.Sprintf(enmus.GetUserRebateFlowLock, userInfo.UserId)
	//isWait := global.G_REDIS.SetNX(context.Background(), lockKey, "1", 8*time.Second).Val()
	//if !isWait { //操作频繁
	//	response.SuccessMsgJSON(c, nil, langs.GetWithLocaleGin(c, "opt_frequent"))
	//	return
	//}

	// 获取用户的洗码
	flowMap, err := global.G_REDIS.HGetAll(context.Background(), flowKey).Result()
	if err != nil {
		global.G_LOG.Errorf("Get RebateApply username=%s key=%s err: %s", userInfo.UserName, flowKey, err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	//totalBetAmount := 0.0
	resp := &vo.FcUserRebateDataResp{
		MinBonusAmount: GetDictConfigRebateMinBonusAmount(userInfo.MerchantCode),
	}
	resp.CurrDetailList = []vo.FcUserRebateValsResult{}
	//resp.TotalBetAmount = totalBetAmount
	if len(flowMap) == 0 {
		response.SuccessMsgJSON(c, resp, "success")
		return
	}

	/*for k, v := range flowMap {
		kArr := strings.Split(k, ":")
		if len(kArr) < 2 {
			global.G_LOG.Errorf("RebateApply username=%s key=%s invalid", userInfo.UserName, k)
			continue
		}
		amount, err := strconv.ParseFloat(v, 64)
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s amount=%s invalid", userInfo.UserName, v)
			continue
		}

		//amount = amount * (1 - GetDictRebateCostRateValue())

		totalBetAmount += amount
	}*/

	rebateLis, totalBetAmount, totalBonusAmount := PreCalculateUserRebateVals(userInfo.UserId) // 预计算返水数据
	resp.TotalBetAmount = totalBetAmount
	resp.CurrDetailList = rebateLis
	resp.TotalBonusAmount = totalBonusAmount

	response.SuccessMsgJSON(c, resp, "success")
}

// 洗码介绍接口
func RebateIntro(c *gin.Context) {
	response.SuccessJSON(c, modules.GetFcGameRebateListAll())
}

// 处理用户的返水报表等数据
func RebateDataReportHandle(userInfo *dos.FcUserMaterial, userRebateRecords []*dos.FcUserRebateRecords, totalBonusAmount float64, userGameReportRows []*dos.FcUserGameReport,
	hisFlowMap map[string]float64, useUpBetAmountMap map[string]float64, gameTypeRebateRateMap map[string]float64) error {

	var err error
	flowHisKey := fmt.Sprintf(enmus.UserHisRebateFlow, userInfo.UserId)
	flowKey := fmt.Sprintf(enmus.UserRebateFlow, userInfo.UserId)

	// 将用户消耗的场馆游戏流水存入 redis, feild 为 venueCode:gameType
	for k, v := range useUpBetAmountMap {
		/*err = global.G_REDIS.HIncrByFloat(context.Background(), flowKey, k, -v).Err()
		if err != nil {
			global.G_LOG.Errorf("RebateApply redis UpdateUserVenueFlow flowKey: %s feild: %v flow: %v err: %s", flowKey, k, -v, err.Error())
		}*/

		val := global.G_REDIS.HGet(context.Background(), flowKey, k).Val()
		amount, err := strconv.ParseFloat(val, 64)
		if err != nil {
			global.G_LOG.Errorf("[RebateDataReportHandle] Parse float failed: %v", err.Error())
			continue
		}

		amount -= v
		if amount < 0 {
			amount = 0
		}

		err = global.G_REDIS.HSet(context.Background(), flowKey, k, amount).Err()
		if err != nil {
			global.G_LOG.Errorf("[RebateDataReportHandle] Set cache failed: flowKey=%s, k=%s, err=%v", flowKey, k, err.Error())
		}
	}
	// 设置 key 的过期时间为 3 个月
	err = global.G_REDIS.Expire(context.Background(), flowKey, 24*time.Hour*90).Err()
	if err != nil {
		global.G_LOG.Errorf("RebateApply redis UpdateUserVenueFlow flowKey: %s expire err: %s", flowHisKey, err.Error())
	}

	// 将用户各个游戏的返水区间值存入 redis
	for k, v := range gameTypeRebateRateMap {
		//global.G_LOG.Infof("RebateApply username=%s flowHisKey =%s feild=%s value=%v", userInfo.UserName, flowHisKey, k, v)
		global.G_REDIS.HSet(context.Background(), flowHisKey, k, v).Val()
	}
	// 设置 key 的过期时间为 3 个月
	err = global.G_REDIS.Expire(context.Background(), flowHisKey, 24*time.Hour*90).Err()
	if err != nil {
		global.G_LOG.Errorf("RebateApply redis UpdateUserHisFlow flowHisKey: %s expire err: %s", flowHisKey, err.Error())
	}

	for i := range userRebateRecords {
		row := userRebateRecords[i]
		if row.BetAmount <= 0.00 {
			continue
		}

		err = global.G_DB.Save(row).Error
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s save err: %s", userInfo.UserName, err.Error())
			continue
		}

		if global.CONFIG.Mq.IsInit {
			promoData := channelData.UserPromotionMessage{}
			promoData.UserId = userInfo.UserId
			promoData.UserName = userInfo.UserName
			promoData.OrderSn = row.Id
			promoData.ForceStatus = 1
			promoData.T = time.Now().UnixMicro()
			promoData.PromotionTime = automaticType.Now().String()
			promoData.PromotionAmount = row.BonusAmount
			promoData.PromotionType = 2 // 返水

			// 发送红利消息给消息队列
			err = channelData.SendUserPromotion(&promoData)
			if err != nil {
				global.G_LOG.Errorf(" SendUserPromotion kafka err: %v , msg: %s", err, tool.String(promoData))
			}
		}

		//global.G_LOG.Infof("RebateApply username=%s HisBetAmount: %v BatchBetAmount: %v betAmount: %v BonusAmount: %v success",
		//	row.UserName, row.HisBetAmount, row.BatchBetAmount, row.BetAmount, row.BonusAmount)
	}

	// 用户返水报表
	rebateReport := dos.FcRebateReport{}
	err = global.G_DB.Model(&dos.FcRebateReport{}).Where("user_id = ?", userInfo.UserId).First(&rebateReport).Error
	if err != nil {
		// 不存在则插入
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tmpData := dos.FcRebateReport{}
			tmpData.UserId = userInfo.UserId
			tmpData.UserName = userInfo.UserName
			tmpData.MerchantCode = userInfo.MerchantCode
			tmpData.Amount = totalBonusAmount

			_, errR := modules.SaveFcRebateReport(&tmpData)
			if errR != nil {
				global.G_LOG.Errorf("RebateApply SaveFcRebateReport username: %v err: %s", userInfo.UserName, errR.Error())
			}
		} else {
			global.G_LOG.Errorf("RebateApply query user FcRebateReport username=%s  err: %v", userInfo.UserName, err)
		}
	} else {
		// 更新用户返水报表
		err = global.G_DB.Model(&dos.FcRebateReport{}).Where("id = ?", rebateReport.Id).Update("amount", gorm.Expr("amount + ?", totalBonusAmount)).Error
		if err != nil {
			global.G_LOG.Errorf("RebateApply update FcRebateReport username=%s  err: %v", userInfo.UserName, err)
		}
	}

	// 更新用户流水累积
	for i := range userGameReportRows {
		row := userGameReportRows[i]
		tmpFlow, ok := hisFlowMap[row.GameType]
		if !ok {
			continue
		}

		err = global.G_DB.Model(&dos.FcUserGameReport{}).Where("id = ?", row.Id).Update("rebate_flow", gorm.Expr("rebate_flow + ?", tmpFlow)).Error
		if err != nil {
			global.G_LOG.Errorf("RebateApply username=%s update FcUserGameReport err: %s", userInfo.UserName, err.Error())
			continue
		}
	}

	return nil
}

// 获取返水全局扣量值
func GetDictRebateCostRateValue() float64 {
	rebateCostRate := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "Server_System_Settings",
		DictsTag:      "RebateCostRate", // 约定如下：10%则字典配置填写0.1
	})

	rebateCostRateDefault := 0.0

	if rebateCostRate != nil {
		val, err := strconv.ParseFloat(rebateCostRate.DictsValue, 64)
		if err != nil {
			global.G_LOG.Errorf("[GetRebateConfigCostRate] ParseFloat failed: %v", rebateCostRate.DictsValue)
			return rebateCostRateDefault
		}

		if val < 0 {
			global.G_LOG.Errorf("[GetRebateConfigCostRate] Configure dict wrong: %v", val)
			return rebateCostRateDefault
		}

		return val
	} else {
		global.G_LOG.Error("[GetRebateConfigCostRate] Can't find Server_System_Settings.RebateCostRate")
		return rebateCostRateDefault
	}
}

// GetDictConfigRebateMinBonusAmount - 获取字典反水领取最低金额
// @param {string} merchantCode 商户码
// @returns float64
func GetDictConfigRebateMinBonusAmount(merchantCode string) float64 {
	minBonusAmountDefault := 0.01

	result := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "min_bonus_amount",
		DictsTag:      merchantCode,
		Status:        1,
	})

	if result == nil {
		global.G_LOG.Errorf("[GetDictRebateMinBonusAmountValue] Find dict min_bonus_amount failed: merchantCode=%s", merchantCode)
		return minBonusAmountDefault
	}

	val, err := strconv.ParseFloat(result.DictsValue, 64)
	if err != nil {
		global.G_LOG.Errorf("[GetDictRebateMinBonusAmountValue] ParseFloat failed: %v", result.DictsValue)
		return minBonusAmountDefault
	}

	return val
}

// PreCalculateUserRebateVals - 预计算用户返水数据
// @param {string} userId 用户Id
// @returns []vo.FcUserRebateValsResult, float64
func PreCalculateUserRebateVals(userId string) ([]vo.FcUserRebateValsResult, float64, float64) {
	result := []vo.FcUserRebateValsResult{}
	totalBetAmount := 0.00
	totalBonusAmount := 0.00

	flowKey := fmt.Sprintf(enmus.UserRebateFlow, userId)

	// 获取用户的洗码
	flowMap, err := global.G_REDIS.HGetAll(context.Background(), flowKey).Result()
	if err != nil {
		global.G_LOG.Errorf("[PreCalculateUserRebateVals] Get flowMap failed: userId=%s key=%s err: %s", userId, flowKey, err.Error())
	}
	if len(flowMap) == 0 {
		return result, totalBetAmount, totalBonusAmount
	}

	gameTypeArr := []string{}            // 获取所有不重复的 gameType
	betAmountMap := map[string]float64{} // 洗码map, key 为 gameType
	rebateDatas := []*dos.FcUserRebateRecords{}
	for k, v := range flowMap {
		kArr := strings.Split(k, ":")
		if len(kArr) < 2 {
			global.G_LOG.Errorf("[PreCalculateUserRebateVals] flowMap format wrong: userId=%s, key=%s", userId, k)
			continue
		}
		amount, err := strconv.ParseFloat(v, 64)
		if err != nil {
			global.G_LOG.Errorf("[PreCalculateUserRebateVals] Parse flowMap amount wrong: userId=%s, amount=%s", userId, v)
			continue
		}

		venueCode := kArr[0]
		gameType := kArr[1]

		tmp := &dos.FcUserRebateRecords{}
		tmp.UserId = userId
		tmp.VenueCode = venueCode
		tmp.GameType = gameType
		tmp.BetAmount = amount
		rebateDatas = append(rebateDatas, tmp)

		tmpBet, ok := betAmountMap[gameType]
		if !ok {
			betAmountMap[gameType] = amount
			gameTypeArr = append(gameTypeArr, gameType)
		} else {
			betAmountMap[gameType] = amount + tmpBet
		}
	}

	// 获取用户游戏累积打码量, 如果没有查到则就直接使用打码过来的数据
	queryCtx, queryCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer queryCancel()
	userGameReportRows := []*dos.FcUserGameReport{}
	err = global.G_DB.WithContext(queryCtx).Model(&dos.FcUserGameReport{}).Where("user_id = ? AND game_type IN ?", userId, gameTypeArr).Find(&userGameReportRows).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			global.G_LOG.Errorf("[PreCalculateUserRebateVals] Query FcUserGameReport failed: gameType=%v, err=%v", gameTypeArr, err.Error())
			return result, totalBetAmount, totalBonusAmount
		}
	}

	// 用户历史累积游戏类型数据分类, key 为 gameType
	userGameReporMap := map[string]*dos.FcUserGameReport{}
	for i, v := range userGameReportRows {
		tmp := userGameReportRows[i]
		userGameReporMap[v.GameType] = tmp
	}

	// 获取所有类型的返水比例数据根据 gameType
	gameRebateRows := []*dos.FcGameRebate{}
	err = global.G_DB.WithContext(queryCtx).Model(&dos.FcGameRebate{}).Where("game_type IN ?", gameTypeArr).Order("bonus_rate asc").Find(&gameRebateRows).Error
	if err != nil {
		global.G_LOG.Errorf("[PreCalculateUserRebateVals] Query FcGameRebate failed: gameType=%v, err=%v", gameTypeArr, err.Error())
		return result, totalBetAmount, totalBonusAmount
	}
	if len(gameRebateRows) == 0 {
		global.G_LOG.Errorf("[PreCalculateUserRebateVals] Query FcGameRebate is empty: gameType=%v", gameTypeArr)
		return result, totalBetAmount, totalBonusAmount
	}
	gameMaxRebateMap := map[string]*dos.FcGameRebate{} // 存储每个游戏类型最大的返水比例和返水上限
	gameMinRebateMap := map[string]*dos.FcGameRebate{} // 存储每个游戏类型最小的返水比例和返水上限
	for i := range gameRebateRows {
		row := gameRebateRows[i]
		tmpMax, ok := gameMaxRebateMap[row.GameType]
		if !ok {
			gameMaxRebateMap[row.GameType] = row
		} else {
			if tmpMax.MaxBetAmount < row.MaxBetAmount {
				gameMaxRebateMap[row.GameType] = row
			}
		}

		tmpMin, ok := gameMinRebateMap[row.GameType]
		if !ok {
			gameMinRebateMap[row.GameType] = row
		} else {
			if tmpMin.MaxBetAmount > row.MaxBetAmount {
				gameMinRebateMap[row.GameType] = row
			}
		}
	}

	// 获取上次用户各个返水区间的打码量
	flowHisKey := fmt.Sprintf(enmus.UserHisRebateFlow, userId)
	flowHisMap, err := global.G_REDIS.HGetAll(context.Background(), flowHisKey).Result()
	if err != nil {
		global.G_LOG.Errorf("[PreCalculateUserRebateVals] Get flowHisMap failed: userId=%s, flowHisKey=%s, err=%s", userId, flowHisKey, err.Error())
	}
	if len(flowHisMap) == 0 {
		flowHisMap = map[string]string{}
	}
	gameTypeRebateRateMap := map[string]float64{} // 游戏类型返水比例流水  key 为 gameType:rebate比例，值为流水
	for k, v := range flowHisMap {
		// key 为 gameType:rebate比例, 值为流水
		amount, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		gameTypeRebateRateMap[k] = amount
	}

	// 获取返水比例以及返水记录
	userRebateRecords := []*dos.FcUserRebateRecords{}
	limitRecords := []*dos.FcUserRebateRecords{} // 超过最大返水的数据
	for k1, v1 := range rebateDatas {
		tmpRebateData := rebateDatas[k1]
		tmpGameType := v1.GameType

		batchFlow, ok := betAmountMap[tmpGameType]
		if !ok {
			continue
		}
		venueFlow := v1.BetAmount // 场馆流水
		if venueFlow == 0 {
			continue
		}

		// 如果历史累积打码量不存在，则用刚获取的数据量进行打码计算
		ugrRow, ok := userGameReporMap[tmpGameType]
		if !ok {
			ugrRow = &dos.FcUserGameReport{}
		}

		// 历史累积的游戏类型流水
		hisFlow := ugrRow.RebateFlow
		totalFlow := hisFlow + batchFlow

		tmpRebateData.HisBetAmount = hisFlow
		tmpRebateData.BatchBetAmount = batchFlow

		// 如果超过最大返水配置的打码量则直接进行计算
		tmpGameMaxRebate, ok := gameMaxRebateMap[tmpGameType]
		if ok {
			// 历史流水超过最大值
			if hisFlow >= tmpGameMaxRebate.MaxBetAmount {
				tmpRebateData.BonusRate = tmpGameMaxRebate.BonusRate
				tmpRebateData.BonusAmount = venueFlow * tmpGameMaxRebate.BonusRate
				userRebateRecords = append(userRebateRecords, tmpRebateData)
				continue
			}
		} else {
			continue
		}

		// 计算阶梯式返水比例
		thSumBet := 0.00
		thMaxRate := 0.00
		for _, v2 := range gameRebateRows {
			if v2.GameType != tmpGameType {
				continue
			}

			// 最小流水不满足
			if v2.MinBetAmount > totalFlow {
				continue
			}
			tmpRate := v2.BonusRate
			if tmpRate > thMaxRate {
				thMaxRate = tmpRate
			}

			// 如果历史流水比该层次最大流水大则直接跳过找再高一层次流水配置
			if hisFlow > v2.MaxBetAmount {
				continue
			}
			venueFlow2 := tmpRebateData.BetAmount
			if venueFlow2 == 0.00 {
				continue
			}

			// 返水区间
			rebateInterval := v2.MaxBetAmount - v2.MinBetAmount
			// 如果是最低的返水区间，则区间值为最低返水区间的最大值
			tmpGameMinRebate, ok := gameMinRebateMap[tmpGameType]
			if ok {
				if v2.BonusRate == tmpGameMinRebate.BonusRate {
					rebateInterval = tmpGameMinRebate.MaxBetAmount
				}
			}

			key1 := tmpGameType + ":" + strconv.FormatFloat(v2.BonusRate, 'f', -1, 64)
			// 游戏类型返水比例对应的累计流水
			gameTypeRebateRateFlow, ok := gameTypeRebateRateMap[key1]
			if !ok {
				gameTypeRebateRateFlow = 0.0
			}

			// 该类型区间流水已返水
			if gameTypeRebateRateFlow >= rebateInterval {
				continue
			}

			// 如果场馆流水+游戏类型返水比例对应的累计流水大于返水区间值
			if venueFlow2+gameTypeRebateRateFlow > rebateInterval {
				// 剩余的区间返水值
				venueAvailFlow := rebateInterval - gameTypeRebateRateFlow

				// 更改数据,剩余的数据接着下一次再算
				tmp2 := &dos.FcUserRebateRecords{}
				tool.JsonMapper(tmpRebateData, tmp2)
				tmpRebateData.BetAmount = venueFlow2 - venueAvailFlow

				tmp2.BetAmount = venueAvailFlow
				tmp2.BonusRate = v2.BonusRate
				tmp2.BonusAmount = tmp2.BetAmount * v2.BonusRate
				userRebateRecords = append(userRebateRecords, tmp2)
				thSumBet += tmp2.BetAmount

				//fmt.Printf("gameTypeRebateRateMap111 key=%v value=%v\n", key1, rebateInterval)
				gameTypeRebateRateMap[key1] = rebateInterval
				continue
			} else { // 如果场馆流水+游戏类型返水比例对应的累计流水小于返水区间值, 直接 continue
				tmpRebateData.BonusRate = v2.BonusRate
				tmpRebateData.BonusAmount = tmpRebateData.BetAmount * v2.BonusRate
				userRebateRecords = append(userRebateRecords, tmpRebateData)

				thSumBet += tmpRebateData.BetAmount

				//fmt.Printf("gameTypeRebateRateMap222 key=%v value=%v\n", key1, gameTypeRebateRateFlow+venueFlow2)
				gameTypeRebateRateMap[key1] = gameTypeRebateRateFlow + venueFlow2
				break
			}
		}

		// 如果到达了最大返水率 且 还有剩余的流水量
		if thMaxRate == tmpGameMaxRebate.BonusRate && thSumBet < venueFlow {
			limitRecords = append(limitRecords, tmpRebateData)
		}
	}

	// 处理超过最大返水的数据
	for i, v := range limitRecords {
		row := limitRecords[i]
		tmpGameType := v.GameType

		// 如果超过最大返水配置的打码量则直接进行计算
		tmpGameMaxRebate, ok := gameMaxRebateMap[tmpGameType]
		if ok {
			row.BonusRate = tmpGameMaxRebate.BonusRate
			row.BonusAmount = row.BetAmount * tmpGameMaxRebate.BonusRate
			userRebateRecords = append(userRebateRecords, row)
		}
	}

	tmpResultMp := map[string]vo.FcUserRebateValsResult{} // game type => FcUserRebateValsResult

	for i := range userRebateRecords {
		row := userRebateRecords[i]

		if row.BetAmount <= 0.00 {
			continue
		}
		totalBetAmount += row.BetAmount
		totalBonusAmount += row.BonusAmount

		if v, ok := tmpResultMp[row.GameType]; ok {
			v.BetAmount += row.BetAmount
			if row.BonusRate > v.BonusRate { // 返水比例最大值应该就是当前值了吧
				v.BonusRate = row.BonusRate
			}
			v.BonusAmount += row.BonusAmount
			tmpResultMp[row.GameType] = v
		} else {
			tmpResultMp[row.GameType] = vo.FcUserRebateValsResult{
				VenueCode:   row.VenueCode,
				GameType:    row.GameType,
				BetAmount:   row.BetAmount,
				BonusRate:   row.BonusRate,
				BonusAmount: row.BonusAmount,
			}
		}
	}

	for _, v := range tmpResultMp {
		v.BonusAmount = tool.TruncateFloat(v.BonusAmount, 2)
		result = append(result, v)
	}

	totalBetAmount = tool.TruncateFloat(totalBetAmount, 2)
	totalBonusAmount = tool.TruncateFloat(totalBonusAmount, 2)

	// modules.SetCacheUserRebateBonusValue(userId, totalBonusAmount)

	return result, totalBetAmount, totalBonusAmount
}
