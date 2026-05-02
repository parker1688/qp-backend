package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserTask struct {
	UserId    string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `             // 用户ID
	Data      string             `gorm:"column:data" json:"data" form:"data" uri:"data" `                         // 用户任务数据
	CreatedAt automaticType.Time `gorm:"column:created_at" json:"created_at" form:"created_at" uri:"created_at" ` // 创建时间
}

func (FcUserTask) TableName() string {
	return "fc_user_task"
}
