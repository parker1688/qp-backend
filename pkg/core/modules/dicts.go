// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveDicts(vo *dos.Dicts) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageDicts(page, pageSize int, vo *dos.Dicts) (ret []*dos.Dicts, total int64) {
	query := global.G_DB.Model(&dos.Dicts{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DictsName) > 0 {
		query = query.Where("dicts_name = ?", vo.DictsName)
	}

	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Dicts
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDicts(vo *dos.Dicts) []*dos.Dicts {
	var data []*dos.Dicts
	query := global.G_DB.Model(&dos.Dicts{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DictsName) > 0 {
		query = query.Where("dicts_name = ?", vo.DictsName)
	}

	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateDicts(vo *dos.Dicts) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"dicts_name":      vo.DictsName,
		"dicts_type_code": vo.DictsTypeCode,
		"remarks":         vo.Remarks,
		"update_by":       vo.UpdateBy,
	}).Error == nil
}

func DeleteDicts(vo *dos.Dicts) bool {
	return global.G_DB.Model(&dos.Dicts{}).Where("id = ?", vo.Id).Delete(vo).Error != nil
}
