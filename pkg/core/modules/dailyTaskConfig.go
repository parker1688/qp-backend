// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var dailyTaskCache = cache.New(10*time.Minute, 15*time.Minute) // 任务缓存

const (
	cacheDailyTaskConfigIndexKey = "cacheDailyTaskConfigIndex" // 任务配置表索引缓存KEY
	cacheDailyTaskConfigKey      = "cacheDailyTaskConfig"      // 任务配置表缓存KEY
	cacheDailyTaskConfigSyncKey  = "cacheDailyTaskConfigSync"  // 任务配置同步缓存KEY（用于后台数据修改同步缓存）

	SyncDailyTaskConfigIndexMark = "i" // 同步索引标识
	SyncDailyTaskConfigMark      = "d" // 同步配置标识
)

var (
	dailyTaskSyncExpiresDuration          = 30 * time.Minute       // 任务配置同步过期时间
	dailyTaskConfigIndexExpiresDurationFn = func() time.Duration { // 任务配置索引过期时间
		return time.Duration(tool.RandInt(3, 6)) * time.Minute
	}
	dailyTaskConfigExpiresDurationFn = func() time.Duration { // 任务配置过期时间
		return time.Duration(tool.RandInt(3, 6)) * time.Minute
	}
)

// getDailyTaskConfigIndexKey - 获取任务配置缓存索引KEY
// @param {string} merchantCode
// @returns string
func getDailyTaskConfigIndexKey(merchantCode string) string {
	return cacheDailyTaskConfigIndexKey + ":" + merchantCode
}

// loadCacheDailyTaskIndexesFromDB - 从库中载入任务配置索引到缓存（排除已结束任务）
// @param {string} merchantCode
// @returns []string
func loadCacheDailyTaskIndexesFromDB(merchantCode string) []string {
	data := []dos.DailyTask{}
	err := global.G_DB.Model(&dos.DailyTask{}).Omit("sort", "create_time", "create_by", "update_time", "update_by").
		Where("merchant_code = ? AND status = ? AND (end_at IS NULL OR end_at > ?)",
			merchantCode, enmus.DailyTaskStats_Normal, time.Now()).Order("sort").Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadCacheDailyTaskIndexesFromDB] Find daily task data failed: %v", err.Error())
		return []string{}
	}

	if len(data) == 0 {
		return []string{}
	}

	isLoad := checkLoadDailyTaskConfigOnceLock() // 任务配置只全部载入一次之后每个任务都单独处理

	lis := []string{}
	for _, v := range data {
		lis = append(lis, v.Id)

		if isLoad {
			SetDailyTaskConfig(v.Id, v)
		}
	}

	// 同步数据到缓存记录（防止重复同步）
	if isLoad {
		doSyncDailyTaskToCacheAll("")
	} else {
		doSyncDailyTaskToCacheAll(SyncDailyTaskConfigIndexMark + "$" + merchantCode) // 仅索引
	}

	setLoadDailyTaskConfigOnceLock()

	SetDailyTaskConfigIndexes(merchantCode, lis)

	return lis
}

// GetDailyTaskConfigIndexes - 获取任务配置索引
// @param {string} merchantCode
// @returns []string
func GetDailyTaskConfigIndexes(merchantCode string) []string {
	data, ret := dailyTaskCache.Get(getDailyTaskConfigIndexKey(merchantCode))
	if !ret {
		// 没有缓存则从库中载入
		return loadCacheDailyTaskIndexesFromDB(merchantCode)
	}

	if v, ok := data.([]string); ok {
		return v
	}

	return []string{}
}

// GetDailyTaskConfigIndexesByType - 根据类型获取任务索引（排除已结束任务）
// @param {string} merchantCode
// @param {int} taskType
// @returns []string
func GetDailyTaskConfigIndexesByType(merchantCode string, taskType int) []string {
	now := time.Now().Unix()

	newTaskIds := []string{}
	taskIds := GetDailyTaskConfigIndexes(merchantCode)
	for _, taskId := range taskIds {
		taskConfig := GetDailyTaskConfig(taskId)
		if taskConfig == nil {
			continue
		}

		// 判断任务是否已结束
		eTime := tool.CovertTimestampFromAutomaticTypeTime(taskConfig.EndAt)
		if eTime > 0 && now >= eTime {
			continue
		}

		// 任务是否为指定类型
		if taskConfig.Type != taskType {
			continue
		}

		newTaskIds = append(newTaskIds, taskId)
	}

	return newTaskIds
}

