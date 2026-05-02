package vo

type AgentAwardReportDayReq struct {
	Grade     int    `json:"grade" form:"grade" uri:"grade" validate:"required"` //等级
	PageSize  int    `json:"page_size" form:"page_size" uri:"page_size"`         //当前页大小
	PageIndex int    `json:"page_index" form:"page_index" uri:"page_index"`      //当前页码
	StartTime string `json:"start_time" form:"start_time" uri:"start_time"`      //开始时间
	EndTime   string `json:"end_time" form:"end_time" uri:"end_time"`            //结束时间
}

type AgentAwardReportDayResp struct {
	ReportDate     string  `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `                     // 统计日期
	DepositUserNum int     `gorm:"column:deposit_user_num" json:"deposit_user_num" form:"deposit_user_num" uri:"deposit_user_num" ` // 充值用户
	DepositAward   float64 `gorm:"column:deposit_award" json:"deposit_award" form:"deposit_award" uri:"deposit_award" `             // 首次充值奖励
	GameBetAmount  float64 `gorm:"column:game_bet_amount" json:"game_bet_amount" form:"game_bet_amount" uri:"game_bet_amount" `     // 游戏投注流水
	GameBetAward   float64 `gorm:"column:game_bet_award" json:"game_bet_award" form:"game_bet_award" uri:"game_bet_award" `         // 游戏交易奖励
	TotalAward     float64 `gorm:"column:total_award" json:"total_award" form:"total_award" uri:"total_award" `
	Grade          int     `gorm:"column:grade" json:"grade" form:"grade" uri:"grade" `                         // 级别
	IssueAward     float64 `gorm:"column:issue_award" json:"issue_award" form:"issue_award" uri:"issue_award" ` // 发放金额
	Status         int     `gorm:"column:status" json:"status" form:"status" uri:"status" `                     // 1:待发放  2：已发放  3：部分发放
}

type AgentReportResp struct {
	RegisterNum            int                           `gorm:"column:register_num" json:"register_num" form:"register_num" uri:"register_num" `
	TotalAward             float64                       `gorm:"column:total_award" json:"total_award" form:"total_award" uri:"total_award" `
	InviteNum              int                           `gorm:"column:invite_num" json:"invite_num" form:"invite_num" uri:"invite_num" `
	IsAgent                bool                          `gorm:"column:is_agent" json:"is_agent" form:"is_agent" uri:"is_agent" `
	ActivityFriendsSetting []*ActivityFriendsSettingResp `gorm:"column:activityFriendsSetting" json:"activityFriendsSetting" form:"activityFriendsSetting" uri:"activityFriendsSetting" `
}

type ActivityFriendsSettingResp struct {
	Name          string  `gorm:"column:name" json:"name" form:"name" uri:"name" `                                         // 活动名称
	ValidRecharge float64 `gorm:"column:valid_recharge" json:"valid_recharge" form:"valid_recharge" uri:"valid_recharge" ` // 最低有效充值
	MinGameTimes  int     `gorm:"column:min_game_times" json:"min_game_times" form:"min_game_times" uri:"min_game_times" ` // 最小游戏次数
	MaxGameTimes  int     `gorm:"column:max_game_times" json:"max_game_times" form:"max_game_times" uri:"max_game_times" ` // 最大游戏次数
	MinPeople     int     `gorm:"column:min_people" json:"min_people" form:"min_people" uri:"min_people" `                 // 最小人数
	MaxPeople     int     `gorm:"column:max_people" json:"max_people" form:"max_people" uri:"max_people" `                 // 最大人数
	Desc          string  `gorm:"column:desc" json:"desc" form:"desc" uri:"desc" `                                         // 说明
	Sort          int     `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `
}

type RedpacketTaskResp struct {
	GameTimes int     `gorm:"column:game_times" json:"game_times" form:"game_times" uri:"game_times" ` // 每人抢红包次数
	CycleType string  `gorm:"column:cycle_type" json:"cycle_type" form:"cycle_type" uri:"cycle_type" `
	InviteNum int     `gorm:"column:invite_num" json:"invite_num" form:"invite_num" uri:"invite_num" `
	FriendNum int     `gorm:"column:friend_num" json:"friend_num" form:"friend_num" uri:"friend_num" `
	Bouns     float64 `gorm:"column:bouns" json:"bouns" form:"bouns" uri:"bouns" `
	Status    int     `gorm:"column:status" json:"status" form:"status" uri:"status" `
	Desc      string  `gorm:"column:desc" json:"desc" form:"desc" uri:"desc" `
}

type RedpacketTaskBonusReq struct {
	CycleType string `gorm:"column:cycle_type" json:"cycle_type" form:"cycle_type" uri:"cycle_type" `
}
