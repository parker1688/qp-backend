package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserReportDay struct {
	BaseDos
	UserId              string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                                         // 用户Id
	UserName            string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                                                 // 用户账号
	RechargeAmount      float64            `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `                         // 充值金额
	RechargeCount       int                `gorm:"column:recharge_count" json:"recharge_count" form:"recharge_count" uri:"recharge_count" `                             // 存款笔数
	RechargeFee         float64            `gorm:"column:recharge_fee" json:"recharge_fee" form:"recharge_fee" uri:"recharge_fee" `                                     // 充值手续费
	BetAmount           float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                                             // 投注金额
	BetCount            int                `gorm:"column:bet_count" json:"bet_count" form:"bet_count" uri:"bet_count" `                                                 // 投注笔数
	ValidBetamount      float64            `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `                         // 有效投注
	WithdrawalAmount    float64            `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" `                 // 提款金额
	WithdrawalCount     int                `gorm:"column:withdrawal_count" json:"withdrawal_count" form:"withdrawal_count" uri:"withdrawal_count" `                     // 提款笔数
	WithdrawalFee       float64            `gorm:"column:withdrawal_fee" json:"withdrawal_fee" form:"withdrawal_fee" uri:"withdrawal_fee" `                             // 提款手续费
	PromotionAmount     float64            `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `                     // 优惠金额
	RebateAmount        float64            `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                                 // 返水金额
	WinAmount           float64            `gorm:"column:win_amount" json:"win_amount" form:"win_amount" uri:"win_amount" `                                             // 输赢金额
	AbsolutelyWinAmount float64            `gorm:"column:absolutely_win_amount" json:"absolutely_win_amount" form:"absolutely_win_amount" uri:"absolutely_win_amount" ` // 绝对输赢（去掉红利）
	ReportDate          string             `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `
	MerchantCode        string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	CreateTime          automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy            string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime          automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy            string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcUserReportDay) TableName() string {
	return "fc_user_report_day"
}

// 玩家账目
type FcUserReportDayBills struct {
	FcUserReportDay
	RechargeWithDrawSub float64 `gorm:"column:recharge_withdrawal_sub" json:"recharge_withdrawal_sub" form:"recharge_withdrawal_sub" uri:"recharge_withdrawal_sub" ` // 充提差
	FinanceOdds         float64 `gorm:"column:finance_odds" json:"finance_odds" form:"finance_odds" uri:"finance_odds" `                                             // 财务杀率
	BetMultiplier       float64 `gorm:"column:bet_multiplier" json:"bet_multiplier" form:"bet_multiplier" uri:"bet_multiplier" `                                     // 下注倍数
	LocalBet            float64 `gorm:"column:local_bet" json:"local_bet" form:"local_bet" uri:"local_bet" `                                                         // 本站投注
	LocalWinAmount      float64 `gorm:"column:local_win_mount" json:"local_win_mount" form:"local_win_mount" uri:"local_win_mount" `                                 // 本站输赢
	WinLose             float64 `gorm:"column:win_lose" json:"win_lose" form:"win_lose" uri:"win_lose" `                                                             // 本站输赢
}
