package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcInviteBetReportDay struct {
	BaseDos
	UserId           string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `         // 会员ID
	UserName         string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 会员名称
	UserType         int                `gorm:"column:user_type" json:"user_type" form:"user_type" uri:"user_type" ` // 1：普通用户 2：代理
	ReportDate       automaticType.Time `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `
	BetAmount        float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                                 // 有效投注金额
	BonusAmount      float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                         // 奖金金额
	BonusAmountIssue float64            `gorm:"column:bonus_amount_issue" json:"bonus_amount_issue" form:"bonus_amount_issue" uri:"bonus_amount_issue" ` // 奖金发放金额
	CreateTime       automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy         string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime       automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy         string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
	MerchantCode     string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                     // 商户code
}

func (FcInviteBetReportDay) TableName() string {
	return "fc_invite_bet_report_day"
}
