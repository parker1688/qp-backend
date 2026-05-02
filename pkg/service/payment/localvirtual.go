package payment

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

const (
	LocalVirtual_RAND_KEY = "LocalVirtual_RAND:%v_%v"
	// 本地虚拟币
	PaymentCode_LOCAL_VIRTUAL = "LOCAL_VIRTUAL"
)

func init() {
	channelPay[PaymentCode_LOCAL_VIRTUAL] = &LocalVirtual{}
}

// LocalVirtual
// @Description: 本地虚拟币通道 LOCAL_VIRTUAL
type LocalVirtual struct {
}

func (b *LocalVirtual) ChannelCallBackGetOrderSn(param CallBackPaymentParam) CallBackPaymentResp {
	//TODO implement me
	panic("implement me")
}

func (*LocalVirtual) ChannelRecharge(payment *dos.FcPayment, other OtherPaymentParam) PaymentChannelResp {
	data := modules.FindByKeyFcVirtualCurrencyDetails(&dos.FcVirtualCurrencyDetails{
		Status:        enmus.PAYMENT_STATUS_OK,
		CurrencyChain: other.CurrencyChain,
		CurrencyName:  other.CurrencyName,
	})
	if len(data) == 0 {
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: "无支付虚拟币地址",
		}
	}

	next := global.G_REDIS.Incr(context.Background(), fmt.Sprintf(LocalVirtual_RAND_KEY, other.CurrencyChain, other.CurrencyName)).Val()
	nextIndex := next % int64(len(data))
	detail := data[nextIndex]
	fx := modules.FindByKeyFcVirtualCurrencyFxFirst(&dos.FcVirtualCurrencyFx{
		OptType:       1,
		CurrencyChain: other.CurrencyChain,
		CurrencyName:  other.CurrencyName,
		CurrencyCode:  other.Currency,
	})
	if fx.FxAmount == 0 {
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: "无有效汇率",
		}
	}
	//向下保留2位
	numD := decimal.NewFromFloat(other.Amount).Div(decimal.NewFromFloat(fx.FxAmount))
	num := numD.Truncate(2).InexactFloat64()
	return PaymentChannelResp{
		Success: true,
		Data: PaymentChannelDataResp{
			Style: STYLE_VirtualCoin,
			VirtualCoin: &PaymentChannelVirtualCoinResp{
				CurrencyName:  detail.CurrencyName,
				CurrencyChain: detail.CurrencyChain,
				ToAddr:        detail.ToAddr,
				ToAddrQrPre:   detail.ToAddrQrPre,
				Fx:            fx.FxAmount,
				Num:           num,
				Expiration:    30,
			},
		},
	}
}

func (b *LocalVirtual) ChannelCallBack(param CallBackPaymentParam) CallBackPaymentResp {
	//TODO implement me
	panic("implement me")
}
