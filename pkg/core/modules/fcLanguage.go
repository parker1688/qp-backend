// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcLanguage(vo *dos.FcLanguage) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcLanguage(page, pageSize int, vo *dos.FcLanguage) (ret []*dos.FcLanguage, total int64) {
	query := global.G_DB.Model(&dos.FcLanguage{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcLanguage
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcLanguage(vo *dos.FcLanguage) []*dos.FcLanguage {
	var data []*dos.FcLanguage
	query := global.G_DB.Model(&dos.FcLanguage{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcLanguageFirst(vo *dos.FcLanguage) *dos.FcLanguage {
	var data *dos.FcLanguage
	query := global.G_DB.Model(&dos.FcLanguage{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcLanguage(vo *dos.FcLanguage) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"language": vo.Language,
		"code":     vo.Code,
	}).Error == nil
}

func DeleteFcLanguage(vo *dos.FcLanguage) bool {
	return global.G_DB.Model(&dos.FcLanguage{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
