// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveDomainDistribution(vo *dos.DomainDistribution) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageDomainDistribution(page, pageSize int, vo *dos.DomainDistribution) (ret []*dos.DomainDistribution, total int64) {
	query := global.G_DB.Model(&dos.DomainDistribution{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DomainLink) > 0 {
		query = query.Where("domain_link = ?", vo.DomainLink)
	}

	if len(vo.DomainType) > 0 {
		query = query.Where("domain_type = ?", vo.DomainType)
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

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.DomainDistribution
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDomainDistribution(vo *dos.DomainDistribution) []*dos.DomainDistribution {
	var data []*dos.DomainDistribution
	query := global.G_DB.Model(&dos.DomainDistribution{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DomainLink) > 0 {
		query = query.Where("domain_link = ?", vo.DomainLink)
	}

	if len(vo.DomainType) > 0 {
		query = query.Where("domain_type = ?", vo.DomainType)
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

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	query.Find(&data)
	return data
}

func FindByKeyDomainDistributionFirst(vo *dos.DomainDistribution) *dos.DomainDistribution {
	var data *dos.DomainDistribution
	query := global.G_DB.Model(&dos.DomainDistribution{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.DomainLink) > 0 {
		query = query.Where("domain_link = ?", vo.DomainLink)
	}

	if len(vo.DomainType) > 0 {
		query = query.Where("domain_type = ?", vo.DomainType)
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

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateDomainDistribution(vo *dos.DomainDistribution) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"domain_link": vo.DomainLink,
		"domain_type": vo.DomainType,
		"update_by":   vo.UpdateBy,
		"sort":        vo.Sort,
		"min_level":   vo.MinLevel,
		"max_level":   vo.MaxLevel,
	}).Error == nil
}

func DeleteDomainDistribution(vo *dos.DomainDistribution) bool {
	return global.G_DB.Model(&dos.DomainDistribution{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
