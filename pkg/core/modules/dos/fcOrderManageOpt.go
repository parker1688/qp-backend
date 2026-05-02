package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcOrderManageOpt struct {
	BaseDos
	Currency          string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                  // 币种
	FlowMultiple      int                `gorm:"column:flow_multiple" json:"flow_multiple" form:"flow_multiple" uri:"flow_multiple" `              // 打码倍数
	Amount            float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                          // 操作金额
	TranscationBefore float64            `gorm:"transcation_before" json:"transcation_before" form:"transcation_before" uri:"transcation_before" ` // 账变前金额
	TranscationAfter  float64            `gorm:"transcation_after" json:"transcation_after" form:"transcation_after" uri:"transcation_after" `     // 账变后金额
	FlowAmount        float64            `gorm:"column:flow_amount" json:"flow_amount" form:"flow_amount" uri:"flow_amount" `                      // 打码流水
	Status            int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                          // 状态 1 完成  2 失败
	TrsType           int                `gorm:"column:trs_type" json:"trs_type" form:"trs_type" uri:"trs_type" `                                  // 交易类型 1 官方赠送 2 客服补偿 3 人工汇款-微信 4 人工汇款-支付宝 5 人工汇款-银行卡 6 人工汇款-U 7 人工汇款-数字人民币 8 一般扣除 9 福利扣除
	ScoreType         int                `gorm:"column:score_type" json:"score_type" form:"score_type" uri:"score_type" `                          // 分类型 1 上分  2 下分
	CreateTime        automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `         // 创建时间
	CreateBy          string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime        automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `         // 修改时间
	UpdateBy          string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                              // 修改人
	MerchantCode      string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `              // 商户code
	UserId            string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                      // 用户ID
	UserName          string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                              // 用户名称
	Remarks           string             `gorm:"column:remarks" json:"remarks" form:"remarks" uri:"remarks" `                                      // 备注
}

func (FcOrderManageOpt) TableName() string {
	return "fc_order_manage_opt"
}
