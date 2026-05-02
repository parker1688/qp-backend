// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcChannel(vo *dos.FcChannel) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcChannel(page, pageSize int, vo *dos.FcChannel) (ret []*dos.FcChannel, total int64) {
	query := global.G_DB.Model(&dos.FcChannel{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.Type > 0 {
		query = query.Where("type = ?", vo.Type)
	}

	if vo.IsBlockchain > 0 {
		query = query.Where("is_blockchain = ?", vo.IsBlockchain)
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
	var dataSlice []*dos.FcChannel
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcChannel(vo *dos.FcChannel) []*dos.FcChannel {
	var data []*dos.FcChannel
	query := global.G_DB.Model(&dos.FcChannel{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.Type > 0 {
		query = query.Where("type = ?", vo.Type)
	}

	if vo.IsBlockchain > 0 {
		query = query.Where("is_blockchain = ?", vo.IsBlockchain)
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

// 根据主键Update
func UpdateFcChannel(vo *dos.FcChannel) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"channel_name":  vo.ChannelName,
		"channel_code":  vo.ChannelCode,
		"currency":      vo.Currency,
		"remark":        vo.Remark,
		"type":          vo.Type,
		"is_blockchain": vo.IsBlockchain,
		"min_amount":    vo.MinAmount,
		"max_amount":    vo.MaxAmount,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcChannel(vo *dos.FcChannel) bool {
	return global.G_DB.Model(&dos.FcChannel{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
