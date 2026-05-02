// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserVipRecord(vo *dos.FcUserVipRecord) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserVipRecord(page, pageSize int, vo *dos.FcUserVipRecord, c *gin.Context) (ret []*dos.FcUserVipRecord, total int64) {
	query := global.G_DB.Model(&dos.FcUserVipRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BeforLevel > 0 {
		query = query.Where("befor_level = ?", vo.BeforLevel)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.BeforVip) > 0 {
		query = query.Where("befor_vip = ?", vo.BeforVip)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserVipRecord
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserVipRecord(vo *dos.FcUserVipRecord, c *gin.Context) []*dos.FcUserVipRecord {
	var data []*dos.FcUserVipRecord
	query := global.G_DB.Model(&dos.FcUserVipRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BeforLevel > 0 {
		query = query.Where("befor_level = ?", vo.BeforLevel)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.BeforVip) > 0 {
		query = query.Where("befor_vip = ?", vo.BeforVip)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
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

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserVipRecordFirst(vo *dos.FcUserVipRecord) *dos.FcUserVipRecord {
	var data *dos.FcUserVipRecord
	query := global.G_DB.Model(&dos.FcUserVipRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.BeforLevel > 0 {
		query = query.Where("befor_level = ?", vo.BeforLevel)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.BeforVip) > 0 {
		query = query.Where("befor_vip = ?", vo.BeforVip)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
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
func UpdateFcUserVipRecord(vo *dos.FcUserVipRecord) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":                vo.UserId,
		"user_name":              vo.UserName,
		"befor_level":            vo.BeforLevel,
		"level":                  vo.Level,
		"befor_vip":              vo.BeforVip,
		"vip":                    vo.Vip,
		"total_recharege_amount": vo.TotalRecharegeAmount,
		"total_bet_amount":       vo.TotalBetAmount,
		"create_by":              vo.CreateBy,
		"update_by":              vo.UpdateBy,
		"merchant_code":          vo.MerchantCode,
		"bonus":                  vo.Bonus,
	}).Error == nil
}

func DeleteFcUserVipRecord(vo *dos.FcUserVipRecord) bool {
	return global.G_DB.Model(&dos.FcUserVipRecord{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
