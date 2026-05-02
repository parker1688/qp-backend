// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcOrderPromotion(vo *dos.FcOrderPromotion) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcOrderPromotion(page, pageSize int, vo *dos.FcOrderPromotion) (ret []*dos.FcOrderPromotion, total int64) {
	query := global.G_DB.Model(&dos.FcOrderPromotion{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.ApplyType > 0 {
		query = query.Where("apply_type = ?", vo.ApplyType)
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.TurnOver > 0 {
		query = query.Where("turn_over = ?", vo.TurnOver)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Order("create_time desc").Count(&count)
	var dataSlice []*dos.FcOrderPromotion
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcOrderPromotion(vo *dos.FcOrderPromotion) []*dos.FcOrderPromotion {
	var data []*dos.FcOrderPromotion
	query := global.G_DB.Model(&dos.FcOrderPromotion{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.ApplyType > 0 {
		query = query.Where("apply_type = ?", vo.ApplyType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.TurnOver > 0 {
		query = query.Where("turn_over = ?", vo.TurnOver)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcOrderPromotionFirst(vo *dos.FcOrderPromotion) *dos.FcOrderPromotion {
	var data *dos.FcOrderPromotion
	query := global.G_DB.Model(&dos.FcOrderPromotion{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.ApplyType > 0 {
		query = query.Where("apply_type = ?", vo.ApplyType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if vo.TurnOver > 0 {
		query = query.Where("turn_over = ?", vo.TurnOver)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcOrderPromotion(vo *dos.FcOrderPromotion) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"order_sn":     vo.OrderSn,
		"apply_amount": vo.ApplyAmount,
		"apple_rate":   vo.AppleRate,
		"apply_type":   vo.ApplyType,
		"amount":       vo.Amount,
		"status":       vo.Status,
		"update_by":    vo.UpdateBy,
		"user_name":    vo.UserName,
		"user_id":      vo.UserId,
		"turn_over":    vo.TurnOver,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcOrderPromotion(vo *dos.FcOrderPromotion) bool {
	return global.G_DB.Model(&dos.FcOrderPromotion{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
