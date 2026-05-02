package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserLevelWeekSetting struct {
	BaseDos
	Level          int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                             // 等级
	BetType        int                `gorm:"column:bet_type" json:"bet_type" form:"bet_type" uri:"bet_type" `                                 // 1: 体育电竞 2：棋牌  3：彩票
	ValidBetAmount float64            `gorm:"column:valid_bet_amount" json:"valid_bet_amount" form:"valid_bet_amount" uri:"valid_bet_amount" ` // 有效投注额
	LevelBouns     float64            `gorm:"column:level_bouns" json:"level_bouns" form:"level_bouns" uri:"level_bouns" `                     // 等级礼金
	WeekBouns      float64            `gorm:"column:week_bouns" json:"week_bouns" form:"week_bouns" uri:"week_bouns" `                         // 每周收益
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy       string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                             // 创建人
	UpdateTime     automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy       string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                             // 修改人
	MerchantCode   string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
}

func (FcUserLevelWeekSetting) TableName() string {
	return "fc_user_level_week_setting"
}
