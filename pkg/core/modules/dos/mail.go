package dos

import (
	"bootpkg/common/expands/automaticType"
)

// 后台邮件表
type Mail struct {
	BaseDos
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code"`      // 商户号
	UserIds      string             `gorm:"column:user_ids" json:"user_ids" form:"user_ids" uri:"user_ids"`                          // 用户ID组（空为全部用户）
	Type         int                `gorm:"column:type" json:"type" form:"type" uri:"type"`                                          // 类型（0 人工发送；1 充值成功；2 提款失败；3 提款成功）
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title"`                                      // 标题
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content"`                              // 内容
	IsPopup      int                `gorm:"column:is_popup" json:"is_popup" form:"is_popup" uri:"is_popup"`                          // 是否弹窗（0 否；1 是）
	IsKeep       int                `gorm:"column:is_keep" json:"is_keep" form:"is_keep" uri:"is_keep"`                              // 是否保留（0 否；1 是）
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time"` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by"`                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time"` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by"`                      // 修改人
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status"`                                  // 状态（0 关闭；1 开启）
}

func (Mail) TableName() string {
	return "mail"
}

// 用户邮件表
type FcUserMail struct {
	BaseDos
	MsgId        string             `gorm:"column:msg_id" json:"-" form:"msg_id" uri:"msg_id"`                                       // 消息ID（后台发送邮件ID）
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id"`                              // 用户ID
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code"`      // 商户号
	Type         int                `gorm:"column:type" json:"type" form:"type" uri:"type"`                                          // 类型（0 人工发送；1 充值成功；2 提款失败；3 提款成功）
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title"`                                      // 标题
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content"`                              // 内容
	IsPopup      int                `gorm:"column:is_popup" json:"is_popup" form:"is_popup" uri:"is_popup"`                          // 是否弹窗（0 否；1 是）
	IsKeep       int                `gorm:"column:is_keep" json:"is_keep" form:"is_keep" uri:"is_keep"`                              // 是否保留（0 否；1 是）：有弹窗弹窗后删除；没弹窗则阅读后删除
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time"` // 创建时间
	ReadStatus   int                `gorm:"column:read_status" json:"read_status" form:"read_status" uri:"read_status"`              // 已读状态（1 未读；2 已读）
	DelStatus    int                `gorm:"column:del_status" json:"-" form:"del_status" uri:"del_status"`                           // 删除状态（0 未删除；1 已删除）
	Status       int                `gorm:"column:status;default:1" json:"status" form:"status" uri:"status" `                       // 状态                                // 状态（保留）
}

func (FcUserMail) TableName() string {
	return "fc_user_mail"
}
