package vo

import "bootpkg/common/expands/automaticType"

// 用户详情返回
type MaterialDetail struct {
	FcUserMaterialVO
	RechargeAmount   float64 `json:"recharge_amount"`   // 总充值
	WithdrawalAmount float64 `json:"withdrawal_amount"` // 总提现
	WinAmount        float64 `json:"win_amount"`        // 总输赢
	ValidBetamount   float64 `json:"valid_betamount"`   // 总有效投注
	BetAmount        float64 `json:"bet_amount"`        // 总投注
	RebateAmount     float64 `json:"rebate_amount"`     // 返水金额
	PromotionAmount  float64 `json:"promotion_amount"`  // 累计福利,不包含返水
	IsLogin          int     `json:"is_login"`          // 是否在线 1 在线 2 不在线
}

type FcUserMaterialVO struct {
	UserId               string                `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `                                                             // 用户Id
	UserName             string                `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                                                     // 用户账号
	NickName             string                `gorm:"nick_name" json:"nick_name" form:"nick_name" uri:"nick_name" `                                                     // 用户昵称
	RealName             string                `gorm:"column:real_name" json:"real_name" form:"real_name" uri:"real_name" `                                              // 真实姓名
	Sex                  int                   `gorm:"column:sex" json:"sex" form:"sex" uri:"sex" `                                                                      // 用户性别
	Currency             string                `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                                         // 货币类型
	TotalAmount          float64               `gorm:"total_amount" json:"total_amount" form:"total_amount" uri:"total_amount" `                                         // 总金额
	AvaAmount            float64               `gorm:"ava_amount" json:"ava_amount" form:"ava_amount" uri:"ava_amount" `                                                 // 可用金额
	FronzenAmount        float64               `gorm:"fronzen_amount" json:"fronzen_amount" form:"fronzen_amount" uri:"fronzen_amount" `                                 // 冻结金额
	IsLock               int                   `gorm:"is_lock" json:"is_lock" form:"is_lock" uri:"is_lock" `                                                             // 钱包状态 1禁用 2正常
	Level                int                   `gorm:"level" json:"level" form:"level" uri:"level" `                                                                     // 1~10 VIP等级
	Birthday             automaticType.Time    `gorm:"column:birthday" json:"birthday" form:"birthday" uri:"birthday" `                                                  // 生日
	IsWithdraw           int                   `gorm:"column:is_withdraw" json:"is_withdraw" form:"is_withdraw" uri:"is_withdraw" `                                      // 是否可提款 1:可以 2:不可以
	IsBonus              int                   `gorm:"column:is_bonus" json:"is_bonus" form:"is_bonus" uri:"is_bonus" `                                                  // 是否可领取福利 1:可以 2:不可以
	MerchantCode         string                `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                     // 商户code
	CreateTime           automaticType.Time    `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                // 创建时间
	ManualTransferWallet automaticType.BitBool `gorm:"manual_transfer_wallet" json:"manual_transfer_wallet" form:"manual_transfer_wallet" uri:"manual_transfer_wallet" ` // 手动转账钱包
	IsFree               automaticType.BitBool `gorm:"is_free" json:"is_free" form:"is_free" uri:"is_free" `                                                             // 0:非试玩   1：试玩
	LastLoginTime        automaticType.Time    `gorm:"last_login_time" json:"last_login_time" form:"-" uri:"-" `                                                         // 注册IP
	InviteCode           string                `gorm:"invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `                                             // 邀请短码
	AgentInviteCode      int                   `gorm:"column:agent_invite_code" json:"agent_invite_code" form:"agent_invite_code" uri:"agent_invite_code" `              // 上级推广码
	Email                string                `gorm:"email" json:"email" form:"email" uri:"email" `                                                                     // 邮件
	Tel                  string                `gorm:"column:tel" json:"tel" form:"tel" uri:"tel" `                                                                      // 电话号码
	WalletPassword       string                `gorm:"column:wallet_password" json:"wallet_password" form:"wallet_password" uri:"wallet_password" `                      // 支付密码
	LoginStatus          int                   `gorm:"login_status" json:"login_status" form:"login_status" uri:"login_status" `                                         // 登录状态 0 正常 1 禁止
	Remark               string                `gorm:"remark" json:"remark" form:"remark" uri:"remark" `                                                                 // 用户备注
	LastLoginIp          string                `gorm:"last_login_ip" json:"last_login_ip" form:"last_login_ip" uri:"last_login_ip" `                                     // 最后登录IP
	LastLoginIpCity      string                `gorm:"last_login_ip_city" json:"last_login_ip_city" form:"last_login_ip_city" uri:"last_login_ip_city" `                 // 最后登录IP地域信息
	AgentId              string                `gorm:"agent_id" json:"agent_id" form:"agent_id" uri:"agent_id" `                                                         // 代理Id
	LastLoginCount       int                   `gorm:"last_login_count" json:"last_login_count" form:"-" uri:"-" `                                                       // 最后登录总数
	RegisterIp           string                `gorm:"register_ip" json:"register_ip" form:"-" uri:"register_ip" `                                                       // 注册登录IP
	RegisterIpCity       string                `gorm:"register_ip_city" json:"register_ip_city" form:"-" uri:"register_ip_city" `                                        // 注册登录IP地域信息
	VisitorId            string                `gorm:"column:visitor_id" json:"visitor_id" form:"visitor_id" uri:"visitor_id" `                                          // 最后一次登录设备号
	RegistVisitorId      string                `gorm:"column:regist_visitor_id" json:"regist_visitor_id" form:"regist_visitor_id" uri:"regist_visitor_id" `              // 注册设备号
	IsActive             int                   `gorm:"is_active" json:"is_active" form:"is_active" uri:"is_active" `                                                     // 是否活跃, 1 活跃 2 不活跃
	//Alipay               string                `gorm:"column:alipay" json:"alipay" form:"alipay" uri:"alipay" `                                                          // 支付宝
	//AlipayRealname       string                `gorm:"column:alipay_realname" json:"alipay_realname" form:"alipay_realname" uri:"alipay_realname" `                      // 支付宝姓名
	Website string `gorm:"column:website" json:"website" form:"website" uri:"website" ` // 网站
}

