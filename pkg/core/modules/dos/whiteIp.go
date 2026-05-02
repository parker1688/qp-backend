package dos

import (
	"bootpkg/common/expands/automaticType"
)

type WhiteIp struct {
	BaseDos
	IpAddr     string             `gorm:"column:ip_addr" json:"ip_addr" form:"ip_addr" uri:"ip_addr" `                              // IPV4/IPV6
	Remarks    string             `gorm:"column:remarks" json:"remarks" form:"remarks" uri:"remarks" `                              // 备注
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
}

func (WhiteIp) TableName() string {
	return "white_ip"
}
