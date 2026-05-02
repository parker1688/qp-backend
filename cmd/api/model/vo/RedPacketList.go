package vo

import "bootpkg/common/expands/automaticType"

type RedPacketListResp struct {
	RoomId     string  `gorm:"column:room_id" json:"room_id" form:"room_id" uri:"room_id" `                 // 房间号
	RoomAmount string  `gorm:"column:room_amount" json:"room_amount" form:"room_amount" uri:"room_amount" ` // 房间显示金额
	Num        int     `gorm:"column:num" json:"num" form:"num" uri:"num" `                                 // 红包数量
	Amount     float64 `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                     // 红包金额
}

type RedPacketGuessResp struct {
	AvaAmount   float64 `gorm:"column:ava_amount" json:"ava_amount" form:"ava_amount" uri:"ava_amount" `         //可用金额
	BonusAmount float64 `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" ` //中奖金额
	BonusUp     bool    `json:"bonus_up"`                                                                        //最佳金额
	Count       int     `json:"count"`                                                                           //剩余红包数
}

type RedPacketJobResp struct {
	Id           string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                 // 标题简写(本月)(本周)
	TitleDesc    string             `gorm:"column:title_desc" json:"title_desc" form:"title_desc" uri:"title_desc" `             // 标题描述
	ExtraContent string             `gorm:"column:extra_content" json:"extra_content" form:"extra_content" uri:"extra_content" ` // 完成度内容
	Amount       float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                             // 奖励金额
	RType        int                `gorm:"column:r_type" json:"r_type" form:"r_type" uri:"r_type" `                             // 红包奖励类型 1. 推荐好友 2. 月推荐好友 3 今天抢红包  4. 充值满金额红包
	EndTime      automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                     // 结束时间
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                             //1、去完成 2、待领取 3、 已领取
}

type RedPacketJobResultResp struct {
	Id           string `gorm:"column:id" json:"id" form:"id" uri:"id" `
	ExtraContent string `gorm:"column:extra_content" json:"extra_content" form:"extra_content" uri:"extra_content" ` //扩展内容
	Status       int    `gorm:"column:status" json:"status" form:"status" uri:"status" `                             //状态 1、去完成 2、待领取 3、 已领取
}
