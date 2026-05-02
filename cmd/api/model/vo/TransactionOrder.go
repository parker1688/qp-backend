package vo

import "bootpkg/common/expands/automaticType"

type TransactionOrderReq struct {
	Currency    string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
	StartAt     string `json:"startAt"`                                  //开始时间
	EndAt       string `json:"endAt"`                                    //结束时间
	PageSize    int    `json:"pageSize"`                                 //当前页大小
	Current     int    `json:"current"`                                  //当前页码
	FundingType int    `json:"funding_type"`                             //转变类型
}

type TransactionOrderResp struct {
	Id          string             `gorm:"id" json:"id" form:"id" uri:"id" `
	Amount      float64            `gorm:"amount" json:"amount" form:"amount" uri:"amount"`                         // 金额
	FundingType int                `gorm:"funding_type" json:"funding_type" form:"funding_type" uri:"funding_type"` // 资金类型：  1：公司入款  2：在线存款 3：手动存款 4：提款 5：优惠 6：返水 7：手续费 8：管理员添加 9：管理员扣除 10：额度转换
	CreateTime  automaticType.Time `gorm:"create_time" json:"create_time" form:"create_time" uri:"create_time"`     // 创建时间
	Currency    string             `gorm:"currency" json:"currency" form:"currency" uri:"currency"`                 // 币种类型
	Status      int                `gorm:"status" json:"status" form:"status" uri:"status"`                         // 状态 0 处理中 1 成功 2 失败
	VenueCode   string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code"`
	OptType     int                `gorm:"opt_type" json:"opt_type" form:"opt_type" uri:"opt_type"`
	TrsType     int                `gorm:"trs_type" json:"trs_type" form:"trs_type" uri:"trs_type"`
}
