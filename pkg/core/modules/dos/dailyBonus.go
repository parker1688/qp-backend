package dos

import (
	"bootpkg/common/expands/automaticType"
)

type DailyBonus struct {
	BaseDos
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Bonus        string             `gorm:"column:bonus" json:"bonus" form:"bonus" uri:"bonus" `                                      // 奖励
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
}

func (DailyBonus) TableName() string {
	return "daily_bonus"
}

type DailyBonusField struct {
	Days        int     `json:"days"`
	Amount      float64 `json:"amount"`
	Reward      float64 `json:"reward"`
	ExtraReward float64 `json:"extraReward"`
	Icon        string  `json:"icon"`
}

type DailyBonusData struct {
	Day     int     `json:"day"`     // 奖励计数天数
	ExtDay  int     `json:"extDay"`  // 额外奖励计数天数
	Amount  float64 `json:"amount"`  // 累计金额（保留）
	Status  int     `json:"status"`  // 状态（0 不可领取；1 可领取；2 已领取）
	Date    string  `json:"date"`    // 日期(年月日，用于签到累计用途)
	ExtDate string  `json:"extDate"` // 额外日期（年月，用于月累计重置用途）
}

type DailyBonusResult struct {
	DailyBonusField
}

type DailyBonusResp struct {
	RewardList      []DailyBonusResult `json:"rewardList"`
	ExtraRewardList []DailyBonusResult `json:"extraRewardList"`
	Days            int                `json:"days"`        // 基础累计天数
	ExtDays         int                `json:"extDays"`     // 额外累计天数
	Status          int                `json:"status"`      // 是否可领取
	TotalAmount     float64            `json:"totalAmount"` // 今日打码金额
}
