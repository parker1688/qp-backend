package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	url2 "net/url"
	"time"
)

type VenueVRCP struct {
	VenueConfig conf.VRCP
}

func NewVRCP(venueConfig *conf.VRCP) IVenues {
	return &VenueVRCP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueVRCP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	param := `{"playerName":"` + request.UserName + `"}`
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(param, v.VenueConfig.AesKey)
	//data := url2.QueryEscape(paramEncryptStr)
	data := paramEncryptStr
	apiUrl := url + "/Account/CreateUser"

	//global.G_LOG.Infof("VRCP CreateUser ----params:%v,apiurl:%v", param, apiUrl)
	reqMap := map[string]string{
		"version": "1.0",
		"id":      appId,
		"data":    data,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("VRCP-CreateUser error :%v", apiUrl)
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("VRCP-CreateUser-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	respStr, err = tool.AesEcbPk7DecryptBase64(respStr, v.VenueConfig.AesKey)
	if err != nil {
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	//global.G_LOG.Infof("VRCP CreateUser info:%v", respStr)
	respBytes := []byte(respStr)

	var respData venuevo.VRCPCreateResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	//global.G_LOG.Infof("VRCP CreateUser info2:%v, %v", respData, respData.Error)
	if respData.Error == 0 || respData.Error == 18 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVRCP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	param := `{"playerName":"` + request.UserName + `"}`
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(param, v.VenueConfig.AesKey)
	//data := url2.QueryEscape(paramEncryptStr)
	data := paramEncryptStr
	apiUrl := url + "/UserWallet/Balance"

	//global.G_LOG.Infof("VRCP Balance ----params:%v,apiurl:%v", param, apiUrl)
	reqMap := map[string]string{
		"version": "1.0",
		"id":      appId,
		"data":    data,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("VRPC-Balance error :%v", apiUrl)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("VRCP-Balance-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respStr, err = tool.AesEcbPk7DecryptBase64(respStr, v.VenueConfig.AesKey)
	if err != nil {
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)
	//global.G_LOG.Infof("VRCP-GetBalance info:%v", respStr)
	var respData venuevo.VRCPBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Balance != -1 {
		amount := respData.Balance
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueVRCP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueVRCP) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueVRCP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	sn := tool.SnowflakeIdByKey("vrcp-sn")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	createTime := time.Now().Format("2006-01-02T15:04:05") + "Z"
	param := `{"serialNumber":"` + sn + `","playerName":"` + request.UserName + `","type":0, "amount":` + amount + `,"createTime":"` + createTime + `"}`
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(param, v.VenueConfig.AesKey)
	//data := url2.QueryEscape(paramEncryptStr)
	data := paramEncryptStr
	apiUrl := url + "/UserWallet/Transaction"

	//global.G_LOG.Infof("VRCP Deposit ----params:%v,apiurl:%v", param, apiUrl)
	reqMap := map[string]string{
		"version": "1.0",
		"id":      appId,
		"data":    data,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("VRCP-Deposit error :%v", apiUrl)
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("VRCP-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	respStr, err = tool.AesEcbPk7DecryptBase64(respStr, v.VenueConfig.AesKey)
	if err != nil {
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)
	//global.G_LOG.Infof("VRCP-Deposit info:%v", respStr)
	var respData venuevo.VRCPDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Amount
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueVRCP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	sn := tool.SnowflakeIdByKey("vrcp-sn")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	createTime := time.Now().Format("2006-01-02T15:04:05") + "Z"
	param := `{"serialNumber":"` + sn + `","playerName":"` + request.UserName + `","type":1, "amount":` + amount + `,"createTime":"` + createTime + `"}`
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(param, v.VenueConfig.AesKey)
	//data := url2.QueryEscape(paramEncryptStr)
	data := paramEncryptStr
	apiUrl := url + "/UserWallet/Transaction"

	//global.G_LOG.Infof("VRCP Withdraw ----params:%v,apiurl:%v", param, apiUrl)
	reqMap := map[string]string{
		"version": "1.0",
		"id":      appId,
		"data":    data,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("VRCP-Withdraw error :%v", apiUrl)
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("VRCP-Withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	respStr, err = tool.AesEcbPk7DecryptBase64(respStr, v.VenueConfig.AesKey)
	if err != nil {
		global.G_LOG.Errorf(apiLog)
		global.G_LOG.Errorf("VRCP-Withdraw error :%v", err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)
	//global.G_LOG.Infof("VRCP-Withdraw info:%v", respStr)
	var respData venuevo.VRCPDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.State == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Amount
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenueVRCP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	appId := v.VenueConfig.AppID
	channelId := request.GameCode
	loginTime := time.Now().Format("2006-01-02T15:04:05")
	param := "playerName=" + request.UserName + "&loginTime=" + loginTime + "&channelId=" + channelId
	paramEncryptStr, err := tool.AesEcbPk7EncryptBase64(param, v.VenueConfig.AesKey)
	if err != nil {
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	data := url2.QueryEscape(paramEncryptStr)
	apiUrl := url + "/Account/LoginValidate?version=1.0&id=" + appId + "&data=" + data

	//global.G_LOG.Infof("VRCP Login ----params:%v,apiurl:%v", param, apiUrl)

	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = apiUrl
	return &venueResp
}
func (v VenueVRCP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenueVRCP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueVRCP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueVRCP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
