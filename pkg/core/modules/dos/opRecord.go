package dos

import (
	"bootpkg/common/expands/automaticType"
)

type OpRecord struct {
	BaseDos
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 操作人
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                              // 被操作用户id
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"user_name" uri:"user_name" `              // 商户号
	IP           string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `                                                  // 操作ip
	Menu1        string             `gorm:"column:menu1" json:"menu1" form:"menu1" uri:"menu1" `                                      // 菜单1
	Menu2        string             `gorm:"column:menu2" json:"menu2" form:"menu2" uri:"menu2" `                                      // 菜单2
	Op           string             `gorm:"column:op" json:"op" form:"op" uri:"op" `                                                  // 操作
	Result       string             `gorm:"column:result" json:"result" form:"result" uri:"result" `                                  // 结果
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 操作时间
}

func (OpRecord) TableName() string {
	return "op_record"
}