// SetDailyTaskConfigIndexes - 更新任务配置索引
// @param {string} merchantCode
// @param {[]string} data
// @returns
func SetDailyTaskConfigIndexes(merchantCode string, data []string) {
	dailyTaskCache.Set(getDailyTaskConfigIndexKey(merchantCode), data,
		dailyTaskConfigIndexExpiresDurationFn())
}

// getDailyTaskConfigKey - 获取任务配置缓存KEY
// @param {string} taskId
// @returns string
func getDailyTaskConfigKey(taskId string) string {
	return cacheDailyTaskConfigKey + ":" + taskId
}

// checkLoadDailyTaskConfigOnceLock - 判断载入任务配置一次锁
// @returns bool
func checkLoadDailyTaskConfigOnceLock() bool {
	_, ret := dailyTaskCache.Get(getDailyTaskConfigKey("onceLock"))
	return !ret
}

// setLoadDailyTaskConfigOnceLock - 设置载入任务配置一次锁
func setLoadDailyTaskConfigOnceLock() {
	dailyTaskCache.Set(getDailyTaskConfigKey("onceLock"), "1", cache.NoExpiration)
}

// loadCacheDailyTaskFromDB - 任务配置从数据库载入到缓存
// @param {string} taskId
// @returns *dos.DailyTask
func loadCacheDailyTaskFromDB(taskId string) *dos.DailyTask {
	var data dos.DailyTask
	err := global.G_DB.Model(&dos.DailyTask{}).Omit("sort", "create_time", "create_by", "update_time", "update_by").
		Where("id = ?", taskId).Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadCacheDailyTaskFromDB] Find daily task data failed: taskId=%s, err=%v", taskId, err.Error())
		return nil
	}

	doSyncDailyTaskToCacheAll(SyncDailyTaskConfigMark + "$" + taskId) // 防止重复同步（仅配置）

	SetDailyTaskConfig(taskId, data)

	return &data
}

// GetDailyTaskConfig - 获取任务配置
// @param {string} taskId
// @returns *dos.DailyTask
func GetDailyTaskConfig(taskId string) *dos.DailyTask {
	data, ret := dailyTaskCache.Get(getDailyTaskConfigKey(taskId))
	if !ret {
		// 没有缓存则从库中载入
		return loadCacheDailyTaskFromDB(taskId)
	}

	if v, ok := data.(dos.DailyTask); ok {
		return &v
	}

	return nil
}

// SetDailyTaskConfig - 更新任务配置
// @param {string} taskId
// @param {dos.DailyTask} data
// @returns
func SetDailyTaskConfig(taskId string, data dos.DailyTask) {
	dailyTaskCache.Set(getDailyTaskConfigKey(taskId), data,
		dailyTaskConfigExpiresDurationFn())
}

// AddSyncDailyTaskConfig - 加入需同步的任务配置（后台使用）
// @param {string} mark 标识（i 索引；d 配置；）
// @param {string} value 值（mark==i 商户码；mark==d 任务id）
// @returns
func AddSyncDailyTaskConfig(mark string, value string) {
	err := global.G_REDIS.SAdd(context.Background(), cacheDailyTaskConfigSyncKey,
		fmt.Sprintf("%s$%s$%d", mark, value, time.Now().Unix())).Err()
	if err != nil {
		global.G_LOG.Errorf("[AddSyncDailyTaskConfig] Add daily task sync config failed: %v", err.Error())
		return
	}

	expIdent := global.G_REDIS.Get(context.Background(), cacheDailyTaskConfigSyncKey+":ExpIdent").Val()
	if expIdent != "1" {
		// 设置过期
		global.G_REDIS.Set(context.Background(), cacheDailyTaskConfigSyncKey+":ExpIdent", "1", dailyTaskSyncExpiresDuration)
		global.G_REDIS.Expire(context.Background(), cacheDailyTaskConfigSyncKey, dailyTaskSyncExpiresDuration)
	}
}

