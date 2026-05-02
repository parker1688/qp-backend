package dos

import (
	"bootpkg/common/expands/automaticType"
)

type MailTemplate struct {
	BaseDos
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 标题
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 内容
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (MailTemplate) TableName() string {
	return "mail_template"
}
