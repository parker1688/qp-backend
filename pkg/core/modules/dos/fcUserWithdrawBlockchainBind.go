package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
)

type FcUserWithdrawBlockchainBind struct {
	BaseDos
	UserId            string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName          string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Blockchain        string             `gorm:"blockchain" json:"blockchain" form:"blockchain" uri:"blockchain" `                                 // 主链名称
	BlockchainAddress string             `gorm:"blockchain_address" json:"blockchain_address" form:"blockchain_address" uri:"blockchain_address" ` // 区块链地址
	ContractType      string             `gorm:"contract_type" json:"contract_type" form:"contract_type" uri:"contract_type" `                     // 合约类型
	IsDefault         int                `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `                                 // 1:不默认   2:默认
	Sort              int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                                         // 排序 值越大越靠前
	CreateTime        automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy          string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime        automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy          string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
	MerchantCode      string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                     // 商户code
	PaymentName       string             `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `                  // 通道名称
	PaymentCode       string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `                  // 通道code
	RealName          string             `gorm:"column:real_name" json:"real_name" form:"real_name" uri:"real_name" `                              // 通道code
}

func (FcUserWithdrawBlockchainBind) TableName() string {
	return "fc_user_withdraw_blockchain_bind"
}

func (m *FcUserWithdrawBlockchainBind) Encrypt() {
	if len(m.BlockchainAddress) > 0 {
		m.BlockchainAddress, _ = tool.EncryptAESPrefixAesKey(m.BlockchainAddress, "fc_user_withdraw_blockchain_bind")
	}

	if len(m.RealName) > 0 {
		m.RealName, _ = tool.EncryptAESPrefixAesKey(m.RealName, "fc_user_withdraw_blockchain_bind")
	}
}

func (m *FcUserWithdrawBlockchainBind) Decrypt() {
	m.BlockchainAddress = tool.DecryptAESPrefixRandKeySaltDefault(m.BlockchainAddress, "fc_user_withdraw_blockchain_bind")
	m.RealName = tool.DecryptAESPrefixRandKeySaltDefault(m.RealName, "fc_user_withdraw_blockchain_bind")
}
