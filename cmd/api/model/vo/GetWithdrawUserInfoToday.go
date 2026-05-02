package vo

type GetWithdrawUserInfoTodayResp struct {
	Amount            float64 `json:"amount"`              //当日剩余金额
	Num               int     `json:"num"`                 //当日剩余次数
	MinRechargeAmount float64 `json:"min_recharge_amount"` // 最小充值金额
	MinWithdrawAmount float64 `json:"min_withdraw_amount"` // 最小提现金额
	Vip               int     `json:"vip"`
}
