package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcLoginLog struct {
	BaseDos
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Ip           string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `
	ClientType   string             `gorm:"column:client_type" json:"client_type" form:"client_type" uri:"client_type" `
	Version      string             `gorm:"column:version" json:"version" form:"version" uri:"version" `
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
	Citys        string             `gorm:"column:citys" json:"citys" form:"citys" uri:"citys" `                                 // 城市解析
	VisitorId    string             `gorm:"column:visitor_id" json:"visitor_id" form:"visitor_id" uri:"visitor_id" `             // 设备号
}

func (FcLoginLog) TableName() string {
	return "fc_login_log"
}
