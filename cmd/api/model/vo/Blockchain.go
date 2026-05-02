package vo

type FcUserWithdrawBlockchainBind struct {
	//Blockchain        string ` json:"blockchain" form:"blockchain" uri:"blockchain" validate:"required" `                        // 主链名称
	BlockchainAddress string ` json:"blockchain_address" form:"blockchain_address" uri:"blockchain_address" validate:"required"` // 区块链地址
	//ContractType      string ` json:"contract_type" form:"contract_type" uri:"contract_type" validate:"required"`                // 合约类型
	VeryCode    string ` json:"veryCode" form:"veryCode" uri:"veryCode" validate:"required"`
	PaymentName string `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" ` // 通道名称
	PaymentCode string `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
	IsDefault   int    `gorm:"column:is_default" json:"is_default" form:"is_default" uri:"is_default" `         // 1:不默认   2:默认
	RealName    string `gorm:"column:real_name" json:"real_name" form:"real_name" uri:"real_name" `             // 1:姓名
}

type FcUserWithdrawOnlineBind struct {
	AccountNumber string `gorm:"column:account_number" json:"account_number" form:"account_number" uri:"account_number" `
	AccountHolder string `gorm:"column:account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `
	IsDefault     int    `gorm:"column:is_default" json:"is_default" form:"is_default" uri:"is_default" `         // 1:不默认   2:默认
	ChannelName   string `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" ` // 渠道名称
	ChannelCode   string `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
	VeryCode      string ` json:"veryCode" form:"veryCode" uri:"veryCode" validate:"required"`
}

type FcUserWithdrawBankBind struct {
	UserId          string `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName        string `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Province        string `gorm:"province" json:"province" form:"province" uri:"province" `
	City            string `gorm:"city" json:"city" form:"city" uri:"city" `
	BankAddress     string `gorm:"bank_address" json:"bank_address" form:"bank_address" uri:"bank_address" `                     // 银行地址
	AccountNumber   string `gorm:"account_number" json:"account_number" form:"account_number" uri:"account_number" `             // 卡号
	NumberHash      string `gorm:"number_hash" json:"number_hash" form:"number_hash" uri:"number_hash" `                         // 卡号Hash
	AccountHolder   string `gorm:"account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `             // 收款人
	AccountBankType string `gorm:"account_bank_type" json:"account_bank_type" form:"account_bank_type" uri:"account_bank_type" ` // 银行类别
	AccountBankCode string `gorm:"account_bank_code" json:"account_bank_code" form:"account_bank_code" uri:"account_bank_code" ` // 银行编码
	IsDefault       int    `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `                             // 1:不默认   2:默认
	Sort            int    `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                                     // 排序 值越大越靠前
	Currency        string `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                     // 币种
	MerchantCode    string `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
	BankType        int    `gorm:"column:bank_type" json:"bank_tye" form:"bank_tye" uri:"bank_tye" `
	VeryCode        string ` json:"veryCode" form:"veryCode" uri:"veryCode" validate:"required"`
	PaymentName     string `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" ` // 通道名称
	PaymentCode     string `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
}
