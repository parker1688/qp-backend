package modules

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"encoding/json"
	"fmt"
	"time"
)

type UserActivityData struct {
	Id     string `json:"id"`
	Value  string `json:"value"`
	Date   string `json:"date"`
	Status int    `json:"status"`
}

type ActivityActionParam struct {
	RechargeBalanceRatio float64 // 充值余额比
	Balance              float64 // 余额
	RegTime              string  // 注册时间
	FirstRechargeAmount  float64 // 首充金额
}

type UserActivityHealthPackResult struct {
	// TODO(活动): 返回字段待完善
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	StageContent string  `json:"stage_content"`
	GiftStyle    int     `json:"gift_style"`
	BonusAmount  float64 `json:"bonus_amount"`
	Status       int     `json:"status"`
}

type UserActivityRedEnvelopeRainResult struct {
	// TODO(活动): 返回字段待完善
	Id        string                  `json:"id"`
	Title     string                  `json:"title"`
	Content   string                  `json:"content"`
	Cd        int64                   `json:"cd"`
	DateRange *dos.DateRangeDataField `json:"-"`
}

// loadUserActivityFromDB - 从库载入用户活动数据
// @param {string} userId
// @returns []UserActivityData
func loadUserActivityFromDB(userId string) []UserActivityData {
	result := []UserActivityData{}

	data := dos.FcUserData{
		ActivityData: "[]",
	}
	err := global.G_DB.Model(&dos.FcUserData{}).Select("user_id", "activity_data").
		Where("user_id = ?", userId).Scan(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadUserActivityFromDB] Find user activity data failed: userId=%s, err=%v", userId, err.Error())
		return result
	}

	if len(data.UserId) == 0 {
		// 没有数据则新增
		data.UserId = userId
		data.ActivityData = "[]"
		err = global.G_DB.Model(&dos.FcUserData{}).Create(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[loadUserActivityFromDB] Create user activity data failed: userId=%s, err=%v", userId, err.Error())
		}

		return result
	}

	if data.ActivityData == "" || data.ActivityData == "[]" {
		return result
	}

	err = tool.JsonUnmarshal([]byte(data.ActivityData), &result)
	if err != nil {
		global.G_LOG.Errorf("[loadUserActivityFromDB] Unmarshal user activity data failed: userId=%s, data=%s",
			userId, data.ActivityData)
		return result
	}

	return result
}

// GetUserActivityData - 获取用户活动数据
// @param {string} userId
// @returns []UserActivityData
func GetUserActivityData(userId string) []UserActivityData {
	return loadUserActivityFromDB(userId)
}

// SetUserActivityData - 更新用户活动数据
// @param {string} userId
// @param {[]UserActivityData} data
// @returns
func SetUserActivityData(userId string, data []UserActivityData) {
	sData, err := tool.JsonMarshalString(data)
	if err != nil {
		global.G_LOG.Errorf("[SetUserActivityData] Marshal user activity data failed: userId=%s, err=%v", userId, err.Error())
		return
	}

	err = global.G_DB.Model(&dos.FcUserData{}).Where("user_id = ?", userId).
		Update("activity_data", sData).Error
	if err != nil {
		global.G_LOG.Errorf("[SetUserActivityData] Update user activity data failed: userId=%s, err=%v", userId, err.Error())
		return
	}
}

// SetUserActivityDataByMap - 根据map更新用户活动数据
// @param {string} userId
// @param {map[string]UserActivityData} data
// @returns
func SetUserActivityDataByMap(userId string, data map[string]UserActivityData) {
	activityLis := []UserActivityData{}
	for _, v := range data {
		activityLis = append(activityLis, v)
	}

	SetUserActivityData(userId, activityLis)
}

// GetUserActivityDataToMap - 获取用户活动数据（map形式）
// @param {string} userId
// @returns map[string]UserActivityData
func GetUserActivityDataToMap(userId string) map[string]UserActivityData {
	isSave := false
	activityMap := map[string]UserActivityData{}
	activityLis := GetUserActivityData(userId)
	for _, v := range activityLis {
		activityMap[v.Id] = v
	}

	if isSave {
		SetUserActivityDataByMap(userId, activityMap)
	}

	return activityMap
}

