package vo

type BannerResp struct {
	BannerLink string `gorm:"column:banner_link" json:"banner_link" form:"banner_link" uri:"banner_link" ` // banner图片地址
	BannerHref string `gorm:"column:banner_href" json:"banner_href" form:"banner_href" uri:"banner_href" ` // banner跳转地址
	Sort       int    `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                             // 排序从大到小
}
