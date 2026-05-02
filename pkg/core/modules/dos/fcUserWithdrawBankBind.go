package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
)

type FcUserWithdrawBankBind struct {
	BaseDos
	UserId          string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName        string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Province        string             `gorm:"province" json:"province" form:"province" uri:"province" `
	City            string             `gorm:"city" json:"city" form:"city" uri:"city" `
	BankAddress     string             `gorm:"bank_address" json:"bank_address" form:"bank_address" uri:"bank_address" `                     // 银行地址
	AccountNumber   string             `gorm:"account_number" json:"account_number" form:"account_number" uri:"account_number" `             // 卡号
	NumberHash      string             `gorm:"number_hash" json:"number_hash" form:"number_hash" uri:"number_hash" `                         // 卡号Hash
	AccountHolder   string             `gorm:"account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `             // 收款人
	AccountBankType string             `gorm:"account_bank_type" json:"account_bank_type" form:"account_bank_type" uri:"account_bank_type" ` // 银行类别
	AccountBankCode string             `gorm:"account_bank_code" json:"account_bank_code" form:"account_bank_code" uri:"account_bank_code" ` // 银行编码
	IsDefault       int                `gorm:"is_default" json:"is_default" form:"is_default" uri:"is_default" `                             // 1:不默认   2:默认
	Sort            int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                                     // 排序 值越大越靠前
	Currency        string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                     // 币种
	CreateTime      automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `            // 创建时间
	CreateBy        string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                                 // 创建人
	UpdateTime      automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `            // 修改时间
	UpdateBy        string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                                 // 修改人
	MerchantCode    string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
	//BankType        int                `gorm:"column:bank_type" json:"bank_tye" form:"bank_tye" uri:"bank_tye" `
}

func (FcUserWithdrawBankBind) TableName() string {
	return "fc_user_withdraw_bank_bind"
}

func (m *FcUserWithdrawBankBind) Encrypt() {
	if len(m.AccountHolder) > 0 {
		// 随机字符串不利于搜索, 暂时注释
		m.AccountHolder, _ = tool.EncryptAESPrefixAesKey(m.AccountHolder, "fc_user_withdraw_bank_bind")
	}
	if len(m.AccountNumber) > 0 {
		m.AccountNumber, _ = tool.EncryptAESPrefixAesKey(m.AccountNumber, "fc_user_withdraw_bank_bind")
	}
	if len(m.BankAddress) > 0 {
		m.BankAddress, _ = tool.EncryptAESPrefixAesKey(m.BankAddress, "fc_user_withdraw_bank_bind")
	}
}

func (m *FcUserWithdrawBankBind) Decrypt() {
	if len(m.AccountHolder) > 0 {
		m.AccountHolder = tool.DecryptAESPrefixRandKeySaltDefault(m.AccountHolder, "fc_user_withdraw_bank_bind")
	}
	if len(m.AccountNumber) > 0 {
		m.AccountNumber = tool.DecryptAESPrefixRandKeySaltDefault(m.AccountNumber, "fc_user_withdraw_bank_bind")
	}
	if len(m.BankAddress) > 0 {
		m.BankAddress = tool.DecryptAESPrefixRandKeySaltDefault(m.BankAddress, "fc_user_withdraw_bank_bind")
	}
}
