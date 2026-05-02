package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVirtualCurrency struct {
	BaseDos
	CurrencyName     string             `gorm:"column:currency_name" json:"currency_name" form:"currency_name" uri:"currency_name" `                 // 币种名称
	CurrencyNameImg  string             `gorm:"column:currency_name_img" json:"currency_name_img" form:"currency_name_img" uri:"currency_name_img" ` // 币种图片
	CurrencyChain    string             `gorm:"column:currency_chain" json:"currency_chain" form:"currency_chain" uri:"currency_chain" `             // 币种所属链
	CurrencyProtocol string             `gorm:"column:currency_protocol" json:"currency_protocol" form:"currency_protocol" uri:"currency_protocol" ` // 币种显示名称
	CreateTime       automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `            // 创建时间
	CreateBy         string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                                         // 创建人
	UpdateTime       automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `            // 修改时间
	UpdateBy         string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                                         // 修改人
	Status           int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                             // 状态 1: 启用 2 禁用
	MerchantCode     string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
}

func (FcVirtualCurrency) TableName() string {
	return "fc_virtual_currency"
}
