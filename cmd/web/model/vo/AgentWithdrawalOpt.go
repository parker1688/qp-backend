package vo

type AgentWithdrawalOptReq struct {
	AgentId string  `json:"agent_id" form:"agent_id" uri:"agent_id" ` // 用户ID
	Amount  float64 `json:"amount" form:"amount" uri:"amount" `       // 存款金额
	Remark  string  `json:"remark" form:"remark" uri:"remark" `       // 备注

	BankAddress     string `gorm:"column:bank_address" json:"bank_address" form:"bank_address" uri:"bank_address" `                     // 银行地址
	AccountNumber   string `gorm:"column:account_number" json:"account_number" form:"account_number" uri:"account_number" `             // 卡号或地址
	AccountHolder   string `gorm:"column:account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `             // 收款人
	AccountBankType string `gorm:"column:account_bank_type" json:"account_bank_type" form:"account_bank_type" uri:"account_bank_type" ` // 银行类型名称
}
