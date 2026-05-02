// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveLoginLog(vo *dos.LoginLog) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageLoginLog(page, pageSize int, vo *dos.LoginLog) (ret []*dos.LoginLog, total int64) {
	query := global.G_DB.Model(&dos.LoginLog{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_by = ?", vo.CreateTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.LoginLog
	query.Offset((page - 1) * pageSize).Order("create_time desc").Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyLoginLog(vo *dos.LoginLog) []*dos.LoginLog {
	var data []*dos.LoginLog
	query := global.G_DB.Model(&dos.LoginLog{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_by = ?", vo.CreateTime)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateLoginLog(vo *dos.LoginLog) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_name": vo.UserName,
		"ip":        vo.Ip,
		"create_by": vo.CreateTime,
	}).Error == nil
}

func DeleteLoginLog(vo *dos.LoginLog) bool {
	return global.G_DB.Model(&dos.LoginLog{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
