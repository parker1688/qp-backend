package vo

import "bootpkg/common/expands/automaticType"

type RedPacketSendHistoryListResp struct {
	Id          string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	CountAmount float64            `gorm:"column:count_amount" json:"count_amount" form:"count_amount" uri:"count_amount" `          // 总红包金额
	Num         int                `gorm:"column:num" json:"num" form:"num" uri:"num" `                                              // 红包数量
	GrantNum    int                `gorm:"column:grant_num" json:"grant_num" form:"grant_num" uri:"grant_num" `                      // 已发放数量
	AmountUp    float64            `gorm:"column:amount_up" json:"amount_up" form:"amount_up" uri:"amount_up" `                      // 最佳金额
	UpNameNick  string             `gorm:"column:up_name_nick" json:"up_name_nick" form:"up_name_nick" uri:"up_name_nick" `          // 最佳用户昵称
	NickName    string             `gorm:"column:nick_name" json:"nick_name" form:"nick_name" uri:"nick_name" `                      // 发放用户昵称
	PhotoLink   string             `gorm:"column:photo_link" json:"photo_link" form:"photo_link" uri:"photo_link" `                  // 发放用户头像地址
	CreateTime  automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
}

type RedPacketRecordResp struct {
	PhotoLink   string                `gorm:"column:photo_link" json:"photo_link" form:"photo_link" uri:"photo_link" `                  // 头像地址
	UserName    string                `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户名称
	BonusUp     automaticType.BitBool `gorm:"column:bonus_up" json:"bonus_up" form:"bonus_up" uri:"bonus_up" `                          // 最高奖金
	BonusLow    automaticType.BitBool `gorm:"column:bonus_low" json:"bonus_low" form:"bonus_low" uri:"bonus_low" `                      // 最低奖金
	Amount      float64               `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 红包金额
	Bonus       float64               `gorm:"column:bonus" json:"bonus" form:"bonus" uri:"bonus" `                                      // 领取金额
	ExtraAmount float64               `gorm:"column:extra_amount" json:"extra_amount" form:"extra_amount" uri:"extra_amount" `          // 额外金额
	UserId      string                `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                              // 用户ID
	UpdateTime  automaticType.Time    `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 抢红包时间
}
