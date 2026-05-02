// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/patrickmn/go-cache"
)

var userVenueEntryCache = cache.New(10*time.Minute, 15*time.Minute)

// getCacheVenueEntryRecordKey - 设置场馆进入记录缓存Key
// @param {string} userId
// @returns string
func getCacheVenueEntryRecordKey(userId string) string {
	return "cacheVenueEntryRecordKey:" + userId
}

// CheckVenueEntryRecordVal - 判断是否有场馆进入记录缓存值
// @param {string} userId
// @returns bool
func CheckVenueEntryRecordVal(userId string) bool {
	return len(GetVenueEntryRecordVal(userId)) > 0
}

// saveUserVenueEntryToDB - 保存数据到库
// @param {string} userId
// @param {string} venueCode
// @returns bool
func saveUserVenueEntryToDB(userId string, venueCode string) {
	data := dos.FcUserVenueEntry{}
	err := global.G_DB.Model(&dos.FcUserVenueEntry{}).Select("user_id").
		Where("user_id = ?", userId).Scan(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[saveUserVenueEntry] Find user venue entry data failed: userId=%s, err=%v", userId, err.Error())
		return
	}

	if len(data.UserId) == 0 {
		data.UserId = userId
		data.VenueCode = venueCode
		data.CreatedAt = automaticType.Now()
		err = global.G_DB.Model(&dos.FcUserVenueEntry{}).Create(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[saveUserVenueEntry] Create user venue entry data failed: userId=%s, err=%v",
				userId, err.Error())
		}
	} else {
		err = global.G_DB.Model(&dos.FcUserVenueEntry{}).Where("user_id = ?", userId).
			Update("venue_code", venueCode).Error
		if err != nil {
			global.G_LOG.Errorf("[saveUserVenueEntry] Update user venue entry data failed: userId=%s, err=%v",
				userId, err.Error())
		}
	}
}

// SetVenueEntryRecordVal - 设置场馆进入记录缓存值
// @param {string} userId
// @param {string} venueCode
// @returns
func SetVenueEntryRecordVal(userId string, venueCode string) {
	//global.G_LOG.Infof("[SetVenueEntryRecordVal] =======> userId=%s, venueCode=%s", userId, venueCode)

	userVenueEntryCache.Set(getCacheVenueEntryRecordKey(userId), venueCode,
		time.Duration(tool.RandInt(1, 3))*time.Minute)
	saveUserVenueEntryToDB(userId, venueCode)
}

// GetVenueEntryRecordVal - 获取场馆进入记录缓存值
// @param {string} userId
// @returns string
func GetVenueEntryRecordVal(userId string) string {
	venueCode, ret := userVenueEntryCache.Get(getCacheVenueEntryRecordKey(userId))
	if !ret {
		data := dos.FcUserVenueEntry{}
		err := global.G_DB.Model(&dos.FcUserVenueEntry{}).Select("venue_code").
			Where("user_id = ?", userId).Scan(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[GetVenueEntryRecordVal] Find user venue entry data failed: userId=%s, err=%v", userId, err.Error())
			return ""
		}

		userVenueEntryCache.Set(getCacheVenueEntryRecordKey(userId), data.VenueCode,
			time.Duration(tool.RandInt(1, 3))*time.Minute)

		return data.VenueCode
	}

	if v, ok := venueCode.(string); ok {
		return v
	}

	return ""
}

// DelVenueEntryRecordVal - 删除场馆进入记录缓存值
// @param {string} userId
// @returns
func DelVenueEntryRecordVal(userId string) {
	//global.G_LOG.Infof("[DelVenueEntryRecordVal] =======> userId=%s", userId)

	userVenueEntryCache.Delete(getCacheVenueEntryRecordKey(userId))
	err := global.G_DB.Where("user_id = ?", userId).Delete(&dos.FcUserVenueEntry{}).Error
	if err != nil {
		global.G_LOG.Errorf("[DelVenueEntryRecordVal] Delete user venue entry failed: userId=%s, err=%v", userId, err.Error())
	}
}

// CheckVenueEntryRecordValNoCache - 判断是否有场馆进入值（非缓存）
// @param {string} userId
// @returns bool
func CheckVenueEntryRecordValNoCache(userId string) bool {
	data := dos.FcUserVenueEntry{}
	err := global.G_DB.Model(&dos.FcUserVenueEntry{}).Select("venue_code").
		Where("user_id = ?", userId).Scan(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckVenueEntryRecordValNoCache] Find user venue entry data failed: userId=%s, err=%v", userId, err.Error())
		return false
	}

	return len(data.VenueCode) > 0
}
