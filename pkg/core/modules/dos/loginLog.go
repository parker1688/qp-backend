package dos

import (
	"bootpkg/common/expands/automaticType"
)

type LoginLog struct {
	BaseDos
	UserName   string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Ip         string             `gorm:"ip" json:"ip" form:"ip" uri:"ip" `
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `
}

func (LoginLog) TableName() string {
	return "login_log"
}
