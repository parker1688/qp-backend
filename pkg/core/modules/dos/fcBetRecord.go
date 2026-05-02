package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcBetRecord struct {
	BaseDos
	UserId              string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                 // 用户ID
	UserName            string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `         // 用户账号
	Account             string             `gorm:"column:account" json:"account" form:"account" uri:"account" `                 // 游戏账号
	PlayerName          string             `gorm:"column:player_name" json:"player_name" form:"player_name" uri:"player_name" ` // 三方游戏code
	OrderSn             string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `             // 三方订单号
	GameCode            string             `gorm:"column:game_code" json:"game_code" form:"game_code" uri:"game_code" `         // 三方游戏code
	BetAmount           float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `
	NetAmount           float64            `gorm:"column:net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `
	ValidBetamount      float64            `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `
	BetTime             automaticType.Time `gorm:"column:bet_time" json:"bet_time" form:"bet_time" uri:"bet_time" `                                                 // 投注时间
	SettlementTime      automaticType.Time `gorm:"column:settlement_time" json:"settlement_time" form:"settlement_time" uri:"settlement_time" `                     // 结算时间
	ThirdBettime        automaticType.Time `gorm:"column:third_bettime" json:"third_bettime" form:"third_bettime" uri:"third_bettime" `                             // 三方投注时间
	ThirdSettlementtime automaticType.Time `gorm:"column:third_settlementtime" json:"third_settlementtime" form:"third_settlementtime" uri:"third_settlementtime" ` // 三方结算时间
	GameType            string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                                             // 1 体育 2 真人 3 棋牌 4 电子 5 捕鱼 6  彩票  7  区块链 8:电竞
	GameName            string             `gorm:"column:game_name" json:"game_name" form:"game_name" uri:"game_name" `                                             // 游戏名称
	TableId             string             `gorm:"table_id" json:"table_id" form:"table_id" uri:"table_id" `                                                        // 牌桌号
	VenueCode           string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                                         // 场馆code
	Odds                float64            `gorm:"column:odds" json:"odds" form:"odds" uri:"odds" `                                                                 // 赔率
	OddsType            int                `gorm:"column:odds_type" json:"odds_type" form:"odds_type" uri:"odds_type" `                                             // 赔率类型
	Version             int64              `gorm:"column:version" json:"version" form:"version" uri:"version" `                                                     // 版本
	CreateTime          automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                        // 创建时间
	CreateBy            string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                             // 创建人
	UpdateTime          automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                        // 修改时间
	UpdateBy            string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                             // 修改人
	MerchantCode        string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                             // 商户code
	MerchantNetAmount   float64            `gorm:"column:merchant_net_amount" json:"merchant_net_amount" form:"merchant_net_amount" uri:"merchant_net_amount" `     // 商家输赢
	AfterBalance        float64            `gorm:"column:after_balance" json:"after_balance" form:"after_balance" uri:"after_balance" `                             // 游戏后金额
	IsSettled           int                `gorm:"column:is_settled" json:"is_settled" form:"is_settled" uri:"is_settled" `                                         //是否结算
}

func (FcBetRecord) TableName() string {
	return "fc_bet_record"
}

type FcBetRecordExt struct {
	FcBetRecord
	Merchant FcMerchant `json:"merchant" gorm:"foreignkey:MerchantCode;references:MerchantCode"`
}

type FcBetRecordListResp struct {
	RecordList          []*FcBetRecordResp `json:"list"`
	TotalBetTime        int64              `json:"total_bet_time"`
	TotalBetAmount      float64            `json:"total_bet_amount"`
	TotalNetAmount      float64            `json:"total_net_amount"`
	TotalValidBetAmount float64            `json:"total_valid_betamount"`
}
type FcBetRecordResp struct {
	FcBetRecord
	MerchantName string `json:"merchant_name"`
}
