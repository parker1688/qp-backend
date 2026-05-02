package venuevo

type PTDZCreateResp struct {
	Error  string `json:"error"`
	Result struct {
		PlayerName string `json:"playername"`
		Password   string `json:"password"`
	} `json:"result"`
}

type PTDZBalanceResp struct {
	Error  string `json:"error"`
	Result struct {
		CUrrencyCode string `json:"currencycode"`
		Balance      string `json:"balance"`
	} `json:"result"`
}

type PTDZDepositResp struct {
	Error  string `json:"error"`
	Result struct {
		Amount string `json:"amount"`
	} `json:"result"`
}

type PTDZWithdrawResp struct {
	Error  string `json:"error"`
	Result struct {
		Amount string `json:"amount"`
	} `json:"result"`
}

type PTDZLoginResp struct {
	ErrorCode    int `json:"errorCode"`
	SessionToken struct {
		SessionToken string `json:"sessionToken"`
	} `json:"sessionToken"`
}
