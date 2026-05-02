package vo

type WalletAmountOptRequest struct {
	UserId       string  `json:"user_id" form:"user_id" uri:"user_id" `                   // 用户ID
	Amount       float64 `json:"amount" form:"amount" uri:"amount" `                      // 存款金额
	Remark       string  `json:"remark" form:"remark" uri:"remark" `                      // 备注
	Currency     string  `json:"currency" form:"currency" uri:"currency" `                // 币种简码
	OptType      int     `json:"opt_type" form:"opt_type" uri:"opt_type" `                // 操作类型
	FlowMultiple int     `json:"flow_multiple" form:"flow_multiple" uri:"flow_multiple" ` //流水要求倍数
}
