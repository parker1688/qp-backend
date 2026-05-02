package vo

import "bootpkg/common/expands/automaticType"

type UserNotifyMessageResp struct {
	Id         string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 消息标题
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 消息内容
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	Status     int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 读取状态 0 未读 1 已读
}

type UserNotifyMessageDetailResp struct {
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 消息标题
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 消息内容
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	Status     int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 读取状态 0 未读 1 已读
}

type NotifyMarqueeResp struct {
	Id        string `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Content   string `gorm:"column:content" json:"content" form:"content" uri:"content" `         // 消息内容
	Status    int    `gorm:"column:status" json:"status" form:"status" uri:"status" `             // 读取状态 0 未读 1 已读
	Frequency int    `gorm:"column:frequency" json:"frequency" form:"frequency" uri:"frequency" ` // 频率
}

type BulletinResp struct {
	Id           string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                 // 公告名称
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                         // 公告内容
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                     // 排序从大到小
	IsDisplay    int                `gorm:"column:is_display" json:"is_display" form:"is_display" uri:"is_display" `             // 是否显示 1:显示 2:不显示
	BulletinType int                `gorm:"column:bulletin_type" json:"bulletin_type" form:"bulletin_type" uri:"bulletin_type" ` // 1:常规 2:临时维护 3:节日定时 4:活动 5:其它
	ContentType  int                `gorm:"column:content_type" json:"content_type" form:"content_type" uri:"content_type" `     // 内容类型 1:文字 2:图片 3:文字+图片
	BulletinImg  string             `gorm:"column:bulletin_img" json:"bulletin_img" form:"bulletin_img" uri:"bulletin_img" `     // 公告图片
	StartTime    automaticType.Time `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `             // 开始时间
	EndTime      automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                     // 结束时间
}

type FcUserSiteMsgList struct {
	Id         string             `gorm:"column:id" json:"id" form:"id" uri:"id" `
	UserId     string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName   string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Title      string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 消息标题
	Content    string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 消息内容
	MsgType    int                `gorm:"column:msg_type" json:"msg_type" form:"msg_type" uri:"msg_type" `                          // 消息类型 1 普通信息 2 赛事  3 充值 4提款 5 红利
	ReadStatus int                `gorm:"column:read_status" json:"read_status" form:"read_status" uri:"read_status" `              // 已读状态, 1:未读, 2:已读
	DelStatus  int                `gorm:"column:del_status" json:"del_status" form:"del_status" uri:"del_status" `                  // 删除状态, 1:未删除, 2:已删除
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
}

type FcUserMailResp struct {
	Id         string             `json:"id"`
	Type       int                `json:"type"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	CreateTime automaticType.Time `json:"create_time"`
	ReadStatus int                `json:"read_status"`
	IsPopup    int                `json:"is_popup"`
}

type FcUserSiteMsgReadReq struct {
	Id string `gorm:"column:id" json:"id" form:"id" uri:"id" `
}
