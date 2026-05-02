package dos

type FcVipCondition struct {
	VipId              string  `gorm:"column:vip_id" json:"vip_id" form:"vip_id" uri:"vip_id" `
	Level              int     `gorm:"column:level" json:"level" form:"level" uri:"level" `                                                             // 层级：1 ~ 10
	VipName            string  `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" `                                                 // VIP名称：VIP1~VIP10
	MinRecharegeAmount float64 `gorm:"column:min_recharege_amount" json:"min_recharege_amount" form:"min_recharege_amount" uri:"min_recharege_amount" ` // 累计最小充值金额
	MinBetAmount       float64 `gorm:"column:min_bet_amount" json:"min_bet_amount" form:"min_bet_amount" uri:"min_bet_amount" `                         // 最小投注金额
	MerchantCode       string  `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                             // 商户code
}

func (FcVipCondition) TableName() string {
	return "fc_vip_condition"
}
