package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type VenuePPDZ struct {
	VenueConfig conf.PPDZ
}

func NewPPDZ(venueConfig *conf.PPDZ) IVenues {
	return &VenuePPDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenuePPDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + "/player/account/create/"
	secureLogin := v.VenueConfig.SecureLogin
	secretKey := v.VenueConfig.SecretKey
	param := "currency=CNY&externalPlayerId=" + request.UserName + "&secureLogin=" + secureLogin + secretKey
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"secureLogin":      secureLogin,
		"externalPlayerId": request.UserName,
		"currency":         "CNY",
		"hash":             sign,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}
	global.G_LOG.Infof("PPDZ-createUser info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
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
	apiLog := fmt.Sprintf("PPDZ-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PPDZUserResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "0" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenuePPDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + "/balance/current/"
	secureLogin := v.VenueConfig.SecureLogin
	secretKey := v.VenueConfig.SecretKey
	param := "externalPlayerId=" + request.UserName + "&secureLogin=" + secureLogin + secretKey
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"secureLogin":      secureLogin,
		"externalPlayerId": request.UserName,
		"hash":             sign,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}
	global.G_LOG.Infof("PPDZ-balance info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
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
	apiLog := fmt.Sprintf("PPDZ-GetBalance-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.PPDZBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Error)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// PGDZ 金额单位分
	if respData.Error == "0" {
		amount := float64(respData.Balance)
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenuePPDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenuePPDZ) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenuePPDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + "/balance/transfer/"
	secureLogin := v.VenueConfig.SecureLogin
	secretKey := v.VenueConfig.SecretKey
	transactionId := tool.SnowflakeIdByKey("ppdz-transactionId")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	param := "amount=" + amount + "&externalPlayerId=" + request.UserName + "&externalTransactionId=" + transactionId + "&secureLogin=" + secureLogin + secretKey
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"amount":                amount,
		"secureLogin":           secureLogin,
		"externalPlayerId":      request.UserName,
		"externalTransactionId": transactionId,
		"hash":                  sign,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}
	global.G_LOG.Infof("PPDZ-deposit info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("PPDZ-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PPDZTransactionResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "0" {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Balance
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenuePPDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + "/balance/transfer/"
	secureLogin := v.VenueConfig.SecureLogin
	secretKey := v.VenueConfig.SecretKey
	transactionId := tool.SnowflakeIdByKey("ppdz-transactionId")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(-request.Amount, 2))
	param := "amount=" + amount + "&externalPlayerId=" + request.UserName + "&externalTransactionId=" + transactionId + "&secureLogin=" + secureLogin + secretKey
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"amount":                amount,
		"secureLogin":           secureLogin,
		"externalPlayerId":      request.UserName,
		"externalTransactionId": transactionId,
		"hash":                  sign,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}
	global.G_LOG.Infof("PPDZ-deposit info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("PPDZ-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PPDZTransactionResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "0" {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Balance
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenuePPDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	apiUrl := url + "/game/start/"
	secureLogin := v.VenueConfig.SecureLogin
	secretKey := v.VenueConfig.SecretKey
	param := "externalPlayerId=" + request.UserName + "&gameId=" + request.GameCode + "&language=cn&secureLogin=" + secureLogin + secretKey
	sign := tool.MD5([]byte(param))
	reqMap := map[string]string{
		"secureLogin":      secureLogin,
		"externalPlayerId": request.UserName,
		"gameId":           request.GameCode,
		"language":         "cn",
		"hash":             sign,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}
	global.G_LOG.Infof("PPDZ-balance info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
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
	apiLog := fmt.Sprintf("PPDZ-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.PPDZLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "0" {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.GameUrl
		return &venueResp
	}
	global.G_LOG.Errorf(apiLog)
	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}
func (v VenuePPDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenuePPDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePPDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePPDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
