package vo

type WithdrawDepositPromotionResp struct {
	Title       string  `json:"title"`                                           //显示名称
	PromotionId string  `json:"promotion_id"`                                    //优惠ID
	Rate        float64 `gorm:"column:rate" json:"rate" form:"rate" uri:"rate" ` // 比率百分比
}
