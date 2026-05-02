// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/patrickmn/go-cache"
)

var gameRebateCache = cache.New(5*time.Minute, 10*time.Minute)

func SaveFcGameRebate(vo *dos.FcGameRebate) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcGameRebate(page, pageSize int, vo *dos.FcGameRebate) (ret []*dos.FcGameRebate, total int64, err error) {
	query := global.G_DB.Model(&dos.FcGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	var count int64
	if err = query.Count(&count).Error; err != nil {
		return ret, total, err
	}
	var dataSlice []*dos.FcGameRebate
	if err = query.Order("create_time desc, id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice).Error; err != nil {
		return ret, total, err
	}
	return dataSlice, count, nil
}

func FindByKeyFcGameRebate(vo *dos.FcGameRebate) []*dos.FcGameRebate {
	var data []*dos.FcGameRebate
	query := global.G_DB.Model(&dos.FcGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcGameRebateFirst(vo *dos.FcGameRebate) *dos.FcGameRebate {
	var data *dos.FcGameRebate
	query := global.G_DB.Model(&dos.FcGameRebate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcGameRebate(vo *dos.FcGameRebate) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"game_type":      vo.GameType,
		"bonus_rate":     vo.BonusRate,
		"min_bet_amount": vo.MinBetAmount,
		"max_bet_amount": vo.MaxBetAmount,
		"describe":       vo.Describe,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"update_time":    vo.UpdateTime,
	}).Error == nil
}

func DeleteFcGameRebate(vo *dos.FcGameRebate) bool {
	return global.G_DB.Model(&dos.FcGameRebate{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetFcGameRebateListAll - 获取全部游戏返水列表
// @returns []dos.FcGameRebateRes
func GetFcGameRebateListAll() []dos.FcGameRebateRes {
	if list, ok := gameRebateCache.Get("gameRebateConfigKey"); ok {
		if v, ret := list.([]dos.FcGameRebateRes); ret {
			return v
		}
	}

	data := []dos.FcGameRebateRes{}
	err := global.G_DB.Model(&dos.FcGameRebate{}).
		Select("game_type", "bonus_rate", "min_bet_amount", "max_bet_amount").
		Order("game_type, bonus_rate").
		Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcGameRebateListAll] Find game rebate configs failed: err=%v", err.Error())
		return data
	}

	gameRebateCache.Set("gameRebateConfigKey", data,
		time.Duration(tool.RandInt(3, 6))*time.Minute)

	return data
}

// GetCacheUserRebateBonusValue - 获取用户反水值（用于判断领取值减少计算）
// @param {string} userId
// @returns float64
func GetCacheUserRebateBonusValue(userId string) float64 {
	if data, ok := gameRebateCache.Get("userRebateBonusValKey:" + userId); ok {
		if v, ret := data.(float64); ret {
			return v
		}
	}

	return 0.0
}

// GetCacheUserRebateBonusValue - 更新用户反水值
// @param {string} userId
// @param {float64} v
// @returns
func SetCacheUserRebateBonusValue(userId string, v float64) {
	gameRebateCache.Set("userRebateBonusValKey:"+userId, v, -1)
}

// DelCacheUserRebateBonusValue - 删除用户反水值
// @param {string} userId
// @returns
func DelCacheUserRebateBonusValue(userId string) {
	gameRebateCache.Delete("userRebateBonusValKey:" + userId)
}
