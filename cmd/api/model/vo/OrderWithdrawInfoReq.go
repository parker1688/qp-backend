package vo

type OrderWithdrawInfoReq struct {
	Currency  string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
	StartTime string `json:"startAt"`                                  //开始时间
	EndTime   string `json:"endAt"`                                    //结束时间
	PageSize  int    `json:"pageSize" `                                //当前页码
	PageIndex int    `json:"current"`                                  //当前页码
	Status    int    `json:"status"`                                   //当前页码
}

type PaymentOutInfoReq struct {
	ChannelCode string `json:"channel_code" form:"channel_code" uri:"channel_code" ` //渠道code
}

type PaymentOutInfoResp struct {
	Id          string `gorm:"column:id" json:"id" form:"id" uri:"id" `                                         // 通道名称
	PaymentName string `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" ` // 通道名称
	PaymentCode string `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
	ChannelName string `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" ` // 渠道名称
	ChannelCode string `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
	Status      int    `gorm:"column:status" json:"status" form:"status" uri:"status" `                         // 1:正常  2:禁止
	Sort        int    `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                 // 排序：值越大越靠前
	Icon        string `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	IconPath    string `gorm:"column:icon_path" json:"icon_path" form:"icon_path" uri:"icon_path" `
	Img         string `gorm:"column:img" json:"img" form:"img" uri:"img" `
	ImgPath     string `gorm:"column:img_path" json:"img_path" form:"img_path" uri:"img_path" `
}
