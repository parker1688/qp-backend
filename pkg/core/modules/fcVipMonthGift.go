// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcVipMonthGift(vo *dos.FcVipMonthGift) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcVipMonthGift(page, pageSize int, vo *dos.FcVipMonthGift, c *gin.Context) (ret []*dos.FcVipMonthGift, total int64) {
	query := global.G_DB.Model(&dos.FcVipMonthGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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
	var dataSlice []*dos.FcVipMonthGift
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVipMonthGift(vo *dos.FcVipMonthGift, c *gin.Context) []*dos.FcVipMonthGift {
	var data []*dos.FcVipMonthGift
	query := global.G_DB.Model(&dos.FcVipMonthGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVipMonthGiftFirst(vo *dos.FcVipMonthGift) *dos.FcVipMonthGift {
	var data *dos.FcVipMonthGift
	query := global.G_DB.Model(&dos.FcVipMonthGift{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Month) > 0 {
		query = query.Where("month = ?", vo.Month)
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
func UpdateFcVipMonthGift(vo *dos.FcVipMonthGift) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"month":              vo.Month,
		"vip_name":           vo.VipName,
		"level":              vo.Level,
		"merchant_code":      vo.MerchantCode,
		"bonus_amount":       vo.BonusAmount,
		"bonus_amount_issue": vo.BonusAmountIssue,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
	}).Error == nil
}

func DeleteFcVipMonthGift(vo *dos.FcVipMonthGift) bool {
	return global.G_DB.Model(&dos.FcVipMonthGift{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
