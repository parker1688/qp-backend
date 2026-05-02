// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVipRebate(vo *dos.FcVipRebate) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVipRebate(page, pageSize int, vo *dos.FcVipRebate) (ret []*dos.FcVipRebate, total int64) {
	query := global.G_DB.Model(&dos.FcVipRebate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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
	query.Count(&count)
	var dataSlice []*dos.FcVipRebate
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVipRebate(vo *dos.FcVipRebate) []*dos.FcVipRebate {
	var data []*dos.FcVipRebate
	query := global.G_DB.Model(&dos.FcVipRebate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

func FindByKeyFcVipRebateFirst(vo *dos.FcVipRebate) *dos.FcVipRebate {
	var data *dos.FcVipRebate
	query := global.G_DB.Model(&dos.FcVipRebate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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
func UpdateFcVipRebate(vo *dos.FcVipRebate) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"level":        vo.Level,
		"vip_name":     vo.VipName,
		"chess":        vo.Chess,
		"elecgame":     vo.Elecgame,
		"zhenren":      vo.Zhenren,
		"sport":        vo.Sport,
		"esport":       vo.Esport,
		"lottery":      vo.Lottery,
		"fish":         vo.Fish,
		"rebate_limit": vo.RebateLimit,
		// "merchant_code": vo.MerchantCode,
		"create_by": vo.CreateBy,
		"update_by": vo.UpdateBy,
	}).Error == nil
}

func DeleteFcVipRebate(vo *dos.FcVipRebate) bool {
	return global.G_DB.Model(&dos.FcVipRebate{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
