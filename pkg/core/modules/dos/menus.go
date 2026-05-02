package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Menus struct {
	BaseDos
	Name       string             `gorm:"name" json:"name" form:"name" uri:"name" `                                // 菜单名称（原数据库字段）
	MenuName   string             `gorm:"menu_name" json:"menu_name" form:"menu_name" uri:"menu_name" `            // 菜单名称
	Icon       string             `gorm:"icon" json:"icon" form:"icon" uri:"icon" `                                // ICON图标
	Sort       int64              `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                // 排序
	RoleFlag   string             `gorm:"role_flag" json:"role_flag" form:"role_flag" uri:"role_flag" `            // 角色前端路由地址
	Address    string             `gorm:"address" json:"address" form:"address" uri:"address" `                    // 组件地址
	ParentId   string             `gorm:"parent_id" json:"parent_id" form:"parent_id" uri:"parent_id" `            // 上一级ID
	Type       int64              `gorm:"type" json:"type" form:"type" uri:"type" `                                // 1. 目录  2.菜单  3.按钮
	ApiRegular string             `gorm:"api_regular" json:"api_regular" form:"api_regular" uri:"api_regular" `    // 权限正则匹配
	Perms      string             `gorm:"perms" json:"perms" form:"perms" uri:"perms" `                            // perms
	ShowStatus int64              `gorm:"show_status" json:"show_status" form:"show_status" uri:"show_status" `    // 1 显示 0 不显示
	OpenCache  int64              `gorm:"open_cache" json:"open_cache" form:"open_cache" uri:"open_cache" `        // 1 显示 0 不显示
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"-" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"create_by" json:"-" form:"create_by" uri:"create_by" `                    // 创建人
	UpdateTime automaticType.Time `gorm:"update_time;default:null" json:"-" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"update_by" json:"-" form:"update_by" uri:"update_by" `                    // 修改人
	Locales    string             `gorm:"locales" json:"locales" form:"locales" uri:"locales" `                    // 国际化标识
}

func (Menus) TableName() string {
	return "menu"
}

type MenusResp struct {
	Menus
	ParentMenuName string `gorm:"parent_menu_name" json:"parent_menu_name" form:"parent_menu_name" uri:"parent_menu_name" `
	ParentIcon     string `gorm:"parent_icon" json:"parent_icon" form:"parent_icon" uri:"parent_icon" `
	ParentRoleFlag string `gorm:"parent_role_flag" json:"parent_role_flag" form:"parent_role_flag" uri:"parent_role_flag" `
}
