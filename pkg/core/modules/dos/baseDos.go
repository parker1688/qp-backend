package dos

import (
	"bootpkg/common/tool"
	"gorm.io/gorm"
)

type BaseDos struct {
	Id string `gorm:"column:id" json:"id" form:"id" uri:"id" `
}

func (u *BaseDos) BeforeCreate(tx *gorm.DB) error {
	if len(u.Id) == 0 {
		u.Id = tool.SnowflakeIdByKey(tx.Statement.Table)
	}
	return nil
}
