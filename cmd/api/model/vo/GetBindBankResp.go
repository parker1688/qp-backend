package vo

type GetBindBankResp struct {
	Id              string `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	BankAddress     string `gorm:"bank_address" json:"bank_address" form:"bank_address" uri:"bank_address" `                     // 银行地址
	AccountNumber   string `gorm:"account_number" json:"account_number" form:"account_number" uri:"account_number" `             // 卡号
	AccountHolder   string `gorm:"account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `             // 收款人
	AccountBankType string `gorm:"account_bank_type" json:"account_bank_type" form:"account_bank_type" uri:"account_bank_type" ` // 银行类别
	AccountBankCode string `gorm:"account_bank_code" json:"account_bank_code" form:"account_bank_code" uri:"account_bank_code" ` // 银行编码
	IsDefault       int    `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `                             // 1:不默认   2:默认
	Currency        string `gorm:"currency" json:"currency" form:"currency" uri:"currency" `
	Icon            string `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	IconPath        string `gorm:"column:icon_path" json:"icon_path" form:"icon_path" uri:"icon_path" `
	Img             string `gorm:"column:img" json:"img" form:"img" uri:"img" `
	ImgPath         string `gorm:"column:img_path" json:"img_path" form:"img_path" uri:"img_path" ` // 币种
}

func (m *GetBindBankResp) Hide() {
	if len(m.AccountNumber) > 3 {
		//m.AccountNumber = strutil.HideString(m.AccountNumber, 3, len(m.AccountNumber), "*")
	}
	if len(m.AccountHolder) > 3 {
		//m.AccountHolder = strutil.HideString(m.AccountHolder, 3, len(m.AccountNumber), "*")
	}
}
