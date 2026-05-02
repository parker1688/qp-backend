package venuevo

type WUGDZResp struct {
	StatusCode int    `json:"status_code"`
	Balance    int    `json:"balance"` // 分为单位
	Message    string `json:"message"`
}

type WUGDZTransferConfirmResp struct {
	TxnId    int    `json:"txn_id"`
	Balance  int    `json:"balance"` // 分为单位
	MemberId int    `json:"member_id"`
	Note     string `json:"note"`
}
