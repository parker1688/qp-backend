package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcFinanceReportDay struct {
	BaseDos
	ReportDate          automaticType.Time `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `                                         // 统计日期
	Online              int                `gorm:"column:online" json:"online" form:"online" uri:"online" `                                                             // 在线人数
	RechargeUserCount   int                `gorm:"column:recharge_user_count" json:"recharge_user_count" form:"recharge_user_count" uri:"recharge_user_count" `         // 充值人数
	BetUserCount        int                `gorm:"column:bet_user_count" json:"bet_user_count" form:"bet_user_count" uri:"bet_user_count" `                             // 投注人数
	PromotionAmount     float64            `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `                     // 优惠
	RebateAmount        float64            `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                                 // 总返水
	GrossProfit         float64            `gorm:"column:gross_profit" json:"gross_profit" form:"gross_profit" uri:"gross_profit" `                                     // 毛利
	RechargeAmount      float64            `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `                         // 充值金额
	RechargeCount       int                `gorm:"column:recharge_count" json:"recharge_count" form:"recharge_count" uri:"recharge_count" `                             // 存款笔数
	AdminRechargeAmount float64            `gorm:"column:admin_recharge_amount" json:"admin_recharge_amount" form:"admin_recharge_amount" uri:"admin_recharge_amount" ` // 后台添加充值金额
	AdminRechargeCount  int                `gorm:"column:admin_recharge_count" json:"admin_recharge_count" form:"admin_recharge_count" uri:"admin_recharge_count" `     // 后台添加充值金额笔数
	WithdrawalAmount    float64            `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" `                 // 提款金额
	WithdrawalCount     int                `gorm:"column:withdrawal_count" json:"withdrawal_count" form:"withdrawal_count" uri:"withdrawal_count" `                     // 提款笔数
	AgentAwardAmount    float64            `gorm:"column:agent_award_amount" json:"agent_award_amount" form:"agent_award_amount" uri:"agent_award_amount" `             // 代理奖励
	MerchantCode        string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                 // 商户code
	CreateTime          automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                            // 创建时间
	CreateBy            string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                 // 创建人
	UpdateTime          automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                            // 修改时间
	UpdateBy            string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                 // 修改人
}

func (FcFinanceReportDay) TableName() string {
	return "fc_finance_report_day"
}
