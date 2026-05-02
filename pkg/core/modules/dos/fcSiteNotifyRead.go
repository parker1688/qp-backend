package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteNotifyRead struct {
	BaseDos
	SiteNotifyId string             `gorm:"column:site_notify_id" json:"site_notify_id" form:"site_notify_id" uri:"site_notify_id" `  // 公告ID
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                              // 用户ID
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 1. 已读取  2. 删除
}

func (FcSiteNotifyRead) TableName() string {
	return "fc_site_notify_read"
}
