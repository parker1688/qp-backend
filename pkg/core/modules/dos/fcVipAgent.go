package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVipAgent struct {
	BaseDos
	VipName               string             `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" `                                                             // VIP1~VIP10
	Level                 int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                                                         // 层级 1~10
	MinRecharegeAmount    float64            `gorm:"column:min_recharege_amount" json:"min_recharege_amount" form:"min_recharege_amount" uri:"min_recharege_amount" `             // 累计最小充值金额
	MinBetAmount          float64            `gorm:"column:min_bet_amount" json:"min_bet_amount" form:"min_bet_amount" uri:"min_bet_amount" `                                     // 流水要求
	RelegationBetAmount   float64            `gorm:"column:relegation_bet_amount" json:"relegation_bet_amount" form:"relegation_bet_amount" uri:"relegation_bet_amount" `         // 保级流水金额
	DailyWithdrawalTimes  int                `gorm:"column:daily_withdrawal_times" json:"daily_withdrawal_times" form:"daily_withdrawal_times" uri:"daily_withdrawal_times" `     // 每日提现次数
	DailyWithdrawalAmount float64            `gorm:"column:daily_withdrawal_amount" json:"daily_withdrawal_amount" form:"daily_withdrawal_amount" uri:"daily_withdrawal_amount" ` // 每日提现额度
	WithdrawalFee         float64            `gorm:"column:withdrawal_fee" json:"withdrawal_fee" form:"withdrawal_fee" uri:"withdrawal_fee" `                                     // 提款手续费
	UpgradeGift           float64            `gorm:"column:upgrade_gift" json:"upgrade_gift" form:"upgrade_gift" uri:"upgrade_gift" `                                             // 升级礼金
	BirthdayGift          float64            `gorm:"column:birthday_gift" json:"birthday_gift" form:"birthday_gift" uri:"birthday_gift" `                                         // 生日礼金
	MonthlyGift           float64            `gorm:"column:monthly_gift" json:"monthly_gift" form:"monthly_gift" uri:"monthly_gift" `                                             // 每月礼金
	MerchantCode          string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                         // 商户code
	CreateTime            automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                    // 创建时间
	CreateBy              string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                         // 创建人
	UpdateTime            automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                    // 修改时间
	UpdateBy              string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                         // 修改人
	MinRechargeAmount     float64            `gorm:"column:min_recharge_amount" json:"min_recharge_amount" form:"min_recharge_amount" uri:"min_recharge_amount" `                 // 最小充值金额
	MinWithdrawAmount     float64            `gorm:"column:min_withdraw_amount" json:"min_withdraw_amount" form:"min_withdraw_amount" uri:"min_withdraw_amount" `                 // 最小提款金额
}

func (FcVipAgent) TableName() string {
	return "fc_vip_agent"
}
