package paymentOut

import "bootpkg/common/tool"

var (
	channelPayOut = make(map[string]IPaymentOut)
)

// GetPaymentOutChannel
//
//	@Description:
//	@param paymentOutCode 代付通道编码
//	@return IPaymentOut
func GetPaymentOutChannel(paymentOutCode string) IPaymentOut {
	value, ok := channelPayOut[paymentOutCode]
	if !ok {
		return nil
	}
	return value
}

// OtherPaymentOutParam
// @Description: 代付通道额外提交信息
type OtherPaymentOutParam struct {
	Amount              float64 //代付金额
	WithdrawName        string  //提款银行卡所属姓名
	WithdrawType        int     // 1 银行  2 虚拟币
	WithdrawAccount     string  //提款账号
	WithdrawBankType    string  //银行名称
	WithdrawBankCode    string  //银行编码
	WithdrawBankAddress string  //银行开户(Cert)

	MerchantCode string //商户Code

	DepositName    string //存款人姓名
	OrderSn        string //订单号
	UserId         string
	VirtualAddress string
	PaymentCode    string
	ThirdCode      string
}

// PreOtherPaymentOutParam
// @Description: 预检测信息
type PreOtherPaymentOutParam struct {
	WithdrawBankType string //银行名称
	WithdrawBankCode string //银行编码
}

// PaymentChannelResp
// @Description: 代付通道返回信息
type PaymentChannelOutResp struct {
	Code            int //200 成功 404 订单不存在 500失败  -200: 不支持
	ErrorMsg        string
	ThirdPayOrderSn string //三方订单号
}

// PrePaymentChannelOutResp
// @Description: 预提交检测
type PrePaymentChannelOutResp struct {
	Code     int //200 成功 404 订单不存在 500失败  -200: 不支持
	ErrorMsg string
}

// PaymentVirtualResp
// @Description: 获取通道支持虚拟信息
type PaymentVirtualResp struct {
	Banks       []string `json:"banks"` //银行卡信息
	VirtualCoin struct {
	} `json:"virtualCoin"`
}

// CallBackPaymentOutParam
// @Description: 回调接收参数
type CallBackPaymentOutParam struct {
	Raw string //接收参数

	MerchantCode string //商户Code
}

// CallBackPaymentResp
// @Description: 通道回调输出参数
type CallBackPaymentOutResp struct {
	Code    int    //状态  200 成功 201 处理中 404 订单不存在 500失败 0 异常不处理
	OrderSn string //订单号

	ReturnRaw string //返回字符串
	ErrorMsg  string //代付失败描述
}

// FindOrderPaymentOutResp
// @Description: 查询订单输出
type FindOrderPaymentOutResp struct {
	Code int //200 成功 404 订单不存在 500失败  -200: 不支持

	ReturnRaw string //返回字符串
	ErrorMsg  string //代付失败描述
}

// IPayment
// @Description: 支付通道接口
type IPaymentOut interface {
	//
	// PreWithdraw
	//  @Description: 预提交检测
	//  @param PreOtherPaymentOutParam 通道提交信息
	//  @return PrePaymentChannelOutResp

	PreWithdraw(PreOtherPaymentOutParam) PrePaymentChannelOutResp

	//
	// ChannelWithdraw
	//  @Description: 通道代付
	//  @param OtherPaymentParam 通道提交信息
	//  @return PaymentResp 通道返回信息

	ChannelWithdraw(OtherPaymentOutParam) PaymentChannelOutResp

	//
	// ChannelCallBack
	//  @Description: 回调确认
	//  @param CallBackPaymentParam 回调参数
	//  @return CallBackPaymentResp 回调输出参数
	//
	ChannelCallBack(CallBackPaymentOutParam) CallBackPaymentOutResp

	//
	// FindOrderStatus
	//  @Description: 查询订单信息
	//  @param merchantOrderId 回调参数
	//  @return FindOrderPaymentOutResp 输出参数
	//
	FindOrderStatus(merchantOrderId string) FindOrderPaymentOutResp

	//
	// ChannelCallBackGetOrderSn
	//  @Description: 获取订单号
	//  @param CallBackPaymentParam 回调参数
	//  @return CallBackPaymentResp 回调输出参数
	//
	ChannelCallBackGetOrderSn(CallBackPaymentOutParam) CallBackPaymentOutResp
}

// GetOrderExtraData - 获取订单额外数据
// @param {string} extra
// @returns string, string
func GetOrderExtraData(extra string) (string, string) {
	var m map[string]interface{}
	tool.JsonUnmarshalFromString(extra, &m)
	var uid string
	if v, ok := m["uid"]; ok {
		uid = v.(string)
	}

	var oid string
	if v, ok := m["oid"]; ok {
		oid = v.(string)
	}

	return uid, oid
}
