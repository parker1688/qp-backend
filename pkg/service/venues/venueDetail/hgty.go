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
)

const (
	CreateMemberMethod = "CreateMember"
	ChkMemberBalance   = "chkMemberBalance"
	Deposit            = "Deposit"
	Withdraw           = "Withdraw"
	AGLogin            = "AGLogin"
	LaunchGame         = "LaunchGame"
	ChkTransInfo       = "ChkTransInfo"

	HGTY_API = "/app/control_API/agents/api_doaction.php"
)

var (
	HGTYTransferProccesCode map[string]string = map[string]string{
		"0016": "系统维护中",
		"0017": "系统流量较高，请重新再试",
		"0019": "系统优化中",
		"9998": "网站维护中",
		"9999": "系统异常",
	}
)

type VenueHGTY struct {
	VenueConfig conf.HGTY
}

func NewHGTY(venueConfig *conf.HGTY) IVenues {
	return &VenueHGTY{
		VenueConfig: *venueConfig,
	}
}

func (v VenueHGTY) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueResponse{}
	if request.Token == "" {
		tmpStr := fmt.Sprintf("HGTY-CreateUser-param username: %s token is empty", request.UserName)
		global.G_LOG.Errorf(tmpStr)
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = tmpStr
		return &venueResp
	}

	paramEncryptOrg := venuevo.HGTYCreatUserEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Currency = request.Currency
	paramEncryptOrg.Password = request.Password
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Method = CreateMemberMethod
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-CreateUser-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptOrgStr := string(paramEncryptOrgBytes)
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-CreateUser-param %s AesEcbPk7EncryptBase64 err: %v", paramEncryptOrgStr, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYCreatUserReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = CreateMemberMethod
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-CreateUser-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-CreateUser-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-CreateUser-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiLog := fmt.Sprintf("HGTY-CreateUser-param: %s paramOrg: %s apiUrl: %s response: %s", paramStr, paramEncryptOrgStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYCreateUserResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Respcode == "0000" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.ThirdCode = respData.Respcode
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueHGTY) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueGetUserBalanceResponse{}

	paramEncryptOrg := venuevo.HGTYBalanceEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Method = ChkMemberBalance
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYBalanceReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = ChkMemberBalance
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	apiLog := fmt.Sprintf("HGTY-GetUserBalance-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYBalanceResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Respcode == "0000" {
		amount, err := strconv.ParseFloat(respData.Balance, 64)
		if err != nil {
			global.G_LOG.Errorf("%s balanceTransport err: %v", apiLog, err.Error())
			venueResp.Code = GetUserBalance_FAIL_CODE
			venueResp.Msg = err.Error()
			return &venueResp
		}

		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.ThirdCode = respData.Respcode
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueHGTY) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueHGTY) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueHGTY) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueDepositResponse{}

	paramEncryptOrg := venuevo.HGTYDepositEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Method = Deposit
	paramEncryptOrg.Amount = fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	paramEncryptOrg.Payno = request.OrderSn
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Deposit-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Deposit-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYBalanceReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = Deposit
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Deposit-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-GetUserBalance-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}
	apiLog := fmt.Sprintf("HGTY-Withdraw-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYDepositResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}

	if respData.Respcode == "0000" {
		venueResp := VenueDepositResponse{}
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.ThirdCode = respData.Respcode
	venueResp.Msg = respData.Status

	global.G_LOG.Errorf("%s StatusCode faild ", apiLog)
	return &venueResp
}

// 转出
func (v VenueHGTY) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueWithdrawResponse{}

	paramEncryptOrg := venuevo.HGTYWithdrawEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Amount = fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	paramEncryptOrg.Payno = request.OrderSn
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Method = Withdraw
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYBalanceReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = Withdraw
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Withdraw-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiLog := fmt.Sprintf("HGTY-Withdraw-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYWithdrawResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Respcode == "0000" {
		return &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.ThirdCode = respData.Respcode
	venueResp.Msg = respData.Status

	global.G_LOG.Errorf("%s StatusCode faild ", apiLog)
	return &venueResp
}

func (v VenueHGTY) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueLoginGameResponse{}

	machine := "MOBILE"
	if request.ClientType == "PC" || request.ClientType == "web" {
		machine = "PC"
	}
	language := strings.ToLower(request.Language)

	paramEncryptOrg := venuevo.HGTYLaunchGameEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Password = request.Password
	paramEncryptOrg.Remoteip = "127.0.0.1"
	paramEncryptOrg.Currency = request.Currency
	paramEncryptOrg.Langx = language
	paramEncryptOrg.Machine = machine
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-LoginGame-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-LoginGame-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}
	paramReqData := venuevo.HGTYBalanceReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = LaunchGame
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-LoginGame-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-LoginGame-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-LoginGame-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}
	apiLog := fmt.Sprintf("HGTY-Login-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYLaunchGameResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Respcode == "0000" {
		venueLoginGameResponse := &VenueLoginGameResponse{
			Code: LoginGame_SUCCESS_CODE,
		}
		oldUrl := respData.Launchgameurl
		newUrl := strings.ReplaceAll(oldUrl, "http", "https")
		venueLoginGameResponse.Data.GameUrl = newUrl
		return venueLoginGameResponse
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.ThirdCode = respData.Respcode
	venueResp.Msg = respData.Status
	return &venueResp
}

