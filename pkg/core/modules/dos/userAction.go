package dos

import (
	"bootpkg/common/expands/automaticType"
)

type UserAction struct {
	BaseDos
	UserName   string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 账号
	Url        string             `gorm:"url" json:"url" form:"url" uri:"url" `                         // 路由
	Method     string             `gorm:"method" json:"method" form:"method" uri:"method" `             // 访问方式
	Ip         string             `gorm:"ip" json:"ip" form:"ip" uri:"ip" `                             // 访问IP
	Body       string             `gorm:"body" json:"body" form:"body" uri:"body" `
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `
}

func (UserAction) TableName() string {
	return "user_action"
}
