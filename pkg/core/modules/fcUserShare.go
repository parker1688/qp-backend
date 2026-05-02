// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcUserShare(vo *dos.FcUserShare) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserShare(page, pageSize int, vo *dos.FcUserShare) (ret []*dos.FcUserShare, total int64) {
	query := global.G_DB.Model(&dos.FcUserShare{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.ShareLink) > 0 {
		query = query.Where("share_link = ?", vo.ShareLink)
	}

	if len(vo.ShareCode) > 0 {
		query = query.Where("share_code = ?", vo.ShareCode)
	}

	if vo.Quantity > 0 {
		query = query.Where("quantity = ?", vo.Quantity)
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
	var dataSlice []*dos.FcUserShare
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserShare(vo *dos.FcUserShare) []*dos.FcUserShare {
	var data []*dos.FcUserShare
	query := global.G_DB.Model(&dos.FcUserShare{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.ShareLink) > 0 {
		query = query.Where("share_link = ?", vo.ShareLink)
	}

	if len(vo.ShareCode) > 0 {
		query = query.Where("share_code = ?", vo.ShareCode)
	}

	if vo.Quantity > 0 {
		query = query.Where("quantity = ?", vo.Quantity)
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
func UpdateFcUserShare(vo *dos.FcUserShare) bool {
	return global.G_DB.Model(vo).Where(`user_id = ?`, vo.UserId).Updates(map[string]interface{}{
		"user_name":  vo.UserName,
		"share_link": vo.ShareLink,
		"share_code": vo.ShareCode,
		"quantity":   vo.Quantity,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserShare(vo *dos.FcUserShare) bool {
	return global.G_DB.Model(&dos.FcUserShare{}).Where("user_id = ?", vo.UserId).Delete(vo).Error == nil
}
