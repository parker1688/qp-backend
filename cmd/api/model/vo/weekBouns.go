package vo

type UserLevelWeekSettingResp struct {
	Level          int     `gorm:"column:level" json:"level" form:"level" uri:"level" `                                             // 等级
	ValidBetAmount float64 `gorm:"column:valid_bet_amount" json:"valid_bet_amount" form:"valid_bet_amount" uri:"valid_bet_amount" ` // 有效投注额
	LevelBouns     float64 `gorm:"column:level_bouns" json:"level_bouns" form:"level_bouns" uri:"level_bouns" `                     // 等级礼金
	WeekBouns      float64 `gorm:"column:week_bouns" json:"week_bouns" form:"week_bouns" uri:"week_bouns" `                         // 每周收益
	TotalBouns     float64 `gorm:"column:totalBouns" json:"totalBouns" form:"totalBouns" uri:"totalBouns" `                         // 每周收益
}
