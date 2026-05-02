package venuevo

type SBTYUserResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type SBTYBalanceResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      []struct {
		VendorMemberID string  `json:"vendor_member_id"`
		Balance        float64 `json:"balance"`
		Outstanding    float64 `json:"outstanding"`
		Currency       int     `json:"currency"`
		ErrorCode      int     `json:"error_code"`
	} `json:"Data"`
}

type SBTYDepositResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      struct {
		TransID      int     `json:"trans_id"`
		BeforeAmount float64 `json:"before_amount"`
		AfterAmount  float64 `json:"after_amount"`
		SystemID     string  `json:"system_id"`
		Status       int     `json:"status"`
	} `json:"Data"`
}

type SBTYWithdrawResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      struct {
		TransID      int     `json:"trans_id"`
		BeforeAmount float64 `json:"before_amount"`
		AfterAmount  float64 `json:"after_amount"`
		SystemID     string  `json:"system_id"`
		Status       int     `json:"status"`
	} `json:"Data"`
}

type SBTYLoginResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      string `json:"Data"`
}

type SBTYTransferConfirmResp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      struct {
		TransID      int     `json:"trans_id"`
		TransferDate string  `json:"transfer_date"`
		Amount       float64 `json:"amount"`
		Currency     int     `json:"currency"`
		BeforeAmount float64 `json:"before_amount"`
		AfterAmount  float64 `json:"after_amount"`
		Status       int     `json:"status"`
	} `json:"Data"`
}
