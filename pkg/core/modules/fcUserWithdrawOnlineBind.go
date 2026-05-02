// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserWithdrawOnlineBind(vo *dos.FcUserWithdrawOnlineBind) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserWithdrawOnlineBind(page, pageSize int, vo *dos.FcUserWithdrawOnlineBind, c *gin.Context) (ret []*dos.FcUserWithdrawOnlineBind, total int64) {
	query := global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserWithdrawOnlineBind
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserWithdrawOnlineBind(vo *dos.FcUserWithdrawOnlineBind) []*dos.FcUserWithdrawOnlineBind {
	var data []*dos.FcUserWithdrawOnlineBind
	query := global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserWithdrawOnlineBind2(vo *dos.FcUserWithdrawOnlineBind, c *gin.Context) []*dos.FcUserWithdrawOnlineBind {
	var data []*dos.FcUserWithdrawOnlineBind
	query := global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
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

func FindByKeyFcUserWithdrawOnlineBindFirst(vo *dos.FcUserWithdrawOnlineBind) *dos.FcUserWithdrawOnlineBind {
	var data *dos.FcUserWithdrawOnlineBind
	query := global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.AccountNumber) > 0 {
		query = query.Where("account_number = ?", vo.AccountNumber)
	}

	if len(vo.AccountHolder) > 0 {
		query = query.Where("account_holder = ?", vo.AccountHolder)
	}

	if vo.IsDefault > 0 {
		query = query.Where("is_default = ?", vo.IsDefault)
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

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserWithdrawOnlineBind(vo *dos.FcUserWithdrawOnlineBind) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":        vo.UserId,
		"user_name":      vo.UserName,
		"account_number": vo.AccountNumber,
		"account_holder": vo.AccountHolder,
		"is_default":     vo.IsDefault,
		"sort":           vo.Sort,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"merchant_code":  vo.MerchantCode,
		"channel_name":   vo.ChannelName,
		"channel_code":   vo.ChannelCode,
	}).Error == nil
}

func DeleteFcUserWithdrawOnlineBind(vo *dos.FcUserWithdrawOnlineBind) bool {
	return global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
