package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcCustomerOrderType struct {
	BaseDos
	FlowMultiple int                `gorm:"column:flow_multiple" json:"flow_multiple" form:"flow_multiple" uri:"flow_multiple" `      // 打码倍数
	BonusType    int                `gorm:"column:bonus_type" json:"bonus_type" form:"bonus_type" uri:"bonus_type" `                  // 交易工单类型
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 名称
	Remark       string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                  // 备注
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcCustomerOrderType) TableName() string {
	return "fc_customer_order_type"
}
