package dos

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
)

type FcUserMaterial struct {
	UserId               string                `gorm:"column:user_id;primary_key" json:"user_id" form:"user_id" uri:"user_id" `                     // 用户Id
	UserName             string                `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                         // 用户账号
	NickName             string                `gorm:"column:nick_name" json:"nick_name" form:"nick_name" uri:"nick_name" `                         // 用户昵称
	RealName             string                `gorm:"column:real_name" json:"real_name" form:"real_name" uri:"real_name" `                         // 真实姓名
	ParentId             string                `gorm:"column:parent_id" json:"parent_id" form:"parent_id" uri:"parent_id" `                         // 上级Id
	AgentId              string                `gorm:"column:agent_id" json:"agent_id" form:"agent_id" uri:"agent_id" `                             // 代理Id
	AgentName            string                `gorm:"column:agent_name" json:"agent_name" form:"agent_name" uri:"agent_name" `                     // 代理账号
	Sex                  int                   `gorm:"column:sex" json:"sex" form:"sex" uri:"sex" `                                                 // 1:男  2:女
	Tel                  string                `gorm:"column:tel" json:"tel" form:"tel" uri:"tel" `                                                 // 电话
	Email                string                `gorm:"column:email" json:"email" form:"email" uri:"email" `                                         // 邮箱
	Qq                   string                `gorm:"column:qq" json:"qq" form:"qq" uri:"qq" `                                                     // qq
	Wx                   string                `gorm:"column:wx" json:"wx" form:"wx" uri:"wx" `                                                     // 微信
	Alipay               string                `gorm:"column:alipay" json:"alipay" form:"alipay" uri:"alipay" `                                     // 支付宝
	AlipayRealname       string                `gorm:"column:alipay_realname" json:"alipay_realname" form:"alipay_realname" uri:"alipay_realname" ` // 支付宝姓名
	Address              string                `gorm:"column:address" json:"address" form:"address" uri:"address" `                                 // 地址
	Birthday             automaticType.Time    `gorm:"column:birthday" json:"birthday" form:"birthday" uri:"birthday" `                             // 生日
	Level                int                   `gorm:"column:level" json:"level" form:"level" uri:"level" `                                         // 1~10 VIP等级
	Vip                  string                `gorm:"column:vip" json:"vip" form:"vip" uri:"vip" `                                                 // VIP显示
	IsWithdraw           int                   `gorm:"column:is_withdraw" json:"is_withdraw" form:"is_withdraw" uri:"is_withdraw" `                 // 是否可提款 1:可以 2:不可以
	IsBonus              int                   `gorm:"column:is_bonus" json:"is_bonus" form:"is_bonus" uri:"is_bonus" `                             // 是否可领取福利 1:可以 2:不可以
	MerchantCode         string                `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `         // 商户code
	Nation               string                `gorm:"column:nation" json:"nation" form:"nation" uri:"nation" `                                     // 国家
	Language             string                `gorm:"column:language" json:"language" form:"language" uri:"language" `                             // 语言
	Avatar               string                `gorm:"column:avatar" json:"avatar" form:"avatar" uri:"avatar" `                                     // 头像路径
	AgentSubId           string                `gorm:"column:agent_sub_id" json:"agent_sub_id" form:"agent_sub_id" uri:"agent_sub_id" `             // 子代理ID
	AgentSubName         string                `gorm:"column:agent_sub_name" json:"agent_sub_name" form:"agent_sub_name" uri:"agent_sub_name" `     // 子代理名称
	CreateTime           automaticType.Time    `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `    // 创建时间
	CreateBy             string                `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                         // 创建人
	UpdateTime           automaticType.Time    `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `    // 修改时间
	UpdateBy             string                `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                         // 修改人
	RegisterIp           string                `gorm:"column:register_ip" json:"register_ip" form:"register_ip" uri:"register_ip" `
	LastLoginIp          string                `gorm:"column:last_login_ip" json:"last_login_ip" form:"last_login_ip" uri:"last_login_ip" `
	LastLoginTime        automaticType.Time    `gorm:"column:last_login_time" json:"last_login_time" form:"last_login_time" uri:"last_login_time" `
	ManualTransferWallet automaticType.BitBool `gorm:"column:manual_transfer_wallet" json:"manual_transfer_wallet" form:"manual_transfer_wallet" uri:"manual_transfer_wallet" `
	WalletPassword       string                `gorm:"column:wallet_password" json:"wallet_password" form:"wallet_password" uri:"wallet_password" `         // 钱包密码
	IsFree               automaticType.BitBool `gorm:"column:is_free" json:"is_free" form:"is_free" uri:"is_free" `                                         // 0:非试玩   1：试玩
	InviteCode           string                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `                         // 用户邀请短码
	AgentInviteCode      int                   `gorm:"column:agent_invite_code" json:"agent_invite_code" form:"agent_invite_code" uri:"agent_invite_code" ` // 上级推广码
	IsOfficialAgent      int                   `gorm:"column:is_official_agent" json:"is_official_agent" form:"is_official_agent" uri:"is_official_agent" ` // 是否官方代理, 1 是 2 不是
	LoginStatus          int                   `gorm:"column:login_status" json:"login_status" form:"login_status" uri:"login_status" `                     // 登录状态 0 正常 1 禁止
	Remark               string                `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                             // 用户备注
	LastLoginCount       int                   `gorm:"column:last_login_count" json:"last_login_count" form:"last_login_count" uri:"last_login_count" `     // 用户总登录次数
	IsVerification       automaticType.BitBool `gorm:"column:is_verification" json:"is_verification" form:"is_verification" uri:"is_verification" `
	VisitorId            string                `gorm:"column:visitor_id" json:"visitor_id" form:"visitor_id" uri:"visitor_id" `                      // 最后登录设备号
	RegistVisitorId      string                `gorm:"column:regist_visitor_id" json:"regist_visitor_id" form:"regist_visitor_id" uri:"visitor_id" ` // 注册设备号
	Website              string                `gorm:"column:website" json:"website" form:"website" uri:"website" `                                  // 网站
	Rank                 string                `gorm:"column:rank" json:"rank" form:"rank" uri:"rank" `                                              // 段位
	RankFlag             string                `gorm:"column:rank_flag" json:"rank_flag" form:"rank_flag" uri:"rank_flag" `                          // 段位标识
	DailyBonusData       string                `gorm:"column:dailybonus_data" json:"dailybonus_data" form:"dailybonus_data" uri:"dailybonus_data" `  // 签到数据
}

func (FcUserMaterial) TableName() string {
	return "fc_user_material"
}

func (m *FcUserMaterial) Encrypt() {
	if len(m.RealName) > 0 {
		m.RealName, _ = tool.EncryptAESPrefixAesKey(m.RealName, "fc_user_material")
	}
	if len(m.Tel) > 0 {
		m.Tel, _ = tool.EncryptAESPrefixAesKey(m.Tel, "fc_user_material")
	}
	if len(m.AlipayRealname) > 0 {
		m.AlipayRealname, _ = tool.EncryptAESPrefixAesKey(m.AlipayRealname, "fc_user_material")
	}
	if len(m.Email) > 0 {
		m.Email, _ = tool.EncryptAESPrefixAesKey(m.Email, "fc_user_material")
	}
}

func (m *FcUserMaterial) EncryptData(data string) string {
	val, _ := tool.EncryptAESPrefixAesKey(data, "fc_user_material")
	return val
}

func (m *FcUserMaterial) Decrypt() {
	m.RealName = tool.DecryptAESPrefixRandKeySaltDefault(m.RealName, "fc_user_material")
	m.Tel = tool.DecryptAESPrefixRandKeySaltDefault(m.Tel, "fc_user_material")
	m.AlipayRealname = tool.DecryptAESPrefixRandKeySaltDefault(m.AlipayRealname, "fc_user_material")
	m.Email = tool.DecryptAESPrefixRandKeySaltDefault(m.Email, "fc_user_material")
}

func (m *FcUserMaterial) DecryptData(data string) string {
	return tool.DecryptAESPrefixRandKeySaltDefault(data, "fc_user_material")
}

func (m *FcUserMaterial) DecryptData2(data string) string {
	return tool.DecryptAESDecodeStr(data, global.CONFIG.General.AesUserKey)
}
