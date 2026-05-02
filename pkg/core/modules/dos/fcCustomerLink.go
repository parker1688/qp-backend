package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcCustomerLink struct {
	BaseDos
	Name         string             `gorm:"column:name" json:"name" form:"name" uri:"name" `                                          // 名称
	Link         string             `gorm:"column:link" json:"link" form:"link" uri:"link" `                                          // 网站
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	MerchantName string             `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `      // 商户名称
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1:正常  2：禁止
}

func (FcCustomerLink) TableName() string {
	return "fc_customer_link"
}
