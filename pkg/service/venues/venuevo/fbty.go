package venuevo

type FbtyRegisterResp struct {
	Code    int64  `json:"code"`
	Data    int64  `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type FbtyTokenResp struct {
	Code int64 `json:"code"`
	Data struct {
		ServerInfo struct {
			APIServerAddress  string `json:"apiServerAddress"`
			H5Address         string `json:"h5Address"`
			PcAddress         string `json:"pcAddress"`
			PushServerAddress string `json:"pushServerAddress"`
			VirtualAddress    string `json:"virtualAddress"`
		} `json:"serverInfo"`
		ThemeBgColor string `json:"themeBgColor"`
		ThemeFgColor string `json:"themeFgColor"`
		Token        string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type FbtyTransferResp struct {
	Code    int64   `json:"code"`
	Data    float64 `json:"data"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
}

type FbtyBalanceResp struct {
	Code int64 `json:"code"`
	Data struct {
		Balance        string `json:"balance"`
		CurrencyID     int64  `json:"currencyId"`
		MerchantUserID string `json:"merchantUserId"`
		OddsLevel      int64  `json:"oddsLevel"`
		UserID         int64  `json:"userId"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type FbtyConfirmResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID                   int64   `json:"id"`
		UserID               int     `json:"userId"`
		MerchantUserID       string  `json:"merchantUserId"`
		BusinessID           string  `json:"businessId"`
		MerchantID           int64   `json:"merchantId"`
		TransferType         string  `json:"transferType"`
		Amount               float64 `json:"amount"`
		BeforeTransferAmount float64 `json:"beforeTransferAmount"`
		AfterTransferAmount  float64 `json:"afterTransferAmount"`
		Status               int     `json:"status"`
		CreateTime           int64   `json:"createTime"`
		CurrencyID           int     `json:"currencyId"`
	} `json:"data"`
	Code int `json:"code"`
}
