package vo

import "bootpkg/common/expands/automaticType"

type OrderDepositInfoReq struct {
	Currency string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
	StartAt  string `json:"startAt"`                                  //开始时间
	EndAt    string `json:"endAt"`                                    //结束时间
	PageSize int    `json:"pageSize"`                                 //当前页大小
	Current  int    `json:"current"`                                  //当前页码
	Status   int    `json:"status"`                                   //状态
}

type OrderDepositInfoResp struct {
	Id          string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	UserId      string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName    string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	OrderSn     string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                          // 订单编号
	Amount      float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 存款金额
	Status      int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0:待处理  2:拒绝 3:已通过
	CreateTime  automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	PayTime     automaticType.Time `gorm:"column:pay_time;default:null" json:"pay_time" form:"pay_time" uri:"pay_time" `             // 支付时间
	Currency    string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种
	ChannelCode string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
	PaymentCode string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `          // 通道code
	OrderType   int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `                  // 订单类型 1 银行卡 2 三方通道 3 虚拟币
}
