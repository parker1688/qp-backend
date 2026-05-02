package dos

import (
	"bootpkg/common/expands/automaticType"
)

type DictsDetail struct {
	BaseDos
	DictsTypeCode string             `gorm:"dicts_type_code" json:"dicts_type_code" form:"dicts_type_code" uri:"dicts_type_code" ` // 字典Code码
	DictsTag      string             `gorm:"dicts_tag" json:"dicts_tag" form:"dicts_tag" uri:"dicts_tag" `                         // 字典标签
	DictsValue    string             `gorm:"dicts_value" json:"dicts_value" form:"dicts_value" uri:"dicts_value" `                 // 字典值
	Sort          int64              `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                             // 排序
	Status        int64              `gorm:"status" json:"status" form:"status" uri:"status" `                                     // 状态 0禁用 1启用
	Remarks       string             `gorm:"remarks" json:"remarks" form:"remarks" uri:"remarks" `                                 // 备注
	CreateTime    automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `    // 创建时间
	CreateBy      string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                         // 创建人
	UpdateTime    automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `    // 修改时间
	UpdateBy      string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                         // 修改人
}

func (DictsDetail) TableName() string {
	return "dicts_detail"
}
