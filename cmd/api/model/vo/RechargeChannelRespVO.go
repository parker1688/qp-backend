package vo

type RechargeChannelRespVO struct {
	Id          string  `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	ChannelName string  `gorm:"channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" ` // 渠道名称
	ChannelCode string  `gorm:"channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
	ChannelType int     `gorm:"channel_type" json:"channel_type" form:"channel_type" uri:"channel_type" ` // 支付渠道类型 1:微信 2:银行卡 3:支付宝 4:钱包 5:数字人民币 6:数字货币
	Currency    string  `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                 // 币种
	Icon        string  `gorm:"icon" json:"icon" form:"icon" uri:"icon"`                                  // 渠道icon
	Sort        int     `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                 // 排序：值越大越靠前
	MinAmount   float64 `gorm:"min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `         // 最小金额
	MaxAmount   float64 `gorm:"max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `         // 最大金额
	Hot         int     `gorm:"hot" json:"hot" form:"hot" uri:"hot" `                                     // 推荐

	AmountRange            string `gorm:"amount_range" json:"amount_range"`                           // 选择金额区间
	InputAmountDisplay     int    `gorm:"input_amount_display" json:"input_amount_display"`           // 输入金额是否显示
	InputNameDisplay       int    `gorm:"input_name_display" json:"input_name_display"`               // 输入存款姓名是否显示
	InputVirtualPayAddress int    `gorm:"input_virtual_pay_address" json:"input_virtual_pay_address"` // 输入虚拟币地址是否显示
	InputVirtualPayShow    int    `gorm:"input_virtual_pay_show" json:"input_virtual_pay_show"`       // 虚拟本地币选择显示
}
