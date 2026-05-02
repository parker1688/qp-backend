package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcPayChannelOut struct {
	BaseDos
	ChannelName  string             `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `          // 渠道名称
	ChannelCode  string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
	Icon         string             `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `                                          // 渠道icon
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1:正常  2:禁止
	MinLevel     int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小VIP等级
	MaxLevel     int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大vip等级
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	Currency     string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code(暂时无用)
	MinAmount    float64            `gorm:"column:min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                  // 最小金额
	MaxAmount    float64            `gorm:"column:max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                  // 最大金额
	FeeRate      float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `
}

func (FcPayChannelOut) TableName() string {
	return "fc_pay_channel_out"
}
