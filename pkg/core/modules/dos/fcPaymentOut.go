package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcPaymentOut struct {
	BaseDos
	PaymentName  string             `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `          // 通道名称
	PaymentCode  string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `          // 通道code
	ChannelName  string             `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `          // 渠道名称
	ChannelCode  string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1:正常  2:禁止
	MinLevel     int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小VIP等级
	MaxLevel     int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大vip等级
	MinAmount    float64            `gorm:"column:min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                  // 单笔最低金额
	MaxAmount    float64            `gorm:"column:max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                  // 单笔最大金额
	DayMaxAmount float64            `gorm:"column:day_max_amount" json:"day_max_amount" form:"day_max_amount" uri:"day_max_amount" `  // 每日最大充值金额
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	FeeRate      float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                          // 手续费比率
	Icon         string             `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	ThirdCode    string             `gorm:"column:third_code" json:"third_code" form:"third_code" uri:"third_code" `
}

func (FcPaymentOut) TableName() string {
	return "fc_payment_out"
}
