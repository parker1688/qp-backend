package venuevo

type VRCPCreateResp struct {
	Error int `json:"errorCode"`
}

type VRCPBalanceResp struct {
	PlayerName string        `json:"playerName"`
	Balance    float64       `json:"balance"`
	Games      []VRCPBalance `json:"games"`
}

type VRCPBalance struct {
	Type    int     `json:"type"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	Error   int     `json:"errorCode"`
}

type VRCPDepositResp struct {
	SerialNumber string  `json:"serialNumber"`
	PlayerName   string  `json:"playerName"`
	Type         int     `json:"type"`
	Amount       float64 `json:"amount"`
	CreateTime   string  `json:"createTime"`
	State        int     `json:"state"`
	Balance      float64 `json:"balance"`
}
