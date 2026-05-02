package venuevo

type PPDZUserResp struct {
	Error       string `json:"error"`
	Description string `json:"description"`
	PlayerId    int64  `json:"playerId"`
}

type PPDZBalanceResp struct {
	Error       string  `json:"error"`
	Description string  `json:"description"`
	Balance     float64 `json:"balance"`
}

type PPDZLoginResp struct {
	Error       string `json:"error"`
	Description string `json:"description"`
	GameUrl     string `json:"gameURL"`
}

type PPDZTransactionResp struct {
	Error         string  `json:"error"`
	Description   string  `json:"description"`
	TransactionId int64   `json:"transactionId"`
	Balance       float64 `json:"balance"`
}
