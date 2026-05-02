package vo

type AgentPopularizeResp struct {
	Id                string  `gorm:"column:id" json:"id" form:"id" uri:"id" `
	AgentName         string  `gorm:"column:agent_name" json:"agent_name" form:"agent_name" uri:"agent_name" `                                 // 代理名称
	BalanceCommission float64 `gorm:"column:balance_commission" json:"balance_commission" form:"balance_commission" uri:"balance_commission" ` // 可提现佣金
}