// setSyncDailyTaskConfigByCache - 更新缓存中的同步任务数据
// @param {[]string} data
// @returns
func setSyncDailyTaskConfigByCache(data []string) {
	dailyTaskCache.Set(cacheDailyTaskConfigSyncKey, data, dailyTaskSyncExpiresDuration)
}

// getSyncDailyTaskConfigByCache - 获取缓存中的同步任务数据
// @returns []string
func getSyncDailyTaskConfigByCache() []string {
	if data, ret := dailyTaskCache.Get(cacheDailyTaskConfigSyncKey); ret {
		if v, ok := data.([]string); ok {
			return v
		}
	}

	return []string{}
}

// doSyncDailyTaskToCacheAll - 处理全部同步数据到缓存
// @param {string} match 匹配字符串
func doSyncDailyTaskToCacheAll(match string) {
	syncLis := getSyncDailyTaskConfigAll()
	if len(syncLis) == 0 {
		return
	}

	if len(match) == 0 {
		setSyncDailyTaskConfigByCache(syncLis)
	} else {
		cacheLis := getSyncDailyTaskConfigByCache()
		isSave := false
		for _, v := range syncLis {
			if strings.Contains(v, match) {
				cacheLis = append(cacheLis, v)
				isSave = true
			}
		}
		if isSave {
			setSyncDailyTaskConfigByCache(cacheLis)
		}
	}
}

// getSyncDailyTaskConfigAll - 获取全部需同步的任务配置
// @returns []string
func getSyncDailyTaskConfigAll() []string {
	return global.G_REDIS.SMembers(context.Background(), cacheDailyTaskConfigSyncKey).Val()
}

// DoSyncDailyTaskConfig - 处理同步任务配置（后台新的或改动的数据同步到缓存）
func DoSyncDailyTaskConfig() {
	syncLis := getSyncDailyTaskConfigAll()
	if len(syncLis) == 0 {
		return
	}

	cacheLis := getSyncDailyTaskConfigByCache()
	isSave := false

	global.G_LOG.Infof("[DoSyncDailyTaskConfig] syncLis=%+v, cacheLis=%+v", syncLis, cacheLis)

	// 待更新的商户对应的任务索引
	upgradeIndexMp := map[string]int{}

	for _, v := range syncLis {
		if !slices.Contains(cacheLis, v) {
			// 未进行同步的数据
			tmps := strings.Split(v, "$")
			if len(tmps) > 1 {
				switch tmps[0] {
				case SyncDailyTaskConfigIndexMark: // 是索引
					upgradeIndexMp[tmps[1]] = 1
				case SyncDailyTaskConfigMark: // 是配置
					loadCacheDailyTaskFromDB(tmps[1])
				}

				cacheLis = append(cacheLis, v)
				isSave = true
			}
		}
	}

	for k := range upgradeIndexMp {
		loadCacheDailyTaskIndexesFromDB(k)
	}

	if isSave {
		setSyncDailyTaskConfigByCache(cacheLis)
	}
}

// CheckDailyTaskResetByCycle - 根据周期判断是否重置
// @param {dos.DailyTask} config
// @param {string} date
// @returns bool
func CheckDailyTaskResetByCycle(config *dos.DailyTask, taskDate string) bool {
	now := time.Now()
	switch config.Cycle {
	case enmus.DailyTaskCycle_Daily: // 每日
		return now.Format(tool.TimeDateLayout) != taskDate
	case enmus.DailyTaskCycle_Week: // 每周
		t1 := tool.GetTimeFromString(taskDate + " 12:00:00")
		t2 := tool.GetTimeFromString(now.Format(tool.TimeLayout))
		return tool.IsDifferentWeekCustom(t1, t2, time.Monday)
	case enmus.DailyTaskCycle_Month: // 每月
		return now.Format(tool.TimeDateYearMonLayout) != taskDate
	default:
		return false
	}
}
