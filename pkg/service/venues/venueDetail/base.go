package venueDetail

// IVenues
// @Description: 场馆公共接口
type IVenues interface {
	//
	// CreateUser
	//  @Description: 创建用户
	//  @param *VenueCreateUserRequest 请求结构体
	//  @return *VenueResponse 返回结构体
	//
	CreateUser(*VenueCreateUserRequest) *VenueResponse
	//
	// GetUserBalance
	//  @Description: 获取用户余额
	//  @param *VenueGetUserBalanceRequest 请求结构体
	//  @return *VenueGetUserBalanceResponse 返回结构体
	//
	GetUserBalance(*VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse
	//
	// UpdateUserPassword
	//  @Description: 修改用户密码
	//  @param *VenueUpdateUserPasswordRequest 请求结构体
	//  @return *VenueResponse 返回结构体
	//
	UpdateUserPassword(*VenueUpdateUserPasswordRequest) *VenueResponse
	//
	// GetOrderNo
	//  @Description: 获取转账和提现订单号(每个场馆订单号不相同)
	//  @return string 生成唯一订单号
	//
	GetOrderNo() string
	//
	// Deposit
	//  @Description: 存款
	//  @param *VenueDepositRequest 存款结构体
	//  @return *VenueResponse 返回结构体
	//
	Deposit(*VenueDepositRequest) *VenueDepositResponse
	//
	// AmountLimitFix
	//  @Description: 金额限制修正(有些场馆只能整数或者货币转换)
	//  @param amount 计算金额
	//  @param currency 币种简码
	//  @return float64 返回修正金额
	//
	AmountLimitFix(amount float64, currency string) float64
	//
	// Withdraw
	//  @Description: 提款
	//  @param *VenueWithdrawRequest
	//  @return *VenueResponse
	//
	Withdraw(*VenueWithdrawRequest) *VenueWithdrawResponse
	//
	// LoginGame
	//  @Description: 登录游戏
	//  @param *VenueLoginGameRequest
	//  @return *VenueResponse
	//
	LoginGame(*VenueLoginGameRequest) *VenueLoginGameResponse
	//
	// TransferConfirm
	//  @Description: 转账确认
	//  @return *VenueResponse
	//
	TransferConfirm(*VenueTransferConfirmRequest) *VenueResponse
	//
	// PullOrder
	//  @Description: 拉取订单
	//  @param *VenuePullOrderRequest
	//  @return *VenueResponse
	//
	PullOrder(*VenuePullOrderRequest) *VenuePullOrderResponse
	//
	// CallBackConfirm
	//  @Description: 回调确认
	//  @param *VenueCallBackConfirmRequest
	//  @return *VenueResponse
	//
	CallBackConfirm(*VenueCallBackConfirmRequest) *VenueResponse
}

type VenueCreateUserRequest struct {
	UserName     string //用户名
	NickName     string //昵称
	Password     string //账号密码
	Currency     string //币种
	Token        string //部分场馆的token
	Ip           string //客户端IP
	Product      string //集成场馆码
	OrderSn      string //部分场馆需要订单号如 lyqp
	MerchantCode string //部分场馆需要如 lyqp
	UserType     int    //账号类型，1:正式，2:试玩
}

type VenueGetUserBalanceRequest struct {
	UserName string //用户名
	Password string //账号密码
	Currency string //币种
	Ip       string //客户端IP
	Token    string //token
	Product  string //集成场馆码
}

type VenueGetUserBalanceResponse struct {
	Code      int    `json:"code"`       //0 成功
	Msg       string `json:"msg"`        //错误信息
	ThirdCode string `json:"third_code"` //部分场馆需要三方返回的 code
	Data      struct {
		Amount float64 `json:"amount"` //剩余金额
	} `json:"data"`
}

type BatchVenueBalanceRequest struct {
	UserName string //用户名
	Password string //账号密码
	Currency string //币种
	Ip       string //客户端IP
}

type BatchVenueBalanceResponse struct {
	VenueCode string  `json:"venueCode"` //场馆Code
	Amount    float64 `json:"amount"`    //剩余金额
}

type VenueUpdateUserPasswordRequest struct {
	UserName    string //用户名
	Password    string //账号密码
	NewPassword string //新密码
	Ip          string //ip
	Token       string //token
}

type VenueDepositRequest struct {
	UserName string  //用户名
	Password string  //账号密码
	Currency string  //币种
	Amount   float64 //存款金额
	OrderSn  string  //订单号
	Ip       string  //客户端IP
	Token    string  //token
	Product  string  //集成场馆码
}

type VenueDepositResponse struct {
	Code      int    `json:"code"`       //0 成功
	Msg       string `json:"msg"`        //错误信息
	ThirdCode string `json:"third_code"` //部分场馆需要三方返回的 code
	Data      struct {
		Amount        float64 `json:"amount"`        //存款金额
		TransactionId string  `json:"transactionId"` //三方订单号
	} `json:"data"`
}

type VenueWithdrawRequest struct {
	UserName string  //用户名
	Password string  //账号密码
	Currency string  //币种
	Amount   float64 //提款金额
	OrderSn  string  //订单号
	Ip       string  //客户端IP
	Token    string  //token
	Product  string  //集成场馆码
}

type VenueWithdrawResponse struct {
	Code      int    `json:"code"`       //0 成功
	Msg       string `json:"msg"`        //错误信息
	ThirdCode string `json:"third_code"` //部分场馆需要三方返回的 code
	Data      struct {
		Amount        float64 `json:"amount"`        //提款金额
		TransactionId string  `json:"transactionId"` //三方订单号
	} `json:"data"`
}

type VenueLoginGameRequest struct {
	UserId       string //用户ID
	UserName     string //用户名
	Password     string //账号密码
	Token        string //部分场馆的token
	Currency     string //部分场馆的币种
	IP           string //用户IP
	GameCode     string //游戏Code
	GType        string //gtype
	ClientType   string //客户端类型 h5、pc、android、 ios
	Language     string //语言
	ReturnUrl    string //回调url
	CashierURL   string //存款url
	TableId      string //桌号
	IsFree       bool   //试玩用户
	Product      string //集成场馆码
	OrderSn      string //部分场馆需要订单号如 lyqp
	MerchantCode string //部分场馆需要如 lyqp
}

type VenueLoginGameResponse struct {
	Code      int                        `json:"code"`       //0 成功
	Msg       string                     `json:"msg"`        //错误信息
	ThirdCode string                     `json:"third_code"` //部分场馆需要三方返回的 code
	Data      VenueLoginGameDataResponse `json:"data"`
}

type VenueLoginGameDataResponse struct {
	GameUrl   string `json:"gameUrl"` //启动地址
	ApiDomain string `json:"apiDomain"`
	ImgDomain string `json:"imgDomain"`
	HtmlBody  string `json:"html_body"`
	Token     string `json:"token"`
}

type VenueTransferConfirmRequest struct {
	UserName      string //用户名
	Password      string //账号密码
	Token         string //部分场馆的token
	IP            string //用户IP
	OrderSn       string //订单号
	TransactionId string //三方订单号
	TransferType  string //IN 存入游戏  OUT 从游戏提款
	Product       string //集成场馆码
	Credit        string //额度
}

type VenuePullOrderResponse struct {
	Code int    `json:"code"` //0 成功
	Msg  string `json:"msg"`  //错误信息
	Data []byte `json:"data"` //返回信息
}

type VenuePullOrderRequest struct {
	Product   string //集成场馆码
	StartTime string //开始时间
	EndTime   string //结束时间
	Page      int
	PageSize  int
	GameType  string
}

type VenueCallBackConfirmRequest struct {
}

type VenueResponse struct {
	Code      int         `json:"code"`       //0 成功
	Msg       string      `json:"msg"`        //错误信息
	ThirdCode string      `json:"third_code"` //部分场馆需要三方返回的 code
	Data      interface{} `json:"data"`       //返回信息
}

const (
	CreateUser_SUCCESS_CODE = 0
	CreateUser_FAIL_CODE    = -1

	GetUserBalance_SUCCESS_CODE = 0
	GetUserBalance_FAIL_CODE    = -1

	Withdraw_SUCCESS_CODE    = 0
	Withdraw_FAIL_CODE       = -1
	Withdraw_Processing_CODE = 2

	Deposit_SUCCESS_CODE    = 0
	Deposit_FAIL_CODE       = -1
	Deposit_Processing_CODE = 2

	LoginGame_SUCCESS_CODE = 0
	LoginGame_FAIL_CODE    = -1

	UpdateUserPassword_SUCCESS_CODE = 0
	UpdateUserPassword_FAIL_CODE    = -1

	TransferConfirm_SUCCESS_CODE    = 0
	TransferConfirm_FAIL_CODE       = -1
	TransferConfirm_Processing_CODE = 2

	PullOrder_SUCCESS_CODE = 0
	PullOrder_Fail_code    = -1
)
