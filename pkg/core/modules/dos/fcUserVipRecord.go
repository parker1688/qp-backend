package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserVipRecord struct {
	BaseDos
	UserId               string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName             string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `         // 用户名
	BeforLevel           int                `gorm:"column:befor_level" json:"befor_level" form:"befor_level" uri:"befor_level" ` // 升级钱层级
	Level                int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                         // 升级后层级
	BeforVip             string             `gorm:"column:befor_vip" json:"befor_vip" form:"befor_vip" uri:"befor_vip" `         // 升级前VIP
	Vip                  string             `gorm:"column:vip" json:"vip" form:"vip" uri:"vip" `
	TotalRecharegeAmount float64            `gorm:"column:total_recharege_amount" json:"total_recharege_amount" form:"total_recharege_amount" uri:"total_recharege_amount" ` // 总充值额度
	TotalBetAmount       float64            `gorm:"column:total_bet_amount" json:"total_bet_amount" form:"total_bet_amount" uri:"total_bet_amount" `                         // 总投注额度
	CreateTime           automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                // 创建时间
	CreateBy             string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                     // 创建人
	UpdateTime           automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                // 修改时间
	UpdateBy             string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                     // 修改人
	MerchantCode         string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                     // 商户code
	Bonus                float64            `gorm:"column:bonus" json:"bonus" form:"bonus" uri:"bonus" `                                                                     // 升级奖金
	IssueBonus           float64            `gorm:"column:issue_bonus" json:"issue_bonus" form:"issue_bonus" uri:"issue_bonus" `                                             // 升级奖金
}

func (FcUserVipRecord) TableName() string {
	return "fc_user_vip_record"
}
