package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Department struct {
	BaseDos
	DepartmentName string             `gorm:"department_name" json:"department_name" form:"department_name" uri:"department_name" ` // 部门名称
	Sort           int64              `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                             // 排序
	ParentId       int64              `gorm:"parent_id" json:"parent_id" form:"parent_id" uri:"parent_id" `                         // 上级ID
	Status         int64              `gorm:"status" json:"status" form:"status" uri:"status" `                                     // 状态: 1 启用 0: 停用
	Remarks        string             `gorm:"remarks" json:"remarks" form:"remarks" uri:"remarks" `                                 // 备注
	CreateTime     automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `    // 创建时间
	CreateBy       string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                         // 创建人
	UpdateTime     automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `    // 修改时间
	UpdateBy       string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                         // 修改人
}

func (Department) TableName() string {
	return "department"
}
