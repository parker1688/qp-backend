package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcBanks struct {
	BaseDos
	BankName     string             `gorm:"column:bank_name" json:"bank_name" form:"bank_name" uri:"bank_name" `                      // 银行名称
	BankCode     string             `gorm:"column:bank_code" json:"bank_code" form:"bank_code" uri:"bank_code" `                      // 银行简码Code
	MinLevel     int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小等级
	MaxLevel     int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大等级
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 状态 1: 启用 2 禁用
	Currency     string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种简码
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcBanks) TableName() string {
	return "fc_banks"
}
