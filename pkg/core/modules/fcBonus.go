// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcBonus(vo *dos.FcBonus) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcBonus(page, pageSize int, vo *dos.FcBonus) (ret []*dos.FcBonus, total int64) {
	query := global.G_DB.Model(&dos.FcBonus{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
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
	var dataSlice []*dos.FcBonus
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBonus(vo *dos.FcBonus) []*dos.FcBonus {
	var data []*dos.FcBonus
	query := global.G_DB.Model(&dos.FcBonus{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
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

//根据主键Update
func UpdateFcBonus(vo *dos.FcBonus) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":    vo.UserId,
		"user_name":  vo.UserName,
		"amount":     vo.Amount,
		"currency":   vo.Currency,
		"bonus_type": vo.BonusType,
		"remark":     vo.Remark,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcBonus(vo *dos.FcBonus) bool {
	return global.G_DB.Model(&dos.FcBonus{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
