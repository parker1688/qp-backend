package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcFirstOrderDeposit struct {
	BaseDos
	UserId          string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName        string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	OrderSn         string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                          // 订单编号
	Amount          float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 存款金额
	FactAmount      float64            `gorm:"column:fact_amount" json:"fact_amount" form:"fact_amount" uri:"fact_amount" `              // 实际金额
	BonusAmount     float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `          // 优惠金额
	BonusRate       float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                  // 优惠比例
	Remark          string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                  // 审核备注
	DepositRemark   string             `gorm:"column:deposit_remark" json:"deposit_remark" form:"deposit_remark" uri:"deposit_remark" `  // 存款备注
	Ip              string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `                                                  // 存款IP
	Level           int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                      // 用户等级
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	PayTime         automaticType.Time `gorm:"column:pay_time" json:"pay_time" form:"pay_time" uri:"pay_time" `                          // 支付时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime      automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy        string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode    string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	InviteCode      int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `
	ChannelId       int                `gorm:"column:channel_id" json:"channel_id" form:"channel_id" uri:"channel_id" `                             // 渠道Id
	ChannelCode     string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `                     // 渠道code
	PaymentId       int                `gorm:"column:payment_id" json:"payment_id" form:"payment_id" uri:"payment_id" `                             // 通道ID
	PaymentCode     string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `                     // 通道code
	PayAliasName    string             `gorm:"column:pay_alias_name" json:"pay_alias_name" form:"pay_alias_name" uri:"pay_alias_name" `             // 通道名称别名
	PaymentName     string             `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `                     // 通道名称
	Currency        string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                     // 币种
	OrderType       int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `                             // 支付渠道类型 1:微信 2:银行卡 3:支付宝 4:钱包 5:数字人民币 6:数字货币 20 人工充值
	OrderSecondType int                `gorm:"column:order_second_type" json:"order_second_type" form:"order_second_type" uri:"order_second_type" ` // 订单第二类型 1:人工存款-微信 2:人工存款-银行卡 3:人工存款-支付宝 4:人工存款-钱包 5:人工存款-数字人民币 6:人工存款-数字货币U
	FeeRate         float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                                     // 手续费比例(百分比)
	Fee             float64            `gorm:"column:fee" json:"fee" form:"fee" uri:"fee" `                                                         // 手续费
	AuthBy          string             `gorm:"column:auth_by" json:"auth_by" form:"auth_by" uri:"auth_by" `                                         // 审核人员
	AuthTime        automaticType.Time `gorm:"column:auth_time" json:"auth_time" form:"auth_time" uri:"auth_time" `                                 // 审核时间
}

func (FcFirstOrderDeposit) TableName() string {
	return "fc_first_order_deposit"
}
