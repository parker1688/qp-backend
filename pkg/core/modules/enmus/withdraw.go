package enmus

const (
	Withdraw_STATUS_wait_check = 0
	Withdraw_STATUS_wait_pay   = 1
	Withdraw_STATUS_refuse     = 2
	Withdraw_STATUS_succese    = 3
	Withdraw_STATUS_cancel     = 4
)

const (
	//重置提款限制
	RESETWITHDRAW_AGENT_KEY = "ResetWithdraw:%s"
)
