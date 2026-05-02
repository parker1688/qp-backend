// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func CheckDepartmentName(departmentName string, id string) bool {
	if len(id) == 0 {
		data := &dos.Department{}
		err := global.G_DB.Model(&dos.Department{}).Select("id").
			Where("department_name = ?", departmentName).First(&data).Error
		return err == nil && len(data.Id) > 0
	} else {
		data := &dos.Department{}
		err := global.G_DB.Model(&dos.Department{}).Select("id").
			Where("id != ? and department_name = ?", id, departmentName).First(&data).Error
		return err == nil && len(data.Id) > 0
	}
}

func SaveDepartment(vo *dos.Department) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageDepartment(page, pageSize int, vo *dos.Department) (ret []*dos.Department, total int64) {
	query := global.G_DB.Model(&dos.Department{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DepartmentName) > 0 {
		query = query.Where("department_name = ?", vo.DepartmentName)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.ParentId > 0 {
		query = query.Where("parent_id = ?", vo.ParentId)
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
	var dataSlice []*dos.Department
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDepartment(vo *dos.Department) []*dos.Department {
	var data []*dos.Department
	query := global.G_DB.Model(&dos.Department{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DepartmentName) > 0 {
		query = query.Where("department_name = ?", vo.DepartmentName)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.ParentId > 0 {
		query = query.Where("parent_id = ?", vo.ParentId)
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

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateDepartment(vo *dos.Department) bool {
	return global.G_DB.Model(vo).Where("id = ?", vo.Id).Updates(map[string]interface{}{
		"department_name": vo.DepartmentName,
		"sort":            vo.Sort,
		"parent_id":       vo.ParentId,
		"status":          vo.Status,
		"remarks":         vo.Remarks,
		"create_by":       vo.CreateBy,
		"update_by":       vo.UpdateBy,
	}).Error == nil
}

func DeleteDepartment(vo *dos.Department) bool {
	return global.G_DB.Model(&dos.Department{}).Where("id = ?", vo.Id).Delete(vo).Error != nil
}

func FindDepartmentAll() []*dos.Department {
	var dataSlice []*dos.Department
	global.G_DB.Model(&dos.Department{}).Find(&dataSlice)
	return dataSlice
}

func GetDepartmentName(departmentId string) string {
	department := dos.Department{}
	err := global.G_DB.Model(&dos.Department{}).Select("department_name").Where("id = ?", departmentId).First(&department).Error
	if err != nil {
		global.G_LOG.Errorf("[GetDepartmentName] find department data failed: id=%s, err=%v", departmentId, err.Error())
		return ""
	}

	return department.DepartmentName
}
