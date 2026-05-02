package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcMerchantMenus struct {
	BaseDos
	MenuName     string             `gorm:"menu_name" json:"menu_name" form:"menu_name" uri:"menu_name" `                      // 菜单名称
	Icon         string             `gorm:"icon" json:"icon" form:"icon" uri:"icon" `                                          // ICON图标
	Sort         int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序
	RoleFlag     string             `gorm:"role_flag" json:"role_flag" form:"role_flag" uri:"role_flag" `                      // 角色前端路由地址
	Address      string             `gorm:"address" json:"address" form:"address" uri:"address" `                              // 组件地址
	ParentId     int                `gorm:"parent_id" json:"parent_id" form:"parent_id" uri:"parent_id" `                      // 上一级ID
	Type         int                `gorm:"type" json:"type" form:"type" uri:"type" `                                          // 1. 目录  2.菜单  3.按钮
	ApiRegular   string             `gorm:"api_regular" json:"api_regular" form:"api_regular" uri:"api_regular" `              // 权限正则匹配
	ShowStatus   int                `gorm:"show_status" json:"show_status" form:"show_status" uri:"show_status" `              // 1 显示 0 不显示
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcMerchantMenus) TableName() string {
	return "fc_merchant_menus"
}
