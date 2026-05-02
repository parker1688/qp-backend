package vo

type UserWalletMoneyRequest struct {
	Currency     string `json:"currency" form:"currency" uri:"currency" `                // 币种简码
	VenueBalance bool   `json:"venue_balance" form:"venue_balance" uri:"venue_balance" ` // 总金额 true 不带场余额 false 带场馆余额
}
