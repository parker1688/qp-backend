package vo

import "bootpkg/common/expands/automaticType"

type AgentWalletRecordResp struct {
	Id         string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Type       int                `gorm:"column:type" json:"type" form:"type" uri:"type" `                                          // 1'手动处理', 2'在线充值', 3'提款', 4'奖金', 5'代充', 6'扶持金'
	Amount     float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 变化额度
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
}
