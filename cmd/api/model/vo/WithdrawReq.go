package vo

import "bootpkg/common/expands/automaticType"

type WithdrawReq struct {
	Amount       float64 `json:"amount" form:"amount" uri:"amount" `                       // 取款金额
	WithdrawType int     `json:"withdraw_type" form:"withdraw_type" uri:"withdraw_type" `  // 1 银行卡 3 虚拟币 4 支付宝微信提款
	Currency     string  `gorm:"currency" json:"currency" form:"currency" uri:"currency" ` // 币种简码
	ActivityId   int     `json:"activity_id" form:"activity_id" uri:"activity_id" `        // 参与活动ID
	ChannelCode  string  `json:"channel_code" form:"channel_code" uri:"channel_code" `
	VirtualId    string  `json:"virtual_id" form:"virtual_id" uri:"virtual_id"` //虚拟币绑定ID
	BankId       string  `json:"bank_id" form:"bank_id" uri:"bank_id"`          //银行卡绑定ID
	OnlineId     string  `json:"online_id" form:"online_id" uri:"online_id"`    //在线支付绑定ID
	VeryCode     string  `json:"veryCode" form:"veryCode" uri:"veryCode"`       //短信验证码
	//WalletPassword string  ` json:"wallet_password" form:"wallet_password" uri:"wallet_password" ` // 钱包密码
}

type WithdrawListReq struct {
	Currency  string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
	StartTime string `json:"start_time"`                               //开始时间
	EndTime   string `json:"end_time"`                                 //结束时间
	PageSize  int    `json:"page_size"`                                //当前页大小
	PageIndex int    `json:"page_index"`                               //当前页码
}

type WithdrawListResp struct {
	OrderSn        string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                             // 订单编号
	Amount         float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                     // 提款金额
	Status         int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                     // 0:  待审核 1:等待出款 2:拒绝 3:已完成 4:取消订单
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `    // 创建时间
	Currency       string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                             // 币种
	CallbackRemark string             `gorm:"column:callback_remark" json:"callback_remark" form:"callback_remark" uri:"callback_remark" ` // 状态备注
}
