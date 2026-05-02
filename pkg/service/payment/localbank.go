package payment

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
)

const (
	LocalBankAmount_KEY_R = "LocalBankAmount_KEY_R:%v"
	//  本地银行卡
	PaymentCode_LOCAL_BANK = "LOCAL_BANK"
)

func init() {
	channelPay[PaymentCode_LOCAL_BANK] = &LocalBank{}
}

// LocalBank
// @Description: 本地银行卡通道 LOCAL_BANK
type LocalBank struct {
}

func (b *LocalBank) ChannelCallBackGetOrderSn(param CallBackPaymentParam) CallBackPaymentResp {
	//TODO implement me
	panic("implement me")
}

func (b *LocalBank) ChannelCallBack(param CallBackPaymentParam) CallBackPaymentResp {
	//TODO implement me
	panic("implement me")
}

func (*LocalBank) ChannelRecharge(payment *dos.FcPayment, other OtherPaymentParam) PaymentChannelResp {
	data := modules.FindByKeyFcBanksDetails(&dos.FcBanksDetails{
		Status: enmus.PAYMENT_STATUS_OK,
	})
	if len(data) == 0 {
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: "无支付银行卡",
		}
	}

	var bank *dos.FcBanksDetails
	for _, v := range data {
		if v.MinLevel <= other.UserInfo.Level && v.MaxLevel >= other.UserInfo.Level && v.MinAmount <= other.Amount && v.MaxAmount >= other.Amount {
			if getLocalBankPaymentChannel(v.Id, v.DayMaxAmount) {
				bank = v
				break
			}
		}
	}
	if bank == nil {
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: "无符合支付银行卡",
		}
	}
	setLocalBankPaymentChannel(bank.Id, other.Amount)
	bank.Decrypt()
	return PaymentChannelResp{
		Success: true,
		Data: PaymentChannelDataResp{
			Style: STYLE_BANK,
			Bank: &PaymentChannelBankResp{
				EntityAccountBankName: bank.EntityAccountBankName,
				EntityAccountHolder:   bank.EntityAccountHolder,
				EntityAccountNumber:   bank.EntityAccountNumber,
				Expiration:            15,
			},
		},
	}
}

// getLocalBankPaymentChannel
//
//	@Description: 银行卡通道每日最大金额
//	@param bankId 银行卡ID
//	@param dayMaxAmount 今日最大金额
//	@return bool true: 可继续存款 / 不可继续存款
func getLocalBankPaymentChannel(bankId string, dayMaxAmount float64) bool {
	value := global.G_REDIS.Get(context.Background(), fmt.Sprintf(LocalBankAmount_KEY_R, bankId))
	valueInt := tool.Int(value)
	if valueInt > 0 && valueInt >= int64(dayMaxAmount) {
		return false
	}
	return true
}

// getLocalBankPaymentChannel
//
//	@Description: 银行卡通道每日最大金额添加
//	@param bankId 银行卡ID
//	@param addAmount 金额
func setLocalBankPaymentChannel(bankId string, addAmount float64) {
	key := fmt.Sprintf(LocalBankAmount_KEY_R, bankId)
	global.G_REDIS.IncrBy(context.Background(), key, int64(addAmount))
	t := tool.TimeTomorrowTime()
	global.G_REDIS.Expire(context.Background(), key, t)
}
