package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcOrderPromotion struct {
	BaseDos
	OrderSn      string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                          // 订单号
	ApplyAmount  float64            `gorm:"column:apply_amount" json:"apply_amount" form:"apply_amount" uri:"apply_amount" `          // 申请金额
	AppleRate    float64            `gorm:"column:apple_rate" json:"apple_rate" form:"apple_rate" uri:"apple_rate" `                  // 计算比例(百分比)
	ApplyType    int                `gorm:"column:apply_type" json:"apply_type" form:"apply_type" uri:"apply_type" `                  // 申请类型 1. 存款活动 2. 好友邀请 3 红包奖金 4 生日礼金 5 邀请好友
	Amount       float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 实际派发奖金额
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0 待处理 2 拒绝  3 通过
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户名称
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                              // 用户id
	TurnOver     int                `gorm:"column:turn_over" json:"turn_over" form:"turn_over" uri:"turn_over" `                      // 流水倍数
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Remake       string             `gorm:"column:remake" json:"remake" form:"remake" uri:"remake" `                                  // 备注
	Currency     string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                 // 货币类型
}

func (FcOrderPromotion) TableName() string {
	return "fc_order_promotion"
}
