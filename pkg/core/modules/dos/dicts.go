package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Dicts struct {
	BaseDos
	DictsName     string             `gorm:"dicts_name" json:"dicts_name" form:"dicts_name" uri:"dicts_name" ` // 菜单名称
	DictsTypeCode string             `gorm:"dicts_type_code" json:"dicts_type_code" form:"dicts_type_code" uri:"dicts_type_code" `
	Remarks       string             `gorm:"remarks" json:"remarks" form:"remarks" uri:"remarks" `                              // 备注
	CreateTime    automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (Dicts) TableName() string {
	return "dicts"
}
