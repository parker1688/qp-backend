package dos

type GameInfo struct {
	Id             string `gorm:"column:Id" json:"Id" form:"Id" uri:"Id" `
	Gamecnname     string `gorm:"column:GameCnName" json:"GameCnName" form:"GameCnName" uri:"GameCnName" `
	Gameenname     string `gorm:"column:GameEnName" json:"GameEnName" form:"GameEnName" uri:"GameEnName" `
	Gamecode       string `gorm:"column:GameCode" json:"GameCode" form:"GameCode" uri:"GameCode" `
	Classid        string `gorm:"column:ClassId" json:"ClassId" form:"ClassId" uri:"ClassId" `
	Orderid        string `gorm:"column:OrderId" json:"OrderId" form:"OrderId" uri:"OrderId" `
	Imgurl         string `gorm:"column:ImgUrl" json:"ImgUrl" form:"ImgUrl" uri:"ImgUrl" `
	Platform       string `gorm:"column:Platform" json:"Platform" form:"Platform" uri:"Platform" `
	Isenable       string `gorm:"column:IsEnable" json:"IsEnable" form:"IsEnable" uri:"IsEnable" `
	Hasdemo        string `gorm:"column:HasDemo" json:"HasDemo" form:"HasDemo" uri:"HasDemo" `
	Createtime     string `gorm:"column:CreateTime" json:"CreateTime" form:"CreateTime" uri:"CreateTime" `
	Updatetime     string `gorm:"column:UpdateTime" json:"UpdateTime" form:"UpdateTime" uri:"UpdateTime" `
	Createoperator string `gorm:"column:CreateOperator" json:"CreateOperator" form:"CreateOperator" uri:"CreateOperator" `
	Updateoperator string `gorm:"column:UpdateOperator" json:"UpdateOperator" form:"UpdateOperator" uri:"UpdateOperator" `
	Issupporthtml5 string `gorm:"column:IsSupportHtml5" json:"IsSupportHtml5" form:"IsSupportHtml5" uri:"IsSupportHtml5" `
	Html5code      string `gorm:"column:Html5Code" json:"Html5Code" form:"Html5Code" uri:"Html5Code" `
	Ishot          string `gorm:"column:IsHot" json:"IsHot" form:"IsHot" uri:"IsHot" `
	Isrecommend    string `gorm:"column:IsRecommend" json:"IsRecommend" form:"IsRecommend" uri:"IsRecommend" `
}

func (GameInfo) TableName() string {
	return "Game_Info"
}
