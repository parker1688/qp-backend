package venuevo

type FGDZCommonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type FGDZCreateUserResp struct {
	FGDZCommonResp
	Data FGDZCreateUserDataResp `json:"data"`
}

type FGDZCreateUserDataResp struct {
	Openid string `json:"openid"`
}

type FGDZGetBalanceResp struct {
	FGDZCommonResp
	Data FGDZGetBalanceDataResp `json:"data"`
}

type FGDZGetBalanceDataResp struct {
	Balance  int    `json:"balance"`  // 单位为分
	Currency string `json:"currency"` // 单位为分
}

type FGDZLoginGameResp struct {
	FGDZCommonResp
	Data FGDZLoginGameDataResp `json:"data"`
}

type FGDZLoginGameDataResp struct {
	GameUrl string `json:"game_url"`
	Name    string `json:"name"`
	Meta    string `json:"meta"`
	Token   string `json:"token"`
}

type FGDZLoginLobbyResp struct {
	FGDZCommonResp
	Data FGDZLoginLobbyDataResp `json:"data"`
}

type FGDZLoginLobbyDataResp struct {
	LobbyUrl string `json:"lobby_url"`
}
