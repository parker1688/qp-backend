package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	url2 "net/url"
)

const (
	LYQP_API = "/channelHandle"
)

var (
	LYQPTransferProccesCode map[int]string = map[int]string{
		27: "数据库异常",
		32: "更新玩家信息失败",
		33: "更新玩家金币失败",
		44: "订单正在处理中",
		98: "进入游戏发生错误",
		99: "未知的错误",
	}
)

type VenueLYQP struct {
	VenueConfig conf.LYQP
}

func NewLYQP(venueConfig *conf.LYQP) IVenues {
	return &VenueLYQP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueLYQP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}

	req := VenueLoginGameRequest{}
	req.UserName = request.UserName
	req.OrderSn = request.OrderSn
	req.Currency = request.Currency
	req.MerchantCode = request.MerchantCode

	loginResp := v.LoginGame(&req)
	if loginResp == nil {
		tmpStr := fmt.Sprintf("TYQP-CreateUser username=%v return nil", req.UserName)
		global.G_LOG.Errorf(tmpStr)
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = tmpStr
		return &venueResp
	}

	if loginResp.Code == LoginGame_SUCCESS_CODE {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = loginResp.Msg
	return &venueResp
}

func (v VenueLYQP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	timestamp := tool.TimeNowTimestampString()

	paramOrg := fmt.Sprintf("s=7&account=%s&currency=%s",
		request.UserName, request.Currency)

	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(paramOrg, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf("LYQP-GetUserBalance username=%s AesEcbPk7EncryptBase64 err: %v", request.UserName, err)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptStr = url2.QueryEscape(paramEncryptStr)
	// key agent+timestamp+md5key
	keyOrg := v.VenueConfig.Agent + timestamp + v.VenueConfig.Md5key
	keyOrgEncrypt := tool.MD5([]byte(keyOrg))

	paramStr := fmt.Sprintf("agent=%s&timestamp=%s&param=%s&key=%s",
		v.VenueConfig.Agent, tool.TimeNowTimestampString(), paramEncryptStr, keyOrgEncrypt)
	apiUrl := url + LYQP_API + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("LYQP-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.LYQPBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 {
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = respData.D.FreeMoney // 可下分余额
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueLYQP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueLYQP) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueLYQP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	timestamp := tool.TimeNowTimestampString()
	amount := request.Amount

	paramOrg := fmt.Sprintf("s=2&account=%s&money=%v&orderid=%s&currency=%s",
		request.UserName, amount, request.OrderSn, request.Currency)

	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(paramOrg, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf("LYQP-Deposit username=%s AesEcbPk7EncryptBase64 err: %v", request.UserName, err)
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptStr = url2.QueryEscape(paramEncryptStr)
	// key agent+timestamp+md5key
	keyOrg := v.VenueConfig.Agent + timestamp + v.VenueConfig.Md5key
	keyOrgEncrypt := tool.MD5([]byte(keyOrg))

	paramStr := fmt.Sprintf("agent=%s&timestamp=%s&param=%s&key=%s",
		v.VenueConfig.Agent, tool.TimeNowTimestampString(), paramEncryptStr, keyOrgEncrypt)
	apiUrl := url + LYQP_API + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("LYQP-Deposit-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.LYQPDepositResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 && amount > 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueLYQP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	timestamp := tool.TimeNowTimestampString()
	amount := request.Amount

	paramOrg := fmt.Sprintf("s=3&account=%s&money=%v&orderid=%s&currency=%s",
		request.UserName, amount, request.OrderSn, request.Currency)

	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(paramOrg, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf("LYQP-Withdraw username=%s AesEcbPk7EncryptBase64 err: %v", request.UserName, err)
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptStr = url2.QueryEscape(paramEncryptStr)

	// key agent+timestamp+md5key
	keyOrg := v.VenueConfig.Agent + timestamp + v.VenueConfig.Md5key
	keyOrgEncrypt := tool.MD5([]byte(keyOrg))

	paramStr := fmt.Sprintf("agent=%s&timestamp=%s&param=%s&key=%s",
		v.VenueConfig.Agent, tool.TimeNowTimestampString(), paramEncryptStr, keyOrgEncrypt)
	apiUrl := url + LYQP_API + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("LYQP-Withdraw-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.LYQPWithdrawResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 && amount > 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueLYQP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	timestamp := tool.TimeNowTimestampString()
	kindId := request.GameCode

	// 代理编号 + yyyyMMddHHmmssSSS+ account
	orderid := request.OrderSn
	paramOrg := fmt.Sprintf("s=0&account=%s&money=0&orderid=%s&ip=127.0.0.1&lineCode=%s&KindID=%s&currency=%s",
		request.UserName, orderid, request.MerchantCode, kindId, request.Currency)
	//global.G_LOG.Infof("LYQP-LoginGame username=%s paramOrg: %s", request.UserName, paramOrg)

	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(paramOrg, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf("LYQP-LoginGame username=%s AesEcbPk7EncryptBase64 err: %v", request.UserName, err)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptStr = url2.QueryEscape(paramEncryptStr)

	// key agent+timestamp+md5key
	keyOrg := v.VenueConfig.Agent + timestamp + v.VenueConfig.Md5key
	keyOrgEncrypt := tool.MD5([]byte(keyOrg))

	paramStr := fmt.Sprintf("agent=%s&timestamp=%s&param=%s&key=%s",
		v.VenueConfig.Agent, tool.TimeNowTimestampString(), paramEncryptStr, keyOrgEncrypt)
	apiUrl := url + LYQP_API + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("LYQP-LoginGame-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)

	//global.G_LOG.Infof(apiLog)

	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.LYQPLaunchGameResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.D.Url
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueLYQP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	timestamp := tool.TimeNowTimestampString()

	paramOrg := fmt.Sprintf("s=4&orderid=%s", request.OrderSn)
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(paramOrg, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf("LYQP-TransferConfirm username=%s AesEcbPk7EncryptBase64 err: %v", request.UserName, err)
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptStr = url2.QueryEscape(paramEncryptStr)

	// key agent+timestamp+md5key
	keyOrg := v.VenueConfig.Agent + timestamp + v.VenueConfig.Md5key
	keyOrgEncrypt := tool.MD5([]byte(keyOrg))

	paramStr := fmt.Sprintf("agent=%s&timestamp=%s&param=%s&key=%s",
		v.VenueConfig.Agent, tool.TimeNowTimestampString(), paramEncryptStr, keyOrgEncrypt)
	apiUrl := url + LYQP_API + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("LYQP-TransferConfirm-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.LYQPTransferConfirmResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 {
		switch respData.D.Status {
		case -1, 2:
			venueResp.Code = TransferConfirm_FAIL_CODE
			return &venueResp
		case 0:
			venueResp.Code = TransferConfirm_SUCCESS_CODE
			return &venueResp
		}
	}

	venueResp.Code = TransferConfirm_Processing_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueLYQP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueLYQP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueLYQP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

// 获取乐游棋牌特殊的订单号
func GetLYQPOrderSn(conf conf.LYQP, req *vo.VenueGetOrderSnReq) string {
	orderSn := conf.Agent + tool.TimeNowYMDHMSMNoSpaceString() + req.UserName
	return orderSn
}
