// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"

	"github.com/gin-gonic/gin"
)

func SaveFcOrderWithdrawPaymentOut(vo *dos.FcOrderWithdrawPaymentOut) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcOrderWithdrawPaymentOut(pageQuery response.PageTimeQuery, vo *dos.FcOrderWithdrawPaymentOut, c *gin.Context) (ret []*dos.FcOrderWithdrawPaymentOutResp, total int64) {
	response.NormalizePageTimeQuery(&pageQuery)
	query := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > -1 {
		//query = query.Where("status = ?", vo.Status)
		switch vo.Status {
		case 0: // 未打款
			query = query.Where("withdraw_status in ?",
				[]int{
					enmus.OrderWithdrawStats_ManualNo,
					enmus.OrderWithdrawStats_No,
				})
		case 1: // 已打款
			query = query.Where("withdraw_status in ?",
				[]int{
					enmus.OrderWithdrawStats_ManualYes,
					enmus.OrderWithdrawStats_Yes,
				})
		case 2: // 打款中
			query = query.Where("status in ?", []int{
				enmus.OrderWithdrawPaymentOutStats_Prepare,
				enmus.OrderWithdrawPaymentOutStats_Progress,
			})
		}
	}

	/*if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}*/

	if pageQuery.StartAt != "" {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
	}
	if pageQuery.EndAt != "" {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
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

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.ChannelId) > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentId) > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	dataSlice := []*dos.FcOrderWithdrawPaymentOutEx{}
	query.Order("create_time desc").
		Offset((pageQuery.PageNo - 1) * pageQuery.PageSize).
		Limit(pageQuery.PageSize).
		Preload("OrderWithdraw").
		Find(&dataSlice)

	list := []*dos.FcOrderWithdrawPaymentOutResp{}
	for _, v := range dataSlice {
		respData := dos.FcOrderWithdrawPaymentOutResp{}
		tool.JsonMapper(v, &respData)
		respData.AuditStatus = v.OrderWithdraw.Status

		(&v.OrderWithdraw).Decrypt()

		respData.AccountBankType = v.OrderWithdraw.AccountBankType
		respData.AccountHolder = v.OrderWithdraw.AccountHolder
		respData.AccountNumber = v.OrderWithdraw.AccountNumber
		switch v.OrderWithdraw.OrderType {
		case enmus.ORDER_TYPE_Virtual:
			respData.AccountBankType = v.OrderWithdraw.VirtualCurrencyChain
			respData.AccountNumber = v.OrderWithdraw.VirtualAddress
		case enmus.ORDER_TYPE_Online:
			respData.AccountBankType = "支付宝"
		}

		list = append(list, &respData)
	}

	/*query.Select(`fc_order_withdraw_payment_out.id`,
	`fc_order_withdraw_payment_out.order_sn`,
	`fc_order_withdraw_payment_out.user_id`,
	`fc_order_withdraw_payment_out.user_name`,
	`fc_order_withdraw_payment_out.amount`,
	`fc_order_withdraw_payment_out.apply_amount`,
	`fc_order_withdraw_payment_out.status`,
	`fc_order_withdraw_payment_out.create_time`,
	`fc_order_withdraw_payment_out.create_by`,
	`fc_order_withdraw_payment_out.update_time`,
	`fc_order_withdraw_payment_out.update_by`,
	`fc_order_withdraw_payment_out.merchant_code`,
	`fc_order_withdraw_payment_out.currency`,
	`fc_order_withdraw_payment_out.order_type`,
	`fc_order_withdraw_payment_out.fee_rate`,
	`fc_order_withdraw_payment_out.fee`,
	`fc_order_withdraw_payment_out.channel_id`,
	`fc_order_withdraw_payment_out.channel_code`,
	`fc_order_withdraw_payment_out.payment_id`,
	`fc_order_withdraw_payment_out.payment_code`,
	`fc_order_withdraw_payment_out.third_code`,
	`fc_order_withdraw_payment_out.remark`,
	`fc_order_withdraw_payment_out.deposit_withdraw_sub_amount`,
	`fc_order_withdraw_payment_out.withdraw_amount`,
	`fc_order_withdraw_payment_out.withdraw_status`,
	`fc_order_withdraw.status as audit_status`).
	Joins("left join fc_order_withdraw on fc_order_withdraw_payment_out.order_sn = fc_order_withdraw.order_sn").
	Order("create_time desc").
	Offset((pageQuery.PageNo - 1) * pageQuery.PageSize).
	Limit(pageQuery.PageSize).
	Find(&dataSlice)*/

	return list, count
}

func FindByKeyFcOrderWithdrawPaymentOut(vo *dos.FcOrderWithdrawPaymentOut, c *gin.Context) []*dos.FcOrderWithdrawPaymentOut {
	var data []*dos.FcOrderWithdrawPaymentOut
	query := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.ChannelId) > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentId) > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)

	for i, v := range data {
		switch v.WithdrawStatus { // 将人工提款类型调整为常态（人工提款类型仅用于内部使用）
		case enmus.OrderWithdrawStats_ManualNo:
			data[i].WithdrawStatus = enmus.OrderWithdrawStats_No
		case enmus.OrderWithdrawStats_ManualYes:
			data[i].WithdrawStatus = enmus.OrderWithdrawStats_Yes
		}
	}

	return data
}

func FindByKeyFcOrderWithdrawPaymentOutFirst(vo *dos.FcOrderWithdrawPaymentOut) *dos.FcOrderWithdrawPaymentOut {
	data := &dos.FcOrderWithdrawPaymentOut{}
	query := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.OrderType > 0 {
		query = query.Where("order_type = ?", vo.OrderType)
	}

	if len(vo.ChannelId) > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if len(vo.PaymentId) > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	err := query.Take(data).Error
	if err != nil {
		return nil
	}
	return data
}

// 根据主键Update
func UpdateFcOrderWithdrawPaymentOut(vo *dos.FcOrderWithdrawPaymentOut) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"order_sn":  vo.OrderSn,
		"user_id":   vo.UserId,
		"user_name": vo.UserName,
		"amount":    vo.Amount,
		"status":    vo.Status,
		"create_by": vo.CreateBy,
		"update_by": vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
		"currency":     vo.Currency,
		"order_type":   vo.OrderType,
		"fee_rate":     vo.FeeRate,
		"fee":          vo.Fee,
		"channel_id":   vo.ChannelId,
		"channel_code": vo.ChannelCode,
		"payment_id":   vo.PaymentId,
		"payment_code": vo.PaymentCode,
	}).Error == nil
}

func DeleteFcOrderWithdrawPaymentOut(vo *dos.FcOrderWithdrawPaymentOut) bool {
	return global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// SetFcOrderWithdrawPaymentoutData - 更新汇款申请数据
// @param {string} id
// @param {map[stirng]interface{}} sets
// @returns error
func SetFcOrderWithdrawPaymentoutData(id string, sets map[string]interface{}) error {
	err := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).
		Where("id = ?", id).
		Updates(sets).Error
	if err != nil {
		global.G_LOG.Errorf("[SetFcOrderWithdrawPaymentoutData] Update FcOrderWithdrawPaymentOut failed: id=%s, sets=%+v, err=%s",
			id, sets, err.Error())
	}

	return err
}
