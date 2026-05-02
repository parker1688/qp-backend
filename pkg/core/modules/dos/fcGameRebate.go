package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcGameRebate struct {
	BaseDos
	GameType     string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                      // 游戏类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	BonusRate    float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                  // 返水比例
	MinBetAmount float64            `gorm:"column:min_bet_amount" json:"min_bet_amount" form:"min_bet_amount" uri:"min_bet_amount" `  // 最少流水
	MaxBetAmount float64            `gorm:"column:max_bet_amount" json:"max_bet_amount" form:"max_bet_amount" uri:"max_bet_amount" `  // 最大流水
	Describe     string             `gorm:"column:describe" json:"describe" form:"describe" uri:"describe" `                          // 描述
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcGameRebate) TableName() string {
	return "fc_game_rebate"
}

type FcGameRebateRes struct {
	GameType     string  `gorm:"column:game_type" json:"game_type"`           // 游戏类型
	BonusRate    float64 `gorm:"column:bonus_rate" json:"bonus_rate"`         // 返水比例
	MinBetAmount float64 `gorm:"column:min_bet_amount" json:"min_bet_amount"` // 最少流水
	MaxBetAmount float64 `gorm:"column:max_bet_amount" json:"max_bet_amount"` // 最大流水
}
