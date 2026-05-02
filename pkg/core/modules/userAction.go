// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveUserAction(vo *dos.UserAction) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageUserAction(page, pageSize int, vo *dos.UserAction) (ret []*dos.UserAction, total int64) {
	query := global.G_DB.Model(&dos.UserAction{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Url) > 0 {
		query = query.Where("url = ?", vo.Url)
	}

	if len(vo.Method) > 0 {
		query = query.Where("method = ?", vo.Method)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if len(vo.Body) > 0 {
		query = query.Where("body = ?", vo.Body)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.UserAction
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyUserAction(vo *dos.UserAction) []*dos.UserAction {
	var data []*dos.UserAction
	query := global.G_DB.Model(&dos.UserAction{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Url) > 0 {
		query = query.Where("url = ?", vo.Url)
	}

	if len(vo.Method) > 0 {
		query = query.Where("method = ?", vo.Method)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if len(vo.Body) > 0 {
		query = query.Where("body = ?", vo.Body)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateUserAction(vo *dos.UserAction) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_name": vo.UserName,
		"url":       vo.Url,
		"method":    vo.Method,
		"ip":        vo.Ip,
		"body":      vo.Body,
	}).Error == nil
}

func DeleteUserAction(vo *dos.UserAction) bool {
	return global.G_DB.Model(&dos.UserAction{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