func (v VenueHGTY) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	url := v.VenueConfig.Url
	apiUrl := url + HGTY_API
	venueResp := VenueResponse{}

	paramEncryptOrg := venuevo.HGTYChkTransInfoEncryptReq{}
	paramEncryptOrg.Memname = request.UserName
	paramEncryptOrg.Transid = request.OrderSn
	paramEncryptOrg.Transidtype = "0" // 0: 三方订单号，1: 平台订单号
	paramEncryptOrg.Token = request.Token
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-TransferConfirm-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-TransferConfirm-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYChkTransInfoReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = ChkTransInfo
	paramReqData.AGID = v.VenueConfig.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-TransferConfirm-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-TransferConfirm-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.VenueConfig.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-TransferConfirm-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())

		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiLog := fmt.Sprintf("HGTY-TransferConfirm-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYChkTransInfoResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if len(respData.Transdata) < 1 {
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = respData.Status
		return &venueResp
	}

	if respData.Status == "success" {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respData.Status
	return &venueResp
}

func (v VenueHGTY) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueHGTY) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueHGTY) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func HGTYLogin(v conf.HGTY) *vo.VenueLoginResponse {
	url := v.Url
	apiUrl := url + HGTY_API
	venueResp := vo.VenueLoginResponse{}

	paramEncryptOrg := venuevo.HGTYLoginEncryptReq{}
	paramEncryptOrg.UserName = v.Account
	paramEncryptOrg.Password = v.Password
	paramEncryptOrg.Timestamp = tool.TimeNowTimestampString()

	paramEncryptOrgBytes, err := tool.JsonMarshal(&paramEncryptOrg)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Login-param JsonMarshalToString1 err: %v", err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramEncryptData, err := tool.AesEcbPk7EncryptBase64(string(paramEncryptOrgBytes), v.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Login-param %s AesEcbPk7EncryptBase64 err: %v", string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramReqData := venuevo.HGTYLoginReq{}
	paramReqData.Request = paramEncryptData
	paramReqData.Method = AGLogin
	paramReqData.AGID = v.AgId
	reqDataBytes, err := tool.JsonMarshal(&paramReqData)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Login-param JsonMarshalToString2 err: %v", err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	paramStr := string(reqDataBytes)

	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(paramStr).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Login-param:%s apiUrl:%s httpReq err: %v", apiUrl, string(paramEncryptOrgBytes), err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respEncryptStr := resp.String()

	// 对数据进行解密
	respStr, err := tool.AesEcbPk7DecryptBase64(respEncryptStr, v.Aeskey)
	if err != nil {
		global.G_LOG.Errorf("HGTY-Login-param:%s apiUrl:%s AesEcbPk7DecryptBase64 body:%s err: %v",
			apiUrl, string(paramEncryptOrgBytes), respEncryptStr, err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiLog := fmt.Sprintf("HGTY-Token-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.HGTYLoginResp
	err = tool.JsonUnmarshal([]byte(respStr), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = vo.Login_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Respcode == "0000" {
		venueResp.Code = vo.Login_SUCCESS_CODE
		venueResp.Data.Token = respData.Token
		return &venueResp
	}

	venueResp.Code = vo.Login_FAIL_CODE
	venueResp.Msg = respData.Status
	return &venueResp
}
