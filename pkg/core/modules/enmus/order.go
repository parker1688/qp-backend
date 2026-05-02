package enmus

const (
	//待处理
	ORDER_PENDING_STATUS = 0
	//等待处理
	ORDER_STATUS_WAIT = 1
	//拒绝通过
	ORDER_NO_STATUS = 2
	//通过
	ORDER_YES_STATUS = 3
	//待支付
	Order_STATUS_PENDING_PAY = 7

	Manual_Recharge_ChannelCode  = "manual"
	Manual_Recharge_PayCode      = "manual"
	Manual_Recharge_PayName      = "人工存款"
	Manual_Recharge_PayAliasName = "人工存款"
	Manual_Recharge_PayId        = 1
)

const (
	// 提款打款状态
	OrderWithdrawStats_No        = 0 // 未打款
	OrderWithdrawStats_Yes       = 1 // 已打款
	OrderWithdrawStats_ManualNo  = 2 // 人工未打款
	OrderWithdrawStats_ManualYes = 3 // 人工已打款

	// 提现订单状态
	OrderWithdrawStats_AuditWait    = 0 // 待审核
	OrderWithdrawStats_AuditApprove = 1 // 审核通过
	OrderWithdrawStats_AuditReject  = 2 // 审核拒绝

	// 提现订单三方状态
	OrderWithdrawAnotherPayStats_None     = 0 // 无代付
	OrderWithdrawAnotherPayStats_Progress = 1 // 代付中
	OrderWithdrawAnotherPayStats_Failed   = 2 // 代付失败
	OrderWithdrawAnotherPayStats_Success  = 3 // 代付成功

	// 提现汇款申请状态
	OrderWithdrawPaymentOutStats_Prepare  = 0 // 下发中（去队列排队状态）
	OrderWithdrawPaymentOutStats_Progress = 1 // 打款中（请求三方但未回调的期间状态）
	OrderWithdrawPaymentOutStats_Failed   = 2 // 打款失败
	OrderWithdrawPaymentOutStats_Success  = 3 // 打款成功
)

const (
	//  银行卡
	ORDER_TYPE_BANK = 1
	//  三方订单
	ORDER_TYPE_Other = 2
	//  虚拟币
	ORDER_TYPE_Virtual = 3

	ORDER_TYPE_Online = 4 //支付宝 微信之类的

	ORDER_TYPE_WX = 1

	Recharge_Order_Type_Wx        = 1
	Recharge_Order_Type_Bank      = 2
	Recharge_Order_Type_Alipay    = 3
	Recharge_Order_Type_Wallet    = 4
	Recharge_Order_Type_NumberCNY = 5
	Recharge_Order_Type_Virtual   = 6
	Recharge_Order_Type_Manual    = 20
)

const (
	//1. 存款活动 2. 好友邀请 3 红包奖金 4 生日礼金
	Promotion_Deposit              = 1
	Promotion_Friend               = 2
	Promotion_Red_Envelope         = 3
	Promotion_Birthday             = 4
	Promotion_Invite_Friend        = 5
	Promotion_Vip_Upgradation      = 6
	Promotion_Register_Redpacket   = 7  //注册红包赠送
	Promotion_Friend_Redpacket     = 9  //邀请好友抢红包
	Promotion_Task_Redpacket       = 10 //任务红包
	Promotion_Vip_Rank_Upgradation = 11 //段位升级奖金
	Promotion_VIP_WeekGift         = 14 //周俸禄
	Promotion_VIP_MonthGift        = 15 //月俸禄
	Promotion_Official_Bonus       = 16 // 官方赠送
	Promotion_Official_Recoup      = 17 // 官方补偿
	Promotion_Official_Deduct      = 18 // 福利扣除

	Promotion_apply_status_wait    = 0
	Promotion_apply_status_refuse  = 1
	Promotion_apply_status_success = 2
)
