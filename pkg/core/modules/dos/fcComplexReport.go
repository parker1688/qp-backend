package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcComplexReport struct {
	BaseDos
	Day                       string             `gorm:"column:day" json:"day" form:"day" uri:"day" `                                                                                                         // 日期
	MerchantCode              string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                                                 // 商户code
	MerchantName              string             `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `                                                                 // 商户名称
	BetWin                    float64            `gorm:"column:bet_win" json:"bet_win" form:"bet_win" uri:"bet_win" `                                                                                         // 累计输赢
	BetAmount                 float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                                                                             // 有效投注
	GameKillRate              float64            `gorm:"column:game_kill_rate" json:"game_kill_rate" form:"game_kill_rate" uri:"game_kill_rate" `                                                             // 游戏杀率
	BetMultiple               float64            `gorm:"column:bet_multiple" json:"bet_multiple" form:"bet_multiple" uri:"bet_multiple" `                                                                     // 投注倍数
	RebateAmount              float64            `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                                                                 // 投注倍数
	RegisterNum               int                `gorm:"column:register_num" json:"register_num" form:"register_num" uri:"register_num" `                                                                     // 注册人数
	FirstDepositNum           int                `gorm:"column:first_deposit_num" json:"first_deposit_num" form:"first_deposit_num" uri:"first_deposit_num" `                                                 // 首充值人数
	RegisterDepositRate       float64            `gorm:"column:register_deposit_rate" json:"register_deposit_rate" form:"register_deposit_rate" uri:"register_deposit_rate" `                                 // 注册充值率
	DepositNum                int                `gorm:"column:deposit_num" json:"deposit_num" form:"deposit_num" uri:"deposit_num" `                                                                         // 充值人数
	DepositCount              int                `gorm:"column:deposit_count" json:"deposit_count" form:"deposit_count" uri:"deposit_count" `                                                                 // 充值笔数
	FirstDepositAmount        float64            `gorm:"column:first_deposit_amount" json:"first_deposit_amount" form:"first_deposit_amount" uri:"first_deposit_amount" `                                     // 首充日累计金额
	NewUserDepositAmount      float64            `gorm:"column:new_user_deposit_amount" json:"new_user_deposit_amount" form:"new_user_deposit_amount" uri:"new_user_deposit_amount" `                         // 新户充值金额
	NewUserDepositCount       int                `gorm:"column:new_user_deposit_count" json:"new_user_deposit_count" form:"new_user_deposit_count" uri:"new_user_deposit_count" `                             // 新户充值笔数
	NewUserFisrtDepositAmount float64            `gorm:"column:new_user_first_deposit_amount" json:"new_user_first_deposit_amount" form:"new_user_first_deposit_amount" uri:"new_user_first_deposit_amount" ` // 新户充值笔数
	LoginNum                  int                `gorm:"column:login_num" json:"login_num" form:"login_num" uri:"login_num" `                                                                                 // 登录人数
	WithdrawNum               int                `gorm:"column:withdraw_num" json:"withdraw_num" form:"withdraw_num" uri:"withdraw_num" `                                                                     // 提现人数
	BetNum                    int                `gorm:"column:bet_num" json:"bet_num" form:"bet_num" uri:"bet_num" `                                                                                         // 投注人数
	PromotionAmount           float64            `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `                                                     // 福利汇总
	AlipayDepositAmount       float64            `gorm:"column:alipay_deposit_amount" json:"alipay_deposit_amount" form:"alipay_deposit_amount" uri:"alipay_deposit_amount" `                                 // 支付宝充值
	WxDepositAmount           float64            `gorm:"column:wx_deposit_amount" json:"wx_deposit_amount" form:"wx_deposit_amount" uri:"wx_deposit_amount" `                                                 // 微信充值
	BankDepositAmount         float64            `gorm:"column:bank_deposit_amount" json:"bank_deposit_amount" form:"bank_deposit_amount" uri:"bank_deposit_amount" `                                         // 银行卡充值
	WalletDepositAmount       float64            `gorm:"column:wallet_deposit_amount" json:"wallet_deposit_amount" form:"wallet_deposit_amount" uri:"wallet_deposit_amount" `                                 // 钱包充值
	NumCnyDepositAmount       float64            `gorm:"column:num_cny_deposit_amount" json:"num_cny_deposit_amount" form:"num_cny_deposit_amount" uri:"num_cny_deposit_amount" `                             // 数字人民币充值
	UsdtDepositAmount         float64            `gorm:"column:usdt_deposit_amount" json:"usdt_deposit_amount" form:"usdt_deposit_amount" uri:"usdt_deposit_amount" `                                         // USDT充值
	AdminDepositAmount        float64            `gorm:"column:admin_deposit_amount" json:"admin_deposit_amount" form:"admin_deposit_amount" uri:"admin_deposit_amount" `                                     // 人工存款
	TotalDepositAmount        float64            `gorm:"column:total_deposit_amount" json:"total_deposit_amount" form:"total_deposit_amount" uri:"total_deposit_amount" `                                     // 充值汇总
	AlipayWithdrawAmount      float64            `gorm:"column:alipay_withdraw_amount" json:"alipay_withdraw_amount" form:"alipay_withdraw_amount" uri:"alipay_withdraw_amount" `                             // 支付宝提现
	BankWithdrawAmount        float64            `gorm:"column:bank_withdraw_amount" json:"bank_withdraw_amount" form:"bank_withdraw_amount" uri:"bank_withdraw_amount" `                                     // 银行卡提现
	WalletWithdrawAmount      float64            `gorm:"column:wallet_withdraw_amount" json:"wallet_withdraw_amount" form:"wallet_withdraw_amount" uri:"wallet_withdraw_amount" `                             // 钱包提现
	UsdtWithdrawAmount        float64            `gorm:"column:usdt_withdraw_amount" json:"usdt_withdraw_amount" form:"usdt_withdraw_amount" uri:"usdt_withdraw_amount" `                                     // usdt提现
	TotalWithdrawAmount       float64            `gorm:"column:total_withdraw_amount" json:"total_withdraw_amount" form:"total_withdraw_amount" uri:"total_withdraw_amount" `                                 // 提现汇总
	DepositWithdrawSubAmount  float64            `gorm:"column:deposit_withdraw_sub_amount" json:"deposit_withdraw_sub_amount" form:"deposit_withdraw_sub_amount" uri:"deposit_withdraw_sub_amount" `         // 充提差
	KillRate                  float64            `gorm:"column:kill_rate" json:"kill_rate" form:"kill_rate" uri:"kill_rate" `                                                                                 // 财务杀率
	CreateTime                automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                                            // 创建时间
	CreateBy                  string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                                                 // 创建人
	UpdateTime                automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                                            // 修改时间
	UpdateBy                  string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                                                 // 修改人
}

func (FcComplexReport) TableName() string {
	return "fc_complex_report"
}
