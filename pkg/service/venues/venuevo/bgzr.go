package venuevo

type BGZRResp struct {
	StatusCode int    `json:"status_code"`
	Balance    int    `json:"balance"` // 分为单位
	Message    string `json:"message"`
}

type BGZRCreateUserResp struct {
	Id     string `json:"id"`
	Result struct {
		LoginId  string `json:"loginId"`
		Success  bool   `json:"success"`
		SnType   int    `json:"snType"`
		NickName string `json:"nickname"`
		RegType  string `json:"regType"`
		UserId   int64  `json:"userId"`
	} `json:"result"` // 分为单位
	Error   string `json:"error"`
	Jsonrpc string `json:"jsonrpc"`
}

type BGZRBalanceResp struct {
	Result float64 `json:"result"` // 分为单位
}

type BGZRLoginResp struct {
	Result string `json:"result"` // 分为单位
}

type BGZRTransferResp struct {
	Result float64 `json:"result"` // 分为单位
	Error  struct {
		Code    string `json:"code"`
		Sn      string `json:"sn"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}

type BGZRTransferConfirmResp struct {
	TxnId    int    `json:"txn_id"`
	Balance  int    `json:"balance"` // 分为单位
	MemberId int    `json:"member_id"`
	Note     string `json:"note"`
}
