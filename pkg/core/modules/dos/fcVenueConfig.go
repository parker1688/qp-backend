package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueConfig struct {
	BaseDos
	Alias      string             `gorm:"alias" json:"alias" form:"alias" uri:"alias" `                     // 别名
	VenueName  string             `gorm:"venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" ` // 场馆名称
	VenueCode  string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆code
	Remark     string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `
	Content    string             `gorm:"content" json:"content" form:"content" uri:"content" `
	CreateBy   string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
}

func (FcVenueConfig) TableName() string {
	return "fc_venue_config"
}
