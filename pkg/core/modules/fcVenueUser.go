// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcVenueUser(vo *dos.FcVenueUser) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenueUser(page, pageSize int, vo *dos.FcVenueUser) (ret []*dos.FcVenueUser, total int64) {
	query := global.G_DB.Model(&dos.FcVenueUser{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Account) > 0 {
		query = query.Where("account = ?", vo.Account)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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
	var dataSlice []*dos.FcVenueUser
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueUser(vo *dos.FcVenueUser) []*dos.FcVenueUser {
	var data []*dos.FcVenueUser
	query := global.G_DB.Model(&dos.FcVenueUser{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Account) > 0 {
		query = query.Where("account = ?", vo.Account)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

	query.Find(&data)
	return data
}

func FindByKeyFcVenueUserFirst(vo *dos.FcVenueUser) *dos.FcVenueUser {
	var data *dos.FcVenueUser
	query := global.G_DB.Model(&dos.FcVenueUser{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Account) > 0 {
		query = query.Where("account = ?", vo.Account)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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
func UpdateFcVenueUser(vo *dos.FcVenueUser) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"venue_id":   vo.VenueId,
		"venue_code": vo.VenueCode,
		"user_id":    vo.UserId,
		"user_name":  vo.UserName,
		"account":    vo.Account,
		"password":   vo.Password,
		"currency":   vo.Currency,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVenueUser(vo *dos.FcVenueUser) bool {
	return global.G_DB.Model(&dos.FcVenueUser{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
