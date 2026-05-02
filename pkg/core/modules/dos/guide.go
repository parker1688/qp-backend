package dos

import (
	"bootpkg/common/expands/automaticType"
)

type Guide struct {
	BaseDos
	Key        string             `gorm:"column:key" json:"key" form:"key" uri:"key"`                                              // 入口key
	Name       string             `gorm:"column:name" json:"name" form:"name" uri:"name"`                                          // 入口名称
	Data       string             `gorm:"column:data" json:"data" form:"data" uri:"data"`                                          // 数据
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time"` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by"`                              // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time"` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by"`                              // 修改人
}

func (Guide) TableName() string {
	return "guides"
}

type GuideInfoResp struct {
	BaseDos
	Key  string `gorm:"column:key" json:"key" form:"key" uri:"key"`
	Name string `gorm:"column:name" json:"name" form:"name" uri:"name"`
	Data string `gorm:"column:data" json:"data" form:"data" uri:"data"`
}
