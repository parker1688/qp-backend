package venuevo

type JDBResp struct {
	StatusCode int    `json:"status_code"`
	Balance    int    `json:"balance"` // 分为单位
	Message    string `json:"message"`
}

type JDBBalanceResp struct {
	Status string       `json:"status"`
	Error  string       `json:"err_text"` // 分为单位
	Data   []JDBBalance `json:"data"`
}

type JDBBalance struct {
	Uid         string  `json:"uid"`
	Balance     float64 `json:"balance"`
	Parent      string  `json:"parent"`
	UserName    string  `json:"username"`
	Curreny     string  `json:"currency"`
	Lvl         int     `json:"lvl"`
	Locked      int     `json:"locked"`
	Closed      int     `json:"closed"`
	JackpotFlag int     `json:"jackpotFlag"`
}

type JDBLoginResp struct {
	Status string `json:"status"`
	Error  string `json:"err_text"` // 分为单位
	Url    string `json:"path"`
}

type JDBDepositResp struct {
	Status string  `json:"status"`
	Error  string  `json:"err_text"` // 分为单位
	Amount float64 `json:"amount"`
}

type JDBTransferConfirmResp struct {
	TxnId    int    `json:"txn_id"`
	Balance  int    `json:"balance"` // 分为单位
	MemberId int    `json:"member_id"`
	Note     string `json:"note"`
}
