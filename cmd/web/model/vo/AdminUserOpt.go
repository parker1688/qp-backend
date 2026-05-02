package vo

type AdminUserAddReq struct {
	UserNick           string  `gorm:"user_nick" json:"user_nick" form:"user_nick" uri:"user_nick" `                                             // 用户昵称
	UserName           string  `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" validate:"required,min=4,max=16"`             // 用户名称-登录名称
	AccountType        int     `gorm:"account_type" json:"account_type" form:"account_type" uri:"account_type" `                                 // 账号类型 1:超级管理员 2:商户管理员 3:普通账号
	Status             int64   `gorm:"status" json:"status" form:"status" uri:"status" `                                                         // 状态: 1 启用 0: 停用
	Pwd                string  `gorm:"pwd" json:"pwd" form:"pwd" uri:"pwd" validate:"required,min=6,max=12"`                                     // 登录密码
	Remarks            string  `gorm:"remarks" json:"remarks" form:"remarks" uri:"remarks" `                                                     // 备注
	MerchantCodes      string  `gorm:"merchant_codes" json:"merchant_codes" form:"merchant_codes" uri:"merchant_codes" `                         // 商户Code集合逗号,分割
	TotalAmount        float64 `gorm:"total_amount" json:"total_amount" form:"total_amount" uri:"total_amount" `                                 // 总额度
	LimitPertimeAmount float64 `gorm:"limit_pertime_amount" json:"limit_pertime_amount" form:"limit_pertime_amount" uri:"limit_pertime_amount" ` // 每次限额
}
