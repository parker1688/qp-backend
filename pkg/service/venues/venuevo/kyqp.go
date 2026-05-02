package venuevo

type KYQPSportRegisterResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code int    `json:"code"`
		URL  string `json:"url"`
	} `json:"d"`
}

type KYQPBalanceResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		TotalMoney float64 `json:"totalMoney"`
		FreeMoney  float64 `json:"freeMoney"`
		Status     int     `json:"status"`
		Code       int     `json:"code"`
	} `json:"d"`
}

type KYQPDepositResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code int `json:"code"`
	} `json:"d"`
}

type KYQPWithdrawResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code int `json:"code"`
	} `json:"d"`
}

type KYQPTransferConfirmResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code   int `json:"code"`
		Status int `json:"status"`
	} `json:"d"`
}
