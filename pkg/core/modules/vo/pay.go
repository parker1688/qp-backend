package vo

type YRPayUpdateResp struct {
	RetCode string `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type YRPayUpdateReq struct {
	MchId      int64  `json:"mchId"`
	AppId      string `json:"appId"`
	ProductId  int    `json:"productId"`
	Status     int    `json:"status"`
	Title      string `json:"title"`
	OrderBy    int    `json:"orderBy"`
	VipLevel   string `json:"vipLevel"`
	Device     string `json:"device"`
	IsSdk      int    `json:"isSdk"`
	IsSuccess  int    `json:"isSuccess"`
	AmountList string `json:"amountList"`
	Sign       string `json:"sign"`
}

type YRPayQueryProductListResp struct {
	RetCode string `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Data    string `json:"data"`
}

type YRPayQueryProductListReq struct {
	MchId int64  `json:"mchId"`
	AppId string `json:"appId"`
	Sign  string `json:"sign"`
}
