// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveDailyBonus(vo *dos.DailyBonus) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageDailyBonus(page, pageSize int, vo *dos.DailyBonus, c *gin.Context) (ret []*dos.DailyBonus, total int64) {
	query := global.G_DB.Model(&dos.DailyBonus{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	dataSlice := []*dos.DailyBonus{}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDailyBonus(vo *dos.DailyBonus, c *gin.Context) []*dos.DailyBonus {
	var data []*dos.DailyBonus
	query := global.G_DB.Model(&dos.DailyBonus{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

func FindByKeyDailyBonusFirst(vo *dos.DailyBonus) *dos.DailyBonus {
	var data *dos.DailyBonus
	query := global.G_DB.Model(&dos.DailyBonus{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	query.Take(&data)
	return data
}

func UpdateDailyBonus(vo *dos.DailyBonus) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_code": vo.MerchantCode,
		"bonus":         vo.Bonus,
		"update_time":   automaticType.Time(time.Now()),
		"update_by":     vo.UpdateBy,
	}).Error == nil
}

func DeleteDailyBonus(vo *dos.DailyBonus) bool {
	return global.G_DB.Model(&dos.DailyBonus{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// CheckDailyBonusMerchant - 判断是否已存在商户
// @param {string} merchantCode 组名称
// @param {string} excludeId 排除id
// @returns bool
func CheckDailyBonusMerchant(merchantCode string, excludeId string) bool {
	data := dos.DailyBonus{}
	err := global.G_DB.Model(&dos.DailyBonus{}).Select("id").Where("merchant_code = ?", merchantCode).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckDailyBonusName] Query user group count failed: %v", err.Error())
		return false
	}

	if len(excludeId) > 0 {
		return data.Id != excludeId
	}

	return len(data.Id) > 0
}

// GetDailyBonusConfig - 根据商户码获取签到配置数据
// @param {string} merchantCode 商户码
// @returns
func GetDailyBonusConfig(merchantCode string) ([]dos.DailyBonusField, error) {
	result := []dos.DailyBonusField{}
	dailyBonus := dos.DailyBonus{}
	err := global.G_DB.Model(&dos.DailyBonus{}).Select("bonus").
		Where("merchant_code = ?", merchantCode).Take(&dailyBonus).Error
	if err != nil {
		global.G_LOG.Errorf("[GetDailyBonusConfig] Find daily bonus failed: merchantCode=%s, err=%v",
			merchantCode, err.Error())
		return result, err
	}

	if len(dailyBonus.Bonus) == 0 {
		global.G_LOG.Errorf("[GetDailyBonusConfig] Daily bonus is empty: merchantCode = %s", merchantCode)
		return result, err
	}

	err = tool.JsonUnmarshal([]byte(dailyBonus.Bonus), &result)
	if err != nil {
		global.G_LOG.Errorf("[GetDailyBonusConfig] Unmarshal daily bonus failed: merchantCode = %s, data=%s",
			merchantCode, dailyBonus.Bonus)
		return result, err
	}

	return result, nil
}

// GetDailyBonusDiffAmount - 获取签到打码量差值
// @param {string} merchantCode 商户码
// @param {dos.DailyBonusData} data 用户签到数据
func GetDailyBonusDiffAmount(merchantCode string, data dos.DailyBonusData) float64 {
	dailyBonusConfigs, _ := GetDailyBonusConfig(merchantCode)
	if len(dailyBonusConfigs) == 0 {
		return 0
	}

	for _, config := range dailyBonusConfigs {
		if config.Days == data.Day {
			return max(config.Amount-data.Amount, 0)
		}
	}

	return 0
}

// DoDailyBonusAction - 触发用户签到
// @param {string} userId 用户Id
// @param {string} merchantCode 商户码
// @returns
func DoDailyBonusAction(userId string, merchantCode string) {
	dailyBonusData, err1 := GetUserDailyBonusData(userId, true)
	if err1 != nil {
		return
	}

	// 统计今日投注金额
	var totalAmount *float64
	err := global.G_DB.Model(&dos.FcBetRecord{}).Select("SUM(valid_betamount) as totalAmount").
		Where("user_id = ? AND DATE(create_time) = DATE(?)", userId, time.Now()).Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[DoDailyBonusAction] Calculate bet amount failed: %v", err.Error())
		return
	}

	if totalAmount == nil { // 没有投注金额
		defaultAmount := 0.0
		totalAmount = &defaultAmount
	}

	// 判断状态是否为已领取（注意：可领取状态下如果后台修改配置比之前高则可能退回到不可领取状态）
	if dailyBonusData.Status == enmus.DailyBonusStats_Rewarded {
		return
	}

	// 判断是否达到指标并设置状态
	isSave := false
	dailyBonusConfigs, err2 := GetDailyBonusConfig(merchantCode)
	if err2 != nil {
		global.G_LOG.Errorf("[DoDailyBonusAction] Find daily bonus configs failed: merchantCode=%s, err=%v",
			merchantCode, err2.Error())
		return
	}
	for _, config := range dailyBonusConfigs {
		if config.Days == dailyBonusData.Day {
			if *totalAmount >= config.Amount {
				// 触发可领取
				dailyBonusData.Status = enmus.DailyBonusStats_Reward
				dailyBonusData.Amount = *totalAmount
				isSave = true
			} else if *totalAmount < config.Amount &&
				dailyBonusData.Status == enmus.DailyBonusStats_Reward {
				// 当只比条件少且状态是可领取需要退回不可领取（注意：可能后台配置改了）
				dailyBonusData.Status = enmus.DailyBonusStats_None
				dailyBonusData.Amount = *totalAmount
				isSave = true
			}
			break
		}
	}

	if isSave {
		SetUserDailyBonusData(userId, dailyBonusData)
	}
}

// DoDailyBonusReset - 重置用户签到数据
// @param {string} userId 用户Id
// @param {dos.DailyBonusData} data 签到数据
// @param {bool} isSet 是否更新
// @returns dos.DailyBonusData
func DoDailyBonusReset(userId string, data dos.DailyBonusData, isSet bool) dos.DailyBonusData {
	now := time.Now()
	sNow := now.Format(tool.TimeDateLayout)
	isSave := false
	if data.Date != sNow {
		switch data.Status {
		case enmus.DailyBonusStats_None:
			data.Amount = 0
			data.Date = sNow
			isSave = true
		case enmus.DailyBonusStats_Reward: // 已隔日且为可领取状态
			data.Amount = 0
			data.Status = enmus.DailyBonusStats_None
			data.Date = sNow
			isSave = true
		case enmus.DailyBonusStats_Rewarded: // 已隔日且为已领取状态
			if data.Day == enmus.DailyBonusResetCycle {
				data.Day = 1
			} else {
				data.Day += 1
			}
			data.ExtDay += 1
			data.Amount = 0
			data.Status = enmus.DailyBonusStats_None
			data.Date = sNow
			isSave = true
		}
	}

	if data.ExtDate != now.Format(tool.TimeDateYearMonLayout) {
		// 月份不同则重置额外奖励计数
		data.ExtDay = 1
		data.ExtDate = now.Format(tool.TimeDateYearMonLayout)
		isSave = true
	}

	if isSave && isSet {
		SetUserDailyBonusData(userId, data)
	}

	return data
}

// GetUserDailyBonusData - 获取用户签到数据
// @param {string} userId 用户ID
// @param {bool} isReset 是否重置
// @returns *dos.DailyBonusData
func GetUserDailyBonusData(userId string, isReset bool) (dos.DailyBonusData, error) {
	now := time.Now()
	result := dos.DailyBonusData{
		Day:     1,
		ExtDay:  1,
		Status:  enmus.DailyBonusStats_None,
		Date:    now.Format(tool.TimeDateLayout),
		ExtDate: now.Format(tool.TimeDateYearMonLayout),
	}
	userData := dos.FcUserMaterial{}
	err := global.G_DB.Model(&dos.FcUserMaterial{}).Select("dailybonus_data").
		Where("user_id = ?", userId).Take(&userData).Error
	if err != nil {
		global.G_LOG.Errorf("[GetUserDailyBonusData] Find user data failed: %v", err.Error())
		return result, err
	}

	if len(userData.DailyBonusData) > 0 {
		err = tool.JsonUnmarshal([]byte(userData.DailyBonusData), &result)
		if err != nil {
			global.G_LOG.Errorf("[GetUserDailyBonusData] Unmarshal user daily bonus data failed: %v", err.Error())
			return result, err
		}
	}

	if isReset {
		result = DoDailyBonusReset(userId, result, true)
	}

	return result, nil
}

// SetUserDailyBonusData - 更新用户签到数据
// @param {string} userId
// @param {dos.DailyBonusData} data
// @returns
func SetUserDailyBonusData(userId string, data dos.DailyBonusData) {
	sData, err := tool.JsonMarshalString(data)
	if err != nil {
		global.G_LOG.Errorf("[SetUserDailyBonusData] Json marshal daily bonus data failed: %v", err.Error())
		return
	}
	err = global.G_DB.Model(&dos.FcUserMaterial{}).
		Where("user_id = ?", userId).Update("dailybonus_data", sData).Error
	if err != nil {
		global.G_LOG.Errorf("[SetUserDailyBonusData] Update user daily bonus data failed: %v", err.Error())
		return
	}
}

// GetDailyBonusInfo - 获取签到信息
// @param {string} userId 用户Id
// @param {string} merchantCode 商户码
// @returns dos.DailyBonusResp
func GetDailyBonusInfo(userId string, merchantCode string) (dos.DailyBonusResp, error) {
	rewardLis := []dos.DailyBonusResult{}
	extraRewardLis := []dos.DailyBonusResult{}

	dailyBonusData, err := GetUserDailyBonusData(userId, false)
	if err != nil {
		return dos.DailyBonusResp{}, err
	}

	dailyBonusConfigs, err1 := GetDailyBonusConfig(merchantCode)
	if err1 != nil {
		return dos.DailyBonusResp{}, err1
	}
	for _, config := range dailyBonusConfigs {
		if config.Days <= enmus.DailyBonusResetCycle {
			rewardLis = append(rewardLis, dos.DailyBonusResult{
				DailyBonusField: config,
			})
		}

		if config.ExtraReward > 0 {
			extraRewardLis = append(extraRewardLis, dos.DailyBonusResult{
				DailyBonusField: config,
			})
		}
	}

	newExtDay := dailyBonusData.ExtDay
	if dailyBonusData.Status != enmus.DailyBonusStats_Rewarded {
		newExtDay = dailyBonusData.ExtDay - 1
	}

	return dos.DailyBonusResp{
		RewardList:      rewardLis,
		ExtraRewardList: extraRewardLis,
		Days:            dailyBonusData.Day,
		ExtDays:         newExtDay,
		Status:          dailyBonusData.Status,
		TotalAmount:     dailyBonusData.Amount,
	}, nil
}

// GetDailyBonusPromotionTotalVal - 获取签到福利统计
// @param {string} userId
func GetDailyBonusPromotionTotalVal(userId string) float64 {
	baseFundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus")
	extFundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus_extra")
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(userId) > 0 {
		query = query.Where("user_id = ? AND (funding_subtype = ? OR funding_subtype = ?)", userId, baseFundingSubType, extFundingSubType)
	} else {
		query = query.Where("funding_subtype = ? OR funding_subtype = ?", baseFundingSubType, extFundingSubType)
	}

	var totalAmount *float64
	err := query.Select("sum(amount) as totalAmount").Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[GetDailyBonusPromotionTotalVal] Accumulate daily bonus promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}

func GetDailyBonusPromotionTotalByDate(merchantCode string,
	sTime automaticType.Time, eTime automaticType.Time) float64 {
	baseFundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus")
	extFundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_bonus_extra")
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

	query = query.Where("funding_subtype = ? OR funding_subtype = ?",
		baseFundingSubType, extFundingSubType)

	var totalAmount *float64
	err := query.Select("sum(amount) as totalAmount").Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[GetDailyBonusPromotionTotalByDate] Accumulate daily bonus promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}
