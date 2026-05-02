package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserRebateRecords struct {
	BaseDos
	UserId         string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                     // 用户Id
	UserName       string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                             // 用户账号
	Level          int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                             // 用户vip等级
	VenueCode      string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                         // 场馆code
	GameType       string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                             // 游戏类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	BetAmount      float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                         // 有效投注
	HisBetAmount   float64            `gorm:"column:his_bet_amount" json:"his_bet_amount" form:"his_bet_amount" uri:"his_bet_amount" `         // 历史累计流水
	BatchBetAmount float64            `gorm:"column:batch_bet_amount" json:"batch_bet_amount" form:"batch_bet_amount" uri:"batch_bet_amount" ` // 批次流水
	RebateType     int                `gorm:"column:rebate_type" json:"rebate_type" form:"rebate_type" uri:"rebate_type" `                     // 发放类型 1: 系统发放  2:手动发放
	Status         int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                         // 状态 1: 已完成  2:发放中 3: 失败
	BonusAmount    float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                 // 奖金金额
	BonusRate      float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                         // 返水比例
	Remarks        string             `gorm:"column:remarks" json:"remarks" form:"remarks" uri:"remarks" `                                     // 备注
	MerchantCode   string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy       string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                             // 创建人
	UpdateTime     automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy       string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                             // 修改人
}

func (FcUserRebateRecords) TableName() string {
	return "fc_user_rebate_records"
}

type FcUserRebateRecordsListRow struct {
	VenueCode   string             `gorm:"column:venue_code" json:"venue_code"`                // 场馆code
	GameType    string             `gorm:"column:game_type" json:"game_type"`                  // 游戏类型
	BonusAmount float64            `gorm:"column:bonus_amount" json:"bonus_amount"`            // 返水金额
	CreateTime  automaticType.Time `gorm:"column:create_time;default:null" json:"create_time"` // 领取时间
}

type FcUserRebateRecordsDetailListResp struct {
	BaseDos
	VenueCode   string  `gorm:"column:venue_code" json:"venue_code"`     // 场馆code
	GameType    string  `gorm:"column:game_type" json:"game_type"`       // 游戏类型
	BetAmount   float64 `gorm:"column:bet_amount" json:"bet_amount"`     // 有效投注
	BonusAmount float64 `gorm:"column:bonus_amount" json:"bonus_amount"` // 返水金额
	BonusRate   float64 `gorm:"column:bonus_rate" json:"bonus_rate"`     // 返水比例
}
