package userTransfer

// 转账类型
type TransactionAmountType int

const (
	TranCompany        TransactionAmountType = 1  //公司入款
	TranOnline         TransactionAmountType = 2  //在线存款
	TranManual         TransactionAmountType = 3  //手动存款
	TranWithdraw       TransactionAmountType = 4  //提款
	TranDiscount       TransactionAmountType = 5  //优惠
	TranRebate         TransactionAmountType = 6  //返水
	TranServiceCharge  TransactionAmountType = 7  //服务费/手续费
	TranManageAdd      TransactionAmountType = 8  //管理员添加
	TranManageReduce   TransactionAmountType = 9  //管理员扣除
	TranAmountConvert  TransactionAmountType = 10 //额度转换
	TranWithdrawReject TransactionAmountType = 11 //提款拒绝
	TranAgentDeposit   TransactionAmountType = 12 //代理充值

	TranSingleVenueBet                TransactionAmountType = 13 //单一钱包投注
	TranSingleVenueWin                TransactionAmountType = 14 //单一钱包输赢
	TranSingleVenueTransIn            TransactionAmountType = 15 //单一钱包转入游戏
	TranSingleVenueTransOut           TransactionAmountType = 16 //单一钱包转出游戏
	TranSingleVenueAmend              TransactionAmountType = 17 //单一钱二次结算
	TranSingleVenueBetCancel          TransactionAmountType = 18 //单一钱包注单取消(加款)
	TranSingleVenueBetBackRollCancel  TransactionAmountType = 19 //单一钱包注单取消回滚(扣款)
	TranSingleVenueSettlementRollback TransactionAmountType = 20 //单一钱包	结算回滚(扣款)
	TranSingleVenueManuallyOptMoney   TransactionAmountType = 21 //单一钱包	人工加减款

	TranSingleVenueUnKnow TransactionAmountType = 100 //单一钱包 未知操作
)
