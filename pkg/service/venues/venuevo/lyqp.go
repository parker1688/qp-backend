package venuevo

type LYQPCommonResp struct {
	S int    `json:"s"`
	M string `json:"m"`
}

type LYQPBalanceResp struct {
	LYQPCommonResp
	D LYQPBalanceData `json:"d"`
}

type LYQPBalanceData struct {
	Code       int     `json:"code"`
	TotalMoney float64 `json:"totalMoney"` // 总分数
	FreeMoney  float64 `json:"freeMoney"`  // 可下分余额
	Status     int     `json:"status"`     // 在线状态（0、不在线,1、 在线）
	GameStatus int     `json:"gameStatus"` // 游戏状态（0、不在游戏中,1、游戏 中）
}

type LYQPDepositResp struct {
	LYQPCommonResp
	D LYQPDepositData `json:"d"`
}

type LYQPDepositData struct {
	Code     int     `json:"code"`
	Money    float64 `json:"money"`
	Currency string  `json:"currency"`
}

type LYQPWithdrawResp struct {
	LYQPCommonResp
	D LYQPWithdrawData `json:"d"`
}

type LYQPWithdrawData struct {
	Code     int     `json:"code"`
	Money    float64 `json:"money"`
	Currency string  `json:"currency"`
}

type LYQPLaunchGameResp struct {
	LYQPCommonResp
	D LYQPLaunchGameData `json:"d"`
}

type LYQPLaunchGameData struct {
	Code int    `json:"code"`
	Url  string `json:"url"`
}

type LYQPTransferConfirmResp struct {
	LYQPCommonResp
	D LYQPTransferConfirmData `json:"d"`
}

type LYQPTransferConfirmData struct {
	Code     int     `json:"code"`
	Money    float64 `json:"money"`
	Status   int     `json:"status"`
	Currency string  `json:"currency"`
}
