package vo

type ImTokenValidateResp struct {
	ValidateSFTokenResult int    `json:"ValidateSFTokenResult"`
	MemberCode            string `json:"memberCode"`
	Currency              string `json:"currency"`
	IPAddress             string `json:"ipAddress"`
	StatusCode            int    `json:"statusCode"`
	StatusDesc            string `json:"statusDesc"`
}

type ImLoginResp struct {
	StatusCode int    `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}

type ImGetBalanceReq struct {
	ActionID   int    `json:"ActionId"`
	MemberCode string `json:"MemberCode"`
}

type ImGetBalanceResp struct {
	PackageID     int     `json:"PackageId"`
	Balance       float64 `json:"Balance"`
	DateReceived  string  `json:"DateReceived"`
	DateSent      string  `json:"DateSent"`
	StatusCode    int     `json:"StatusCode"`
	StatusMessage string  `json:"StatusMessage"`
}

type GetEventInfoMBTReq struct {
	SportId  int `json:"SportId"`
	Market   int `json:"Market"`
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
}
