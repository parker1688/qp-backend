package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
	url2 "net/url"
)

type VenueBGZR struct {
	VenueConfig conf.BGZR
}

func NewBGZR(venueConfig *conf.BGZR) IVenues {
	return &VenueBGZR{
		VenueConfig: *venueConfig,
	}
}

func (v VenueBGZR) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	nickName := request.UserName
	sn := v.VenueConfig.Sn
	password := v.VenueConfig.Password
	loginId := v.VenueConfig.LoginId
	url := v.VenueConfig.Url

	secretCode, err := sha1AndBase64Encode(password)
	if err != nil {
		global.G_LOG.Errorf("BGZR-create-user, err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	ip := request.Ip
	if ip == "" {
		ip = "127.0.0.1"
	}
	id := tool.SnowflakeIdByKey("bgzr")
	rand := tool.SnowflakeIdByKey("bgzr-rand")
	method := "open.user.create"
	apiUrl := url + method
	param := rand + sn + secretCode
	digest := tool.MD5([]byte(param))
	paramData := map[string]string{}
	paramData["random"] = rand
	paramData["digest"] = digest
	paramData["sn"] = sn
	paramData["loginId"] = nickName
	paramData["nickname"] = nickName
	paramData["agentLoginId"] = loginId
	paramData["fromIp"] = ip

	reqMap := map[string]interface{}{}
	reqMap["id"] = id
	reqMap["method"] = method
	reqMap["params"] = paramData
	reqMap["jsonrpc"] = "2.0"
	reqParams, _ := json.Marshal(reqMap)
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	//global.G_LOG.Infof("BGZR-CreateUser info :%v,  %v", apiUrl, reqMap)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetBody(reqParams).
		Post(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BGZR-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("BGZR-CreateUser info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BGZRCreateUserResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result.Success {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof("BGZR-CreateUser info :%v", apiLog)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueBGZR) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	sn := v.VenueConfig.Sn
	password := v.VenueConfig.Password
	loginId := request.UserName
	url := v.VenueConfig.Url

	secretCode, err := sha1AndBase64Encode(password)
	if err != nil {
		global.G_LOG.Errorf("BGZR-create-user, err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	id := tool.SnowflakeIdByKey("bgzr")
	rand := tool.SnowflakeIdByKey("bgzr-rand")
	method := "open.balance.get"
	apiUrl := url + method
	param := rand + sn + loginId + secretCode
	digest := tool.MD5([]byte(param))
	paramData := map[string]interface{}{}
	paramData["random"] = id
	paramData["digest"] = digest
	paramData["sn"] = sn
	paramData["loginId"] = loginId

	reqMap := map[string]interface{}{}
	reqMap["id"] = id
	reqMap["method"] = method
	reqMap["params"] = paramData
	reqMap["jsonrpc"] = "2.0"
	reqParams, _ := json.Marshal(reqMap)
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetBody(reqParams).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("BGZR httpReq err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BGZR-Balance-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("BGZR-Balance info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BGZRBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("BGZR-Get balance: %s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Result)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	amount := float64(respData.Result)
	venueResp.Code = GetUserBalance_SUCCESS_CODE
	venueResp.Data.Amount = amount
	return &venueResp
}

func (v VenueBGZR) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueBGZR) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueBGZR) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	amount := fmt.Sprintf("%0.2f", math.Abs(request.Amount))
	sn := v.VenueConfig.Sn
	password := v.VenueConfig.Password
	loginId := request.UserName
	url := v.VenueConfig.Url

	secretCode, err := sha1AndBase64Encode(password)
	if err != nil {
		global.G_LOG.Errorf("BGZR-create-user, err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	id := tool.SnowflakeIdByKey("bgzr")
	rand := tool.SnowflakeIdByKey("bgzr-rand")
	method := "open.balance.transfer"
	apiUrl := url + method
	param := rand + sn + loginId + amount + secretCode
	digest := tool.MD5([]byte(param))
	paramData := map[string]interface{}{}
	paramData["random"] = id
	paramData["digest"] = digest
	paramData["sn"] = sn
	paramData["loginId"] = loginId
	paramData["amount"] = amount

	reqMap := map[string]interface{}{}
	reqMap["id"] = id
	reqMap["method"] = method
	reqMap["params"] = paramData
	reqMap["jsonrpc"] = "2.0"
	reqParams, _ := json.Marshal(reqMap)
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetBody(reqParams).
		Post(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("BGZR httpReq err: %v", err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BGZR-Deposit-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("BGZR-deposit info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BGZRTransferResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("BGZR-deposit: %s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Result)
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	venueResp.Code = Deposit_SUCCESS_CODE
	venueResp.Data.Amount = request.Amount
	return &venueResp
}

// 转出
func (v VenueBGZR) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	amount := fmt.Sprintf("%0.2f", -math.Abs(request.Amount))
	sn := v.VenueConfig.Sn
	password := v.VenueConfig.Password
	loginId := request.UserName
	url := v.VenueConfig.Url

	secretCode, err := sha1AndBase64Encode(password)
	if err != nil {
		global.G_LOG.Errorf("BGZR-withdraw, err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	id := tool.SnowflakeIdByKey("bgzr")
	rand := tool.SnowflakeIdByKey("bgzr-rand")
	method := "open.balance.transfer"
	apiUrl := url + method
	param := rand + sn + loginId + amount + secretCode
	digest := tool.MD5([]byte(param))
	paramData := map[string]interface{}{}
	paramData["random"] = id
	paramData["digest"] = digest
	paramData["sn"] = sn
	paramData["loginId"] = loginId
	paramData["amount"] = amount

	reqMap := map[string]interface{}{}
	reqMap["id"] = id
	reqMap["method"] = method
	reqMap["params"] = paramData
	reqMap["jsonrpc"] = "2.0"
	reqParams, _ := json.Marshal(reqMap)
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetBody(reqParams).
		Post(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("BGZR httpReq err: %v", err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()

	apiLog := fmt.Sprintf("BGZR-Withdraw-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("BGZR-Withdraw info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BGZRTransferResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("BGZR-Get balance: %s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Result)
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	venueResp.Code = Withdraw_SUCCESS_CODE
	venueResp.Data.Amount = request.Amount
	return &venueResp
}

func (v VenueBGZR) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	sn := v.VenueConfig.Sn
	password := v.VenueConfig.Password
	loginId := request.UserName
	url := v.VenueConfig.Url
	returnUrl := request.ReturnUrl

	secretCode, err := sha1AndBase64Encode(password)
	if err != nil {
		global.G_LOG.Errorf("BGZR-create-user, err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	id := tool.SnowflakeIdByKey("bgzr")
	rand := tool.SnowflakeIdByKey("bgzr-rand")
	method := "open.video.game.url"
	apiUrl := url + method
	param := rand + sn + loginId + secretCode
	digest := tool.MD5([]byte(param))
	paramData := map[string]interface{}{}
	paramData["random"] = rand
	paramData["digest"] = digest
	paramData["sn"] = sn
	paramData["loginId"] = loginId
	paramData["local"] = "zh_CN"
	paramData["isMobileUrl"] = 1
	paramData["isHttpsUrl"] = 1
	paramData["returnUrl"] = url2.QueryEscape(returnUrl)
	paramData["fromIp"] = request.IP

	reqMap := map[string]interface{}{}
	reqMap["id"] = id
	reqMap["method"] = method
	reqMap["params"] = paramData
	reqMap["jsonrpc"] = "2.0"

	reqParams, _ := json.Marshal(reqMap)
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetBody(reqParams).
		Post(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("BGZR-Login-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof("BGZR-Login info :%v", apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	var respData venuevo.BGZRLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result == "" {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}
	venueResp.Code = LoginGame_SUCCESS_CODE
	newUrl := respData.Result + "&cui=8208" //8192+16,8192是上面广告栏去掉，16是下面推广去掉
	venueResp.Data.GameUrl = newUrl
	return &venueResp
}

func (v VenueBGZR) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}

func (v VenueBGZR) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueBGZR) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueBGZR) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func sha1AndBase64Encode(input string) (string, error) {
	hasher := sha1.New()
	if _, err := hasher.Write([]byte(input)); err != nil {
		return "", err
	}
	sha1Bytes := hasher.Sum(nil)                               // 获取哈希值
	encodedStr := base64.StdEncoding.EncodeToString(sha1Bytes) // Base64编码
	return encodedStr, nil
}
