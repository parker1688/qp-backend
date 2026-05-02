package vo

import "bootpkg/common/expands/automaticType"

type NotifyTemplateRequest struct {
	Title           string             `gorm:"column:title" json:"title" form:"title" uri:"title" validate:"required" `                                             // 模板主题
	NotifyFlag      string             `gorm:"column:notify_flag" json:"notify_flag" form:"notify_flag" uri:"notify_flag" validate:"required"`                      // 模板标识
	TemplateContent string             `gorm:"column:template_content" json:"template_content" form:"template_content" uri:"template_content" validate:"required" ` // 模板内容
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                            // 创建时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                 // 创建人
}
