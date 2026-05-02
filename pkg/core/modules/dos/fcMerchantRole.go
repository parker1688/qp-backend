package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcMerchantRole struct {
	BaseDos
	RoleName     string             `gorm:"role_name" json:"role_name" form:"role_name" uri:"role_name" `                      // 角色名称
	Sort         int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序
	MeusIds      string             `gorm:"meus_ids" json:"meus_ids" form:"meus_ids" uri:"meus_ids" `                          // 菜单ID集合,分割
	Status       int                `gorm:"status" json:"status" form:"status" uri:"status" `                                  // 状态: 1 启用 0: 停用
	Remarks      string             `gorm:"remarks" json:"remarks" form:"remarks" uri:"remarks" `                              // 备注
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcMerchantRole) TableName() string {
	return "fc_merchant_role"
}
