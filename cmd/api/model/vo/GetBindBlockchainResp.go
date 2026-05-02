package vo

type GetBindBlockchainResp struct {
	Id                string `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	Blockchain        string `gorm:"blockchain" json:"blockchain" form:"blockchain" uri:"blockchain" `                                 // 主链名称
	BlockchainAddress string `gorm:"blockchain_address" json:"blockchain_address" form:"blockchain_address" uri:"blockchain_address" ` // 区块链地址
	ContractType      string `gorm:"contract_type" json:"contract_type" form:"contract_type" uri:"contract_type" `                     // 合约类型
	IsDefault         int    `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `                                 // 1:不默认   2:默认
	RealName          string `gorm:"real_name" json:"real_name" form:"real_name" uri:"real_name" `                                     // 1:不默认   2:默认
	PaymentCode       string `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `                  // 通道code
	ChannelName       string `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `                  // 渠道code
	Icon              string `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	IconPath          string `gorm:"column:icon_path" json:"icon_path" form:"icon_path" uri:"icon_path" `
	Img               string `gorm:"column:img" json:"img" form:"img" uri:"img" `
	ImgPath           string `gorm:"column:img_path" json:"img_path" form:"img_path" uri:"img_path" ` // 币种
}

type GetBindOnlieResp struct {
	Id            string `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	AccountNumber string `gorm:"column:account_number" json:"account_number" form:"account_number" uri:"account_number" `
	AccountHolder string `gorm:"column:account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `
	IsDefault     int    `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `
	Icon          string `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	IconPath      string `gorm:"column:icon_path" json:"icon_path" form:"icon_path" uri:"icon_path" `
	Img           string `gorm:"column:img" json:"img" form:"img" uri:"img" `
	ImgPath       string `gorm:"column:img_path" json:"img_path" form:"img_path" uri:"img_path" `                 // 币种
	ChannelCode   string `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
}
