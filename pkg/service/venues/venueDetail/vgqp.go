package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
	"time"
)

var (
	VGQPTransferProccesCode = map[int]string{
		642: "玩家在游戏中",
		105: "错误",
	}
)

type VenueVGQP struct {
	VenueConfig conf.VGQP
}

func NewVGQP(venueConfig *conf.VGQP) IVenues {
	return &VenueVGQP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueVGQP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token
	agent := ""
	merchantCodeLen := len(request.MerchantCode)
	if merchantCodeLen > 1 && merchantCodeLen < 7 {
		agent = request.MerchantCode
	}

	url := v.VenueConfig.Url
	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s&agent=%s",
		request.UserName, agent)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/CreateUser?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-CreateUser-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPCommonResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVGQP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token

	url := v.VenueConfig.Url
	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s",
		request.UserName)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/GetBalance?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPCommonVResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.Money
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVGQP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueVGQP) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueVGQP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token

	url := v.VenueConfig.Url
	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s&amount=%v&serial=%s",
		request.UserName, request.Amount, request.OrderSn)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/Deposit?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-Deposit-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPCommonVResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueVGQP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token

	url := v.VenueConfig.Url
	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s&amount=%v&serial=%s",
		request.UserName, request.Amount, request.OrderSn)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/Withdraw?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-Withdraw-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPCommonVResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVGQP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token

	url := v.VenueConfig.Url
	gameType := request.GameCode
	if gameType == "" {
		gameType = "1000" // 默认为游戏大厅
	}

	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s&gameType=%s",
		request.UserName, gameType)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/LoginWithChannel?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-LoginGame-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPLoginGameResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.Value
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVGQP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	loginResp := VGQPLogin(v.VenueConfig)
	if loginResp.Code != vo.Login_SUCCESS_CODE {
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = loginResp.Msg
		return &venueResp
	}
	token := loginResp.Data.Token

	url := v.VenueConfig.Url

	headerParams := map[string]string{
		"apitoken": token,
	}

	paramStr := fmt.Sprintf("username=%s&serial=%s",
		request.UserName, request.OrderSn)
	apiUrl := url + "/ChannelApi/API/" + v.VenueConfig.Channel + "/GetTransRecord?" + paramStr

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-LoginGame-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPCommonResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

func (v VenueVGQP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueVGQP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueVGQP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

// VGQP 的 token校验涉及到 ip 校验，如果分布式部署的话会有一定问题，故最好每次请求都获取token
func VGQPLogin(v conf.VGQP) *vo.VenueLoginResponse {
	url := v.Url
	venueResp := vo.VenueLoginResponse{}
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	keyOrg := v.Channel + timestampStr + v.Password
	keyEncrypt := strings.ToUpper(tool.MD5([]byte(keyOrg)))

	paramStr := fmt.Sprintf("channel=%s&timestamp=%s&verifycode=%s",
		v.Channel, timestampStr, keyEncrypt)
	apiUrl := url + "/ChannelApi/Security/GetToken?" + paramStr

	client := resty.New()
	resp, err := client.R().
		Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("VGQP-VGQPLogin-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.VGQPLoginResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = vo.Login_SUCCESS_CODE
		venueResp.Data.Token = respData.Value
		return &venueResp
	}

	venueResp.Code = vo.Login_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}
