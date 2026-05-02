package dos

type FcMerchant struct {
	BaseDos
	MerchantName    string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `                 // 商户名称
	MerchantCode    string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户编码
	Logo            string `gorm:"column:logo" json:"logo" form:"logo" uri:"logo" `                                                     // logo地址
	Status          int    `gorm:"column:status" json:"status" form:"status" uri:"status" `                                             // 1:开启  2:关闭
	AgentInviteCode int    `gorm:"column:agent_invite_code" json:"agent_invite_code" form:"agent_invite_code" uri:"agent_invite_code" ` // 官方代理ID
	Currency        string `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                     // 货币：多个用英文逗号（,）隔开
	Prefix          string `gorm:"column:prefix" json:"prefix" form:"prefix" uri:"prefix" `                                             // 游戏账号前缀
}

func (FcMerchant) TableName() string {
	return "fc_merchant"
}
