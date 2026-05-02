package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVipRank struct {
	BaseDos
	Rank                     string             `gorm:"column:rank" json:"rank" form:"rank" uri:"rank" `                                                                                                 // 段位名称
	RankFlag                 string             `gorm:"column:rank_flag" json:"rank_flag" form:"rank_flag" uri:"rank_flag" `                                                                             // 段位英文表示
	StarNum                  int                `gorm:"column:star_num" json:"star_num" form:"star_num" uri:"star_num" `                                                                                 // 最高星数
	UpBonus                  float64            `gorm:"column:up_bonus" json:"up_bonus" form:"up_bonus" uri:"up_bonus" `                                                                                 // 晋级彩金
	HelpRate                 float64            `gorm:"column:help_rate" json:"help_rate" form:"help_rate" uri:"help_rate" `                                                                             // 救援比例
	HelpMaxAmount            float64            `gorm:"column:help_max_amount" json:"help_max_amount" form:"help_max_amount" uri:"help_max_amount" `                                                     // 每日救援金最大值
	ElectRebate              float64            `gorm:"column:elect_rebate" json:"elect_rebate" form:"elect_rebate" uri:"elect_rebate" `                                                                 // 电子返水百分比
	OtherRebate              float64            `gorm:"column:other_rebate" json:"other_rebate" form:"other_rebate" uri:"other_rebate" `                                                                 // 其他返水百分比
	EveryDepositRebate       float64            `gorm:"column:every_deposit_rebate" json:"every_deposit_rebate" form:"every_deposit_rebate" uri:"every_deposit_rebate" `                                 // 笔笔存百分比
	WeekFirstDepositRebate   float64            `gorm:"column:week_first_deposit_rebate" json:"week_first_deposit_rebate" form:"week_first_deposit_rebate" uri:"week_first_deposit_rebate" `             // 首存百分比
	WeekFirstDepositMaxBonus float64            `gorm:"column:week_first_deposit_max_bonus" json:"week_first_deposit_max_bonus" form:"week_first_deposit_max_bonus" uri:"week_first_deposit_max_bonus" ` // 周首存奖金上限
	YearBouns                float64            `gorm:"column:year_bouns" json:"year_bouns" form:"year_bouns" uri:"year_bouns" `                                                                         // 周年礼
	DayCouponMax             int                `gorm:"column:day_coupon_max" json:"day_coupon_max" form:"day_coupon_max" uri:"day_coupon_max" `                                                         // 盛世宝卷每日领取
	MerchantCode             string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                                             // 商户code
	CreateTime               automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                                        // 创建时间
	CreateBy                 string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                                             // 创建人
	UpdateTime               automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                                        // 修改时间
	UpdateBy                 string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                                             // 修改人
	Sort                     int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                                                                                 // 排序 升序
}

func (FcVipRank) TableName() string {
	return "fc_vip_rank"
}
