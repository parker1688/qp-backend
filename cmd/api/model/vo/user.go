package vo

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/pkg/core/modules/dos"
)

type RegisterReq struct {
	UserName        string `json:"username" form:"username" validate:"required,min=6,max=15"`
	Password        string `json:"password" form:"password" validate:"required,min=7,max=14"`
	ConfirmPassword string `json:"confirmPassword"  form:"confirmPassword" validate:"required,min=7,max=14"`
	Phone           string `json:"phone" validate:"required,min=3,max=16"`
	RealName        string `json:"real_name" validate:"required,min=1,max=32"`
	AgentCode       int    `json:"code"`              //代理
	InviteCode      int    `json:"invite_code"`       //用户邀请
	AgentInviteCode int    `json:"agent_invite_code"` //上级代理邀请码
	VeryCode        string `json:"veryCode"`          //验证码
	VeryCodeRandom  string `json:"veryCodeRandom"`    //验证码随机数
	VisitorId       string `json:"visitorId"`         //设备ID
	RequestId       string `json:"requestId"`         //请求ID
	Website         string `json:"website"`           //网站域名
}

type RegisterResp struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
	Level    int    `json:"level"`
}

type UserIdNameVO struct {
	UserId   string `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `         // 用户Id
	UserName string `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 用户账号
}

type LoginReq struct {
	UserName       string `json:"username"  form:"username" label:"用户名" validate:"required,min=4,max=16"`
	Password       string `json:"password" form:"password" validate:"required,min=4,max=16"`
	VeryCode       string `json:"veryCode"`       //验证码
	VeryCodeRandom string `json:"veryCodeRandom"` //验证码随机数
	VisitorId      string `json:"visitorId"`      //设备ID
}

type LoginAbnormalRecoverReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	//Phone     string `json:"phone"`
	//VeryCode  string `json:"veryCode"`
	VisitorId string `json:"visitorId"`
}

type EmailVeryCodeReq struct {
	Email string `json:"email"  form:"email" validate:"required,email"`
	Tag   string `json:"tag"  form:"tag" validate:"required"`
}

type PhoneVeryCodeReq struct {
	Phone string `json:"phone"  form:"phone" `
	Tag   string `json:"tag"  form:"tag" validate:"required"`
}

type EmailResetVeryCodeReq struct {
	Email string `json:"email"  form:"email" validate:"required,email"`
}

type PhoneResetVeryCodeReq struct {
	Phone    string `json:"phone"  form:"phone" validate:"required"`
	UserName string `json:"userName"  form:"userName" validate:"required"`
}

type BindEmailReq struct {
	Email    string `json:"email"  form:"email" validate:"required,email"`
	VeryCode string `json:"veryCode"  form:"veryCode" validate:"min=4,max=10"`
}

type PhoneVerificationReq struct {
	Code string `json:"code"  form:"code" validate:"required"`
}

type RebateRecordListReq struct {
}

type FcUserRebateRecordsListResp struct {
	TotalBonusAmount float64                           `json:"total_bonus_amount"`
	RebateData       []*dos.FcUserRebateRecordsListRow `json:"list"`
}

type FcUserRebateRecordsVO struct {
	Id             string             `gorm:"column:id" json:"id" form:"id" uri:"id" `                                                         // id
	UserId         string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                                     // 用户Id
	UserName       string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                             // 用户账号
	Level          int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                             // 用户vip等级
	VenueCode      string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                         // 场馆code
	GameType       string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                             // 游戏类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	BetAmount      float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                         // 总有效投注
	HisBetAmount   float64            `gorm:"column:his_bet_amount" json:"his_bet_amount" form:"his_bet_amount" uri:"his_bet_amount" `         // 历史累计流水
	BatchBetAmount float64            `gorm:"column:batch_bet_amount" json:"batch_bet_amount" form:"batch_bet_amount" uri:"batch_bet_amount" ` // 批次流水
	RebateType     int                `gorm:"column:rebate_type" json:"rebate_type" form:"rebate_type" uri:"rebate_type" `                     // 发放类型 1: 系统发放  2:手动发放
	Status         int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                         // 状态 1: 已完成  2:发放中 3: 失败
	BonusAmount    float64            `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                 // 奖金金额
	BonusRate      float64            `gorm:"column:bonus_rate" json:"bonus_rate" form:"bonus_rate" uri:"bonus_rate" `                         // 返水比例
	MerchantCode   string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
}

type FcUserRebateApplyReq struct {
	UserId   string `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `         // 用户Id
	UserName string `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 用户账号
}

type FcUserRebateApplyResp struct {
	BonusAmount float64 `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" ` // 返水金额
}

type FcUserRebateDataResp struct {
	TotalBetAmount   float64                  `gorm:"column:total_bet_amount" json:"total_bet_amount" form:"total_bet_amount" uri:"total_bet_amount" ` // 洗码量
	CurrDetailList   []FcUserRebateValsResult `json:"curr_detail_list"`                                                                                // 洗码详情
	TotalBonusAmount float64                  `json:"total_bonus_amount"`                                                                              // 总返水金额
	MinBonusAmount   float64                  `json:"min_bonus_amount"`
}

type FcUserRebateValsResult struct {
	VenueCode   string  `json:"-"`            // 场馆码
	GameType    string  `json:"game_type"`    // 游戏类型
	BetAmount   float64 `json:"bet_amount"`   // 有效投注
	BonusRate   float64 `json:"bonus_rate"`   // 洗码比例
	BonusAmount float64 `json:"bonus_amount"` // 返水金额
}

type UpdatePhoneReq struct {
	Phone    string `json:"phone"  form:"phone" `
	VeryCode string `json:"code"  form:"code" validate:"code"`
}

type FcUserInviteReq struct {
	UserId     string `json:"user_id"`
	InviteCode int    `json:"invite_code"`
}

type FcInviteDomainReq struct {
	InviteCode int `json:"invite_code"`
}

type FcInviteDomainResp struct {
	InviteCode   int    `gorm:"column:invite_code" json:"invite_code"`          // 推广ID
	Domain       string `gorm:"column:domain" json:"domain"`                    // 代理专属域名短码
	JumpLink     string `gorm:"column:jump_link" json:"jump_link"`              // 跳转后域名
	Type         int    `gorm:"column:type" json:"type" form:"type" uri:"type"` // 类型（1 官网 2 推广）
	CustomerLink string `gorm:"column:customer_link" json:"customer_link"`      // 客服链接
	IosLink      string `gorm:"column:ios_link" json:"ios_link"`                // ios链接
	IosLink2     string `gorm:"column:ios_link2" json:"ios_link2"`              // ios备用链接
	AndroidLink  string `gorm:"column:android_link" json:"android_link"`        // 安卓链接
	AndroidLink2 string `gorm:"column:android_link2" json:"android_link2"`      // 安卓备用链接
	BannerImg    string `gorm:"column:banner_img" json:"banner_img"`            // banner图
	LogoImg      string `gorm:"column:logo_img" json:"logo_img"`                // logo图
}

type CustomerLinkReq struct {
	MerchantCode string `json:"merchant_code" form:"merchant_code"`
}
