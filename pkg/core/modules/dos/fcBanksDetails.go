package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
)

type FcBanksDetails struct {
	BaseDos
	BankName              string             `gorm:"column:bank_name" json:"bank_name" form:"bank_name" uri:"bank_name" `                                                             // 银行名称
	BankCode              string             `gorm:"column:bank_code" json:"bank_code" form:"bank_code" uri:"bank_code" `                                                             // 银行简码Code
	MinLevel              int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                                                             // 最小等级
	MaxLevel              int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                                                             // 最大等级
	Sort                  int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                                                                 // 排序：值越大越靠前
	CreateTime            automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                        // 创建时间
	CreateBy              string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                             // 创建人
	UpdateTime            automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                        // 修改时间
	UpdateBy              string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                             // 修改人
	EntityAccountHolder   string             `gorm:"column:entity_account_holder" json:"entity_account_holder" form:"entity_account_holder" uri:"entity_account_holder" `             // 收款人
	EntityAccountBankName string             `gorm:"column:entity_account_bank_name" json:"entity_account_bank_name" form:"entity_account_bank_name" uri:"entity_account_bank_name" ` // 收款人银行名字
	EntityAccountNumber   string             `gorm:"column:entity_account_number" json:"entity_account_number" form:"entity_account_number" uri:"entity_account_number" `             // 收款卡号或地址
	Status                int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                                         // 状态 1: 启用 2 禁用
	DayMaxAmount          float64            `gorm:"column:day_max_amount" json:"day_max_amount" form:"day_max_amount" uri:"day_max_amount" `                                         // 每日最大充值金额
	MinAmount             float64            `gorm:"column:min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                                                         // 单笔最低金额
	MaxAmount             float64            `gorm:"column:max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                                                         // 单笔最大金额
	MerchantCode          string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                             // 商户code
}

func (FcBanksDetails) TableName() string {
	return "fc_banks_details"
}

func (m *FcBanksDetails) Decrypt() {
	entityAccountHolder := tool.DecryptAESPrefixRandKeySalt(m.EntityAccountHolder, "fc_banks_details")
	if len(entityAccountHolder) > 0 {
		m.EntityAccountHolder = entityAccountHolder
	}
	entityAccountNumber := tool.DecryptAESPrefixRandKeySalt(m.EntityAccountNumber, "fc_banks_details")
	if len(entityAccountNumber) > 0 {
		m.EntityAccountNumber = entityAccountNumber
	}
}

func (m *FcBanksDetails) Encrypt() {
	if len(m.EntityAccountHolder) > 0 {
		m.EntityAccountHolder, _ = tool.EncryptAESPrefixRandKeySalt(m.EntityAccountHolder, "fc_banks_details")
	}
	if len(m.EntityAccountNumber) > 0 {
		m.EntityAccountNumber, _ = tool.EncryptAESPrefixRandKeySalt(m.EntityAccountNumber, "fc_banks_details")
	}
}
