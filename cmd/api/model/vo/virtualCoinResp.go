package vo

type VirtualCoinResp struct {
	CurrencyName     string  `gorm:"column:currency_name" json:"currency_name" form:"currency_name" uri:"currency_name" `                 // 币种名称
	CurrencyNameImg  string  `gorm:"column:currency_name_img" json:"currency_name_img" form:"currency_name_img" uri:"currency_name_img" ` // 币种图片
	CurrencyChain    string  `gorm:"column:currency_chain" json:"currency_chain" form:"currency_chain" uri:"currency_chain" `             // 币种所属链
	CurrencyProtocol string  `gorm:"column:currency_protocol" json:"currency_protocol" form:"currency_protocol" uri:"currency_protocol" ` // 币种显示名称
	FxAmount         float64 `gorm:"column:fx_amount" json:"fx_amount" form:"fx_amount" uri:"fx_amount" `                                 // 汇率金额/个
	//OptType          int     `gorm:"column:opt_type" json:"opt_type" form:"opt_type" uri:"opt_type" `                                     // 1: 存款  2: 提款
}
