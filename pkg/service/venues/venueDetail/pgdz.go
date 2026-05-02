package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm/utils"
	"math"
	"strconv"
)

type VenuePGDZ struct {
	VenueConfig conf.PGDZ
}

func NewPGDZ(venueConfig *conf.PGDZ) IVenues {
	return &VenuePGDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenuePGDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	appId := v.VenueConfig.AppID
	appSecret := v.VenueConfig.AppSecret
	url := v.VenueConfig.Url
	apiUrl := url + "/v2/player/create"
	param := "UserID=" + request.UserName + "&" + appSecret
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"UserID": request.UserName,
		"Sign":   sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"AppID":        appId,
		"AppSecret":    appSecret,
	}
	//global.G_LOG.Infof("PGDZ-Login info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("PGDZ-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("PGDZ-Login info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PGDZResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error != "" {
		global.G_LOG.Errorf("%s create_user error", apiLog)
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}

	if respData.Data.Pid > 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenuePGDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	appSecret := v.VenueConfig.AppSecret
	apiUrl := url + "/v2/player/balance"
	param := "UserID=" + request.UserName + "&" + appSecret
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"UserID": request.UserName,
		"Sign":   sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"AppID":        appId,
		"AppSecret":    appSecret,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("PGDZ-GetUserBalance-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.PGDZBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Error)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// PGDZ 金额单位分
	if respData.Code == 0 {
		amount := float64(respData.Data.Balance)
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenuePGDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenuePGDZ) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenuePGDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	//global.G_LOG.Infof("deposit ---------------1:%v", request.Amount)
	amount := strconv.FormatFloat(math.Floor(request.Amount), 'f', -1, 64)
	amount1 := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	//global.G_LOG.Infof("deposit ---------------2:%v, %v", amount, amount1)
	appId := v.VenueConfig.AppID
	appSecret := v.VenueConfig.AppSecret
	traceId := tool.SnowflakeIdByKey("pg-traceid")
	apiUrl := url + "/v2/player/transferIn"
	param := "UserID=" + request.UserName + "&Amount=" + utils.ToString(amount) + "&TraceId=" + traceId + "&" + appSecret
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"UserID":  request.UserName,
		"Amount":  amount1,
		"TraceId": traceId,
		"Sign":    sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"AppID":        appId,
		"AppSecret":    appSecret,
	}
	global.G_LOG.Infof("PGDZ Deposit ----params:%v, head:%v, apiurl:%v", param, headerParams, apiUrl)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s PGDZ-Deposit-httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("PGDZ-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PGDZDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.AfterBalance
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenuePGDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	amount := strconv.FormatFloat(math.Floor(request.Amount), 'f', -1, 64)
	amount1 := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))

	appId := v.VenueConfig.AppID
	appSecret := v.VenueConfig.AppSecret
	traceId := tool.SnowflakeIdByKey("pg")
	apiUrl := url + "/v2/player/transferOut"
	param := "UserID=" + request.UserName + "&Amount=" + amount + "&TraceId=" + traceId + "&" + appSecret
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"UserID":  request.UserName,
		"Amount":  amount1,
		"TraceId": traceId,
		"Sign":    sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"AppID":        appId,
		"AppSecret":    appSecret,
	}
	global.G_LOG.Infof("PGDZ Withdraw ----params:%v, head:%v, apiurl:%v", param, headerParams, apiUrl)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("PGDZ-Withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	var respData venuevo.PGDZWithdrawResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.AfterBalance
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenuePGDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	gameId := "pg_" + request.GameCode
	appId := v.VenueConfig.AppID
	appSecret := v.VenueConfig.AppSecret
	apiUrl := url + "/v2/game/launch"
	param := "UserID=" + request.UserName + "&GameID=" + gameId + "&Language=zh&" + appSecret
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"UserID":   request.UserName,
		"GameID":   gameId,
		"Language": "zh",
		"Sign":     sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"AppID":        appId,
		"AppSecret":    appSecret,
	}
	//global.G_LOG.Infof("PGDZ Login ----params:%v, head:%v, apiurl:%v", param, headerParams, apiUrl)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("PGDZ-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	respBytes := []byte(respStr)

	if err != nil {
		global.G_LOG.Errorf("PGDZ-login error :%v", apiLog)
	}
	var respData venuevo.PGDZLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code != 0 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}
	//global.G_LOG.Infof("PGDZ-Logingame-----:%v, url:%v", apiLog, respData.Data.Url)
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = respData.Data.Url
	return &venueResp
}
func (v VenuePGDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenuePGDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePGDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePGDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
