package dos

type FcAgentReportDetailRep struct {
	BaseDos
	InviteCode int    `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" ` // 推广ID=邀请码
	StartAt    string `gorm:"column:start_at" json:"start_at" form:"start_at" uri:"start_at" `             // 开始时间
	EndAt      string `gorm:"column:end_at" json:"end_at" form:"end_at" uri:"end_at" `                     //结束时间
}

type FcAgentReportDetail struct {
	BaseDos
	InviteCode int    `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" ` // 推广ID=邀请码
	UserId     string `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                 //用户id

	CreateTime    string `gorm:"column:create_time" json:"create_time" form:"create_time" uri:"create_time" `                 // 注册时间
	LastLoginTime string `gorm:"column:last_login_time" json:"last_login_time" form:"last_login_time" uri:"last_login_time" ` //登录时间

	RechargeAmount      float64 `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `                         //总充值
	WithdrawalAmount    float64 `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" `                 //总提现
	ValidBetamount      float64 `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `                         //总有效投注
	WinAmount           float64 `gorm:"column:win_amount" json:"win_amount" form:"win_amount" uri:"win_amount" `                                             //总赢
	RebateAmount        float64 `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                                 //总反水
	PromotionAmount     float64 `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `                     //总福利
	SubRechargeWithDraw float64 `gorm:"column:sub_recharge_withdraw" json:"sub_recharge_withdraw" form:"sub_recharge_withdraw" uri:"sub_recharge_withdraw" ` //冲提差
}

func (FcAgentReportDetail) TableName() string {
	return "fc_agent_report_detail"
}
