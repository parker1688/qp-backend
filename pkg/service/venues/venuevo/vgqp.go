package venuevo

type VGQPCommonResp struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

type VGQPCommonVResp struct {
	VGQPCommonResp
	Data VGQPCommonDataResp `json:"value"`
}

type VGQPCommonDataResp struct {
	Username string  `json:"Username"`
	Channel  string  `json:"Channel"`
	Money    float64 `json:"Money"`   // 用户余额, 单位为元，精确到小数点 2 位
	Balance  float64 `json:"Balance"` // 用户余额, 单位为分
	InGame   int     `json:"InGame"`
	GameType int     `json:"GameType"`
}

type VGQPLoginGameResp struct {
	VGQPCommonResp
	Value string `json:"value"`
}

type VGQPLoginResp struct {
	VGQPCommonResp
	Value string `json:"value"`
}
