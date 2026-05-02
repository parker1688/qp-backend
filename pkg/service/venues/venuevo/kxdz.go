package venuevo

type KXDZLoginResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code int    `json:"code"`
		Url  string `json:"url"`
	}
}

type KXDZBalanceResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code     int     `json:"code"`
		Money    float64 `json:"money"`
		Account  string  `json:"account"`
		Currency string  `json:"currency"`
	}
}

type KXDZDepositResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code     int     `json:"code"`
		Money    float64 `json:"money"`
		Currency string  `json:"currency"`
	}
}

type KXDZWithdrawResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code     int     `json:"code"`
		Money    float64 `json:"money"`
		Currency string  `json:"currency"`
	}
}
