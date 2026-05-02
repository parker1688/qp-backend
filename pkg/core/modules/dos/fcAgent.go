package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcAgent struct {
	BaseDos
	AgentName    string             `gorm:"column:agent_name" json:"agent_name" form:"agent_name" uri:"agent_name" `                  // 代理名称
	InviteCode   int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `              // 用户邀请短码
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 读取状态 1正常 2 停用
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string             `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
}

func (FcAgent) TableName() string {
	return "fc_agent"
}
