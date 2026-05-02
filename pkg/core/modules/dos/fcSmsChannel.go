package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSmsChannel struct {
	BaseDos
	SmsName      string             `gorm:"column:sms_name" json:"sms_name" form:"sms_name" uri:"sms_name" `                          // 通道名称
	SmsCode      string             `gorm:"column:sms_code" json:"sms_code" form:"sms_code" uri:"sms_code" `                          // 通道code
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1:正常  2:禁止
	MinLevel     int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小VIP等级
	MaxLevel     int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大vip等级
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	FeeRate      float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                          // 手续费比率
}

func (FcSmsChannel) TableName() string {
	return "fc_sms_channel"
}
