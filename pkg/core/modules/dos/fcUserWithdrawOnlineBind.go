package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
)

type FcUserWithdrawOnlineBind struct {
	BaseDos
	UserId        string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName      string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	AccountNumber string             `gorm:"column:account_number" json:"account_number" form:"account_number" uri:"account_number" `
	AccountHolder string             `gorm:"column:account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `
	IsDefault     int                `gorm:"column:is_default" json:"is_default" form:"is_default" uri:"is_default" `                  // 1:不默认   2:默认
	Sort          int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序 值越大越靠前
	CreateTime    automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode  string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	ChannelName   string             `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `          // 渠道名称
	ChannelCode   string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
}

func (FcUserWithdrawOnlineBind) TableName() string {
	return "fc_user_withdraw_online_bind"
}

func (m *FcUserWithdrawOnlineBind) Encrypt() {
	if len(m.AccountHolder) > 0 {
		m.AccountHolder, _ = tool.EncryptAESPrefixAesKey(m.AccountHolder, "fc_user_withdraw_online_bind")
	}
	if len(m.AccountNumber) > 0 {
		m.AccountNumber, _ = tool.EncryptAESPrefixAesKey(m.AccountNumber, "fc_user_withdraw_online_bind")
	}
}

func (m *FcUserWithdrawOnlineBind) Decrypt() {
	m.AccountHolder = tool.DecryptAESPrefixRandKeySaltDefault(m.AccountHolder, "fc_user_withdraw_online_bind")
	m.AccountNumber = tool.DecryptAESPrefixRandKeySaltDefault(m.AccountNumber, "fc_user_withdraw_online_bind")
}
