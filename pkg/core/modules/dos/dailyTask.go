package dos

import (
	"bootpkg/common/expands/automaticType"
)

type DailyTask struct {
	BaseDos
	MerchantCode     string              `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                     // 商户code
	Type             int                 `gorm:"column:type" json:"type" form:"type" uri:"type" `                                                         // 任务大类
	Subtype          int                 `gorm:"column:subtype" json:"subtype" form:"subtype" uri:"subtype" `                                             // 任务目标
	GroupId          string              `gorm:"column:groupid" json:"groupid" form:"groupid" uri:"groupid" `                                             // 任务分组
	Name             string              `gorm:"column:name" json:"name" form:"name" uri:"name" `                                                         // 任务名称
	Sort             int                 `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                                         // 任务排序
	Intro            string              `gorm:"column:intro" json:"intro" form:"intro" uri:"intro" `                                                     // 任务简介
	Detail           string              `gorm:"column:detail" json:"detail" form:"detail" uri:"detail" `                                                 // 任务详情
	Amount           float64             `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                                 // 任务额度
	BonusAmount      float64             `gorm:"column:bonus_amount" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                         // 奖励额度
	GameType         string              `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                                     // 参与游戏
	VenueCode        string              `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                                 // 参与厂商
	ChannelCode      string              `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `                         // 充值渠道
	StartAt          *automaticType.Time `gorm:"column:start_at;default:null" json:"start_at" form:"start_at" uri:"start_at" `                            // 任务开始时间
	EndAt            *automaticType.Time `gorm:"column:end_at;default:null" json:"end_at" form:"end_at" uri:"end_at" `                                    // 任务结束时间
	CreateTime       automaticType.Time  `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                // 创建时间
	CreateBy         string              `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime       automaticType.Time  `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                // 修改时间
	UpdateBy         string              `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                     // 修改人
	Status           int                 `gorm:"column:status;default:1" json:"status" form:"status" uri:"status" `                                       // 任务状态
	IncludeGameCodes string              `gorm:"column:include_game_codes" json:"include_game_codes" form:"include_game_codes" uri:"include_game_codes" ` // 包含游戏
	ExcludeGameCodes string              `gorm:"column:exclude_game_codes" json:"exclude_game_codes" form:"exclude_game_codes" uri:"exclude_game_codes" ` // 屏蔽游戏
	Cycle            int                 `gorm:"column:cycle;default:0" json:"cycle" form:"cycle" uri:"cycle" `                                           // 任务周期
}

func (DailyTask) TableName() string {
	return "daily_task"
}

type DailyTaskEx struct {
	DailyTask
	Merchant FcMerchant `json:"merchant" gorm:"foreignkey:MerchantCode;references:MerchantCode"`
}

type DailyTaskResp struct {
	DailyTask
	MerchantName string `json:"merchant_name"`
}

type DailyTaskMultiField struct {
	Id          string  `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Sort        int     `json:"sort"`
	Amount      float64 `json:"amount"`
	BonusAmount float64 `json:"bonus_amount"`
}
