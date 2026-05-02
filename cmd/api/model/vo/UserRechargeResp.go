package vo

type UserRechargeResp struct {
	Id          int     `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	PaymentCode string  `gorm:"payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
	MinAmount   float64 `gorm:"min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `         // 单笔最低金额
	MaxAmount   float64 `gorm:"max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `         // 单笔最大金额
}
