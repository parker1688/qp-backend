package dos

type FcBindBankType struct {
	BaseDos
	BankName     string `gorm:"column:bank_name" json:"bank_name" form:"bank_name" uri:"bank_name" `                 // 银行卡名称
	BankCode     string `gorm:"column:bank_code" json:"bank_code" form:"bank_code" uri:"bank_code" `                 // 银行卡简码code
	Status       int    `gorm:"column:status" json:"status" form:"status" uri:"status" `                             // 状态 1 可用 2 不可用
	Sort         int    `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                     // 排序：值越大越靠前
	BankImg      string `gorm:"column:bank_img" json:"bank_img" form:"bank_img" uri:"bank_img" `                     // 银行卡图标
	Currency     string `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                     // 币种简码
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
}

func (FcBindBankType) TableName() string {
	return "fc_bind_bank_type"
}
