package vo

type VenueRegisterReq struct {
	VenueCode string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	Currency  string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
}

type VenueBalanceReq struct {
	VenueCode string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	Currency  string `json:"currency" form:"currency" uri:"currency" ` // 币种简码
}

type VenueTransferReq struct {
	VenueCode string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	//TransferType string  `json:"transferType" form:"transferType" uri:"transferType" validate:"required"`
	Amount   float64 `json:"amount" form:"amount" uri:"amount" validate:"required"` //金额
	Currency string  `json:"currency" form:"currency" uri:"currency" `              // 币种简码
}

type VenueTransferConfirmReq struct {
	VenueCode string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	OrderSn   string `json:"order_sn" form:"order_sn" uri:"order_sn" validate:"required"`
}

type VenueLaunchReq struct {
	VenueCode  string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	GameCode   string `json:"game_code" form:"game_code" uri:"game_code" `    //游戏
	Currency   string `json:"currency" form:"currency" uri:"currency" `       // 币种简码
	ReturnUrl  string `json:"returnUrl" form:"returnUrl" uri:"returnUrl" `    //回调url
	CashierURL string `json:"cashierURL" form:"cashierURL" uri:"cashierURL" ` //存款地址
	TableId    string `json:"tableId" form:"tableId" uri:"tableId" `          //桌号
	Gtype      *int   `json:"gtype" form:"gtype" uri:"gtype" `                // 游戏种类
}

type VenueLaunchResp struct {
	GameUrl string `json:"gameUrl" form:"gameUrl" uri:"gameUrl" `
}

type VenueBalanceResp struct {
	Balance float64 `json:"balance" form:"balance" uri:"balance" `
}

type FreeVenueLaunchReq struct {
	VenueCode  string `json:"venue_code" form:"venue_code" uri:"venue_code" validate:"required"`
	GameCode   string `json:"game_code" form:"game_code" uri:"game_code" validate:"required"`    //游戏
	Currency   string `json:"currency" form:"currency" uri:"currency" validate:"required"`       // 币种简码
	ReturnUrl  string `json:"returnUrl" form:"returnUrl" uri:"returnUrl" validate:"required"`    //回调url
	CashierURL string `json:"cashierURL" form:"cashierURL" uri:"cashierURL" validate:"required"` //存款地址
	UserName   string `json:"username" form:"username" uri:"username" validate:"required"`       //用户账号
	Password   string `json:"password" form:"password" uri:"password" validate:"required"`       //用户
	Token      string `json:"token" form:"token" uri:"token" validate:"required"`                //用户
	Language   string `json:"language" form:"language" uri:"language" validate:"required"`       //语言
	IP         string `json:"ip" form:"ip" uri:"ip" validate:"required"`                         //语言
}
