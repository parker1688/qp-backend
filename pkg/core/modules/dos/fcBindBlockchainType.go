package dos

type FcBindBlockchainType struct {
	BaseDos
	Blockchain    string `gorm:"column:blockchain" json:"blockchain" form:"blockchain" uri:"blockchain" `                 // 币种所属链
	ContractType  string `gorm:"column:contract_type" json:"contract_type" form:"contract_type" uri:"contract_type" `     // 合约类型
	Sort          int    `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                         // 排序：值越大越靠前
	BlockchainImg string `gorm:"column:blockchain_img" json:"blockchain_img" form:"blockchain_img" uri:"blockchain_img" ` // 合约图标
	ContractName  string `gorm:"column:contract_name" json:"contract_name" form:"contract_name" uri:"contract_name" `     // 合约显示名称
}

func (FcBindBlockchainType) TableName() string {
	return "fc_bind_blockchain_type"
}
