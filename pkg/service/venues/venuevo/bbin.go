package venuevo

type BBINUserResp struct {
	LoginId  string `json:"loginId"`
	UserName string `json:"userName"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
	Result   string `json:"result"`
	Code     string `json:"code"`
	Msg      string `json:"msg"`
}

type BBINBalanceResp struct {
	WalletType string  `json:"walletType"`
	UserName   string  `json:"userName"`
	Result     string  `json:"result"`
	Balance    float64 `json:"balance"`
	Code       string  `json:"code"`
	Msg        string  `json:"msg"`
}

type BBINTransactionResp struct {
	Result          string `json:"result"`
	Amount          string `json:"amount"`
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	WalletType      string `json:"walletType"`
	UserName        string `json:"userName"`
	TransactionId   string `json:"transactionId"`
	TransactionType string `json:"transactionType"`
	Creditlimit     int    `json:"creditlimit"`
}

type BBINLoginResp struct {
	Result   string `json:"result"`
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	UserName string `json:"userName"`
	GameUrl  string `json:"gameUrl"`
}
