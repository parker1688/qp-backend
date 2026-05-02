package dos

import (
	"bootpkg/common/expands/automaticType"
)

type UserGroup struct {
	BaseDos
	GroupName    string             `gorm:"column:group_name" json:"group_name" form:"group_name" uri:"group_name" `                  // 组名
	Data         string             `gorm:"column:data" json:"data" form:"data" uri:"data" `                                          // 数据
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
}

func (UserGroup) TableName() string {
	return "user_group"
}
