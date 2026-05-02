package vo

type PasswordUpdateReq struct {
	OldPassword     string `json:"oldPassword" validate:"required,min=4,max=10"`
	NewPassword     string `json:"newPassword" validate:"required,min=4,max=10"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=4,max=10"`
}

type PasswordResetReq struct {
	//VeryCode        string `json:"veryCode"  form:"veryCode" validate:"min=4,max=10"`
	OldPassword     string `json:"oldPassword" validate:"required,min=4,max=10"`
	NewPassword     string `json:"newPassword" validate:"required,min=4,max=10"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=4,max=10"`
	//Phone           string `json:"phone"  form:"phone" validate:"required"`
	UserName string `json:"userName"  form:"userName" validate:"required"`
}

type PasswordForgotReq struct {
	UserName string `json:"userName"  form:"userName" validate:"required"`
	Phone    string `json:"phone"  form:"phone" validate:"required"`
	VeryCode string `json:"veryCode"  form:"veryCode" validate:"min=4,max=10"`
}

type PasswordForgotUpdateReq struct {
	UserName        string `json:"userName"  form:"userName" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=4,max=10"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=4,max=10"`
}

type WalletPasswordUpdateReq struct {
	VeryCode              string `json:"veryCode"  form:"veryCode" validate:"min=4,max=10"`
	WalletPassword        string `json:"wallet_password" validate:"required,min=4,max=10"`
	WalletConfirmPassword string `json:"wallet_confirm_password" validate:"required,min=4,max=10"`
	//Email          string `json:"email"  form:"email"`
	//Phone          string `json:"phone"  form:"phone"`
}

type VipProgressResp struct {
	NowVip        string  `json:"nowVip"`
	NextVip       string  `json:"nextVip"`
	NextNeedBet   float64 `json:"nextNeedBet"`
	NowBetAmount  float64 `json:"nowBetAmount"`
	NextBetAmount float64 `json:"nextBetAmount"`
	//NowRechargeAmount  float64 `json:"nowRechargeAmount"`
	//NextRechargeAmount float64 `json:"nextRechargeAmount"`
	Progress float64 `json:"progress"`
	//RechargeProgress   float64 `json:"rechargeProgress"`
	//BetProgress        float64 `json:"betProgress"`
	//MinWithdrawAmount  float64 `json:"minWithdrawAmount"`
	//WithdrawalFee      string  `json:"withdrawalFee"`
	//MinRecharegeAmount float64 `json:"minRecharegeAmount"`
	NowLevel int `json:"nowLevel"`
	//TotalBet           float64 `json:"totalBet"`
	//TotalRecharge      float64 `json:"totalRecharge"`
	//NowWeekBetAmount   float64 `json:"nowWeekBetAmount"` //
	//ToWeekBetAmount    float64 `json:"toWeekBetAmount"`
	//NowMonthBetAmount  float64 `json:"nowMonthBetAmount"`
	//ToMonthBetAmount   float64 `json:"toMonthBetAmount"`
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

type UserReportResp struct {
	UserId           string  `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                         // 用户Id
	UserName         string  `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                                 // 用户账号
	Amount           float64 `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                             // 用于余额
	Level            int     `gorm:"column:level" json:"level" form:"level" uri:"level" `                                                 // 1~10 VIP等级
	Vip              string  `gorm:"column:vip" json:"vip" form:"vip" uri:"vip" `                                                         // VIP显示
	RebateFlow       float64 `gorm:"column:rebate_flow" json:"rebate_flow" form:"rebate_flow" uri:"rebate_flow" `                         // 返水流水累计
	MerchantCode     string  `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
	ValidBetamount   float64 `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" `         // 有效投注
	RechargeAmount   float64 `gorm:"column:recharge_amount" json:"recharge_amount" form:"recharge_amount" uri:"recharge_amount" `         // 充值金额
	WithdrawalAmount float64 `gorm:"column:withdrawal_amount" json:"withdrawal_amount" form:"withdrawal_amount" uri:"withdrawal_amount" ` // 提款金额
	PromotionAmount  float64 `gorm:"column:promotion_amount" json:"promotion_amount" form:"promotion_amount" uri:"promotion_amount" `     // 优惠金额
	RebateAmount     float64 `gorm:"column:rebate_amount" json:"rebate_amount" form:"rebate_amount" uri:"rebate_amount" `                 // 返水金额
	WinAmount        float64 `gorm:"column:win_amount" json:"win_amount" form:"win_amount" uri:"win_amount" `                             // 输赢金额

	GameList []*UserGameDataResp ` json:"list"`
}

type UserGameDataResp struct {
	GameType       string  `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                         // 游戏类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	GameCount      int     `gorm:"column:game_count" json:"game_count" form:"game_count" uri:"game_count" `                     // 游戏局数
	BetAmount      float64 `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                     // 投注量
	ValidBetamount float64 `gorm:"column:valid_betamount" json:"valid_betamount" form:"valid_betamount" uri:"valid_betamount" ` // 有效投注量
	NetAmount      float64 `gorm:"column:net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `                     // 输赢

}
