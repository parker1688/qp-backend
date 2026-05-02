package venuevo

type TYQPSportRegisterResp struct {
	Info       string `json:"info"`
	Code       string `json:"code"`
	Loginurl   string `json:"loginurl"`
	LoginurlPc string `json:"loginurl_pc"`
}

type TYQPBalanceResp struct {
	Code     string `json:"code"`
	BagMoney string `json:"bag_money"`
}

type TYQPDepositResp struct {
	Code     string `json:"code"`
	Info     string `json:"info"`
	BagMoney string `json:"bag_money"`
}

type TYQPWithdrawResp struct {
	Code     string `json:"code"`
	Info     string `json:"info"`
	BagMoney string `json:"bag_money"`
}

type TYQPTransferConfirmResp struct {
	Code    string  `json:"code"`
	Status  int     `json:"status"`
	Balance float64 `json:"balance"`
}

type TYQPPlaybackResp struct {
	Code        string `json:"code"`
	Info        string `json:"info"`
	DeskUuidUrl string `json:"desk_uuid_url"`
}
