package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcOrderWithdrawPaymentOut struct {
	BaseDos
	OrderSn                  string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" ` // 订单号
	UserId                   string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName                 string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Amount                   float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 代付金额
	ApplyAmount              float64            `gorm:"column:apply_amount" json:"apply_amount" form:"apply_amount" uri:"apply_amount" `          // 提款申请金额
	Status                   int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0 下发中；1 打款中；2 打款失败；3 打款成功
	CreateTime               automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy                 string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime               automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy                 string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode             string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Currency                 string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种简码
	OrderType                int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `                  // 订单类型 1 银行卡 3 虚拟币
	FeeRate                  float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                          // 手续费比例(百分比)
	Fee                      float64            `gorm:"column:fee" json:"fee" form:"fee" uri:"fee" `                                              // 手续费
	ChannelId                string             `gorm:"column:channel_id" json:"channel_id" form:"channel_id" uri:"channel_id" `                  // 渠道Id
	ChannelCode              string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 渠道code
	PaymentId                string             `gorm:"column:payment_id" json:"payment_id" form:"payment_id" uri:"payment_id" `                  // 通道ID
	PaymentCode              string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `          // 通道code
	ThirdCode                string             `gorm:"column:third_code" json:"third_code" form:"third_code" uri:"third_code" `
	Remark                   string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                                                                     // 备注
	DepositWithdrawSubAmount float64            `gorm:"column:deposit_withdraw_sub_amount" json:"deposit_withdraw_sub_amount" form:"deposit_withdraw_sub_amount" uri:"deposit_withdraw_sub_amount" ` // 提存差
	WithdrawAmount           float64            `gorm:"column:withdraw_amount" json:"withdraw_amount" form:"withdraw_amount" uri:"withdraw_amount" `                                                 // 提存差
	WithdrawStatus           int                `gorm:"column:withdraw_status" json:"withdraw_status" form:"withdraw_status" uri:"withdraw_status" `                                                 // 打款状态（0 未打款；1 已打款）
}

func (FcOrderWithdrawPaymentOut) TableName() string {
	return "fc_order_withdraw_payment_out"
}

type FcOrderWithdrawPaymentOutEx struct {
	FcOrderWithdrawPaymentOut
	OrderWithdraw FcOrderWithdraw `json:"order_withdraw" gorm:"foreignkey:OrderSn;references:OrderSn"`
}

type FcOrderWithdrawPaymentOutResp struct {
	FcOrderWithdrawPaymentOut
	AuditStatus     int    `json:"audit_status"`
	AccountNumber   string `json:"account_number"`
	AccountHolder   string `json:"account_holder"`
	AccountBankType string `json:"account_bank_type"`
}
