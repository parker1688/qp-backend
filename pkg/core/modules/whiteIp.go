// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveWhiteIp(vo *dos.WhiteIp) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageWhiteIp(page, pageSize int, vo *dos.WhiteIp) (ret []*dos.WhiteIp, total int64) {
	query := global.G_DB.Model(&dos.WhiteIp{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.IpAddr) > 0 {
		query = query.Where("ip_addr = ?", vo.IpAddr)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
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
	var dataSlice []*dos.WhiteIp
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyWhiteIp(vo *dos.WhiteIp) []*dos.WhiteIp {
	var data []*dos.WhiteIp
	query := global.G_DB.Model(&dos.WhiteIp{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.IpAddr) > 0 {
		query = query.Where("ip_addr = ?", vo.IpAddr)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
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

func FindByKeyWhiteIpFirst(vo *dos.WhiteIp) *dos.WhiteIp {
	var data *dos.WhiteIp
	query := global.G_DB.Model(&dos.WhiteIp{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.IpAddr) > 0 {
		query = query.Where("ip_addr like ?", "%"+vo.IpAddr+"%")
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
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
func UpdateWhiteIp(vo *dos.WhiteIp) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"ip_addr":   vo.IpAddr,
		"remarks":   vo.Remarks,
		"create_by": vo.CreateBy,
		"update_by": vo.UpdateBy,
	}).Error == nil
}

func DeleteWhiteIp(vo *dos.WhiteIp) bool {
	return global.G_DB.Model(&dos.WhiteIp{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