// CheckActivityResetByCycle - 根据周期判断是否重置活动
// @param {dos.FcPromotionInfo} config
// @param {string} date
// @returns bool
func CheckActivityResetByCycle(config *dos.FcPromotionInfo, date string) bool {
	now := time.Now()
	switch config.Cycle {
	case enmus.ActivityCycle_Daily: // 每日
		return now.Format(tool.TimeDateLayout) != date
	case enmus.ActivityCycle_Week: // 每周
		t1 := tool.GetTimeFromString(date + " 12:00:00")
		t2 := tool.GetTimeFromString(now.Format(tool.TimeLayout))
		return tool.IsDifferentWeekCustom(t1, t2, time.Monday)
	case enmus.ActivityCycle_Month: // 每月
		return now.Format(tool.TimeDateYearMonLayout) != date
	default:
		return false
	}
}

// DoUserActivityAction - 触发用户活动
// @param {int} activityType
// @param {string} userId
// @param {[]ActivityActionParam} params
// @returns
func DoUserActivityAction(activityType int, userId string,
	params []ActivityActionParam) {

	isSave := false

	now := time.Now().Unix()
	nowDate := time.Now().Format(tool.TimeDateLayout)

	activityMap := GetUserActivityDataToMap(userId)
	activityIds := GetActivityConfigIndexes(getCacheUserMerchantCode(userId),
		activityType) // 指定活动类型的索引

	for _, param := range params {
		for _, activityId := range activityIds {
			activityConfig := GetActivityConfig(activityId)
			if activityConfig == nil {
				continue
			}

			// if activityConfig.PromotionType != activityType {
			// 	continue
			// }

			// 判断是否被关闭
			if activityConfig.Status != enmus.ActivityStats_Opening {
				continue
			}

			// 判断活动时间
			sTime := tool.CovertTimestampFromAutomaticTypeTime(&activityConfig.StartTime)
			eTime := tool.CovertTimestampFromAutomaticTypeTime(&activityConfig.EndTime)
			if sTime > 0 && eTime > 0 && (now < sTime || now >= eTime) {
				if now >= eTime {
					delete(activityMap, activityId) // 活动结束则删除用户相关活动数据
				}
				continue
			}

			switch activityType {
			case enmus.ActivityTypes_HealthPack: // 回血包触发处理
				// 处理周期
				if v, ok := activityMap[activityId]; ok {
					if CheckActivityResetByCycle(activityConfig, v.Date) {
						v.Status = enmus.UserActivityStats_None
						v.Date = nowDate
						activityMap[activityId] = v
						isSave = true
					}
				}

				// 判断注册时间
				if activityConfig.RegStartTime != nil && activityConfig.RegEndTime != nil {
					sRegTime := tool.CovertTimestampFromAutomaticTypeTime(activityConfig.RegStartTime)
					eRegTime := tool.CovertTimestampFromAutomaticTypeTime(activityConfig.RegEndTime)
					regTime := tool.GetTimeFromString(param.RegTime).Unix()
					if sRegTime > 0 && eRegTime > 0 &&
						(regTime < sRegTime || regTime > eRegTime) {
						continue
					}
				}

				// 判断余额比
				if activityConfig.RechargeBalanceRatio > 0 &&
					activityConfig.RechargeBalanceRatio < param.RechargeBalanceRatio {
					continue
				}

				// 判断余额
				if activityConfig.Balance > 0 &&
					activityConfig.Balance < param.Balance {
					continue
				}

				// 判断首充金额
				if activityConfig.FirstRechargeAmount > 0 &&
					activityConfig.FirstRechargeAmount > param.FirstRechargeAmount {
					continue
				}
			}

			// 更新数据
			if v, ok := activityMap[activityId]; ok {
				if v.Status == enmus.UserActivityStats_None { // 仅不可领状态
					v.Status = enmus.UserActivityStats_Reward
					activityMap[activityId] = v
					isSave = true
				}
			} else {
				activityData := UserActivityData{
					Id:     activityId,
					Date:   nowDate,
					Status: enmus.UserActivityStats_Reward,
				}

				activityMap[activityId] = activityData

				isSave = true
			}
		}
	}

	if isSave {
		SetUserActivityDataByMap(userId, activityMap)
	}
}

