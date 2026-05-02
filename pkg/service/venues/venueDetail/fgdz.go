package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
)

const (
	FGDZ_CreateUser_API      = "/v3/players"
	FGDZ_GetBalance_API      = "/v3/player_chips/member_code/"
	FGDZ_Deposit_API         = "/v3/player_uchips/member_code/"
	FGDZ_Withdraw_API        = "/v3/player_uchips/member_code/"
	FGDZ_LoginGame_API       = "/v3/launch_game"
	FGDZ_LoginLobby_API      = "/v3/launch_lobby"
	FGDZ_TransferConfirm_API = "/v3/player_uchips_check/"
)

var (
	FGDZTransferProccesCode = map[int]string{
		116: "代理商请求过多,被阻止",
		119: "单号不存在或者该注单失败",
		201: "api 内部错误",
		202: "重试",
		206: "超时",
		208: "充值转账正在处理中",
	}
)

type VenueFGDZ struct {
	VenueConfig conf.FGDZ
}

func NewFGDZ(venueConfig *conf.FGDZ) IVenues {
	return &VenueFGDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenueFGDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_CreateUser_API
	reqMap := map[string]string{
		"member_code": request.UserName,
		"password":    request.Password,
	}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-CreateUser-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZCreateUserResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueFGDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_GetBalance_API + request.UserName

	reqMap := map[string]string{}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZGetBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		amount := float64(respData.Data.Balance) / 100
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueFGDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueFGDZ) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueFGDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_Deposit_API + request.UserName
	amount := int(request.Amount * 100)

	reqMap := map[string]string{
		"amount":                strconv.Itoa(amount),
		"externaltransactionid": request.OrderSn,
	}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-Deposit-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZCommonResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueFGDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_Withdraw_API + request.UserName
	amount := -int(request.Amount * 100)

	reqMap := map[string]string{
		"amount":                strconv.Itoa(amount),
		"externaltransactionid": request.OrderSn,
	}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-Withdraw-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZCommonResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

func (v VenueFGDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	// 判断 gameCode 是否为空
	if request.GameCode == "" {
		return FGDZLoginLobby(v.VenueConfig, request)
	}

	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_LoginGame_API
	gameType := "app"
	if request.ClientType == "web" || request.ClientType == "h5" {
		gameType = "h5"
	}
	language := strings.ToLower(request.Language)

	reqMap := map[string]string{
		"member_code": request.UserName,
		//"game_code":   request.GameCode, // 长度不超过 20
		"game_id":    request.GameCode, //int 3307
		"game_type":  gameType,
		"language":   language,
		"ip":         "127.0.0.1",
		"return_url": request.ReturnUrl,
	}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-LoginGame-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZLoginGameResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.Data.GameUrl + "&token=" + respData.Data.Token
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueFGDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}

	url := v.VenueConfig.Url
	apiUrl := url + FGDZ_TransferConfirm_API + request.OrderSn

	reqMap := map[string]string{}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": v.VenueConfig.Merchantname,
		"merchantcode": v.VenueConfig.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-TransferConfirm-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZCommonResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respStr
	_, ok := FGDZTransferProccesCode[respData.Code]
	if ok {
		venueResp.Code = TransferConfirm_Processing_CODE
	}

	return &venueResp
}

func (v VenueFGDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueFGDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueFGDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func FGDZLoginLobby(conf conf.FGDZ, request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := conf.Url
	apiUrl := url + FGDZ_LoginLobby_API
	language := strings.ToLower(request.Language)

	reqMap := map[string]string{
		"member_code": request.UserName,
		"language":    language,
		"lobby_code":  "chess", // 目前该场馆只支持棋牌大厅
	}
	paramBytes, _ := tool.JsonMarshal(reqMap)
	paramStr := string(paramBytes)

	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"merchantname": conf.Merchantname,
		"merchantcode": conf.Merchantcode,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("FGDZ-FGDZLoginLobby-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.FGDZLoginLobbyResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.Data.LobbyUrl
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}
