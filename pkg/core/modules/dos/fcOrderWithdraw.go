package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
)

type FcOrderWithdraw struct {
	BaseDos
	UserId                   string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName                 string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Amount                   float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                                                                     // 提款金额
	PreAmount                float64            `gorm:"column:pre_amount" json:"pre_amount" form:"pre_amount" uri:"pre_amount" `                                                                     // 提款金额-除手续费的金额
	Status                   int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                                                     // 0 待审核；1 审核通过；2 审核拒绝
	Province                 string             `gorm:"column:province" json:"province" form:"province" uri:"province" `                                                                             // 身份
	City                     string             `gorm:"column:city" json:"city" form:"city" uri:"city" `                                                                                             // 城市
	BankAddress              string             `gorm:"column:bank_address" json:"bank_address" form:"bank_address" uri:"bank_address" `                                                             // 银行地址
	AccountNumber            string             `gorm:"column:account_number" json:"account_number" form:"account_number" uri:"account_number" `                                                     // 卡号或地址
	AccountHolder            string             `gorm:"column:account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `                                                     // 收款人
	AccountBankType          string             `gorm:"column:account_bank_type" json:"account_bank_type" form:"account_bank_type" uri:"account_bank_type" `                                         // 银行类型名称
	AccountBankCode          string             `gorm:"column:account_bank_code" json:"account_bank_code" form:"account_bank_code" uri:"account_bank_code" `                                         // 银行类型编码
	Remark                   string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                                                                     // 备注
	Ip                       string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `                                                                                                     // 提款IP
	CreateTime               automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                                    // 创建时间
	CreateBy                 string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                                         // 创建人
	UpdateTime               automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                                    // 修改时间
	UpdateBy                 string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                                         // 修改人
	MerchantCode             string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                                         // 商户code
	OrderSn                  string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                                                                             // 订单编号
	Currency                 string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                                                             // 币种简码
	VirtualAddress           string             `gorm:"column:virtual_address" json:"virtual_address" form:"virtual_address" uri:"virtual_address" `                                                 // 虚拟币地址
	VirtualType              string             `gorm:"column:virtual_type" json:"virtual_type" form:"virtual_type" uri:"virtual_type" `                                                             // 虚拟币币种名称
	VirtualNum               float64            `gorm:"column:virtual_num" json:"virtual_num" form:"virtual_num" uri:"virtual_num" `                                                                 // 虚拟币数量
	VirtualFx                float64            `gorm:"column:virtual_fx" json:"virtual_fx" form:"virtual_fx" uri:"virtual_fx" `                                                                     // 虚拟币存款汇率
	VirtualPayNo             string             `gorm:"column:virtual_pay_no" json:"virtual_pay_no" form:"virtual_pay_no" uri:"virtual_pay_no" `                                                     // 虚拟币支付交易ID号
	VirtualPayAddress        string             `gorm:"column:virtual_pay_address" json:"virtual_pay_address" form:"virtual_pay_address" uri:"virtual_pay_address" `                                 // 虚拟币支付地址
	VirtualPayAmount         float64            `gorm:"column:virtual_pay_amount" json:"virtual_pay_amount" form:"virtual_pay_amount" uri:"virtual_pay_amount" `                                     // 虚拟币到账数量
	OrderType                int                `gorm:"column:order_type" json:"order_type" form:"order_type" uri:"order_type" `                                                                     // 订单类型 1 银行卡 2 三方通道 3 虚拟币
	VirtualCurrencyChain     string             `gorm:"column:virtual_currency_chain" json:"virtual_currency_chain" form:"virtual_currency_chain" uri:"virtual_currency_chain" `                     // 虚拟币所属链
	Fee                      float64            `gorm:"column:fee" json:"fee" form:"fee" uri:"fee" `                                                                                                 // 手续费
	FeeRate                  float64            `gorm:"column:fee_rate" json:"fee_rate" form:"fee_rate" uri:"fee_rate" `                                                                             // 手续费比例(百分比)
	AnotherPayStatus         int                `gorm:"column:another_pay_status" json:"another_pay_status" form:"another_pay_status" uri:"another_pay_status" `                                     // 0 无代付 1 代付中 2 代付失败 3 代付成功
	CallbackRemark           string             `gorm:"column:callback_remark" json:"callback_remark" form:"callback_remark" uri:"callback_remark" `                                                 // 状态备注
	AuthBy                   string             `gorm:"column:auth_by" json:"auth_by" form:"auth_by" uri:"auth_by" `                                                                                 // 审核人员
	AuthTime                 automaticType.Time `gorm:"column:auth_time" json:"auth_time" form:"auth_time" uri:"auth_time" `                                                                         // 审核时间
	AnotherPayTime           automaticType.Time `gorm:"column:another_pay_time" json:"another_pay_time" form:"another_pay_time" uri:"another_pay_time" `                                             // 代付时间
	DepositWithdrawSubAmount float64            `gorm:"column:deposit_withdraw_sub_amount" json:"deposit_withdraw_sub_amount" form:"deposit_withdraw_sub_amount" uri:"deposit_withdraw_sub_amount" ` // 提存差
}

func (FcOrderWithdraw) TableName() string {
	return "fc_order_withdraw"
}

func (m *FcOrderWithdraw) Decrypt() {

	if m.OrderType == 1 {
		accountHolder := tool.DecryptAESPrefixRandKeySalt(m.AccountHolder, "fc_user_withdraw_bank_bind")
		if len(accountHolder) > 0 {
			m.AccountHolder = accountHolder
		}
		accountNumber := tool.DecryptAESPrefixRandKeySalt(m.AccountNumber, "fc_user_withdraw_bank_bind")
		if len(accountNumber) > 0 {
			m.AccountNumber = accountNumber
		}
		bankAddress := tool.DecryptAESPrefixRandKeySalt(m.BankAddress, "fc_user_withdraw_bank_bind")
		if len(bankAddress) > 0 {
			m.BankAddress = bankAddress
		}
	} else if m.OrderType == 3 {
		accountHolder := tool.DecryptAESPrefixRandKeySalt(m.AccountHolder, "fc_user_withdraw_blockchain_bind")
		if len(accountHolder) > 0 {
			m.AccountHolder = accountHolder
		}
	} else {
		accountHolder := tool.DecryptAESPrefixRandKeySalt(m.AccountHolder, "fc_user_withdraw_online_bind")
		if len(accountHolder) > 0 {
			m.AccountHolder = accountHolder
		}
		accountNumber := tool.DecryptAESPrefixRandKeySalt(m.AccountNumber, "fc_user_withdraw_online_bind")
		if len(accountNumber) > 0 {
			m.AccountNumber = accountNumber
		}
		bankAddress := tool.DecryptAESPrefixRandKeySalt(m.BankAddress, "fc_user_withdraw_online_bind")
		if len(bankAddress) > 0 {
			m.BankAddress = bankAddress
		}
	}

	virtualAddress := tool.DecryptAESPrefixRandKeySalt(m.VirtualAddress, "fc_user_withdraw_blockchain_bind")
	if len(virtualAddress) > 0 {
		m.VirtualAddress = virtualAddress
	}
}

type FcOrderWithdrawResp struct {
	FcOrderWithdraw
	WithdrawStatus int `json:"withdraw_status"`
}
