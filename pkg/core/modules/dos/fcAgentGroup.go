package dos

import "bootpkg/common/expands/automaticType"

type FcAgentGroup struct {
	BaseDos
	InviteCode   int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code"`
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code"`
	GroupName    string             `gorm:"column:group_name" json:"group_name" form:"group_name" uri:"group_name"`
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `
}

func (FcAgentGroup) TableName() string {
	return "fc_agent_group"
}
