package vo

type UserRechargeReq struct {
	ChannelCode   string  `json:"channel_code" form:"channel_code" uri:"channel_code" `                     // 渠道code
	PaymentCode   string  `gorm:"payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
	PayId         string  `gorm:"pay_id" json:"pay_id" form:"pay_id" uri:"pay_id" `                         // 通道产品id
	Amount        float64 `json:"amount" form:"amount" uri:"amount" `                                       // 存款金额
	DepositRemark string  `json:"deposit_remark" form:"deposit_remark" uri:"deposit_remark" `               // 存款备注
	DepositName   string  `json:"deposit_name" form:"deposit_name" uri:"deposit_name" `                     // 存款人姓名
	Currency      string  `json:"currency" form:"currency" uri:"currency" `                                 // 币种简码

	VirtualPayAddress string `json:"virtual_pay_address" form:"virtual_pay_address" uri:"virtual_pay_address" ` // 虚拟币支付地址
	ActivityId        string `json:"activity_id" form:"activity_id" uri:"activity_id" `                         // 参与活动ID

	ReturnUrl string `json:"return_url"` //支付成功后返回地址
}

type GetPaymentReq struct {
	ChannelCode string `json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
}

type GetPaymentDetailReq struct {
	Id string `json:"id" form:"id" uri:"id" ` // id
}

type GetPaymentResp struct {
	Id           string  `gorm:"id" json:"id" form:"id" uri:"id" `
	PaymentName  string  `gorm:"payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `         // 通道名称
	PayAliasName string  `gorm:"pay_alias_name" json:"pay_alias_name" form:"pay_alias_name" uri:"pay_alias_name" ` // 通道别名
	PaymentCode  string  `gorm:"payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `         // 通道code
	PayId        string  `gorm:"pay_id" json:"pay_id" form:"pay_id" uri:"pay_id" `                                 // 通道产品id
	ChannelName  string  `gorm:"channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `         // 渠道名称
	ChannelCode  string  `gorm:"channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `         // 渠道code
	Status       int     `gorm:"status" json:"status" form:"status" uri:"status" `                                 // 1:正常  2:禁止
	MinLevel     int     `gorm:"min_level" json:"min_level" form:"min_level" uri:"min_level" `                     // 最小VIP等级
	MaxLevel     int     `gorm:"max_level" json:"max_level" form:"max_level" uri:"max_level" `                     // 最大vip等级
	MinAmount    float64 `gorm:"min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                 // 单笔最低金额
	MaxAmount    float64 `gorm:"max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                 // 单笔最大金额
	DayMaxAmount float64 `gorm:"day_max_amount" json:"day_max_amount" form:"day_max_amount" uri:"day_max_amount" ` // 每日最大充值金额
	BonusRate    float64 `gorm:"bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                 // 优惠比例
	Sort         int     `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                         // 排序：值越大越靠前
	MerchantCode string  `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `     // 商户code
	AmountRange  string  `gorm:"amount_range" json:"amount_range"`                                                 // 固定金额
	Remark       string  `gorm:"remark" json:"remark"`                                                             // 备注
}

type UserPayChannelReq struct {
	Currency string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
}

type UserPayChannelSettingReq struct {
	ChannelCode string `json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
}

type UserPayChannelSettingResp struct {
	AmountRange            []float64 `json:"amount_range"`              // 选择金额区间
	InputAmountDisplay     bool      `json:"input_amount_display"`      // 输入金额是否显示
	InputNameDisplay       bool      `json:"input_name_display"`        // 输入存款姓名是否显示
	InputVirtualPayAddress bool      `json:"input_virtual_pay_address"` // 输入虚拟币地址是否显示

	InputVirtualPayShow     bool `json:"input_virtual_pay_show"`   // 虚拟本地币选择
	InputVirtualPayShowList bool `json:"input_virtual_pay_show_l"` // 虚拟币快速充值(每个账户一个币种地址)
}
