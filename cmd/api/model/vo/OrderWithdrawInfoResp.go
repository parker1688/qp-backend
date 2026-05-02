package vo

import "bootpkg/common/expands/automaticType"

type OrderWithdrawInfoResp struct {
	Amount         float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                     // 提款金额
	Status         int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                     // 0:  待审核 1:等待出款 2:拒绝 3:已完成 4:取消订单
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `    // 创建时间
	Fee            float64            `gorm:"column:fee" json:"fee" form:"fee" uri:"fee" `                                                 // 手续费
	OrderSn        string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                             // 订单编号
	Currency       string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                             // 币种简码
	CallbackRemark string             `gorm:"column:callback_remark" json:"callback_remark" form:"callback_remark" uri:"callback_remark" ` // 状态备注
	OrderType      int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `
}
