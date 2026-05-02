package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Blacklist struct {
	BaseDos
	Type       int                `gorm:"column:type" json:"type" form:"type" uri:"type" `                                          // 类型（1 ip地址；2 设备码）
	Value      string             `gorm:"column:value" json:"value" form:"value" uri:"value" `                                      // 数据
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
}

func (Blacklist) TableName() string {
	return "blacklist"
}
