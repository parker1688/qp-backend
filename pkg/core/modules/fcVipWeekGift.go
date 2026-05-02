// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcVipWeekGift(vo *dos.FcVipWeekGift) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcVipWeekGift(page, pageSize int, vo *dos.FcVipWeekGift, c *gin.Context) (ret []*dos.FcVipWeekGift, total int64) {
	query := global.G_DB.Model(&dos.FcVipWeekGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Week) > 0 {
		query = query.Where("week = ?", vo.Week)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVipWeekGift
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVipWeekGift(vo *dos.FcVipWeekGift, c *gin.Context) []*dos.FcVipWeekGift {
	var data []*dos.FcVipWeekGift
	query := global.G_DB.Model(&dos.FcVipWeekGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Week) > 0 {
		query = query.Where("week = ?", vo.Week)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVipWeekGiftFirst(vo *dos.FcVipWeekGift) *dos.FcVipWeekGift {
	var data *dos.FcVipWeekGift
	query := global.G_DB.Model(&dos.FcVipWeekGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Week) > 0 {
		query = query.Where("week = ?", vo.Week)
	}

	if len(vo.VipName) > 0 {
		query = query.Where("vip_name = ?", vo.VipName)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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
func UpdateFcVipWeekGift(vo *dos.FcVipWeekGift) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"week":               vo.Week,
		"vip_name":           vo.VipName,
		"level":              vo.Level,
		"merchant_code":      vo.MerchantCode,
		"bonus_amount":       vo.BonusAmount,
		"bonus_amount_issue": vo.BonusAmountIssue,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
	}).Error == nil
}

func DeleteFcVipWeekGift(vo *dos.FcVipWeekGift) bool {
	return global.G_DB.Model(&dos.FcVipWeekGift{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
