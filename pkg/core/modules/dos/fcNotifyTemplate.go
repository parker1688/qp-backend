package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcNotifyTemplate struct {
	BaseDos
	Title           string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                             // 模板主题
	TemplateContent string             `gorm:"column:template_content" json:"template_content" form:"template_content" uri:"template_content" ` // 模板内容
	NotifyFlag      string             `gorm:"column:notify_flag" json:"notify_flag" form:"notify_flag" uri:"notify_flag" `                     // 模板标识
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                             // 创建人
	UpdateTime      automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy        string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                             // 修改人
	MerchantCode    string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
}

func (FcNotifyTemplate) TableName() string {
	return "fc_notify_template"
}
