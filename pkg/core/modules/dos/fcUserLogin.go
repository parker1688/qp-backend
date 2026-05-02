package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserLogin struct {
	BaseDos
	UserName     string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 账号
	Password     string             `gorm:"password" json:"password" form:"password" uri:"password" `                          // 密码
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Status       int                `gorm:"status" json:"status" form:"status" uri:"status" `                                  //状态 0 正常  1 禁用
}

func (FcUserLogin) TableName() string {
	return "fc_user_login"
}
