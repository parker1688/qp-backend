package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVipWeekGift struct {
	BaseDos
	UserId           string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName         string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Week             string             `gorm:"column:week" json:"week" form:"week" uri:"week" `                                                         // 周 2025-2-3 年-月-周
	VipName          string             `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" `                                         // VIP名称
	Level            int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                                     // 层级
	MerchantCode     string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                     // 商户code
	BonusAmount      float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                         // 奖金金额
	BonusAmountIssue float64            `gorm:"column:bonus_amount_issue" json:"bonus_amount_issue" form:"bonus_amount_issue" uri:"bonus_amount_issue" ` // 已发奖金金额
	CreateTime       automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy         string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime       automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy         string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
}

func (FcVipWeekGift) TableName() string {
	return "fc_vip_week_gift"
}
