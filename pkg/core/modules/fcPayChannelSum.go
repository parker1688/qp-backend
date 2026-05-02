// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcPayChannelSum(vo *dos.FcPayChannelSum) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcPayChannelSum(page, pageSize int, vo *dos.FcPayChannelSum) (ret []*dos.FcPayChannelSum, total int64) {
	query := global.G_DB.Model(&dos.FcPayChannelSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if vo.ChannelType > 0 {
		query = query.Where("channel_type = ?", vo.ChannelType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcPayChannelSum
	query.Offset((page - 1) * pageSize).Order("sort asc").Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPayChannelSum(vo *dos.FcPayChannelSum) []*dos.FcPayChannelSum {
	var data []*dos.FcPayChannelSum
	query := global.G_DB.Model(&dos.FcPayChannelSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if vo.ChannelType > 0 {
		query = query.Where("channel_type = ?", vo.ChannelType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcPayChannelSumFirst(vo *dos.FcPayChannelSum) *dos.FcPayChannelSum {
	var data *dos.FcPayChannelSum
	query := global.G_DB.Model(&dos.FcPayChannelSum{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if vo.ChannelType > 0 {
		query = query.Where("channel_type = ?", vo.ChannelType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
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

	if len(vo.AmountRange) > 0 {
		query = query.Where("amount_range = ?", vo.AmountRange)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcPayChannelSum(vo *dos.FcPayChannelSum) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"channel_name":              vo.ChannelName,
		"channel_code":              vo.ChannelCode,
		"icon":                      vo.Icon,
		"channel_type":              vo.ChannelType,
		"status":                    vo.Status,
		"sort":                      vo.Sort,
		"hot":                       vo.Hot,
		"currency":                  vo.Currency,
		"create_by":                 vo.CreateBy,
		"update_by":                 vo.UpdateBy,
		"update_time":               vo.UpdateTime,
		"min_amount":                vo.MinAmount,
		"max_amount":                vo.MaxAmount,
		"amount_range":              vo.AmountRange,
		"min_level":                 vo.MinLevel,
		"max_level":                 vo.MaxLevel,
		"input_amount_display":      vo.InputAmountDisplay,
		"input_name_display":        vo.InputNameDisplay,
		"input_virtual_pay_address": vo.InputVirtualPayAddress,
		"input_virtual_pay_show":    vo.InputVirtualPayShow,
	}).Error == nil
}

func DeleteFcPayChannelSum(vo *dos.FcPayChannelSum) bool {
	return global.G_DB.Model(&dos.FcPayChannelSum{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
