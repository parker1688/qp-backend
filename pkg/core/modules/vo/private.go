package vo

type FcPrivate struct {
	Tel            string `gorm:"column:tel" json:"tel" form:"tel" uri:"tel" `                                                 // 电话
	RealName       string `gorm:"column:real_name" json:"real_name" form:"real_name" uri:"real_name" `                         // 真实姓名
	Email          string `gorm:"email" json:"email" form:"email" uri:"email" `                                                // 邮箱
	Alipay         string `gorm:"column:alipay" json:"alipay" form:"alipay" uri:"alipay" `                                     // 支付宝                                // 支付宝
	AlipayRealname string `gorm:"column:alipay_realname" json:"alipay_realname" form:"alipay_realname" uri:"alipay_realname" ` // 支付宝姓名
	WalletPassword string `gorm:"column:wallet_password" json:"wallet_password" form:"wallet_password" uri:"wallet_password" ` // 支付密码
	AccountNumber  string `gorm:"account_number" json:"account_number" form:"account_number" uri:"account_number" `            // 卡号
	AccountHolder  string `gorm:"account_holder" json:"account_holder" form:"account_holder" uri:"account_holder" `            // 收款人
}
