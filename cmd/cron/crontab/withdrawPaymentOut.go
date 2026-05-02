package crontab

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/paymentOut"
	"time"
)

func init() {
	cronFunc = append(cronFunc, cronTabEvery{
		spec: "@every 30s",
		cmd:  WithdrawPaymentOut,
	})
}

// WithdrawPaymentOut
//
//	@Description: 提款代付
func WithdrawPaymentOut() {
	// 订单排队节点时间
	preNow := time.Now().Unix() - paymentOut.GetDictConfigPaymentOutWaitDuration()

	// 获取符合条件订单列表
	orders := []*dos.FcOrderWithdrawPaymentOut{}
	global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).
		Where("status = ? AND withdraw_status = ? AND create_time <= ?",
			enmus.OrderWithdrawPaymentOutStats_Prepare,
			enmus.OrderWithdrawStats_No,
			time.Unix(preNow, 0)).
		Find(&orders)

	// t := time.Now().Add(-48 * time.Hour).Format(tool.TimeLayout)

	//var ms []*dos.FcOrderWithdrawPaymentOut
	// global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where(" status = ? and create_time >= ?", 0, t).Find(&ms)
	// global.G_LOG.Infof("有%d订单需要代付", len(ms))

	// 下发第三方
	for _, v := range orders {
		var AccountBankType, AccountHolder, AccountNumber, BankAddress, AccountBankCode, VirtualAddress string

		w := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{
			OrderSn: v.OrderSn,
		})

		if len(w.Id) == 0 {
			global.G_LOG.Errorf("[WithdrawPaymentOut] Find order withdraw data failed: orderSn=%s", v.OrderSn)
			continue
		}
		w.Decrypt()

		AccountBankType = tool.StringReplaceAll(w.AccountBankType, "", true, " ")
		AccountHolder = tool.StringReplaceAll(w.AccountHolder, "", true, " ")
		AccountNumber = tool.StringReplaceAll(w.AccountNumber, "", true, " ")
		BankAddress = tool.StringReplaceAll(w.BankAddress, "", true, " ")

		//global.G_LOG.Infof("WithdrawPaymentOut:USER: %v  old:%v new: %v", w.UserId, w.AccountNumber, AccountNumber)
		AccountBankCode = w.AccountBankCode
		VirtualAddress = w.VirtualAddress

		//如果接其他三方  就应该用渠道code
		m := paymentOut.GetPaymentOutChannel(v.ChannelCode)
		if m == nil {
			global.G_LOG.Errorf("[WithdrawPaymentOut] Can't find payment out channel: orderSn=%s, channelCode=%s", v.OrderSn, v.ChannelCode)
			continue
		}

		r := m.ChannelWithdraw(paymentOut.OtherPaymentOutParam{
			Amount:              v.Amount,
			OrderSn:             v.OrderSn,
			WithdrawType:        1,
			WithdrawBankType:    AccountBankType,
			WithdrawName:        AccountHolder,
			WithdrawAccount:     AccountNumber,
			WithdrawBankCode:    AccountBankCode,
			WithdrawBankAddress: BankAddress,
			UserId:              v.UserId,
			MerchantCode:        v.MerchantCode,
			VirtualAddress:      VirtualAddress,
			PaymentCode:         v.PaymentCode,
			ThirdCode:           v.ThirdCode,
		})

		// global.G_LOG.Infof("提款单号 OrderSn  %s申请代付 r.Code:%d", v.OrderSn, r.Code)
		switch r.Code {
		case 200:
			//修改为代付中
			/*err := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where(" status = ? and id = ?", 0, v.Id).Update("status", 1).Error
			if err != nil {
				global.G_LOG.Errorf("提款单号 OrderSn  %s申请代付 r.Code:%d update status err:%s", v.OrderSn, r.Code, err.Error())
			}*/
			eRow := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).
				Where("id = ? AND status = ?", v.Id, enmus.OrderWithdrawPaymentOutStats_Prepare).
				Update("status", enmus.OrderWithdrawPaymentOutStats_Progress)
			if eRow.Error != nil {
				global.G_LOG.Errorf("[WithdrawPaymentOut] Update order withdraw payment out failed: id=%s, orderSn=%s, err=%s",
					v.Id, v.OrderSn, eRow.Error.Error())
			} else if eRow.RowsAffected != 1 {
				global.G_LOG.Errorf("[WithdrawPaymentOut] Update order withdraw payment out conflict: id=%s, orderSn=%s",
					v.Id, v.OrderSn)
			}
		case 500: //代付失败
			paymentOut.OrderWithdrawAnotherPayFailRemark(v.Id, r.ErrorMsg, enmus.OrderWithdrawPaymentOutStats_Prepare, "system") //未提交的状态
		}
	}
}
