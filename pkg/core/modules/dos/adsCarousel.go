package dos

import (
	"bootpkg/common/expands/automaticType"
)

type AdsCarousel struct {
	BaseDos
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Key          string             `gorm:"column:key" json:"key" form:"key" uri:"key" `                                              // 广告栏key
	Name         string             `gorm:"column:name" json:"name" form:"name" uri:"name" `                                          // 广告栏中文名
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 广告栏排序
	IsCarousel   int                `gorm:"column:is_carousel;default:1" json:"is_carousel" form:"is_carousel" uri:"is_carousel" `    // 是否轮播
	Sources      string             `gorm:"column:sources" json:"sources" form:"sources" uri:"sources" `                              // 轮播图集数据
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
	Status       int                `gorm:"column:status;default:1" json:"status" form:"status" uri:"status" `                        // 状态
	Jumpto       string             `gorm:"column:jumpto" json:"jumpto" form:"jumpto" uri:"jumpto" `                                  // 活动跳转
}

func (AdsCarousel) TableName() string {
	return "ads_carousel"
}

type AdsCarouselResp struct {
	AdsCarousel
	SourceNum int `gorm:"-" json:"source_num"`
}

type AdsCarouselRes struct {
	BaseDos
	Key        string `json:"key"`
	Name       string `json:"name"`
	IsCarousel int    `json:"is_carousel"`
	Sources    string `json:"sources"`
	Jumpto     string `json:"jumpto"`
}
