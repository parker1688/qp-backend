package vo

type WalletMoneyResp struct {
	TotalAmount   float64 `gorm:"total_amount" json:"total_amount" form:"total_amount" uri:"total_amount" `         // 总金额
	AvaAmount     float64 `gorm:"ava_amount" json:"ava_amount" form:"ava_amount" uri:"ava_amount" `                 // 可用金额
	FronzenAmount float64 `gorm:"fronzen_amount" json:"fronzen_amount" form:"fronzen_amount" uri:"fronzen_amount" ` // 冻结金额
}
