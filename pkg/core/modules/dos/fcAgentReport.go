package dos

type FcAgentReport struct {
	BaseDos
	MerchantCode      string  `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	InviteCode        int     `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `
	GroupName         string  `gorm:"column:group_name" json:"group_name" form:"group_name" uri:"group_name" `
	RegistCount       int     `gorm:"column:regist_count" json:"regist_count" form:"regist_count" uri:"regist_count"`
	FirstPayPlayerNum int     `gorm:"column:first_pay_player_num" json:"first_pay_player_num" form:"first_pay_player_num" uri:"first_pay_player_num"`
	PayPlayerNum      int     `gorm:"column:pay_player_num" json:"pay_player_num" form:"pay_player_num" uri:"pay_player_num"`
	FirstPayMoney     float64 `gorm:"column:first_pay_money" json:"first_pay_money" form:"first_pay_money" uri:"first_pay_money"`
	PayMoney          float64 `gorm:"column:pay_money" json:"pay_money" form:"pay_money" uri:"pay_money"`
	PayCount          int     `gorm:"column:pay_count" json:"pay_count" form:"pay_count" uri:"pay_count"`
	VisitCount        int     `gorm:"column:visit_count" json:"visit_count" form:"visit_count" uri:"visit_count"`
	DownloadCount     int     `gorm:"column:download_count" json:"download_count" form:"download_count" uri:"download_count"`
	StartAt           string  `gorm:"column:start_at" json:"start_at" form:"start_at" uri:"start_at"`
	EndAt             string  `gorm:"column:end_at" json:"end_at" form:"end_at" uri:"end_at"`
	Days              int     `gorm:"column:days" json:"days" form:"days" uri:"days"`
}

func (FcAgentReport) TableName() string {
	return "fc_agent_report"
}