// GetUserActivityHealthPackInfo - 获取用户活动回血包信息
// @param {string} userId
// @returns []UserActivityResult
func GetUserActivityHealthPackInfo(userId string) []UserActivityHealthPackResult {
	var result []UserActivityHealthPackResult

	data := GetUserActivityData(userId)

	now := time.Now().Unix()
	nowDate := time.Now().Format(tool.TimeDateLayout)

	isSave := false

	newData := []UserActivityData{}

	for i, v := range data {
		activityConfig := GetActivityConfig(v.Id)
		if activityConfig == nil { // 活动不存在
			continue
		}

		// 判断活动是否结束
		eTime := tool.CovertTimestampFromAutomaticTypeTime(&activityConfig.EndTime)
		if eTime > 0 && now >= eTime {
			continue
		}

		// 仅回血包类型
		if activityConfig.PromotionType == enmus.ActivityTypes_HealthPack {
			// 处理重置数据
			if CheckActivityResetByCycle(activityConfig, v.Date) {
				data[i].Status = enmus.UserActivityStats_None
				data[i].Date = nowDate
				isSave = true
			}

			if v.Status == enmus.UserActivityStats_Reward { // 仅可领取回血包数据给前端
				result = append(result, UserActivityHealthPackResult{
					Id:           activityConfig.Id,
					Title:        activityConfig.Title,
					Content:      activityConfig.Content,
					StageContent: activityConfig.StageContent,
					GiftStyle:    activityConfig.GiftStyle,
					BonusAmount:  activityConfig.BonusAmount,
					Status:       v.Status,
				})
			}
		}

		newData = append(newData, data[i])
	}

	if isSave {
		SetUserActivityData(userId, newData)
	}

	return result
}

