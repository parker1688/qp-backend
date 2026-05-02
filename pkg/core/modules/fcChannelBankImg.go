// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcChannelBankImg(vo *dos.FcChannelBankImg) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcChannelBankImg(page, pageSize int, vo *dos.FcChannelBankImg, c *gin.Context) (ret []*dos.FcChannelBankImg, total int64) {
	query := global.G_DB.Model(&dos.FcChannelBankImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if len(vo.IconPath) > 0 {
		query = query.Where("icon_path = ?", vo.IconPath)
	}

	if len(vo.Img) > 0 {
		query = query.Where("img = ?", vo.Img)
	}

	if len(vo.ImgPath) > 0 {
		query = query.Where("img_path = ?", vo.ImgPath)
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

	/*if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}*/

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcChannelBankImg
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcChannelBankImg(vo *dos.FcChannelBankImg, c *gin.Context) []*dos.FcChannelBankImg {
	var data []*dos.FcChannelBankImg
	query := global.G_DB.Model(&dos.FcChannelBankImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if len(vo.IconPath) > 0 {
		query = query.Where("icon_path = ?", vo.IconPath)
	}

	if len(vo.Img) > 0 {
		query = query.Where("img = ?", vo.Img)
	}

	if len(vo.ImgPath) > 0 {
		query = query.Where("img_path = ?", vo.ImgPath)
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

	/*if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}*/

	query.Find(&data)
	return data
}

func FindByKeyFcChannelBankImgFirst(vo *dos.FcChannelBankImg) *dos.FcChannelBankImg {
	var data *dos.FcChannelBankImg
	query := global.G_DB.Model(&dos.FcChannelBankImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if len(vo.IconPath) > 0 {
		query = query.Where("icon_path = ?", vo.IconPath)
	}

	if len(vo.Img) > 0 {
		query = query.Where("img = ?", vo.Img)
	}

	if len(vo.ImgPath) > 0 {
		query = query.Where("img_path = ?", vo.ImgPath)
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
func UpdateFcChannelBankImg(vo *dos.FcChannelBankImg) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"id":            vo.Id,
		"status":        vo.Status,
		"channel_code":  vo.ChannelCode,
		"payment_code":  vo.PaymentCode,
		"channel_name":  vo.ChannelName,
		"payment_name":  vo.PaymentName,
		"icon":          vo.Icon,
		"icon_path":     vo.IconPath,
		"img":           vo.Img,
		"img_path":      vo.ImgPath,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"merchant_code": vo.MerchantCode,
		"sort":          vo.Sort,
	}).Error == nil
}

func DeleteFcChannelBankImg(vo *dos.FcChannelBankImg) bool {
	return global.G_DB.Model(&dos.FcChannelBankImg{}).Delete(vo).Error == nil
}
