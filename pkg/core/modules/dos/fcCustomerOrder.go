package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcCustomerOrder struct {
	BaseDos
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" ` // 1:待处理 2：拒绝  3：解决
	Amount       float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `
	FlowMultiple int                `gorm:"column:flow_multiple" json:"flow_multiple" form:"flow_multiple" uri:"flow_multiple" ` // 流水倍数
	FlowAmount   float64            `gorm:"column:flow_amount" json:"flow_amount" form:"flow_amount" uri:"flow_amount" `         // 所需流水
	BonusType    int                `gorm:"column:bonus_type" json:"bonus_type" form:"bonus_type" uri:"bonus_type" `             // 交易工单类型
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `
	Remark       string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                             // 备注
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
	SolveRemark  string             `gorm:"column:solve_remark" json:"solve_remark" form:"solve_remark" uri:"solve_remark" `
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcCustomerOrder) TableName() string {
	return "fc_customer_order"
}
