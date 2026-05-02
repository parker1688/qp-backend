package payment

import "strings"

type PayChannelSettingResp struct {
	AmountRange             []float64 `json:"amount_range"`                // 选择金额区间
	InputAmountDisplay      bool      `json:"input_amount_display"`        // 输入金额是否显示
	InputNameDisplay        bool      `json:"input_name_display"`          // 输入存款姓名是否显示
	InputVirtualPayAddress  bool      `json:"input_virtual_pay_address"`   // 输入虚拟币地址是否显示
	InputVirtualPayShow     bool      `json:"input_virtual_pay_show"`      // 虚拟本地币选择
	InputVirtualPayShowList bool      `json:"input_virtual_pay_show_list"` // 虚拟币快速充值(每个账户一个币种地址)
}

// 渠道类型
const (
	PayChannel_Bank               = "bank"               //银行
	PayChannel_FAST               = "fast"               //快捷支付
	PayChannel_Virtual            = "virtual"            //虚拟币
	PayChannel_Virtua_FAST        = "virtual_fast"       //虚拟币快捷支付
	PayChannel_UnionPay_QR        = "union_qr"           //银联扫码
	PayChannel_Number_CNY         = "number_cny"         //数字人民币
	PayChannel_AliPay_QR          = "alipay_qr"          //支付宝扫码
	PayChannel_Phone_Bank_QR      = "phone_bank_qr"      //手机银行扫码
	PayChannel_Phone_Bank_H5      = "phone_bank_h5"      //手机银行H5
	PayChannel_Union_Cloud        = "union_cloud"        //云闪付
	PayChannel_Union_Cloud_Red    = "union_cloud_red"    //云闪付红包
	PayChannel_Bank_To_Bank       = "bank_to_bank"       //银行卡转银行卡
	PayChannel_AliPay_UID_MIN     = "alipay_uid_min"     //支付宝小额UID
	PayChannel_AliPay_UID_MAX     = "alipay_uid_max"     //支付宝大额UID
	PayChannel_AliPay_UID_MAX_MAX = "alipay_uid_max_max" //支付宝超大额UID
	PayChannel_AliPay_NONE        = "alipay_uid_none"    //支付宝固定UID
	PayChannel_AliPay_CODE_QR     = "alipay_code_qr"     //支付宝固码扫码
	PayChannel_AliPay_Sales       = "alipay_sales"       //支付宝售货机
	PayChannel_AliPay_yb          = "alipay_yb"          //支付宝压宝
	PayChannel_AliPay_tb          = "alipay_tb"          //淘宝零钱
	PayChannel_AliPay_shop        = "alipay_shop"        //支付宝旗舰店
	PayChannel_AliPay_fixed       = "alipay_fixed"       //支付宝固额
	PayChannel_AliPay_tm          = "alipay_tm"          //天猫游戏
	PayChannel_AliPay_tm_shop     = "alipay_tm_shop"     //天猫购物
	PayChannel_DY_transfer        = "dy_transfer"        //抖音转账
	PayChannel_Wx_Jd_Game         = "wx_jd_game"         //微信京东游戏
	PayChannel_AliPay_Call        = "alipay_call"        //支付宝话费
)

// GetPayChannelSetting
//
//	@Description: 获取支付渠道配置
func GetPayChannelSetting(channelCode string) PayChannelSettingResp {

	if channelCode == PayChannel_Virtua_FAST {
		return PayChannelSettingResp{
			AmountRange:             []float64{},
			InputVirtualPayShowList: true,
		}
	} else if strings.Contains(channelCode, PayChannel_Virtual) {
		return PayChannelSettingResp{
			AmountRange:            []float64{},
			InputVirtualPayAddress: true,
			InputVirtualPayShow:    true,
		}
	} else if strings.Contains(channelCode, PayChannel_Bank) {
		return PayChannelSettingResp{
			AmountRange:            []float64{},
			InputAmountDisplay:     true,
			InputNameDisplay:       true,
			InputVirtualPayAddress: false,
		}
	} else if strings.Contains(channelCode, PayChannel_FAST) {
		return PayChannelSettingResp{
			AmountRange:            []float64{103, 105, 106, 108, 202, 303},
			InputAmountDisplay:     false,
			InputNameDisplay:       false,
			InputVirtualPayAddress: false,
		}
	}
	return PayChannelSettingResp{
		AmountRange:            []float64{103, 105, 106, 108, 202, 303},
		InputAmountDisplay:     true,
		InputNameDisplay:       false,
		InputVirtualPayAddress: false,
	}
}
