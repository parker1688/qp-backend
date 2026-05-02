// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveDictsDetail(vo *dos.DictsDetail) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageDictsDetail(page, pageSize int, vo *dos.DictsDetail) (ret []*dos.DictsDetail, total int64) {
	query := global.G_DB.Model(&dos.DictsDetail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}

	if len(vo.DictsTag) > 0 {
		query = query.Where("dicts_tag = ?", vo.DictsTag)
	}

	if len(vo.DictsValue) > 0 {
		query = query.Where("dicts_value = ?", vo.DictsValue)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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
	var dataSlice []*dos.DictsDetail
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDictsDetail(vo *dos.DictsDetail) []*dos.DictsDetail {
	var data []*dos.DictsDetail
	query := global.G_DB.Model(&dos.DictsDetail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}
	query.Take(&data)
	return data
}

func FindByKeyDictsDetailFirst(vo *dos.DictsDetail) *dos.DictsDetail {
	var data *dos.DictsDetail
	query := global.G_DB.Model(&dos.DictsDetail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.DictsTag) > 0 {
		query = query.Where("dicts_tag = ?", vo.DictsTag)
	}
	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateDictsDetail(vo *dos.DictsDetail) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"dicts_type_code": vo.DictsTypeCode,
		"dicts_tag":       vo.DictsTag,
		"dicts_value":     vo.DictsValue,
		"sort":            vo.Sort,
		"status":          vo.Status,
		"remarks":         vo.Remarks,
		"create_by":       vo.CreateBy,
		"update_by":       vo.UpdateBy,
	}).Error == nil
}

func DeleteDictsDetail(vo *dos.DictsDetail) bool {
	return global.G_DB.Model(&dos.DictsDetail{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func FindByKeyDictsDetailAll(vo *dos.DictsDetail) []*dos.DictsDetail {
	var data []*dos.DictsDetail
	query := global.G_DB.Model(&dos.DictsDetail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.DictsTypeCode) > 0 {
		query = query.Where("dicts_type_code = ?", vo.DictsTypeCode)
	}
	//query = query.Where("status = ?", enmus.STATUS_ENABLE).Order(" sort desc")
	query.Order("sort desc").Find(&data)
	return data
}
