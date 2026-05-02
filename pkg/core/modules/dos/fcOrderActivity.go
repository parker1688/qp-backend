package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcOrderActivity struct {
	BaseDos
	ActivityType int                `gorm:"column:activity_type" json:"activity_type" form:"activity_type" uri:"activity_type" `      // 1 存款活动  2  提款活动
	ActivityName string             `gorm:"column:activity_name" json:"activity_name" form:"activity_name" uri:"activity_name" `      // 活动名称
	Rate         float64            `gorm:"column:rate" json:"rate" form:"rate" uri:"rate" `                                          // 比率百分比
	MaxAmount    float64            `gorm:"column:max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                  // 最高奖励金额
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
	Currency     string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1 上线 2 下线
	TurnOver     int                `gorm:"column:turn_over" json:"turn_over" form:"turn_over" uri:"turn_over" `                      // 流水倍數

}

func (FcOrderActivity) TableName() string {
	return "fc_order_activity"
}
