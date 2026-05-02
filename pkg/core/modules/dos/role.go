package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Role struct {
	BaseDos
	RoleName   string             `gorm:"column:name" json:"role_name" form:"role_name" uri:"role_name" `                    // 角色名称
	Sort       int64              `gorm:"-" json:"sort" form:"sort" uri:"sort" `                                              // 当前表无 sort 列
	MeusIds    string             `gorm:"column:meus_ids" json:"meus_ids" form:"meus_ids" uri:"meus_ids" `                  // 菜单ID集合,分割
	Status     int64              `gorm:"column:status" json:"status" form:"status" uri:"status" `                           // 状态: 1 启用 0: 停用
	Remarks    string             `gorm:"column:remark" json:"remarks" form:"remarks" uri:"remarks" `                        // 备注
	PermsList  string             `gorm:"column:perms_list" json:"perms_list" form:"perms_list" uri:"perms_list" `          // 权限列表
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `              // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `              // 修改人
}

func (Role) TableName() string {
	return "role"
}
