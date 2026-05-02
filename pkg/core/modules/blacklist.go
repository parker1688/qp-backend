// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveBlacklist(vo *dos.Blacklist) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageBlacklist(page, pageSize int, vo *dos.Blacklist) (ret []*dos.Blacklist, total int64) {
	query := global.G_DB.Model(&dos.Blacklist{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type >= 0 {
		query = query.Where("`type` = ?", vo.Type)
	}

	if len(vo.Value) > 0 {
		query = query.Where("`value` like ?", "%"+vo.Value+"%")
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Blacklist
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyBlacklist(vo *dos.Blacklist) []*dos.Blacklist {
	var data []*dos.Blacklist
	query := global.G_DB.Model(&dos.Blacklist{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type >= 0 {
		query = query.Where("type = ?", vo.Type)
	}

	if len(vo.Value) > 0 {
		query = query.Where("`value` like ?", "%"+vo.Value+"%")
	}

	query.Find(&data)
	return data
}

func FindByKeyBlacklistFirst(vo *dos.Blacklist) *dos.Blacklist {
	var data *dos.Blacklist
	query := global.G_DB.Model(&dos.Blacklist{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type >= 0 {
		query = query.Where("type = ?", vo.Type)
	}

	if len(vo.Value) > 0 {
		query = query.Where("`value` like ?", "%"+vo.Value+"%")
	}

	query.Take(&data)
	return data
}

func UpdateBlacklist(vo *dos.Blacklist) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"value":     vo.Value,
		"update_by": vo.UpdateBy,
	}).Error == nil
}

func DeleteBlacklist(vo *dos.Blacklist) bool {
	return global.G_DB.Model(&dos.Blacklist{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// IsBlacklisted - 是否被黑名单
// @param {string} ip IP地址
// @param {string} deviceId 设备码
// @returns bool
func IsBlacklisted(ip string, deviceId string) bool {
	data := dos.Blacklist{}
	query := global.G_DB.Model(&dos.Blacklist{})
	if len(ip) > 0 && len(deviceId) > 0 {
		query = query.Where("`value` like ? OR `value` like ?", "%"+ip+"%", "%"+deviceId+"%")
	} else if len(ip) > 0 {
		query = query.Where("`value` like ?", "%"+ip+"%")
	} else if len(deviceId) > 0 {
		query = query.Where("`value` like ?", "%"+deviceId+"%")
	}

	err := query.Select("id").First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[IsBlacklisted] Find blacklist data failed: %v", err.Error())
		return false
	}

	return len(data.Id) > 0
}
