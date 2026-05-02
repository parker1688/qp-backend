package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserNotify struct {
	BaseDos
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 消息标题
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 消息内容
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Language     string             `gorm:"column:language" json:"language" form:"language" uri:"language" `                          // 语言简码
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 读取状态 0 未读 1 已读
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcUserNotify) TableName() string {
	return "fc_user_notify"
}
