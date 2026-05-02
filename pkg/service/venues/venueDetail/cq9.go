package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type VenueCQ9 struct {
	VenueConfig conf.CQ9
}

func NewCQ9(venueConfig *conf.CQ9) IVenues {
	return &VenueCQ9{
		VenueConfig: *venueConfig,
	}
}

func (v VenueCQ9) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	Token := v.VenueConfig.TOKEN
	url := v.VenueConfig.Url
	apiUrl := url + "/gameboy/player"
	reqMap := map[string]string{
		"account":  request.UserName,
		"password": request.Password,
		"nickname": request.NickName,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}

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
	apiLog := fmt.Sprintf("CQ9-CreateUser-apiUrl: %s reqmap:%v, header:%v, response: %s", apiUrl, reqMap, headerParams, respStr)
	global.G_LOG.Infof("CQ9-Login info :%v", apiLog)

	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respData venuevo.CQ9CreateUserResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status.Code == "0" || respData.Status.Code == "6" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueCQ9) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	account := request.UserName
	Token := v.VenueConfig.TOKEN

	apiUrl := url + "/gameboy/player/balance/" + account

	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headerParams).Get(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("CQ9-GetUserBalance-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	var respData venuevo.CQ9BalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Status.Code == "0" {
		amount := respData.Data.Balance
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueCQ9) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueCQ9) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueCQ9) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	Token := v.VenueConfig.TOKEN
	mtcode := tool.SnowflakeIdByKey("cq9-mtcode")
	apiUrl := url + "/gameboy/player/deposit"
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	reqMap := map[string]string{
		"account": request.UserName,
		"mtcode":  mtcode,
		"amount":  amount,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("Cq9-deposit error :%v, repMap:%v, error:%v", apiUrl, reqMap, err)
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()

	apiLog := fmt.Sprintf("CQ9-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.CQ9DepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status.Code == "0" {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.Balance
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueCQ9) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	Token := v.VenueConfig.TOKEN
	mtcode := tool.SnowflakeIdByKey("cq9-mtcode")
	apiUrl := url + "/gameboy/player/withdraw"
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	reqMap := map[string]string{
		"account": request.UserName,
		"mtcode":  mtcode,
		"amount":  amount,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("CQ9-Withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("withdraw error :%v", apiLog)
	}
	respBytes := []byte(respStr)

	var respData venuevo.CQ9DepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status.Code == "0" {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.Balance
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueCQ9) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	Token := v.VenueConfig.TOKEN
	apiUrl := url + "/gameboy/player/login"
	reqMap := map[string]string{
		"account":  request.UserName,
		"password": request.Password,
	}
	headerParams := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		global.G_LOG.Errorf("CQ9-login request error apiUrl:%s err:%v", apiUrl, err)
		return &venueResp
	}

	respStr := resp.String()
	apiLog := fmt.Sprintf("CQ9-LoginGame-apiUrl:%s , params:%v, response:%s", apiUrl, reqMap, respStr)
	//global.G_LOG.Infof(apiLog)
	respBytes := []byte(respStr)

	var respData venuevo.CQ9LoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	UserToken := respData.Data.UserToken
	apiUrl2 := url + "/gameboy/player/gamelink"

	reqMap2 := map[string]string{
		"usertoken": UserToken,
		"gamehall":  "cq9",
		"gamecode":  request.GameCode,
		"gameplat":  "web",
		"lang":      "zh-cn",
	}
	headerParams2 := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": Token,
	}

	client2 := resty.New()
	resp2, err2 := client2.R().
		SetHeaders(headerParams2).
		SetFormData(reqMap2).
		Post(apiUrl2)

	respStr2 := resp2.String()
	apiLog2 := fmt.Sprintf("CQ9-LoginGame-apiUrl:%s response:%s", apiUrl2, respStr2)
	//global.G_LOG.Infof(apiLog2)

	if err2 != nil {
		global.G_LOG.Errorf("login error :%v", apiLog2)
	}
	respBytes2 := []byte(respStr2)

	var respData2 venuevo.CQ9Login2Resp
	err = tool.JsonUnmarshal(respBytes2, &respData2)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = respData2.Data.Url
	venueResp.Data.Token = respData2.Data.Token
	return &venueResp
}

func (v VenueCQ9) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}

func (v VenueCQ9) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueCQ9) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueCQ9) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
