package dos

import (
	"bootpkg/common/expands/automaticType"
)

type AdminUser struct {
	BaseDos
	UserName           string             `gorm:"column:username" json:"username" form:"username" uri:"username" `                                           // 用户名称-登录名称
	UserNick           string             `gorm:"column:user_nick" json:"user_nick" form:"user_nick" uri:"user_nick" `                                     // 用户昵称
	AccountType        int                `gorm:"column:account_type" json:"account_type" form:"account_type" uri:"account_type" `                         // 账号类型 1:超级管理员 2:商户管理员 3:普通账号
	Mobile             string             `gorm:"column:mobile" json:"mobile" form:"mobile" uri:"mobile" `                                                 // 手机号
	DepartmentId       string             `gorm:"column:department_id" json:"department_id" form:"department_id" uri:"department_id" `                     // 部门ID
	RoleIds            string             `gorm:"column:role_ids" json:"role_ids" form:"role_ids" uri:"role_ids" `                                         // 角色ID集合,分割
	TotalAmount        float64            `gorm:"column:total_amount" json:"total_amount" form:"total_amount" uri:"total_amount" `                         // 总额度
	LimitPertimeAmount float64            `gorm:"column:limit_pertime_amount" json:"limit_pertime_amount" form:"limit_pertime_amount" uri:"limit_pertime_amount" ` // 每次限额
	CurAmount          float64            `gorm:"column:cur_amount" json:"cur_amount" form:"cur_amount" uri:"cur_amount" `                                 // 当前使用额度
	Status             int64              `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                 // 状态: 1 启用 0: 停用
	Pwd                string             `gorm:"column:pwd" json:"pwd" form:"pwd" uri:"pwd" `                                                             // 登录密码
	Remarks            string             `gorm:"-" json:"remarks" form:"remarks" uri:"remarks" `                                                           // 当前表结构无 remarks 列
	CreateTime         automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy           string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime         automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy           string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
	Token              string             `gorm:"-" sql:"-" json:"token"`                                                                                   // token
	EnforcePwd         int                `gorm:"column:enforce_pwd" json:"enforce_pwd" form:"enforce_pwd" uri:"enforce_pwd" `                             // 0 不强制 1 强制修改
	Mfa                string             `gorm:"column:mfa" json:"mfa" form:"mfa" uri:"mfa" `                                                             // MFA
	MfaHour            int                `gorm:"column:mfa_hour" json:"mfa_hour" form:"mfa_hour" uri:"mfa_hour" `                                         // MFA
	MerchantCodes      string             `gorm:"column:merchant_codes" json:"merchant_codes" form:"merchant_codes" uri:"merchant_codes" `                 // 商户Code集合逗号,分割
}

// 增加枚举状态
const (
	USER_STATUS_DISABLED = 0
	USER_STATUS_ENABLED  = 1
)

func (AdminUser) TableName() string {
	return "admin_user"
}
