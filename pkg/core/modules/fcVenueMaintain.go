// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVenueMaintain(vo *dos.FcVenueMaintain) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenueMaintain(page, pageSize int, vo *dos.FcVenueMaintain) (ret []*dos.FcVenueMaintain, total int64) {
	query := global.G_DB.Model(&dos.FcVenueMaintain{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if !vo.MaintainStart.Timer().IsZero() {
		query = query.Where("maintain_start = ?", vo.MaintainStart)
	}

	if !vo.MaintainEnd.Timer().IsZero() {
		query = query.Where("maintain_end = ?", vo.MaintainEnd)
	}

	if len(vo.CilentType) > 0 {
		query = query.Where("cilent_type = ?", vo.CilentType)
	}

	if vo.AllowTransfer > 0 {
		query = query.Where("allow_transfer = ?", vo.AllowTransfer)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVenueMaintain
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueMaintain(vo *dos.FcVenueMaintain) []*dos.FcVenueMaintain {
	var data []*dos.FcVenueMaintain
	query := global.G_DB.Model(&dos.FcVenueMaintain{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if !vo.MaintainStart.Timer().IsZero() {
		query = query.Where("maintain_start = ?", vo.MaintainStart)
	}

	if !vo.MaintainEnd.Timer().IsZero() {
		query = query.Where("maintain_end = ?", vo.MaintainEnd)
	}

	if len(vo.CilentType) > 0 {
		query = query.Where("cilent_type = ?", vo.CilentType)
	}

	if vo.AllowTransfer > 0 {
		query = query.Where("allow_transfer = ?", vo.AllowTransfer)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVenueMaintain(vo *dos.FcVenueMaintain) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"venue_id":       vo.VenueId,
		"venue_code":     vo.VenueCode,
		"maintain_start": vo.MaintainStart,
		"maintain_end":   vo.MaintainEnd,
		"cilent_type":    vo.CilentType,
		"allow_transfer": vo.AllowTransfer,
		"remark":         vo.Remark,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"merchant_code":  vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVenueMaintain(vo *dos.FcVenueMaintain) bool {
	return global.G_DB.Model(&dos.FcVenueMaintain{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
