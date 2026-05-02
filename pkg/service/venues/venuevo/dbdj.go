package venuevo

type DBDJCreateUserResp struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type DBDJGetUserBalanceResp struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type DBDJTransferResp struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	Balance string `json:"balance"`
}

type DBDJLoginGameResp struct {
	Status string                `json:"status"`
	Data   DBDJLoginGameDataResp `json:"data"`
}

type DBDJLoginGameDataResp struct {
	H5    string `json:"h5"`
	PC    string `json:"pc"`
	Token string `json:"token"`
}

type DBDJTransferConfirmResp struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	Balance string `json:"balance"`
}
