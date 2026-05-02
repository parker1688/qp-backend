package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserGameRebate struct {
	BaseDos
	Day              automaticType.Time `gorm:"column:day" json:"day" form:"day" uri:"day" `
	UserId           string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                             // 用户Id
	UserName         string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                                     // 用户账号
	GameType         string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                                     // chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	NetAmount        float64            `gorm:"column:net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `                                 // 总输赢
	BetAmount        float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                                 // 总投注
	ValidBetamount   float64            `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `             // 总有效投注
	MerchantCode     string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                     // 商户code
	CreateTime       automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy         string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime       automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy         string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
	BonusAmount      float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                         // 奖金金额
	BonusAmountIssue float64            `gorm:"column:bonus_amount_issue" json:"bonus_amount_issue" form:"bonus_amount_issue" uri:"bonus_amount_issue" ` // 已发奖金金额
	BonusRate        float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                                 // 返水比例
}

func (FcUserGameRebate) TableName() string {
	return "fc_user_game_rebate"
}
