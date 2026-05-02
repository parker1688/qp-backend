package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVirtualCurrencyDetails struct {
	BaseDos
	CurrencyName  string             `gorm:"column:currency_name" json:"currency_name" form:"currency_name" uri:"currency_name" `      // 币种名称
	CurrencyChain string             `gorm:"column:currency_chain" json:"currency_chain" form:"currency_chain" uri:"currency_chain" `  // 币种所属链
	CreateTime    automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Status        int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 状态 1: 启用 2 禁用
	ToAddr        string             `gorm:"column:to_addr" json:"to_addr" form:"to_addr" uri:"to_addr" `                              // 接收地址
	ToAddrQrPre   string             `gorm:"column:to_addr_qr_pre" json:"to_addr_qr_pre" form:"to_addr_qr_pre" uri:"to_addr_qr_pre" `  // 接收地址二维码前缀
	MerchantCode  string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcVirtualCurrencyDetails) TableName() string {
	return "fc_virtual_currency_details"
}
