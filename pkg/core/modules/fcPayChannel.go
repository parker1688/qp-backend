// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcPayChannel(vo *dos.FcPayChannel) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcPayChannel(page, pageSize int, vo *dos.FcPayChannel, c *gin.Context) (ret []*dos.FcPayChannelExt, total int64) {
	query := global.G_DB.Model(&dos.FcPayChannel{})

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

	if vo.MinLevel > 0 {
		query = query.Where("min_level >= ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level <= ?", vo.MaxLevel)
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
	var dataSlice []*dos.FcPayChannelExt
	query.Offset((page - 1) * pageSize).Order("sort asc").Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPayChannel(vo *dos.FcPayChannel, c *gin.Context) []*dos.FcPayChannel {
	var data []*dos.FcPayChannel
	query := global.G_DB.Model(&dos.FcPayChannel{})

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

	if vo.MinLevel > 0 {
		query = query.Where("min_level = ?", vo.MinLevel)
	}

	if vo.MaxLevel > 0 {
		query = query.Where("max_level = ?", vo.MaxLevel)
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

func FindByKeyFcPayChannelFirst(vo *dos.FcPayChannel) *dos.FcPayChannel {
	var data *dos.FcPayChannel
	query := global.G_DB.Model(&dos.FcPayChannel{})

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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
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
func UpdateFcPayChannel(vo *dos.FcPayChannel) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"channel_name":              vo.ChannelName,
		"channel_code":              vo.ChannelCode,
		"icon":                      vo.Icon,
		"status":                    vo.Status,
		"min_level":                 vo.MinLevel,
		"max_level":                 vo.MaxLevel,
		"sort":                      vo.Sort,
		"currency":                  vo.Currency,
		"create_by":                 vo.CreateBy,
		"update_by":                 vo.UpdateBy,
		"update_time":               vo.UpdateTime,
		"merchant_code":             vo.MerchantCode,
		"min_amount":                vo.MinAmount,
		"max_amount":                vo.MaxAmount,
		"hot":                       vo.Hot,
		"amount_range":              vo.AmountRange,
		"input_amount_display":      vo.InputAmountDisplay,
		"input_name_display":        vo.InputNameDisplay,
		"input_virtual_pay_address": vo.InputVirtualPayAddress,
		"input_virtual_pay_show":    vo.InputVirtualPayShow,
		"strategy":                  vo.Strategy,
	}).Error == nil
}

func DeleteFcPayChannel(vo *dos.FcPayChannel) bool {
	return global.G_DB.Model(&dos.FcPayChannel{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetFcPayChannelStrategyList - 获取支付渠道限单列表
// @param {string} channelCode
// @param {string} merchantCode
// @returns []dos.FcPayChannelStrategyField
func GetFcPayChannelStrategyList(channelCode string, merchantCode string) []dos.FcPayChannelStrategyField {
	result := []dos.FcPayChannelStrategyField{}

	data := dos.FcPayChannel{}
	err := global.G_DB.Model(&dos.FcPayChannel{}).
		Select("strategy").
		Where("channel_code = ? AND merchant_code = ?", channelCode, merchantCode).
		Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcPayChannelStrategyList] Find fc pay channel failed: channelCode=%s, merchantCode=%s, err=%v",
			channelCode, merchantCode, err.Error())
		return result
	}

	if len(data.Strategy) > 0 {
		err = tool.JsonUnmarshalFromString(data.Strategy, &result)
		if err != nil {
			global.G_LOG.Errorf("[GetFcPayChannelStrategyList] JsonUnmarshal strategy data failed: channelCode=%s, merchantCode=%s, err=%v",
				channelCode, merchantCode, err.Error())
		}
		return result
	}

	return result
}
