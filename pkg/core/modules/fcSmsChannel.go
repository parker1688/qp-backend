// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSmsChannel(vo *dos.FcSmsChannel) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcSmsChannel(page, pageSize int, vo *dos.FcSmsChannel) (ret []*dos.FcSmsChannel, total int64) {
	query := global.G_DB.Model(&dos.FcSmsChannel{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SmsName) > 0 {
		query = query.Where("sms_name = ?", vo.SmsName)
	}

	if len(vo.SmsCode) > 0 {
		query = query.Where("sms_code = ?", vo.SmsCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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
	var dataSlice []*dos.FcSmsChannel
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSmsChannel(vo *dos.FcSmsChannel) []*dos.FcSmsChannel {
	var data []*dos.FcSmsChannel
	query := global.G_DB.Model(&dos.FcSmsChannel{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SmsName) > 0 {
		query = query.Where("sms_name = ?", vo.SmsName)
	}

	if len(vo.SmsCode) > 0 {
		query = query.Where("sms_code = ?", vo.SmsCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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

func FindByKeyFcSmsChannelFirst(vo *dos.FcSmsChannel) *dos.FcSmsChannel {
	var data *dos.FcSmsChannel
	query := global.G_DB.Model(&dos.FcSmsChannel{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.SmsName) > 0 {
		query = query.Where("sms_name = ?", vo.SmsName)
	}

	if len(vo.SmsCode) > 0 {
		query = query.Where("sms_code = ?", vo.SmsCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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
func UpdateFcSmsChannel(vo *dos.FcSmsChannel) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"id":            vo.Id,
		"sms_name":      vo.SmsName,
		"sms_code":      vo.SmsCode,
		"status":        vo.Status,
		"min_level":     vo.MinLevel,
		"max_level":     vo.MaxLevel,
		"sort":          vo.Sort,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"merchant_code": vo.MerchantCode,
		"fee_rate":      vo.FeeRate,
	}).Error == nil
}

func DeleteFcSmsChannel(vo *dos.FcSmsChannel) bool {
	return global.G_DB.Model(&dos.FcSmsChannel{}).Delete(vo).Error == nil
}
