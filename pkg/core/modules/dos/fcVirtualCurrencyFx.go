package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVirtualCurrencyFx struct {
	BaseDos
	CurrencyName  string             `gorm:"column:currency_name" json:"currency_name" form:"currency_name" uri:"currency_name" `      // 币种名称
	CurrencyChain string             `gorm:"column:currency_chain" json:"currency_chain" form:"currency_chain" uri:"currency_chain" `  // 币种所属链
	CreateTime    automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	FxAmount      float64            `gorm:"column:fx_amount" json:"fx_amount" form:"fx_amount" uri:"fx_amount" `                      // 汇率金额/个
	OptType       int                `gorm:"column:opt_type" json:"opt_type" form:"opt_type" uri:"opt_type" `                          // 1: 存款  2: 提款
	CurrencyCode  string             `gorm:"column:currency_code" json:"currency_code" form:"currency_code" uri:"currency_code" `      // 法币类型Code
	MerchantCode  string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcVirtualCurrencyFx) TableName() string {
	return "fc_virtual_currency_fx"
}
