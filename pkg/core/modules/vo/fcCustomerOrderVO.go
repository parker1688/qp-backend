package vo

type CustomerOrderUpdateStatusVO struct {
	Id          string `gorm:"id" json:"id" form:"id" uri:"id" ` // 用户Id
	Status      int    `gorm:"status" json:"status" form:"status" uri:"status" `
	SolveRemark string `gorm:"column:solve_remark" json:"solve_remark" form:"solve_remark" uri:"solve_remark" `
	Currency    string `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `
}

type CustomerOrderTypeVO struct {
	Flag  string `gorm:"flag" json:"flag" form:"flag" uri:"flag" `
	Title string `gorm:"column:title" json:"title" form:"title" uri:"title" `
}
