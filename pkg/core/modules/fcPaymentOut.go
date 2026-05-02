// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"

	"github.com/gin-gonic/gin"
)

func SaveFcPaymentOut(vo *dos.FcPaymentOut) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcPaymentOut(page, pageSize int, vo *dos.FcPaymentOut, c *gin.Context) (ret []*dos.FcPaymentOut, total int64) {
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
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

	/*if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/
	if len(vo.ThirdCode) > 0 {
		query = query.Where("third_code = ?", vo.ThirdCode)
	}

	/*if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}*/

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcPaymentOut
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPaymentOut(vo *dos.FcPaymentOut, c *gin.Context) []*dos.FcPaymentOut {
	var data []*dos.FcPaymentOut
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
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

	/*if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/
	if len(vo.ThirdCode) > 0 {
		query = query.Where("third_code = ?", vo.ThirdCode)
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

func FindByKeyFcPaymentOutFirst(vo *dos.FcPaymentOut) *dos.FcPaymentOut {
	data := &dos.FcPaymentOut{}
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.PaymentName) > 0 {
		query = query.Where("payment_name = ?", vo.PaymentName)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(vo.ChannelName) > 0 {
		query = query.Where("channel_name = ?", vo.ChannelName)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
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
	if len(vo.ThirdCode) > 0 {
		query = query.Where("third_code = ?", vo.ThirdCode)
	}
	err := query.Take(data).Error
	if err != nil {
		return nil
	}
	return data
}

// 根据主键Update
func UpdateFcPaymentOut(vo *dos.FcPaymentOut) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"payment_name":   vo.PaymentName,
		"payment_code":   vo.PaymentCode,
		"channel_name":   vo.ChannelName,
		"channel_code":   vo.ChannelCode,
		"status":         vo.Status,
		"min_level":      vo.MinLevel,
		"max_level":      vo.MaxLevel,
		"min_amount":     vo.MinAmount,
		"max_amount":     vo.MaxAmount,
		"day_max_amount": vo.DayMaxAmount,
		"sort":           vo.Sort,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"merchant_code":  vo.MerchantCode,
		"fee_rate":       vo.FeeRate,
		"icon":           vo.Icon,
		"third_code":     vo.ThirdCode,
	}).Error == nil
}

func DeleteFcPaymentOut(vo *dos.FcPaymentOut) bool {
	return global.G_DB.Model(&dos.FcPaymentOut{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetFcPaymentOutData - 获取支付通道数据
// @param {string} channelCode
// @param {string} thirdCode
// @returns dos.FcPaymentOut, error
func GetFcPaymentOutData(channelCode string, thirdCode string) (dos.FcPaymentOut, error) {
	var err error
	data := dos.FcPaymentOut{}
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	switch channelCode {
	case enmus.Another_Bank, enmus.Another_AliPay: // 随机一条数据
		err = query.Where("channel_code = ? AND status = 1", channelCode).
			Order("RAND()").Limit(1).Find(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[GetFcPaymentOutData] Find payment out data failed: channelCode=%s, err=%s",
				channelCode, err.Error())
			return data, err
		}
	case enmus.Another_Virtual: // 指定三方code
		err = query.Where("channel_code = ? AND third_code = ? AND status = 1", channelCode, thirdCode).
			Find(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[GetFcPaymentOutData] Find payment out data failed: channelCode=%s, thirdCode=%s, err=%s",
				channelCode, thirdCode, err.Error())
			return data, err
		}
	}

	return data, nil
}

// GetFcPaymentOutThirdCode - 获取支付通道三方码
// @param {string} paymentCode （是唯一的）
// @returns string
func GetFcPaymentOutThirdCode(paymentCode string) string {
	data := dos.FcPaymentOut{}
	err := global.G_DB.Model(&dos.FcPaymentOut{}).Select("third_code").
		Where("payment_code = ?", paymentCode).Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcPaymentOutThirdCode] Find payment out data failed: paymentCode=%s, err=%s", paymentCode, err.Error())
		return ""
	}

	return data.ThirdCode
}
