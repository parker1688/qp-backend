package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVisit struct {
	BaseDos
	UserId     string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                       // 用户ID
	InviteCode int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `       // 邀请码
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
}

func (FcVisit) TableName() string {
	return "fc_visit"
}
