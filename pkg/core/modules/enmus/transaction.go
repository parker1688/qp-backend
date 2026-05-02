package enmus

const (
	FundingTypeDesposit  = 1 // 充值
	FundingTypeWithdraw  = 2 // 提现
	FundingTypePlatform  = 3 // 平台资金切换
	FundingTypePromotion = 4 // 福利
	FundingTypeOther     = 5 // 其他
)

var (
	ConstTransacationStatisDesposit = []string{
		"支付宝充值", "钱包充值", "微信充值", "数字人民币充值", "银行卡充值", "USDT充值",
	}
	ConstTransacationStatisWithdraw = []string{
		"支付宝提现", "银行卡提现", "钱包提现", "财务退回",
	}
	ConstTransacationStatisPromotion = []string{
		"VIP晋级彩金",
		"VIP周彩金",
		"VIP月彩金",
		"存款活动充值自动优惠",
		"签到基础奖励",
		"签到额外奖励",
		"任务奖励",
	}
	ConstTransacationStatisManual = []string{
		"人工存款-微信",
		"人工存款-支付宝",
		"人工存款-银行卡",
		"人工存款-钱包",
		"人工存款-数字人民币",
		"人工存款-数字货币U",
	}
)

var (
	FundingSubTypeEnums = map[int]map[string]string{
		FundingTypeDesposit: {
			"alipay":     "支付宝充值",
			"wallet":     "钱包充值",
			"wx":         "微信充值",
			"number_cny": "数字人民币充值",
			"USDT-CR20":  "USDT充值",
			"bank":       "银行卡充值",
		},
		FundingTypeWithdraw: {
			"ordertype_1": "银行卡提现",
			"ordertype_3": "钱包提现",
			"ordertype_4": "支付宝提现",
			"orderrefuse": "审核退回",
			"orderreturn": "财务退回",
		},
		FundingTypePlatform: {
			"deposit":  "转入",
			"withdraw": "转出",
		},
		FundingTypePromotion: {
			"rebate":            "平台洗码",
			"vip_up_gift":       "VIP晋级彩金",
			"vip_week_gift":     "VIP周彩金",
			"vip_month_gift":    "VIP月彩金",
			"deposit_bonus_1":   "存款活动充值自动优惠",
			"daily_bonus":       "签到基础奖励",
			"daily_bonus_extra": "签到额外奖励",
			"daily_task":        "任务奖励",
			"activity":          "活动奖励",
		},
		FundingTypeOther: {
			"trstype_1":  "人工存款-微信",
			"trstype_2":  "人工存款-银行卡",
			"trstype_3":  "人工存款-支付宝",
			"trstype_4":  "人工存款-钱包",
			"trstype_5":  "人工存款-数字人民币",
			"trstype_6":  "人工存款-数字货币U",
			"trstype_30": "官方赠送",
			"trstype_31": "客服补偿",
			"trstype_40": "一般扣除",
			"trstype_41": "福利扣除",
		},
	}
)
