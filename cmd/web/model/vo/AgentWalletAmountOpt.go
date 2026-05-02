package vo

type AgentWalletAmountOptReq struct {
	AgentId string  `json:"agent_id" form:"agent_id" uri:"agent_id" ` // 用户ID
	Amount  float64 `json:"amount" form:"amount" uri:"amount" `       // 存款金额
	Remark  string  `json:"remark" form:"remark" uri:"remark" `       // 备注
	OptType int     `json:"opt_type" form:"opt_type" uri:"opt_type" ` // 操作类型
}
