package venuevo

type WALICreateResp struct {
	Code int            `json:"error"`
	Msg  string         `json:"msg"`
	Data WALICreateData `json:"data"`
}

type WALICreateData struct {
	Text string `json:"text"`
}

type WALIBalanceResp struct {
	Code int             `json:"error"`
	Msg  string          `json:"msg"`
	Data WALIBalanceData `json:"data"`
}

type WALIBalanceData struct {
	Status       int    `json:"status"`
	Balance      string `json:"balance"`
	Transferable string `json:"transferable"`
}

type WALIEnterGameResp struct {
	Code int           `json:"error"`
	Msg  string        `json:"msg"`
	Data WALIEnterGame `json:"data"`
}

type WALIEnterGame struct {
	GameUrl    string `json:"gameUrl"`
	GameReason string `json:"gameReason"`
}

type WALITransferResp struct {
	Code int          `json:"error"`
	Msg  string       `json:"msg"`
	Data WALITransfer `json:"data"`
}

type WALITransfer struct {
	OrderId string `json:"orderId"`
	Status  int    `json:"status"`
	Reason  string `json:"reason"`
	Balance string `json:"balance"`
}