type FcVipListVO struct {
	Id               string  `gorm:"column:id" json:"id" form:"id" uri:"id" `                                                                 // id
	VipName          string  `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" `                                         // VIP1~VIP10
	Level            int     `gorm:"column:level" json:"level" form:"level" uri:"level" `                                                     // 层级 1~10
	MinBetAmount     float64 `gorm:"column:min_bet_amount" json:"min_bet_amount" form:"min_bet_amount" uri:"min_bet_amount" `                 // 流水要求
	UpgradeGift      float64 `gorm:"column:upgrade_gift" json:"upgrade_gift" form:"upgrade_gift" uri:"upgrade_gift" `                         // 升级礼金
	WeeklyGift       float64 `gorm:"column:weekly_gift" json:"weekly_gift" form:"weekly_gift" uri:"weekly_gift" `                             // 每周礼金
	MonthlyGift      float64 `gorm:"column:monthly_gift" json:"monthly_gift" form:"monthly_gift" uri:"monthly_gift" `                         // 每月礼金
	YearlyGift       float64 `gorm:"column:yearly_gift" json:"yearly_gift" form:"yearly_gift" uri:"yearly_gift" `                             // 每年礼金
	UpgradeGiftApply bool    `gorm:"column:upgrade_gift_apply" json:"upgrade_gift_apply" form:"upgrade_gift_apply" uri:"upgrade_gift_apply" ` //升级奖金是领取
	WeeklyGiftApply  bool    `gorm:"column:weekly_gift_apply" json:"weekly_gift_apply" form:"weekly_gift_apply" uri:"weekly_gift_apply" `     //周奖金是领取
	MonthlyApply     bool    `gorm:"column:monthly_gift_apply" json:"monthly_gift_apply" form:"monthly_gift_apply" uri:"monthly_gift_apply" ` //月奖金是领取
}

type UserIdNameVO struct {
	UserId   string `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `         // 用户Id
	UserName string `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 用户账号
}
