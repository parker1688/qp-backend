package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcPaymentSum struct {
	BaseDos
	PaymentName  string             `gorm:"payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `          // 通道名称
	PayAliasName string             `gorm:"pay_alias_name" json:"pay_alias_name" form:"pay_alias_name" uri:"pay_alias_name" `  // 通道别名
	PaymentCode  string             `gorm:"payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `          // 通道code
	PayId        string             `gorm:"pay_id" json:"pay_id" form:"pay_id" uri:"pay_id" `                                  // 通道产品id
	ChannelName  string             `gorm:"channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `          // 渠道名称
	ChannelCode  string             `gorm:"channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
	Status       int                `gorm:"status" json:"status" form:"status" uri:"status" `                                  // 1:正常  2:禁止
	MinLevel     int                `gorm:"min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小VIP等级
	MaxLevel     int                `gorm:"max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大vip等级
	MinAmount    float64            `gorm:"min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                  // 单笔最低金额
	MaxAmount    float64            `gorm:"max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                  // 单笔最大金额
	DayMaxAmount float64            `gorm:"day_max_amount" json:"day_max_amount" form:"day_max_amount" uri:"day_max_amount" `  // 每日最大充值金额
	BonusRate    float64            `gorm:"bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                  // 优惠比例
	Sort         int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	FeeRate      float64            `gorm:"fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                          // 手续费比列
	AmountRange  string             `gorm:"amount_range" json:"amount_range" form:"amount_range" uri:"amount_range" `          // 固定金额
	Remark       string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `                                  // 备注
}

func (FcPaymentSum) TableName() string {
	return "fc_payment_sum"
}
