package vo

type OrderWithdrawAnotherPayReq struct {
	Id string `gorm:"column:id" json:"id" form:"id" uri:"id" ` //提款订单ID
	//ChannelCode string `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` //选择代付渠道
	//PaymentCode string `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` //选择代付通道
}
