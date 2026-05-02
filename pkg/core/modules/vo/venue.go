package vo

const (
	Login_SUCCESS_CODE = 0
	Login_FAIL_CODE    = -1

	Playback_SUCCESS_CODE = 0
	Playback_FAIL_CODE    = -1
)

type VenueLoginRequest struct {
	UserName string //用户名
	Password string //账号密码
}

type VenueLoginResponse struct {
	Code int                    `json:"code"` //0 成功
	Msg  string                 `json:"msg"`  //错误信息
	Data VenueLoginDataResponse `json:"data"`
}

type VenueLoginDataResponse struct {
	Token string `json:"token"`
}

type VenueGetOrderSnReq struct {
	VenueCode string
	UserName  string
}

type VenuePlaybackReq struct {
	Id        string `json:"id"` // 游戏记录的 id
	VenueCode string `json:"venue_code"`
	TableId   string `json:"table_id"` // 牌桌号
}

type VenuePlaybackResp struct {
	Code int                   `json:"code"` //0 成功
	Msg  string                `json:"msg"`  //错误信息
	Data VenuePlaybackDataResp `json:"data"`
}

type VenuePlaybackDataResp struct {
	PlaybackUrl string `json:"playbackUrl"`
}

type UserGameTypeData struct {
	UserId         string  `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                 // 用户Id
	UserName       string  `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                         // 用户账号
	GameType       string  `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                         // 游戏类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	BetAmount      float64 `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                     // 投注量
	ValidBetamount float64 `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" ` // 有效投注量
	NetAmount      float64 `gorm:"column:net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `                     // 输赢
	MerchantCode   string  `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `         // 商户code
}
