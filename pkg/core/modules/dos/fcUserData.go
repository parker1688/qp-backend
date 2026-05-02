package dos

type FcUserData struct {
	UserId       string `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `                         // 用户ID
	ActivityData string `gorm:"column:activity_data" json:"activity_data" form:"activity_data" uri:"activity_data" ` // 活动数据
}

func (FcUserData) TableName() string {
	return "fc_user_data"
}
