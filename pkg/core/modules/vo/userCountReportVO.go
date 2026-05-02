package vo

type UserCountReportVO struct {
	Today   UserCountReportTagVO `json:"today"`
	History UserCountReportTagVO `json:"history"`
}

type UserCountReportTagVO struct {
	RechargeAmount       float64 `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `         // 充值金额
	Tax                  float64 `gorm:"column:tax" json:"tax" form:"tax" uri:"tax" `                                                         // 税收
	WithdrawalAmount     float64 `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" ` // 提款金额
	PromotionAmount      float64 `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `     // 优惠金额
	RebateAmount         float64 `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                 // 返水金额
	PlatformProfitAmount float64 `json:"platform_profit_amount" form:"platform_profit_amount" uri:"platform_profit_amount" `                  // 平台盈利金额
}
