package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcOrderDeposit struct {
	BaseDos
	UserId                  string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName                string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	OrderSn                 string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                                                                         // 订单编号
	Amount                  float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                                                                 // 存款金额
	BonusRate               float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                                                                 // 优惠比例
	BonusAmount             float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                                                         // 优惠金额
	FactAmount              float64            `gorm:"column:fact_amount" json:"fact_amount" form:"fact_amount" uri:"fact_amount" `                                                             // 实际存款
	Status                  int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                                                 // 0:待处理  2:拒绝 3:已通过
	EntityAccountHolder     string             `gorm:"column:entity_account_holder" json:"entity_account_holder" form:"entity_account_holder" uri:"entity_account_holder" `                     // 收款人
	EntityAccountBankName   string             `gorm:"column:entity_account_bank_name" json:"entity_account_bank_name" form:"entity_account_bank_name" uri:"entity_account_bank_name" `         // 收款人银行名字
	EntityAccountNumber     string             `gorm:"column:entity_account_number" json:"entity_account_number" form:"entity_account_number" uri:"entity_account_number" `                     // 收款卡号或地址
	RemitterAccountHolder   string             `gorm:"column:remitter_account_holder" json:"remitter_account_holder" form:"remitter_account_holder" uri:"remitter_account_holder" `             // 汇款人
	RemitterAccountBankName string             `gorm:"column:remitter_account_bank_name" json:"remitter_account_bank_name" form:"remitter_account_bank_name" uri:"remitter_account_bank_name" ` // 快款人银行名字
	RemitterAccountNumber   string             `gorm:"column:remitter_account_number" json:"remitter_account_number" form:"remitter_account_number" uri:"remitter_account_number" `             // 汇款卡号或地址
	Remark                  string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `
	DepositRemark           string             `gorm:"column:deposit_remark" json:"deposit_remark" form:"deposit_remark" uri:"deposit_remark" `
	Ip                      string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `
	Level                   int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                      // 用户等级
	CreateTime              automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	PayTime                 automaticType.Time `gorm:"column:pay_time;default:null" json:"pay_time" form:"pay_time" uri:"pay_time" `             // 支付时间
	CreateBy                string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime              automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy                string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode            string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	InviteCode              int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `
	ChannelId               int                `gorm:"column:channel_id" json:"channel_id" form:"channel_id" uri:"channel_id" `                                                 // 渠道Id
	ChannelCode             string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `                                         // 渠道code
	PaymentId               int                `gorm:"column:payment_id" json:"payment_id" form:"payment_id" uri:"payment_id" `                                                 // 通道ID
	PaymentCode             string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `                                         // 通道code
	PaymentName             string             `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" `                                         // 通道名称
	PayAliasName            string             `gorm:"column:pay_alias_name" json:"pay_alias_name" form:"pay_alias_name" uri:"pay_alias_name" `                                 // 通道名称别名
	Currency                string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                                         // 币种
	VirtualAddress          string             `gorm:"column:virtual_address" json:"virtual_address" form:"virtual_address" uri:"virtual_address" `                             // 虚拟币地址
	VirtualType             string             `gorm:"column:virtual_type" json:"virtual_type" form:"virtual_type" uri:"virtual_type" `                                         // 虚拟币币种名称
	VirtualCurrencyChain    string             `gorm:"column:virtual_currency_chain" json:"virtual_currency_chain" form:"virtual_currency_chain" uri:"virtual_currency_chain" ` // 虚拟币所属链
	VirtualNum              float64            `gorm:"column:virtual_num" json:"virtual_num" form:"virtual_num" uri:"virtual_num" `                                             // 虚拟币数量
	VirtualFx               float64            `gorm:"column:virtual_fx" json:"virtual_fx" form:"virtual_fx" uri:"virtual_fx" `                                                 // 虚拟币汇率
	VirtualPayNo            string             `gorm:"column:virtual_pay_no" json:"virtual_pay_no" form:"virtual_pay_no" uri:"virtual_pay_no" `                                 // 虚拟币支付交易ID号
	VirtualPayAddress       string             `gorm:"column:virtual_pay_address" json:"virtual_pay_address" form:"virtual_pay_address" uri:"virtual_pay_address" `             // 虚拟币支付地址
	VirtualPayAmount        float64            `gorm:"column:virtual_pay_amount" json:"virtual_pay_amount" form:"virtual_pay_amount" uri:"virtual_pay_amount" `                 // 虚拟币到账数量
	OrderType               int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `                                                 // 订单类型 1 银行卡 2 三方通道 3 虚拟币
	OrderSecondType         int                `gorm:"column:order_second_type" json:"order_second_type" form:"order_second_type" uri:"order_second_type" `                     // 订单第二类型 1:人工存款-微信 2:人工存款-银行卡 3:人工存款-支付宝 4:人工存款-钱包 5:人工存款-数字人民币 6:人工存款-数字货币U
	ActivityId              string             `gorm:"column:activity_id" json:"activity_id" form:"activity_id" uri:"activity_id" `                                             // 参与赠送活动ID
	Fee                     float64            `gorm:"column:fee" json:"fee" form:"fee" uri:"fee" `                                                                             // 手续费
	FeeRate                 float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                                                         // 手续费比例(百分比)
	AuthBy                  string             `gorm:"column:auth_by" json:"auth_by" form:"auth_by" uri:"auth_by" `                                                             // 审核人员
	AuthTime                automaticType.Time `gorm:"column:auth_time" json:"auth_time" form:"auth_time" uri:"auth_time" `                                                     // 审核时间
}

func (FcOrderDeposit) TableName() string {
	return "fc_order_deposit"
}
