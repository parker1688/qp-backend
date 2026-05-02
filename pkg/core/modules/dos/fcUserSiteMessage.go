package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserSiteMessage struct {
	BaseDos
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	MsgId        string             `gorm:"column:msg_id" json:"msg_id" form:"msg_id" uri:"msg_id" `                     // 消息ID
	MsgIdType    int                `gorm:"column:msg_id_type" json:"msg_id_type" form:"msg_id_type" uri:"msg_id_type" ` // 消息id类型, 1:人工发送全局信息 2: 模板信息
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                         // 消息标题
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                 // 消息内容
	MsgType      int                `gorm:"column:msg_type" json:"msg_type" form:"msg_type" uri:"msg_type" `             // 消息类型 1 普通信息 2 赛事  3 充值 4提款 5 红利
	NotifyType   int                `gorm:"column:notify_type" json:"notify_type" form:"notify_type" uri:"notify_type" ` // 通知类型 1 全局消息 2 部分通知
	ReadStatus   int                `gorm:"column:read_status" json:"read_status" form:"read_status" uri:"read_status" ` // 已读状态, 1:未读, 2:已读
	DelStatus    int                `gorm:"column:del_status" json:"del_status" form:"del_status" uri:"del_status" `     // 删除状态, 1:未删除, 2:已删除
	Language     string             `gorm:"column:language" json:"language" form:"language" uri:"language" `             // 语言简码
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcUserSiteMessage) TableName() string {
	return "fc_user_site_message"
}
