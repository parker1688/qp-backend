package router

import (
	"bootpkg/cmd/api/controller/walletControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, walletRouter)
}

func walletRouter() {
	r := routers.Group("/api/wallet").Use(handler.AuthApiMiddleware())
	r.POST("/center", walletControl.GetUserWalletMoney) //获取中心钱包余额
	r.POST("/withdraw", walletControl.Withdraw)         //提款

	r.POST("/list/deposit", walletControl.OrderDepositInfo)     //获取存款订单信息
	r.POST("/list/withdraw", walletControl.OrderWithdrawInfo)   //获取提款信息
	r.POST("/list/transaction", walletControl.TransactionOrder) //获取转账信息
	r.POST("/list/promotion", walletControl.OrderPromotionInfo) //获取优惠信息
	r.POST("/channelOut", walletControl.GetWithdrawChannel)     //提款渠道
	r.POST("/paymentOut", walletControl.GetWithdrawPaymentOut)  //提款通道

	rRecharge := routers.Group("/api/recharge").Use(handler.AuthApiMiddleware())
	rRecharge.POST("/channel", walletControl.GetRechargeChannel)                //获取充值渠道
	rRecharge.POST("/channel/setting", walletControl.GetRechargeChannelSetting) //获取充值渠道配置
	rRecharge.POST("/payment", walletControl.UserPaymentChannel)                //用户充值
	rRecharge.POST("/payment/get", walletControl.GetPayment)                    //获取通道列表
	rRecharge.POST("/payment/detail", walletControl.GetPaymentDetail)           //通道详情

	rRecharge.POST("/deposit/orderPay", walletControl.UpdateDepositOrderStatus) //修改订单为已支付
	rRecharge.POST("/deposit/orderStatus", walletControl.GetDepositOrderStatus) //查询订单状态

	bindUser := routers.Group("/api/bind").Use(handler.AuthApiMiddleware())
	bindUser.POST("/bank", walletControl.BindBank)                            //绑定银行卡
	bindUser.POST("/blockchain", walletControl.BindBlockchain)                //绑定虚拟币
	bindUser.POST("/online", walletControl.BindOnline)                        //绑定在线提款
	bindUser.POST("/bank/get", walletControl.GetBindBank)                     //获取绑定银行
	bindUser.POST("/blockchain/get", walletControl.GetBindBlockchain)         //获取绑定虚拟币
	bindUser.POST("/online/get", walletControl.GetBindOnlie)                  //获取绑定虚拟币
	bindUser.POST("/bank/default", walletControl.BindBankDefault)             //设置默认银行卡
	bindUser.POST("/blockchain/default", walletControl.BindBlockchainDefault) //设置默认虚拟币
	bindUser.POST("/online/default", walletControl.BindOnlineDefault)         //设置默认虚拟币
	bindUser.POST("/blockchain/type", walletControl.BindBlockchainType)       //获取可绑定虚拟币类型
	bindUser.POST("/bank/type", walletControl.BindBankType)                   //获取银行卡可绑定类型
	bindUser.POST("/channelBank/img", walletControl.ChannelBankImg)           //代付相关美术资源表
}
