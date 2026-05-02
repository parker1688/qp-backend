package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcBulletin struct {
	BaseDos
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 公告名称
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 公告内容
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序从大到小
	IsDisplay    int                `gorm:"column:is_display" json:"is_display" form:"is_display" uri:"is_display" `                  // 是否显示 1:显示 2:不显示
	BulletinType int                `gorm:"column:bulletin_type" json:"bulletin_type" form:"bulletin_type" uri:"bulletin_type" `      // 1:常规 2:临时维护 3:节日定时 4:活动 5:其它
	ContentType  int                `gorm:"column:content_type" json:"content_type" form:"content_type" uri:"content_type" `          // 内容类型 1:文字 2:图片 3:文字+图片
	BulletinImg  string             `gorm:"column:bulletin_img" json:"bulletin_img" form:"bulletin_img" uri:"bulletin_img" `          // 公告图片
	StartTime    automaticType.Time `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `                  // 开始时间
	EndTime      automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                          // 结束时间
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcBulletin) TableName() string {
	return "fc_bulletin"
}
