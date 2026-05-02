package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var activityCache = cache.New(10*time.Minute, 15*time.Minute) // 活动缓存

const (
	activityConfigIndexCacheKey = "activityConfigIndexCacheKey" // 活动配置索引缓存键
	activityConfigCacheKey      = "activityConfigCacheKey"      // 活动配置缓存键
	activityUserCacheKey        = "activityUserCacheKey"        // 活动用户缓存键
)

// ====================================================
// 活动配置索引方法
// ====================================================
// getActivityConfigIndexCacheKey - 获取活动配置索引缓存键值
// @param {string} merchantCode
// @param {int} activityType
// @returns string
func getActivityConfigIndexCacheKey(merchantCode string, activityType int) string {
	return fmt.Sprintf("%s:%s:%d", activityConfigIndexCacheKey, merchantCode, activityType)
}

// loadActivityIndexesCacheFromDB - 从库中载入活动配置索引到缓存（排除已结束活动）
// @param {string} merchantCode
// @param {int} activityType
// @returns []string
func loadActivityIndexesCacheFromDB(merchantCode string, activityType int) []string {
	data := []dos.FcPromotionInfo{}
	err := global.G_DB.Model(&dos.FcPromotionInfo{}).Omit("sort", "create_time", "create_by", "update_time", "update_by").
		Where("merchant_code = ? AND promotion_type = ? AND status = ? AND (end_time IS NULL OR end_time > ?)",
			merchantCode, activityType, enmus.ActivityStats_Opening, time.Now()).Order("sort").Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadActivityIndexesCacheFromDB] Find activity indexes failed: %v", err.Error())
		return []string{}
	}

	if len(data) == 0 {
		return []string{}
	}

	isLoad := checkLoadActivityConfigOnceLock(merchantCode) // 活动配置只全部载入一次之后每个活动都单独处理

	lis := []string{}
	for _, v := range data {
		lis = append(lis, v.Id)

		if isLoad {
			SetActivityConfig(v.Id, v)
		}
	}

	setLoadActivityConfigOnceLock(merchantCode)

	SetActivityConfigIndexes(merchantCode, activityType, lis)

	return lis
}

// GetActivityConfigIndexes - 获取活动配置索引
// @param {string} merchantCode
// @param {int} activityType
// @returns []string
func GetActivityConfigIndexes(merchantCode string, activityType int) []string {
	data, ret := activityCache.Get(getActivityConfigIndexCacheKey(merchantCode, activityType))
	if !ret {
		// 没有缓存则从库中载入
		return loadActivityIndexesCacheFromDB(merchantCode, activityType)
	}

	if v, ok := data.([]string); ok {
		return v
	}

	return []string{}
}

// SetActivityConfigIndexes - 更新活动配置索引
// @param {string} merchantCode
// @param {int} activityType
// @param {[]string} data
// @returns
func SetActivityConfigIndexes(merchantCode string, activityType int, data []string) {
	activityCache.Set(getActivityConfigIndexCacheKey(merchantCode, activityType), data,
		time.Duration(tool.RandInt(1, 4))*time.Minute)
}

// ====================================================
// 活动配置方法
// ====================================================
// getActivityConfigCacheKey - 获取活动配置缓存键值
// @param {string} activityId
// @returns string
func getActivityConfigCacheKey(activityId string) string {
	return activityConfigCacheKey + ":" + activityId
}

// checkLoadActivityConfigOnceLock - 判断载入活动配置一次锁
// @param {string} merchantCode
// @returns bool
func checkLoadActivityConfigOnceLock(merchantCode string) bool {
	_, ret := activityCache.Get(getActivityConfigCacheKey("onceLock::" + merchantCode))
	return !ret
}

// setLoadActivityConfigOnceLock - 设置载入活动配置一次锁
// @param {string} merchantCode
// @returns
func setLoadActivityConfigOnceLock(merchantCode string) {
	activityCache.Set(getActivityConfigCacheKey("onceLock::"+merchantCode), "1", cache.NoExpiration)
}

// loadActivityCacheFromDB - 活动配置从数据库载入到缓存
// @param {string} activityId
// @returns *dos.Activity
func loadActivityCacheFromDB(activityId string) *dos.FcPromotionInfo {
	var data dos.FcPromotionInfo
	err := global.G_DB.Model(&dos.FcPromotionInfo{}).Omit("sort", "create_time", "create_by", "update_time", "update_by").
		Where("id = ?", activityId).Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[loadActivityCacheFromDB] Find activity data failed: activityId=%s, err=%v", activityId, err.Error())
		return nil
	}

	SetActivityConfig(activityId, data)

	return &data
}

// GetActivityConfig - 获取活动配置
// @param {string} activityId
// @returns *dos.FcPromotionInfo
func GetActivityConfig(activityId string) *dos.FcPromotionInfo {
	data, ret := activityCache.Get(getActivityConfigCacheKey(activityId))
	if !ret {
		// 没有缓存则从库中载入
		return loadActivityCacheFromDB(activityId)
	}

	if v, ok := data.(dos.FcPromotionInfo); ok {
		return &v
	}

	return nil
}

// SetActivityConfig - 更新活动配置
// @param {string} activityId
// @param {dos.FcPromotionInfo} data
// @returns
func SetActivityConfig(activityId string, data dos.FcPromotionInfo) {
	activityCache.Set(getActivityConfigCacheKey(activityId), data,
		time.Duration(tool.RandInt(3, 6))*time.Minute)
}