// GetUserActivityRedEnvelopeRainInfo - 获取用户活动红包雨信息
// @param {string} userId
// @param {int64} bufTime
// @returns *UserActivityRedEnvelopeRainResult
func GetUserActivityRedEnvelopeRainInfo(userId string, bufTime int64) *UserActivityRedEnvelopeRainResult {
	var result *UserActivityRedEnvelopeRainResult

	isSave := false

	now := time.Now().Unix()
	nowDate := time.Now().Format(tool.TimeDateLayout)

	activityMap := GetUserActivityDataToMap(userId)

	activityIds := GetActivityConfigIndexes(getCacheUserMerchantCode(userId),
		enmus.ActivityTypes_RedEnvelopeRain)

	for _, activityId := range activityIds {
		activityConfig := GetActivityConfig(activityId)
		// 判断活动是否存在
		if activityConfig == nil {
			if _, ok := activityMap[activityId]; ok {
				// 删除已不存在配置中的活动
				delete(activityMap, activityId)
				isSave = true
			}
			continue
		}

		// 判断是否被关闭
		if activityConfig.Status != enmus.ActivityStats_Opening {
			continue
		}

		// 判断活动时间
		sTime := tool.CovertTimestampFromAutomaticTypeTime(&activityConfig.StartTime)
		eTime := tool.CovertTimestampFromAutomaticTypeTime(&activityConfig.EndTime)
		if sTime > 0 && eTime > 0 && (now < sTime || now >= eTime) {
			if now >= eTime {
				if _, ok := activityMap[activityId]; ok {
					delete(activityMap, activityId) // 活动结束则删除用户相关活动数据
					isSave = true
				}
			}
			continue
		}

		if !json.Valid([]byte(activityConfig.DateRangeData)) {
			continue
		}

		if v, ok := activityMap[activityId]; ok {
			if v.Date != nowDate {
				v.Status = enmus.UserActivityStats_None
				v.Value = "0"
				v.Date = nowDate
				activityMap[activityId] = v
				isSave = true
			}
		} else {
			activityMap[activityId] = UserActivityData{
				Id:     activityId,
				Value:  "0",
				Date:   nowDate,
				Status: enmus.UserActivityStats_None,
			}
			isSave = true
		}

		scopeLis := []dos.DateRangeDataField{}
		tool.JsonUnmarshalFromString(activityConfig.DateRangeData, &scopeLis)
		for i, v := range scopeLis {
			if len(v.TimeScope) == 2 {
				if tool.Atoi(activityMap[activityId].Value) > i {
					continue
				}

				st := tool.GetTimeFromString(nowDate + " " + v.TimeScope[0]).Unix()
				et := tool.GetTimeFromString(nowDate + " " + v.TimeScope[1]).Unix()

				if now > (et + bufTime) { // 该时间段已过时
					if (i + 1) < len(scopeLis) {
						if vv, ok := activityMap[activityId]; ok {
							vv.Value = fmt.Sprintf("%d", i+1)
							vv.Status = enmus.UserActivityStats_None
							activityMap[activityId] = vv
							isSave = true
						}
					}
					continue
				}

				if i == tool.Atoi(activityMap[activityId].Value) &&
					activityMap[activityId].Status == enmus.UserActivityStats_Rewarded {
					// 当前时间段已领则忽略转下一个时间段
					if (i + 1) < len(scopeLis) {
						if vv, ok := activityMap[activityId]; ok {
							vv.Value = fmt.Sprintf("%d", i+1)
							vv.Status = enmus.UserActivityStats_None
							activityMap[activityId] = vv
							isSave = true
						}
					}
					continue
				}

				if now < st {
					// 未到时
					result = &UserActivityRedEnvelopeRainResult{
						Id:        activityId,
						Title:     activityConfig.Title,
						Content:   activityConfig.Content,
						Cd:        max(st-now, 0),
						DateRange: &v,
					}
					break
				} else if now >= st && now <= (et+bufTime) {
					// 在期间
					result = &UserActivityRedEnvelopeRainResult{
						Id:        activityId,
						Title:     activityConfig.Title,
						Content:   activityConfig.Content,
						Cd:        0,
						DateRange: &v,
					}
					break
				}
			}
		}

		break
	}

	if isSave && bufTime == 0 {
		SetUserActivityDataByMap(userId, activityMap)
	}

	return result // result = nul 前端不显示图标；cd > 0 倒计时；cd = 0 则可以领取
}

// GetUserActivityPromotionTotalVal - 获取活动福利统计
// @param {string} userId
// @returns float64
func GetUserActivityPromotionTotalVal(userId string) float64 {
	fundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "activity")
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(userId) > 0 {
		query = query.Where("user_id = ? AND funding_subtype = ?", userId, fundingSubType)
	} else {
		query = query.Where("funding_subtype = ?", fundingSubType)
	}

	var totalAmount *float64
	err := query.Select("sum(amount) as totalAmount").Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[GetUserActivityPromotionTotalVal] Accumulate activity promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}

// GetUserActivityPromotionTotalByDate - 获取活动福利统计
// @param {string} merchantCode
// @param {automaticType.Time} sTime
// @param {automaticType.Time} eTime
// @returns float64
func GetUserActivityPromotionTotalByDate(merchantCode string,
	sTime automaticType.Time, eTime automaticType.Time) float64 {
	fundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "activity")
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(merchantCode) > 0 {
		query = query.Where("merchant_code = ?", merchantCode)
	}

	if !sTime.Timer().IsZero() {
		query = query.Where("create_time >= ?", sTime)
	}
	if !eTime.Timer().IsZero() {
		query = query.Where("create_time < ?", eTime)
	}

	query = query.Where("funding_subtype = ?", fundingSubType)

	var totalAmount *float64
	err := query.Select("sum(amount) as totalAmount").Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[GetUserActivityPromotionTotalByDate] Accumulate activity promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}
