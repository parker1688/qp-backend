package modules

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var userTaskCache = cache.New(10*time.Minute, 15*time.Minute) // 任务缓存

const (
	cacheUserTaskKey         = "cacheUserTask"         // 用户任务缓存KEY
	cacheUserMerchantCodeKey = "cacheUserMerchantCode" // 用户商户码缓存KEY
	cacheUserTaskReloadKey   = "cacheUserTaskReload"   // 用户任务重载入缓存KEY（rds）
)

var (
	userTaskExpiresDurationFn = func() time.Duration { // 用户任务数据缓存失效时间
		return time.Duration(tool.RandInt(3, 6)) * time.Minute
	}
	userMerchantCodeExpiresDurationFn = func() time.Duration { // 用户商户码缓存失效时间
		return time.Duration(tool.RandInt(1, 5)) * time.Hour
	}
)

// 用户任务
type UserTaskData struct {
	Id     string  `json:"id"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Status int     `json:"status"`
}

type TaskActionParam struct {
	Type        int
	Subtype     int
	Amount      float64
	GameType    string
	VenueCode   string
	GameCode    string
	ChannelCode string
}

type UserTaskResult struct {
	Id          string              `json:"id"`
	Type        int                 `json:"type"`
	Subtype     int                 `json:"subtype"`
	Name        string              `json:"name"`
	Intro       string              `json:"intro"`
	Detail      string              `json:"detail"`
	CurrAmount  float64             `json:"curr_amount"`
	Amount      float64             `json:"amount"`
	BonusAmount float64             `json:"bonus_amount"`
	GameType    string              `json:"game_type"`
	VenueCode   string              `json:"venue_code"`
	VenueName   string              `json:"venue_name"`
	StartAt     *automaticType.Time `json:"start_at"`
	EndAt       *automaticType.Time `json:"end_at"`
	Status      int                 `json:"status"`
	Gtype       *int                `json:"gtype"`
	TableId     string              `json:"tableId"`
	GameCode    string              `json:"game_code"`
	GameName    string              `json:"game_name"`
}

// getCacheUserMerchantCodeKey - 获取用户商户码缓存KEY
// @param {string} userId
// @returns string
func getCacheUserMerchantCodeKey(userId string) string {
	return cacheUserMerchantCodeKey + ":" + userId
}

// getCacheUserMerchantCode - 获取用户商户码缓存
// @param {string} userId
// @returns string
func getCacheUserMerchantCode(userId string) string {
	merchantCode, ret := userTaskCache.Get(getCacheUserMerchantCodeKey(userId))
	if !ret {
		data := dos.FcUserMaterial{}
		err := global.G_DB.Model(&dos.FcUserMaterial{}).Select("merchant_code").
			Where("user_id = ?", userId).Take(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[getCacheUserMerchantCode] Find user merchant code failed: userId=%s, err=%v", userId, err.Error())
			return ""
		}

		setCacheUserMerchantCode(userId, data.MerchantCode)

		return data.MerchantCode
	}

	if v, ok := merchantCode.(string); ok {
		return v
	}

	return ""
}

// SetCacheUserMerchantCode - 更新用户商户码缓存数据
// @param {string} userId
// @param {string} merchantCode
// @returns
func setCacheUserMerchantCode(userId string, merchantCode string) {
	userTaskCache.Set(getCacheUserMerchantCodeKey(userId), merchantCode,
		userMerchantCodeExpiresDurationFn())
}

// getCacheUserTaskKey - 获取用户任务数据缓存KEY
// @param {string} userId
// @returns string
func getCacheUserTaskKey(userId string) string {
	return cacheUserTaskKey + ":" + userId
}

// loadUserTaskFromDB - 从库载入用户任务数据
// @param {string} userId
// @returns []UserTaskData
func loadUserTaskFromDB(userId string) []UserTaskData {
	result := []UserTaskData{}

	data := dos.FcUserTask{
		Data: "[]",
	}
	err := global.G_DB.Model(&dos.FcUserTask{}).Select("user_id", "data").
		Where("user_id = ?", userId).Scan(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadUserTaskFromDB] Find user task data failed: taskId=%s, err=%v", userId, err.Error())
		return result
	}

	if len(data.UserId) == 0 {
		// 没有数据则新增
		data.UserId = userId
		data.Data = "[]"
		data.CreatedAt = automaticType.Time(time.Now())
		err = global.G_DB.Model(&dos.FcUserTask{}).Create(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[loadUserTaskFromDB] Create user task data failed: userId=%s, err=%v", userId, err.Error())
		}

		return result
	}

	if data.Data == "[]" {
		return result
	}

	err = tool.JsonUnmarshal([]byte(data.Data), &result)
	if err != nil {
		global.G_LOG.Errorf("[loadUserTaskFromDB] Unmarshal user task data failed: userId=%s, data=%s",
			userId, data.Data)
		return result
	}

	// 更新用户任务数据缓存
	//userTaskCache.Set(getCacheUserTaskKey(userId), result,
	//	userTaskExpiresDurationFn())

	return result
}

// GetUserTaskData - 获取用户任务数据（用户任务数据不使用缓存）
// @param {string} userId
// @returns []UserTaskData
func GetUserTaskData(userId string) []UserTaskData {
	// data, ret := userTaskCache.Get(getCacheUserTaskKey(userId))
	// if !ret { // 缓存不存在
	return loadUserTaskFromDB(userId)
	// }

	// if v, ok := data.([]UserTaskData); ok {
	// 	return v
	// }

	// return []UserTaskData{}
}

// GetUserTaskDataToMap - 获取用户任务数据返回MAP
// @param {string} userId
// @returns map[string]UserTaskData
func GetUserTaskDataToMap(userId string) map[string]UserTaskData {
	isSave := false
	nowDate := time.Now().Format(tool.TimeDateLayout)

	taskMap := map[string]UserTaskData{}
	taskLis := GetUserTaskData(userId)
	for _, v := range taskLis {
		config := GetDailyTaskConfig(v.Id)
		if config != nil {
			// 处理每日任务
			if config.Type == enmus.DailyTaskType_Daily && v.Date != nowDate {
				// 需重置每日任务
				v.Amount = 0
				v.Status = enmus.UserTaskStats_None
				v.Date = nowDate
				isSave = true
			}

			// 处理状态回退（条件值变为未达到且为可领取状态）
			if v.Amount < config.Amount && v.Status == enmus.UserTaskStats_Reward {
				v.Status = enmus.UserTaskStats_None
				isSave = true
			}
		}

		taskMap[v.Id] = v
	}

	if isSave {
		SetUserTaskDataByMap(userId, taskMap)
	}

	return taskMap
}

/*
// getCacheUserTaskReloadKey - 获取用户任务缓存重载入KEY
// @param {string} userId
// @returns string
func getCacheUserTaskReloadKey(userId string) string {
	return cacheUserTaskReloadKey + ":" + userId
}

// isUserTaskReload - 用户任务缓存是否重新载入
// @param {string} userId
// @returns bool
func isUserTaskReload(userId string) bool {
	// return global.G_REDIS.Get(context.Background(), getCacheUserTaskReloadKey(userId)).Val() == "1"
	data := dos.FcUserTask{}
	err := global.G_DB.Model(&dos.FcUserTask{}).Select("is_reload").Where("user_id = ?", userId).Error
	if err != nil {
		global.G_LOG.Errorf("[isUserTaskReload] Find user task reload flag failed: userId=%s, err=%v", userId, err.Error())
		return false
	}

	return data.IsReload == 1
}

// setUserTaskReload - 更新用户任务缓存重载入标识
// @param {string} userId
// @returns
func setUserTaskReload(userId string) {
	//global.G_REDIS.Set(context.Background(), getCacheUserTaskReloadKey(userId), "1", -1)
	err := global.G_DB.Model(&dos.FcUserTask{}).Where("user_id = ?", userId).
		Update("is_reload", 1).Error
	if err != nil {
		global.G_LOG.Errorf("[setUserTaskReload] Update user task reload flag failed: userId=%s, err=%v", userId, err.Error())
	}
}

// delUserTaskReload - 删除用户任务缓存重载入标识
// @param {string} userId
// @returns
func delUserTaskReload(userId string) {
	//global.G_REDIS.Del(context.Background(), getCacheUserTaskReloadKey(userId))
	err := global.G_DB.Model(&dos.FcUserTask{}).Where("user_id = ?", userId).
		Update("is_reload", 0).Error
	if err != nil {
		global.G_LOG.Errorf("[delUserTaskReload] Update user task reload flag failed: userId=%s, err=%v", userId, err.Error())
	}
}

// doUserTaskReload - 执行用户任务重载入
// @param {string} userId
func doUserTaskReload(userId string) {
	if isUserTaskReload(userId) {
		loadUserTaskFromDB(userId)
		delUserTaskReload(userId)
	}
}*/

// SetUserTaskData - 更新用户任务数据
// @param {string} userId
// @param {[]UserTaskData} data
// @returns
func SetUserTaskData(userId string, data []UserTaskData) {
	//userTaskCache.Set(getCacheUserTaskKey(userId), data,
	//	userTaskExpiresDurationFn())

	sData, err := tool.JsonMarshalString(data)
	if err != nil {
		global.G_LOG.Errorf("[SetUserTaskData] Marshal user task data failed: userId=%s, err=%v", userId, err.Error())
		return
	}

	err = global.G_DB.Model(&dos.FcUserTask{}).Where("user_id = ?", userId).
		Update("data", sData).Error
	if err != nil {
		global.G_LOG.Errorf("[SetUserTaskData] Update user task data failed: userId=%s, err=%v", userId, err.Error())
		return
	}
}

// SetUserTaskDataByMap - 根据MAP更新用户任务数据
// @param {string} userId
// @param {map[string]UserTaskData} data
// @returns
func SetUserTaskDataByMap(userId string, data map[string]UserTaskData) {
	taskLis := []UserTaskData{}
	for _, v := range data {
		taskLis = append(taskLis, v)
	}

	SetUserTaskData(userId, taskLis)
}

// ToUserTaskParamsByBetRecord - 根据注单获取用户任务参数表
// @param {map[string][]TaskActionParam} userParamsMp
// @param {*dos.FcBetRecord} record
// @returns map[string][]TaskActionParam
func ToUserTaskParamsByBetRecord(userParamsMp map[string][]TaskActionParam,
	record *dos.FcBetRecord) map[string][]TaskActionParam {
	taskType, ok := enmus.EnumDailyTaskTypeMap[record.GameType]
	if !ok {
		return userParamsMp
	}

	taskSubType := enmus.DailyTaskSubType_Win
	if record.NetAmount < 0 {
		taskSubType = enmus.DailyTaskSubType_Los
	}
	if _, ok := userParamsMp[record.UserId]; ok {
		userParamsMp[record.UserId] = append(userParamsMp[record.UserId], TaskActionParam{
			Type:      taskType,
			Subtype:   taskSubType,
			Amount:    math.Abs(record.NetAmount), // 累计盈亏
			GameType:  record.GameType,
			VenueCode: record.VenueCode,
			GameCode:  record.GameCode,
		}, TaskActionParam{
			Type:      taskType,
			Subtype:   enmus.DailyTaskSubType_Bet,
			Amount:    record.ValidBetamount, // 累计投注
			GameType:  record.GameType,
			VenueCode: record.VenueCode,
			GameCode:  record.GameCode,
		})
	} else {
		userParamsMp[record.UserId] = []TaskActionParam{
			{
				Type:      taskType,
				Subtype:   taskSubType,
				Amount:    math.Abs(record.NetAmount), // 累计盈亏
				GameType:  record.GameType,
				VenueCode: record.VenueCode,
				GameCode:  record.GameCode,
			},
			{
				Type:      taskType,
				Subtype:   enmus.DailyTaskSubType_Bet,
				Amount:    record.ValidBetamount, // 累计投注
				GameType:  record.GameType,
				VenueCode: record.VenueCode,
				GameCode:  record.GameCode,
			},
		}
	}

	return userParamsMp
}

// DoUserTaskAction - 触发用户任务
// @param {string} userId
// @param {[]TaskActionParam} params
// @param {bool} isReload 是否载入（主要触发可能在不同进程，需要设置个标识让用户任务数据重新从库中载入）
// @returns
func DoUserTaskAction(userId string, params []TaskActionParam, isReload bool) {
	//global.G_LOG.Infof("[DoUserTaskAction] userId=%s, params=%+v", userId, params)

	// 先同步一次任务配置
	//DoSyncDailyTaskConfig()

	taskMap := GetUserTaskDataToMap(userId)

	now := time.Now().Unix()

	isSave := false

	nowDate := time.Now().Format(tool.TimeDateLayout)

	taskIds := GetDailyTaskConfigIndexes(getCacheUserMerchantCode(userId))
	for _, param := range params {
		for _, taskId := range taskIds {
			taskConfig := GetDailyTaskConfig(taskId)
			if taskConfig != nil {
				// 排除结束任务
				if taskConfig.Status == enmus.DailyTaskStats_Over {
					continue
				}

				// 判断周期重置
				if v, ok := taskMap[taskId]; ok {
					if CheckDailyTaskResetByCycle(taskConfig, v.Date) {
						v.Amount = 0
						v.Status = enmus.UserTaskStats_None
						v.Date = nowDate
						taskMap[taskId] = v
						isSave = true
					}
				}

				// 判断任务类型（排除每日任务类型）
				if (taskConfig.Type != enmus.DailyTaskType_Daily &&
					taskConfig.Type != param.Type) ||
					taskConfig.Subtype != param.Subtype {
					continue
				}

				// 判断任务起止时间（未开始或已结束）
				sTime := tool.CovertTimestampFromAutomaticTypeTime(taskConfig.StartAt)
				eTime := tool.CovertTimestampFromAutomaticTypeTime(taskConfig.EndAt)
				if sTime > 0 && eTime > 0 && (now < sTime || now >= eTime) {
					if now >= eTime {
						// 如果是已过期任务则删除对应用户任务节省数据空间
						delete(taskMap, taskId)
					}
					continue
				}

				// 判断其他条件
				if (len(taskConfig.GameType) > 0 && !slices.Contains(strings.Split(taskConfig.GameType, ","), param.GameType)) ||
					(len(taskConfig.VenueCode) > 0 && !slices.Contains(strings.Split(taskConfig.VenueCode, ","), param.VenueCode)) ||
					(len(taskConfig.ChannelCode) > 0 && taskConfig.ChannelCode != param.ChannelCode) {
					continue
				}

				// 判断包含游戏
				if len(taskConfig.IncludeGameCodes) > 0 {
					gameCodes := strings.Split(taskConfig.IncludeGameCodes, ",")
					if !slices.Contains(gameCodes, param.GameCode) {
						continue
					}
				}

				// 判断屏蔽游戏
				if len(taskConfig.ExcludeGameCodes) > 0 {
					gameCodes := strings.Split(taskConfig.ExcludeGameCodes, ",")
					if slices.Contains(gameCodes, param.GameCode) {
						continue
					}
				}

				// 更新任务进度及状态
				if v, ok := taskMap[taskId]; ok {
					if v.Status == enmus.UserTaskStats_None { // 仅不可领状态
						v.Amount += param.Amount
						v.Amount = min(v.Amount, taskConfig.Amount)
						if v.Amount == taskConfig.Amount {
							v.Status = enmus.UserTaskStats_Reward // 更新为可领取
						}
						taskMap[taskId] = v
						isSave = true
					}
				} else {
					taskData := UserTaskData{
						Id:     taskId,
						Amount: min(param.Amount, taskConfig.Amount),
						Date:   "",
						Status: enmus.UserTaskStats_None,
					}

					//if taskConfig.Type == enmus.DailyTaskType_Daily { // 仅每日任务设置时间（节省存储空间）
					taskData.Date = time.Now().Format(tool.TimeDateLayout)
					//}

					if taskData.Amount == taskConfig.Amount {
						taskData.Status = enmus.UserTaskStats_Reward // 更新为可领取
					}

					taskMap[taskId] = taskData

					isSave = true
				}

			}
		}
	}

	if isSave {
		// if isReload {
		// 	setUserTaskReload(userId) // 用户任务缓存数据需要重新载入（主要触发在不同进程间进行）
		// }
		SetUserTaskDataByMap(userId, taskMap)
	}
}

// GetUserTaskList - 获取用户任务列表
// @param {string} merchantCode
// @param {string} userId
// @param {int} taskType
// @param {int} page
// @param {int} pageSize
// @returns []dos.DailyTask
func GetUserTaskList(merchantCode string, userId string, taskType int,
	page, pageSize int) ([]UserTaskResult, int64) {
	result := []UserTaskResult{}

	// 先同步一次任务配置
	//DoSyncDailyTaskConfig()

	// 获取任务索引（剔除已结束和非指定任务类型任务）
	taskIds := GetDailyTaskConfigIndexesByType(merchantCode, taskType)

	s := (page - 1) * pageSize
	e := min(pageSize*page, len(taskIds))
	if s > e {
		return result, 0
	}

	venueNamesMap := GetFcVenueNameMap([]string{})

	// 获取用户任务数据
	userTaskMap := GetUserTaskDataToMap(userId)

	// taskIds = doUserTaskSortEndByStats(userTaskMap, taskIds)

	lmtIds := taskIds[s:e]
	for _, taskId := range lmtIds {
		taskConfig := GetDailyTaskConfig(taskId)
		if taskConfig != nil {
			res := UserTaskResult{}
			tool.JsonMapper(taskConfig, &res)
			if venueName, ret := venueNamesMap[taskConfig.VenueCode]; ret {
				res.VenueName = venueName
			}

			if v, ok := userTaskMap[taskId]; ok {
				// 存在用户任务数据则替换进度、状态
				res.CurrAmount = v.Amount
				res.Status = v.Status
			} else {
				res.Status = enmus.UserTaskStats_None
			}

			if len(taskConfig.IncludeGameCodes) > 0 &&
				!strings.Contains(taskConfig.IncludeGameCodes, ",") {
				// 只处理单包含游戏
				venueGame := FindByKeyFcVenueGameFirst(&dos.FcVenueGame{
					VenueCode: taskConfig.VenueCode,
					GameCode:  taskConfig.IncludeGameCodes,
				})
				res.TableId = venueGame.Id
				res.Gtype = venueGame.Gtype
				res.GameCode = taskConfig.IncludeGameCodes
				res.GameName = venueGame.GameName
			}

			result = append(result, res)
		}
	}

	return result, int64(len(taskIds))
}

// doUserTaskSortEndByStats - 根据用户任务状态排序（将已完成放在最后）
// @param {map[string]UserTaskData} userTaskMp
// @param {[]string} taskIds
// @returns []string
func doUserTaskSortEndByStats(userTaskMp map[string]UserTaskData, taskIds []string) []string {
	rewardedIds := []string{}
	newTaskIds := []string{}
	for _, taskId := range taskIds {
		if v, ok := userTaskMp[taskId]; ok {
			if v.Status == enmus.UserTaskStats_Rewarded {
				rewardedIds = append(rewardedIds, taskId)
				continue
			}
		}

		newTaskIds = append(newTaskIds, taskId)
	}

	return append(newTaskIds, rewardedIds...)
}

// GetUserTaskPromotionTotalVal - 获取任务福利统计
// @param {string} userId
// @returns float64
func GetUserTaskPromotionTotalVal(userId string) float64 {
	fundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_task")
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(userId) > 0 {
		query = query.Where("user_id = ? AND funding_subtype = ?", userId, fundingSubType)
	} else {
		query = query.Where("funding_subtype = ?", fundingSubType)
	}

	var totalAmount *float64
	err := query.Select("sum(amount) as totalAmount").Scan(&totalAmount).Error
	if err != nil {
		global.G_LOG.Errorf("[GetUserTaskPromotionTotalVal] Accumulate daily bonus promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}

// GetUserTaskPromotionTotalByDate - 获取任务福利统计
// @param {string} merchantCode
// @param {automaticType.Time} sTime
// @param {automaticType.Time} eTime
// @returns float64
func GetUserTaskPromotionTotalByDate(merchantCode string,
	sTime automaticType.Time, eTime automaticType.Time) float64 {
	fundingSubType := GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_task")
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
		global.G_LOG.Errorf("[GetUserTaskPromotionTotalByDate] Accumulate daily bonus promotion total failed: %v", err.Error())
		return 0.00
	}

	if totalAmount == nil {
		return 0.00
	}

	return *totalAmount
}
