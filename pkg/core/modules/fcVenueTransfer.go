// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcVenueTransfer(vo *dos.FcVenueTransfer) (bool, ecode.Code) {
	if vo.CreateTime.Timer().IsZero() {
		vo.CreateTime = automaticType.Now()
	}
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenueTransfer(page, pageSize int, vo *dos.FcVenueTransfer, pageQuery response.PageTimeQuery, c *gin.Context) (ret []*dos.FcVenueTransfer, total int64) {
	query := global.G_DB.Model(&dos.FcVenueTransfer{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.VenueLine > 0 {
		query = query.Where("venue_line = ?", vo.VenueLine)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("(create_time <= ? OR create_time IS NULL)", vo.CreateTime)
	}

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("(create_time >= ? OR create_time IS NULL)", pageQuery.StartAt)
	}
	if len(pageQuery.EndAt) > 0 {
		query = query.Where("(create_time <= ? OR create_time IS NULL)", pageQuery.EndAt)
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

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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
	var dataSlice []*dos.FcVenueTransfer
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueTransfer(vo *dos.FcVenueTransfer, c *gin.Context) []*dos.FcVenueTransfer {
	var data []*dos.FcVenueTransfer
	query := global.G_DB.Model(&dos.FcVenueTransfer{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.VenueLine > 0 {
		query = query.Where("venue_line = ?", vo.VenueLine)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
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

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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

func FindByKeyFcVenueTransferFirst(vo *dos.FcVenueTransfer) *dos.FcVenueTransfer {
	var data *dos.FcVenueTransfer
	query := global.G_DB.Model(&dos.FcVenueTransfer{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.VenueLine > 0 {
		query = query.Where("venue_line = ?", vo.VenueLine)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OptType > 0 {
		query = query.Where("opt_type = ?", vo.OptType)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVenueTransfer(vo *dos.FcVenueTransfer) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"order_sn":   vo.OrderSn,
		"venue_code": vo.VenueCode,
		"venue_line": vo.VenueLine,
		"user_name":  vo.UserName,
		"currency":   vo.Currency,
		"amount":     vo.Amount,
		"opt_type":   vo.OptType,
		"ip":         vo.Ip,
		"update_by":  vo.UpdateBy,
		"status":     vo.Status,
	}).Error == nil
}

func DeleteFcVenueTransfer(vo *dos.FcVenueTransfer) bool {
	return global.G_DB.Model(&dos.FcVenueTransfer{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
