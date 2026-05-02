package vo

import "bootpkg/common/expands/automaticType"

type AnnouncementResp struct {
	Id         string             `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 公告名称
	TitleImg   string             `gorm:"column:title_img" json:"title_img" form:"title_img" uri:"title_img" `                      // 公告圖片
	ClassType  int                `gorm:"column:class_type" json:"class_type" form:"class_type" uri:"class_type" `                  // 平台公告分类类型 1. 公告 2.  赛事  3. 充提
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	Status     int                `gorm:"column:Status" json:"Status" form:"Status" uri:"Status" `
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" ` // 平台公告分类类型 1. 已阅读 0 未阅读
}

type AnnouncementDetailResp struct {
	Id         string             `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 公告名称
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 公告内容
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
}
