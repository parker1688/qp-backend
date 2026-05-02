package venuevo

type CQ9Resp struct {
	StatusCode int    `json:"status_code"`
	Balance    int    `json:"balance"` // 分为单位
	Message    string `json:"message"`
}

type CQ9CreateUserResp struct {
	Data struct {
		Account  string `json:"account"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	} `json:"data"`
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Datetime  string `json:"datetime"`
		TraceCode string `json:"traceCode"`
	} `json:"status"` // 分为单位
}

type CQ9BalanceResp struct {
	Data struct {
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
	} `json:"data"`
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Datetime  string `json:"datetime"`
		TraceCode string `json:"traceCode"`
	} `json:"status"` // 分为单位
}

type CQ9LoginResp struct {
	Data struct {
		UserToken string `json:"usertoken"`
	} `json:"data"`
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Datetime  string `json:"datetime"`
		TraceCode string `json:"traceCode"`
	} `json:"status"` // 分为单位
}

type CQ9Login2Resp struct {
	Data struct {
		Url   string `json:"url"`
		Token string `json:"token"`
	} `json:"data"`
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Datetime  string `json:"datetime"`
		TraceCode string `json:"traceCode"`
	} `json:"status"` // 分为单位
}

type CQ9DepositResp struct {
	Data struct {
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
	} `json:"data"`
	Status struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Datetime  string `json:"datetime"`
		TraceCode string `json:"traceCode"`
	} `json:"status"` // 分为单位
}

type CQ9TransferConfirmResp struct {
	TxnId    int    `json:"txn_id"`
	Balance  int    `json:"balance"` // 分为单位
	MemberId int    `json:"member_id"`
	Note     string `json:"note"`
}
