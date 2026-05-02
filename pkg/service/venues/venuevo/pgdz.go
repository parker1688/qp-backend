package venuevo

type PGDZResp struct {
	Error string       `json:"error"`
	Data  PGDZRespSub1 `json:"data"`
}

type PGDZRespSub1 struct {
	Pid int64 `json:"Pid"`
}

type PGDZLoginResp struct {
	Code  int       `json:"code"`
	Error string    `json:"error"`
	Data  PGDZLogin `json:"data"`
}

type PGDZLogin struct {
	Url string `json:"Url"`
}

type PGDZBalanceResp struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  PGDZBalance `json:"data"`
}

type PGDZBalance struct {
	Balance         float64 `json:"Balance"`
	Unsettled       float64 `json:"Unsettled"`
	UnsettledDetail string  `json:"UnsettledDetail"`
}

type PGDZDepositResp struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  PGDZDeposit `json:"data"`
}

type PGDZDeposit struct {
	AfterBalance float64 `json:"AfterBalance"`
}

type PGDZWithdrawResp struct {
	Code  int          `json:"code"`
	Error string       `json:"error"`
	Data  PGDZWithdraw `json:"data"`
}

type PGDZWithdraw struct {
	AfterBalance float64 `json:"AfterBalance"`
}

type PGDZTransferConfirmResp struct {
	TxnId    int    `json:"txn_id"`
	Balance  int    `json:"balance"` // 分为单位
	MemberId int    `json:"member_id"`
	Note     string `json:"note"`
}
