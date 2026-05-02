package vo

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/pkg/core/modules/vo"
	"github.com/duke-git/lancet/v2/strutil"
)

type MaterialResponse struct {
	UserId               string                   `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `                                                // 用户账号
	UserName             string                   `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                                        // 用户账号
	NickName             string                   `gorm:"nick_name" json:"nick_name" form:"nick_name" uri:"nick_name" `                                        // 用户昵称
	RealName             string                   `gorm:"real_name" json:"real_name" form:"real_name" uri:"real_name" `                                        // 真实姓名
	Sex                  int                      `gorm:"sex" json:"sex" form:"sex" uri:"sex" `                                                                // 1:男  2:女 0:未知
	Tel                  string                   `gorm:"tel" json:"tel" form:"tel" uri:"tel" `                                                                // 电话
	Email                string                   `gorm:"email" json:"email" form:"email" uri:"email" `                                                        // 邮箱
	Address              string                   `gorm:"address" json:"address" form:"address" uri:"address" `                                                // 地址
	Birthday             string                   `gorm:"birthday" json:"birthday" form:"birthday" uri:"birthday" `                                            // 生日
	Level                int                      `gorm:"level" json:"level" form:"level" uri:"level" `                                                        // 1~10 VIP等级
	Vip                  string                   `gorm:"vip" json:"vip" form:"vip" uri:"vip" `                                                                // VIP显示
	AgentInviteCode      int                      `gorm:"column:agent_invite_code" json:"agent_invite_code" form:"agent_invite_code" uri:"agent_invite_code" ` // 上级推广码
	IsWithdraw           int                      `gorm:"column:is_withdraw" json:"is_withdraw" form:"is_withdraw" uri:"is_withdraw" `                         // 是否可提款 1:可以 2:不可以
	IsBonus              int                      `gorm:"column:is_bonus" json:"is_bonus" form:"is_bonus" uri:"is_bonus" `                                     // 是否可领取福利 1:可以 2:不可以
	Nation               string                   `gorm:"nation" json:"nation" form:"nation" uri:"nation" `                                                    // 国家
	Language             string                   `gorm:"language" json:"language" form:"language" uri:"language" `                                            // 语言
	Avatar               string                   `gorm:"avatar" json:"avatar" form:"avatar" uri:"avatar" `                                                    // 头像路径
	RegisterDay          int                      `gorm:"-" json:"registerDay" form:"registerDay" uri:"registerDay" `
	InviteCode           string                   `gorm:"-" json:"invite_code" form:"invite_code" uri:"invite_code" `                                                       //邀请短码
	ManualTransferWallet automaticType.BitBool    `gorm:"manual_transfer_wallet" json:"manual_transfer_wallet" form:"manual_transfer_wallet" uri:"manual_transfer_wallet" ` // 手动转账钱包
	IsFree               automaticType.BitBool    `gorm:"is_free" json:"is_free" form:"is_free" uri:"" `                                                                    // 0：非试玩   1：试玩用户
	WalletPassword       string                   `gorm:"column:wallet_password" json:"wallet_password" form:"wallet_password" uri:"wallet_password" `
	IsVerification       automaticType.BitBool    `gorm:"column:is_verification" json:"is_verification" form:"is_verification" uri:"is_verification" `
	RechargeInfo         vo.RechargeSuccessInfoVO `gorm:"-" json:"recharge_info" form:"recharge_info" uri:"recharge_info" `                         //充值成功信息
	CreateTime           automaticType.Time       `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	//Qq                   string                   `gorm:"qq" json:"qq" form:"qq" uri:"qq" `                                                                    // qq
	//Wx                   string                   `gorm:"wx" json:"wx" form:"wx" uri:"wx" `                                                                    // 微信
}

// Hide
//
//	@Description: 解密
//	@receiver m
func (m *MaterialResponse) Hide() {
	if len(m.RealName) > 3 {
		startStr := m.RealName[:3]
		m.RealName = startStr + "***"
	}
	if len(m.Tel) > 4 {
		m.Tel = strutil.HideString(m.Tel, 3, len(m.Tel), "*")
	}
	if len(m.Birthday) > 4 {
		m.Birthday = strutil.HideString(m.Birthday, 3, len(m.Birthday), "*")
	}
	if len(m.Email) > 4 {
		m.Email = strutil.HideString(m.Email, 3, len(m.Email), "*")
	}
}
