package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserRankReportDay struct {
	BaseDos
	ReportDate                 automaticType.Time `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" ` // 统计日期
	UserId                     string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName                   string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	UserType                   int                `gorm:"column:user_type" json:"user_type" form:"user_type" uri:"user_type" `                                                                                     // 1:会员  2：代理
	Business                   string             `gorm:"column:business" json:"business" form:"business" uri:"business" `                                                                                         // 商务
	SupportAmount              float64            `gorm:"column:support_amount" json:"support_amount" form:"support_amount" uri:"support_amount" `                                                                 // 扶植金
	GradeOneNum                int                `gorm:"column:grade_one_num" json:"grade_one_num" form:"grade_one_num" uri:"grade_one_num" `                                                                     // 一级人数
	GradeOneRechargeNum        int                `gorm:"column:grade_one_recharge_num" json:"grade_one_recharge_num" form:"grade_one_recharge_num" uri:"grade_one_recharge_num" `                                 // 一级充值人数
	GradeOneRechargeFee        float64            `gorm:"column:grade_one_recharge_fee" json:"grade_one_recharge_fee" form:"grade_one_recharge_fee" uri:"grade_one_recharge_fee" `                                 // 一级充值成本
	GradeOneRechargeAmount     float64            `gorm:"column:grade_one_recharge_amount" json:"grade_one_recharge_amount" form:"grade_one_recharge_amount" uri:"grade_one_recharge_amount" `                     // 一级充值
	GradeOneRechargeAvAmount   float64            `gorm:"column:grade_one_recharge_av_amount" json:"grade_one_recharge_av_amount" form:"grade_one_recharge_av_amount" uri:"grade_one_recharge_av_amount" `         // 一级均充
	GradeOneRechargePerce      float64            `gorm:"column:grade_one_recharge_perce" json:"grade_one_recharge_perce" form:"grade_one_recharge_perce" uri:"grade_one_recharge_perce" `                         // 一级充值率
	GradeTwoNum                int                `gorm:"column:grade_two_num" json:"grade_two_num" form:"grade_two_num" uri:"grade_two_num" `                                                                     // 二级人数
	GradeTwoRechargeNum        int                `gorm:"column:grade_two_recharge_num" json:"grade_two_recharge_num" form:"grade_two_recharge_num" uri:"grade_two_recharge_num" `                                 // 二级充值人数
	GradeTwoRechargeAmount     float64            `gorm:"column:grade_two_recharge_amount" json:"grade_two_recharge_amount" form:"grade_two_recharge_amount" uri:"grade_two_recharge_amount" `                     // 二级充值
	GradeTwoRechargeAvAmount   float64            `gorm:"column:grade_two_recharge_av_amount" json:"grade_two_recharge_av_amount" form:"grade_two_recharge_av_amount" uri:"grade_two_recharge_av_amount" `         // 二级均充
	GradeTwoRechargePerce      float64            `gorm:"column:grade_two_recharge_perce" json:"grade_two_recharge_perce" form:"grade_two_recharge_perce" uri:"grade_two_recharge_perce" `                         // 二级充值率
	GradeThreeNum              int                `gorm:"column:grade_three_num" json:"grade_three_num" form:"grade_three_num" uri:"grade_three_num" `                                                             // 三级人数
	GradeThreeRechargeNum      int                `gorm:"column:grade_three_recharge_num" json:"grade_three_recharge_num" form:"grade_three_recharge_num" uri:"grade_three_recharge_num" `                         // 三级充值人数
	GradeThreeRechargeAmount   float64            `gorm:"column:grade_three_recharge_amount" json:"grade_three_recharge_amount" form:"grade_three_recharge_amount" uri:"grade_three_recharge_amount" `             // 三级充值
	GradeThreeRechargeAvAmount float64            `gorm:"column:grade_three_recharge_av_amount" json:"grade_three_recharge_av_amount" form:"grade_three_recharge_av_amount" uri:"grade_three_recharge_av_amount" ` // 三级平均充值
	GradeThreeRechargePerce    float64            `gorm:"column:grade_three_recharge_perce" json:"grade_three_recharge_perce" form:"grade_three_recharge_perce" uri:"grade_three_recharge_perce" `                 // 三级充值率
	UserInviteBonusAmount      float64            `gorm:"column:user_invite_bonus_amount" json:"user_invite_bonus_amount" form:"user_invite_bonus_amount" uri:"user_invite_bonus_amount" `                         // 邀请奖金
	MerchantCode               string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                                                     // 商户code
	CreateTime                 automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                                                // 创建时间
	CreateBy                   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                                                     // 创建人
	UpdateTime                 automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                                                // 修改时间
	UpdateBy                   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                                                     // 修改人
	TotalNum                   int                `gorm:"column:total_num" json:"total_num" form:"total_num" uri:"total_num" `                                                                                     // 共计人数
	TotalRechargeNum           int                `gorm:"column:total_recharge_num" json:"total_recharge_num" form:"total_recharge_num" uri:"total_recharge_num" `                                                 // 综合充值人数
	TotalRechargeAmount        float64            `gorm:"column:total_recharge_amount" json:"total_recharge_amount" form:"total_recharge_amount" uri:"total_recharge_amount" `                                     // 综合充值
	TotalRechargeAvAmount      float64            `gorm:"column:total_recharge_av_amount" json:"total_recharge_av_amount" form:"total_recharge_av_amount" uri:"total_recharge_av_amount" `                         // 综均充
	TotalRechargeFee           float64            `gorm:"column:total_recharge_fee" json:"total_recharge_fee" form:"total_recharge_fee" uri:"total_recharge_fee" `                                                 // 综合充值成本
	TotalRechargePerce         float64            `gorm:"column:total_recharge_perce" json:"total_recharge_perce" form:"total_recharge_perce" uri:"total_recharge_perce" `                                         // 综合充值率
	Remark                     string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                                                                                 // 备注
}

func (FcUserRankReportDay) TableName() string {
	return "fc_user_rank_report_day"
}
