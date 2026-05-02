// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcCustomerOrderType(vo *dos.FcCustomerOrderType) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcCustomerOrderType(page, pageSize int, vo *dos.FcCustomerOrderType) (ret []*dos.FcCustomerOrderType, total int64) {
	query := global.G_DB.Model(&dos.FcCustomerOrderType{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if vo.FlowMultiple > -1 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
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
	var dataSlice []*dos.FcCustomerOrderType
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcCustomerOrderType(vo *dos.FcCustomerOrderType) []*dos.FcCustomerOrderType {
	var data []*dos.FcCustomerOrderType
	query := global.G_DB.Model(&dos.FcCustomerOrderType{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if vo.FlowMultiple > -1 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
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

func FindByKeyFcCustomerOrderTypeFirst(vo *dos.FcCustomerOrderType) *dos.FcCustomerOrderType {
	var data *dos.FcCustomerOrderType
	query := global.G_DB.Model(&dos.FcCustomerOrderType{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if vo.FlowMultiple > -1 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
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

//根据主键Update
func UpdateFcCustomerOrderType(vo *dos.FcCustomerOrderType) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"bonus_type":    vo.BonusType,
		"flow_multiple": vo.FlowMultiple,
		"title":         vo.Title,
		"remark":        vo.Remark,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
	}).Error == nil
}

func DeleteFcCustomerOrderType(vo *dos.FcCustomerOrderType) bool {
	return global.G_DB.Model(&dos.FcCustomerOrderType{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
