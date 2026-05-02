package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserRemark struct {
	BaseDos
	UserId     string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                              // 用户ID
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Remark     string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                  // 备注
}

func (FcUserRemark) TableName() string {
	return "fc_user_remark"
}
