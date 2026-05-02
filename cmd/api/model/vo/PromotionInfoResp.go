package vo

import "bootpkg/common/expands/automaticType"

type PromotionInfoResp struct {
	Id            string             `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	PromotionType int                `gorm:"column:promotion_type" json:"promotion_type" form:"promotion_type" uri:"promotion_type" ` // 优惠类型 1. 限时活动 2. 存款活动 3 日常活动
	GameType      int                `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                     // 游戏类型 0 默认显示 1 体育 2 真人 3 棋牌 4 电子 5 捕鱼  6  彩票  7  区块链
	Title         string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                     // 优惠标题
	Status        int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                 // 状态 1:开启 2:关闭
	H5Img         string             `gorm:"column:h5_img" json:"h5_img" form:"h5_img" uri:"h5_img" `                                 // H5优惠图片
	PromotionImg  string             `gorm:"column:promotion_img" json:"promotion_img" form:"promotion_img" uri:"promotion_img" `     // 优惠图片
	StartTime     automaticType.Time `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `                 // 开始时间
	EndTime       automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                         // 结束时间
	Link          string             `gorm:"column:link" json:"link" form:"link" uri:"link" `                                         // 优惠链接
	H5Link        string             `gorm:"column:h5_link" json:"h5_link" form:"h5_link" uri:"h5_link" `                             // h5优惠链接
	Content       string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                             // 优惠类型详情
	Sort          int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                         // 排序(最大越靠前)
}
