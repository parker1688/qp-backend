package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	url5 "net/url"
	"strconv"
	"strings"
	"time"
)

type VenuePTDZ struct {
	VenueConfig conf.PTDZ
}

func NewPTDZ(venueConfig *conf.PTDZ) IVenues {
	return &VenuePTDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenuePTDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	//global.G_LOG.Infof("pt create user --------:%v", request.UserName)
	url := v.VenueConfig.Url
	entityKey := v.VenueConfig.AppSecret
	apiUrl := url + "/player/create"
	formData := url5.Values{
		"playername": {request.UserName},
		"password":   {"123456"},
		"currency":   {"CNY"},
	}
	req, err := http.NewRequest("POST", apiUrl,
		strings.NewReader(formData.Encode()))
	if err != nil {
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X_ENTITY_KEY", entityKey)
	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	defer resp.Body.Close()
	respStr, _ := io.ReadAll(resp.Body)

	apiLog := fmt.Sprintf("PTDZ-Active-apiUrl: %s, formData:%v response: %s", apiUrl, formData, string(respStr))
	global.G_LOG.Infof(apiLog)

	//respBytes := []byte(respStr)
	respBytes := respStr

	var respData venuevo.PTDZCreateResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	venueResp.Code = CreateUser_SUCCESS_CODE
	venueResp.Msg = respData.Result.Password
	return &venueResp
}

func (v VenuePTDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	entityKey := v.VenueConfig.AppSecret
	apiUrl := url + "/player/balance/"
	formData := url5.Values{
		"playername": {request.UserName},
	}
	req, err := http.NewRequest("POST", apiUrl,
		strings.NewReader(formData.Encode()))
	if err != nil {
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X_ENTITY_KEY", entityKey)
	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	defer resp.Body.Close()
	respStr, _ := io.ReadAll(resp.Body)

	apiLog := fmt.Sprintf("PTDZ-GetUserBalance-apiUrl: %s response: %s", apiUrl, string(respStr))
	global.G_LOG.Infof(apiLog)

	respBytes := respStr
	var respData venuevo.PTDZBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Error == "" {
		amount, err2 := strconv.ParseFloat(respData.Result.Balance, 64)
		if err2 != nil {
			venueResp.Code = GetUserBalance_FAIL_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = string(respStr)
	return &venueResp
}

func (v VenuePTDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenuePTDZ) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenuePTDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	entityKey := v.VenueConfig.AppSecret
	apiUrl := url + "/player/deposit/"
	amount := tool.String(request.Amount)
	formData := url5.Values{
		"playername": {request.UserName},
		"amount":     {amount},
	}
	global.G_LOG.Infof("PT-deposit :%v, %v", apiUrl, formData.Encode())
	req, err := http.NewRequest("POST", apiUrl,
		strings.NewReader(formData.Encode()))
	if err != nil {
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X_ENTITY_KEY", entityKey)

	client := NewClient()
	resp, err := client.Do(req)

	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	defer resp.Body.Close()
	respStr, _ := io.ReadAll(resp.Body)

	apiLog := fmt.Sprintf("PTDZ-Deposit-apiUrl:%s response:%s", apiUrl, string(respStr))
	global.G_LOG.Infof(apiLog)

	respBytes := respStr

	var respData venuevo.PTDZDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "" {
		amount2, err2 := strconv.ParseFloat(respData.Result.Amount, 64)
		if err2 != nil {
			venueResp.Code = Deposit_FAIL_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = amount2
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = string(respStr)

	return &venueResp
}

// 转出
func (v VenuePTDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	entityKey := v.VenueConfig.AppSecret
	apiUrl := url + "/player/withdraw/"
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))

	formData := url5.Values{
		"playername": {request.UserName},
		"amount":     {amount},
	}
	req, err := http.NewRequest("POST", apiUrl,
		strings.NewReader(formData.Encode()))
	if err != nil {
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X_ENTITY_KEY", entityKey)
	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		global.G_LOG.Errorf("PTDZ-pullBetRecord error :%v", apiUrl)
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	defer resp.Body.Close()
	respStr, _ := io.ReadAll(resp.Body)
	apiLog := fmt.Sprintf("PTDZ-Withdraw-apiUrl:%s response:%s", apiUrl, string(respStr))
	global.G_LOG.Infof(apiLog)
	respBytes := respStr

	var respData venuevo.PTDZWithdrawResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Error == "" {
		amount2, err2 := strconv.ParseFloat(respData.Result.Amount, 64)
		if err2 != nil {
			venueResp.Code = Withdraw_FAIL_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = amount2
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = string(respStr)

	return &venueResp
}
func (v VenuePTDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{} //这个PTDZ需要客户端独立接入游戏大厅
	entityKey := v.VenueConfig.AppSecret
	apiUrl := "https://login-ag.agdragonbc.com/LoginAndGetTempToken.php?casinoname=agdragon&realMode=1&serviceType=GamePlay&systemId=77&clientType=casino&clientPlatform=mobile&clientSkin=agdragon&languageCode=ZH-CN&messagesSupported=1"
	formData := url5.Values{
		"username": {strings.ToUpper(request.UserName)},
		"password": {"123456"},
		"realMode": {"1"},
		"language": {"ZH-CN"},
		"systemid": {"77"},
		"currency": {"CNYC"},
	}
	req, err := http.NewRequest("POST", apiUrl,
		strings.NewReader(formData.Encode()))
	if err != nil {
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X_ENTITY_KEY", entityKey)
	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	defer resp.Body.Close()
	respStr, _ := io.ReadAll(resp.Body)
	apiLog := fmt.Sprintf("PTDZ-Login-apiUrl: %s formdata:%v, response: %s", apiUrl, formData, string(respStr))
	global.G_LOG.Infof(apiLog)

	var respData venuevo.PTDZLoginResp
	err = tool.JsonUnmarshal(respStr, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	global.G_LOG.Infof("token:%v", respData.SessionToken.SessionToken)
	//https://login-ag.agdragonbc.com/GameLauncher?gameCodeName=gpas_ldblava_pop&username=sss0011&casino=agdragon&clientPlatform=nptgp&language=ZH-CN&playMode=1&deposit=&lobby=&swipeUpOff=true&tempToken=s-Kr-T9gFK8H4wkGfn9ZIHBhbwhRAEhW8EmH5hJnRDgiRh1YyXEryRWZuu9T0R0XrbYRw28TvCdVgdL12Bb79tVQ
	//https://login-ag.agdragonbc.com/GameLauncher?gameCodeName=gpas_ldblava_pop&username=sss0011&casino=agdragon&clientPlatform=nptgp&language=ZH-CN&playMode=1&deposit=&lobby=&swipeUpOff=true&tempToken=s-w4NN_S0XGCP5FEf_yOVeJsoURY-qMoK_ploZGyFd4FWY4hkvD-r77gtoF_EDuRS7DtqSyAGipOtLU17KL_Bp0A
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = string(respStr)
	return &venueResp
	//"PTDZ-Login-apiUrl: https://login-ag.agdragonbc.com/LoginAndGetTempToken.php?casinoname=agdragon&realMode=1&serviceType=GamePlay&systemId=77&clientType=casino&clientPlatform=mobile&clientSkin=agdragon&languageCode=ZH-CN&messagesSupported=1 formdata:map[currency:[CNYC] language:[ZH-CN] password:[123456] realMode:[1] systemid:[77] username:[DVTONYTEST2]], response: {\"flow\":\"258030326194116570:4\",\"errorCode\":0,\"messageId\":\"24117927\",\"actions\":{\"PlayerActionShowMessage\":[{\"message\":\"Access denied\\nUnfortunately we cannot service customers from your country of residence.\",\"displayType\":null,\"templateId\":\"restricted_countries\",\"delegateDisplay\":\"false\"}]},\"playerMessage\":\"Access denied\\nUnfortunately we cannot service customers from your country of residence.\",\"username\":\"DVTONYTEST2\",\"ssoStartTime\":{\"timestamp\":\"2025-07-29T14:42:58.656+00:00\"},\"previousLoginTime\":{\"timestamp\":\"2025-07-29T14:42:58.000+00:00\"},\"realMode\":1,\"perm\":\"AOE6w6boRLX4KIhLzsWKkEBwsBBAgOCI\",\"playerCode\":\"173979784\",\"currencyCode\":\"CNY\",\"playerSessionId\":\"777094234156\",\"ssoLoginCount\":\"1\",\"inactivityTimeout\":\"30\",\"ipCountryCode\":\"HK\",\"loginName\":\"DVTONYTEST2\",\"userId\":\"X0pw/+1FfZAtjGB0B//UvFmwEhsA==\",\"rememberMeToken\":null,\"sessionToken\":{\"sessionToken\":\"s-w4NN_S0XGCP5FEf_yOVeJsoURY-qMoK_ploZGyFd4FWY4hkvD-r77gtoF_EDuRS7DtqSyAGipOtLU17KL_Bp0A\",\"issuerSystemId\":\"77\",\"creationTime\":{\"timestamp\":\"2025-07-29T14:47:31.559+00:00\"},\"expirationTime\":{\"timestamp\":\"2025-07-29T14:52:31.559+00:00\"}}}"}
}
func (v VenuePTDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenuePTDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePTDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenuePTDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func NewClient() *http.Client {
	// 1. 加载客户端证书和私钥
	cert, err := tls.LoadX509KeyPair("./ptzs/pt.pem", "./ptzs/pt.key")
	if err != nil {
		global.G_LOG.Errorf("PTDZ 加载证书失败: %v", err)
		return &http.Client{Timeout: 10 * time.Second}
	}
	// 3. 配置TLS客户端
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	// 4. 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 20,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	return client

}
