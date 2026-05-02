package vo

import "bootpkg/common/expands/automaticType"

type BetRecordUnsetteledReq struct {
	VenueCode string `json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆code
	Status    int    `json:"status" form:"status" uri:"status" `             // 状态
	GameType  string `json:"game_type" form:"game_type" uri:"game_type" `    // 游戏类型
	DateStart string `json:"date_start" form:"date_start" uri:"date_start" ` // 投注开始日期
	DateEnd   string `json:"date_end" form:"date_end" uri:"date_end" `       // 投注结束日期
}

type BetRecordUnsettledResp struct {
	PlayerName     string             `gorm:"player_name" json:"player_name" form:"player_name" uri:"player_name" ` // 三方游戏名称
	OrderSn        string             `gorm:"order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `             // 三方订单号
	GameCode       string             `gorm:"game_code" json:"game_code" form:"game_code" uri:"game_code" `         // 三方游戏code
	BetAmount      float64            `gorm:"bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `
	NetAmount      float64            `gorm:"net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `
	ValidBetamount float64            `gorm:"valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `
	BetTime        automaticType.Time `gorm:"bet_time" json:"bet_time" form:"bet_time" uri:"bet_time" `                             // 投注时间
	SettlementTime automaticType.Time `gorm:"settlement_time" json:"settlement_time" form:"settlement_time" uri:"settlement_time" ` // 结算时间
	GameType       string             `gorm:"game_type" json:"game_type" form:"game_type" uri:"game_type" `
	GameName       string             `gorm:"game_name" json:"game_name" form:"game_name" uri:"game_name" `     // 游戏名称
	VenueCode      string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆code
}

type BetRecordStatisReq struct {
	DateStart string `json:"date_start" form:"date_start" uri:"date_start" ` // 投注开始日期
	DateEnd   string `json:"date_end" form:"date_end" uri:"date_end" `       // 投注结束日期
}

type BetRecordStatisResp struct {
	Data []BetRecordStatisData `json:"data" form:"data" uri:"data"`
}

type BetRecordStatisData struct {
	VenueCode string  `json:"venue_code" form:"venue_code" uri:"venue_code"`
	VenueName string  `json:"venue_name" form:"venue_name" uri:"venue_name"`
	GameType  string  `json:"game_type" form:"game_type" uri:"game_type"`
	Win       float64 `json:"net_amount" form:"net_amount" uri:"net_amount"`
	BetAmount float64 `json:"bet_amount" form:"bet_amount" uri:"bet_amount"`
	BetCount  int     `json:"bet_count" form:"bet_count" uri:"bet_count"`
	IsSettled int     `json:"is_settled" form:"is_settled" uri:"is_settled"`
}
