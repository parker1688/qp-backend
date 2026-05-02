package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserReport struct {
	BaseDos
	UserId              string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                                         // 用户Id
	UserName            string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                                                 // 用户账号
	RechargeAmount      float64            `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `                         // 充值金额
	RechargeCount       int                `gorm:"column:recharge_count" json:"recharge_count" form:"recharge_count" uri:"recharge_count" `                             // 存款笔数
	Tax                 float64            `gorm:"column:tax" json:"tax" form:"tax" uri:"tax" `                                                                         // 税收
	BetAmount           float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                                             // 投注金额
	BetCount            int                `gorm:"column:bet_count" json:"bet_count" form:"bet_count" uri:"bet_count" `                                                 // 投注笔数
	ValidBetamount      float64            `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `                         // 有效投注
	WithdrawalAmount    float64            `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" `                 // 提款金额
	WithdrawalCount     int                `gorm:"column:withdrawal_count" json:"withdrawal_count" form:"withdrawal_count" uri:"withdrawal_count" `                     // 提款笔数
	PromotionAmount     float64            `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `                     // 优惠金额
	RebateAmount        float64            `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                                 // 返水金额
	WinAmount           float64            `gorm:"column:win_amount" json:"win_amount" form:"win_amount" uri:"win_amount" `                                             // 输赢金额
	AbsolutelyWinAmount float64            `gorm:"column:absolutely_win_amount" json:"absolutely_win_amount" form:"absolutely_win_amount" uri:"absolutely_win_amount" ` // 绝对输赢（去掉红利）
	FriendsBonusAmount  float64            `gorm:"column:friends_bonus_amount" json:"friends_bonus_amount" form:"friends_bonus_amount" uri:"friends_bonus_amount" `     // 邀请佣金
	MerchantCode        string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                 // 商户code
	CreateTime          automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                            // 创建时间
	CreateBy            string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                 // 创建人
	UpdateTime          automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                            // 修改时间
	UpdateBy            string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                 // 修改人
}

func (FcUserReport) TableName() string {
	return "fc_user_report"
}

type FcUserReportListResp struct {
	FcUserReport
	SubRechargeWithDraw float64 `gorm:"column:sub_recharge_with_draw" json:"sub_recharge_with_draw" form:"sub_recharge_with_draw" uri:"sub_recharge_with_draw" ` // 充提差
	FdOdds              float64 `gorm:"column:fd_odds" json:"fd_odds" form:"fd_odds" uri:"fd_odds" `                                                             // 财务杀率
	BetMultiple         float64 `gorm:"column:bet_multiple" json:"bet_multiple" form:"bet_multiple" uri:"bet_multiple" `                                         // 下注倍数

}
