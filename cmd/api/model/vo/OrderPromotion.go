package vo

import "bootpkg/common/expands/automaticType"

type OrderPromotionReq struct {
	Currency  string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
	StartTime string `json:"start_time"`                               //开始时间
	EndTime   string `json:"end_time"`                                 //结束时间
	PageSize  int    `json:"page_size"`                                //当前页大小
	PageIndex int    `json:"page_index"`                               //当前页码
}

type OrderPromotionResp struct {
	OrderSn    string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                          // 订单号
	Amount     float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 实际派发奖金额
	Status     int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0 待处理 2 拒绝  3 通过
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	Currency   string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `
	ApplyType  int                `gorm:"column:apply_type" json:"apply_type" form:"apply_type" uri:"apply_type" ` // 货币类型
}
