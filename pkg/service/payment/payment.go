package payment

import (
	"bootpkg/pkg/core/modules/dos"
)

var (
	channelPay = make(map[string]IPayment)
)

const (
	STYLE_LINKE       = "Link"        //跳转类型
	STYLE_LINKE_OPEN  = "Link_Open"   //跳转类型(新打开网页)
	STYLE_BANK        = "Bank"        //显示银行卡类型
	STYLE_VirtualCoin = "VirtualCoin" //虚拟币类型
	STYLE_Qr          = "Qr"          //扫码
)

func GetRechargeChannel(paymentCode string) IPayment {
	value, ok := channelPay[paymentCode]
	if !ok {
		return nil
	}
	return value
}

// OtherPaymentParam
// @Description: 通道支付额外提交信息
type OtherPaymentParam struct {
	Amount      float64 //存款金额
	BonusRate   float64 `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `         // 优惠比例
	BonusAmount float64 `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" ` // 优惠金额
	FactAmount  float64 `gorm:"column:fact_amount" json:"fact_amount" form:"fact_amount" uri:"fact_amount" `     // 实际存款
	DepositName string  //存款人姓名
	OrderSn     string  //订单号
	ClientIp    string  //IP地址

	CurrencyName  string              // 虚拟币币种名称
	CurrencyChain string              // 虚拟币币种所属链
	Currency      string              // 货币简码
	UserInfo      *dos.FcUserMaterial // 存款用户信息

	MerchantCode string //商户Code 根据不同商户获取不同配置
	ReturnUrl    string //支付成功后返回地址
}

// PaymentChannelResp
// @Description: 支付通道返回信息
type PaymentChannelResp struct {
	Success  bool
	ErrorMsg string
	Data     PaymentChannelDataResp
}

type PaymentChannelDataResp struct {
	Style       string                         `json:"style"`                 // STYLE_LINKE | STYLE_BANK | STYLE_VirtualCoin | Qr
	Link        string                         `json:"link"`                  // 跳转URL
	OrderSn     string                         `json:"order_sn"`              //订单号
	Amount      float64                        `json:"amount"`                //充值金额
	Bank        *PaymentChannelBankResp        `json:"bank,omitempty"`        //银行卡信息
	VirtualCoin *PaymentChannelVirtualCoinResp `json:"virtualCoin,omitempty"` //虚拟币信息
	Qr          *PaymentChannelQrResp          `json:"qr,omitempty"`          //二维码显示
}

type PaymentChannelBankResp struct {
	EntityAccountHolder   string `json:"entity_account_holder" form:"entity_account_holder" uri:"entity_account_holder" `          // 收款人
	EntityAccountBankName string `json:"entity_account_bank_name" form:"entity_account_bank_name" uri:"entity_account_bank_name" ` // 收款人银行名字
	EntityAccountNumber   string `json:"entity_account_number" form:"entity_account_number" uri:"entity_account_number" `          // 收款卡号
	QrImg                 string `json:"qr_img" form:"qr_img" uri:"qr_img" `                                                       // 二维码生成内容
	BankRemark            string `json:"bank_remark" form:"bank_remark" uri:"bank_remark" `                                        // 银行卡存款备注码
	Expiration            int    `json:"expiration" form:"entity_account_number" uri:"entity_account_number"`                      // 过期时间分钟
}

type PaymentChannelVirtualCoinResp struct {
	CurrencyName  string  `json:"currency_name" form:"currency_name" uri:"currency_name" `             // 币种名称
	CurrencyChain string  `json:"currency_chain" form:"currency_chain" uri:"currency_chain" `          // 币种所属链
	ToAddr        string  `json:"to_addr" form:"to_addr" uri:"to_addr" `                               // 接收地址
	ToAddrQrPre   string  `json:"to_addr_qr_pre" form:"to_addr_qr_pre" uri:"to_addr_qr_pre" `          // 接收地址二维码前缀
	Fx            float64 `json:"fx" form:"fx" uri:"fx" `                                              // 汇率
	Num           float64 `json:"num" form:"num" uri:"num" `                                           // 应付数量
	Expiration    int     `json:"expiration" form:"entity_account_number" uri:"entity_account_number"` // 过期时间分钟
}

type PaymentChannelQrResp struct {
	QrImg      string `json:"qr_img" form:"qr_img" uri:"qr_img" `                                  // 二维码生成内容
	ToPayLink  string `json:"to_pay_link" form:"to_pay_link" uri:"to_pay_link" `                   // 去支付跳转地址
	Expiration int    `json:"expiration" form:"entity_account_number" uri:"entity_account_number"` // 过期时间分钟
}

// PaymentVirtualResp
// @Description: 获取通道支持虚拟信息
type PaymentVirtualResp struct {
	Banks       []string `json:"banks"` //银行卡信息
	VirtualCoin struct {
	} `json:"virtualCoin"`
}

// CallBackPaymentParam
// @Description: 回调接收参数
type CallBackPaymentParam struct {
	Raw string //接收参数

	MerchantCode string //商户Code
}

// CallBackPaymentResp
// @Description: 通道回调输出参数
type CallBackPaymentResp struct {
	Code    int    //200 成功
	OrderSn string //订单号

	ReturnRaw string //返回字符串
}

// IPayment
// @Description: 支付通道接口
type IPayment interface {
	//
	// ChannelRecharge
	//  @Description: 通道充值
	//  @param *dos.FcPayment 通道信息
	//  @param OtherPaymentParam 通道额外信息
	//  @return PaymentResp 通道返回信息

	ChannelRecharge(*dos.FcPayment, OtherPaymentParam) PaymentChannelResp

	//
	// ChannelCallBack
	//  @Description: 回调确认
	//  @param CallBackPaymentParam 回调参数
	//  @return CallBackPaymentResp 回调输出参数
	//
	ChannelCallBack(CallBackPaymentParam) CallBackPaymentResp

	//
	// ChannelCallBackGetOrderSn
	//  @Description: 获取订单号
	//  @param CallBackPaymentParam 回调参数
	//  @return CallBackPaymentResp 回调输出参数
	//
	ChannelCallBackGetOrderSn(CallBackPaymentParam) CallBackPaymentResp
}
