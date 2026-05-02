package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"runtime"
	"strconv"
)

type VenueAGZR struct {
	VenueConfig conf.AGZR
}

func NewAGZR(venueConfig *conf.AGZR) IVenues {
	return &VenueAGZR{
		VenueConfig: *venueConfig,
	}
}

func (v VenueAGZR) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey

	paramStr := "cagent=" + cagent + "/\\\\/loginname=" + loginName + "/\\\\/method=lg/\\\\/actype=1/\\\\/password=" + password + "/\\\\/oddtype=A/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/doBusiness.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-CreateUser-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respDataMap venuevo.AGZRResp
	err = xml.Unmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	//global.G_LOG.Infof("create-1:%v", respDataMap.Info)
	if respDataMap.Info == "0" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	//global.G_LOG.Infof("create-2:%v", respDataMap.Msg)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueAGZR) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	//global.G_LOG.Infof("AGZR-GetUserBalance-param: ----------------------------------1")
	defer func() {
		err := recover()
		if err != nil {
			var buf [2048]byte
			n := runtime.Stack(buf[:], false)
			global.G_LOG.Errorf(string(buf[:n]))
		}
	}()
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey

	paramStr := "cagent=" + cagent + "/\\\\/loginname=" + loginName + "/\\\\/method=gb/\\\\/actype=1/\\\\/password=" + password + "/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/doBusiness.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respDataMap venuevo.AGZRResp
	err = xml.Unmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	f, err1 := strconv.ParseFloat(respDataMap.Info, 64)
	if err1 != nil {
		venueResp.Code = GetUserBalance_FAIL_CODE
		return &venueResp
	}
	venueResp.Code = GetUserBalance_SUCCESS_CODE
	venueResp.Data.Amount = f
	return &venueResp
}

func (v VenueAGZR) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueAGZR) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueAGZR) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	//global.G_LOG.Infof("AGZR-Deposit-param: ----------------------------------1")
	defer func() {
		err := recover()
		if err != nil {
			var buf [2048]byte
			n := runtime.Stack(buf[:], false)
			global.G_LOG.Errorf(string(buf[:n]))
		}
	}()
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey
	randnum := tool.RandInt(10000000000000, 9999999999999999)
	billno := cagent + tool.String(randnum)
	credit := tool.String(request.Amount)

	paramStr := "cagent=" + cagent + "/\\\\/method=tc/\\\\/loginname=" + loginName + "/\\\\/billno=" + billno + "/\\\\/type=IN/\\\\/credit=" + credit + "/\\\\/actype=1/\\\\/password=" + password + "/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/doBusiness.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-Deposit-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respDataMap venuevo.AGZRResp
	err = xml.Unmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respDataMap.Info == "0" {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = request.Amount
		venueResp.Data.TransactionId = billno

		v.TransferConfirm(&VenueTransferConfirmRequest{
			UserName:     loginName,
			Password:     password,
			IP:           request.Ip,
			OrderSn:      billno,
			Credit:       credit,
			TransferType: "IN",
		})
		return &venueResp
	}
	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

// 转出
func (v VenueAGZR) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey
	randnum := tool.RandInt(10000000000000, 9999999999999999)
	billno := cagent + tool.String(randnum)
	credit := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))

	//paramStr := "cagent=" + cagent + "/\\\\/loginname=" + loginName + "/\\\\/method=tc/\\\\/actype=1/\\\\/password=" + password + "/\\\\/oddtype=A/\\\\/cur=CNY\n"
	paramStr := "cagent=" + cagent + "/\\\\/method=tc/\\\\/loginname=" + loginName + "/\\\\/billno=" + billno + "/\\\\/type=OUT/\\\\/credit=" + credit + "/\\\\/actype=1/\\\\/password=" + password + "/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/doBusiness.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-Withdraw-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respDataMap venuevo.AGZRResp
	err = xml.Unmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respDataMap.Info == "0" {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = request.Amount
		venueResp.Data.TransactionId = billno
		v.TransferConfirm(&VenueTransferConfirmRequest{
			UserName:     loginName,
			Password:     password,
			IP:           request.Ip,
			OrderSn:      billno,
			Credit:       credit,
			TransferType: "OUT",
		})
		return &venueResp
	}
	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueAGZR) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.GameUrl
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey
	randnum := tool.RandInt(10000000000000, 9999999999999999)
	sid := cagent + tool.String(randnum)

	paramStr := "cagent=" + cagent + "/\\\\/loginname=" + loginName + "/\\\\/actype=1/\\\\/password=" + password + "/\\\\/dm=NO_RETURN/\\\\/sid=" + sid + "/\\\\/lang=1/\\\\/gameType=1/\\\\/oddtype=A/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/forwardGame.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-Login-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	//respBytes := []byte(respStr)
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = apiUrl
	return &venueResp
}

func (v VenueAGZR) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	//global.G_LOG.Infof("AGZR-TransferConfirm-param: ----------------------------------1")
	defer func() {
		err := recover()
		if err != nil {
			var buf [2048]byte
			n := runtime.Stack(buf[:], false)
			global.G_LOG.Errorf(string(buf[:n]))
		}
	}()
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	cagent := v.VenueConfig.Agent
	loginName := request.UserName
	password := request.UserName
	md5Key := v.VenueConfig.Md5Key
	desKey := v.VenueConfig.DesKey
	billno := request.OrderSn
	trType := request.TransferType
	credit := request.Credit

	paramStr := "cagent=" + cagent + "/\\\\/loginname=" + loginName + "/\\\\/method=tcc/\\\\/billno=" + billno + "/\\\\/type=" + trType + "/\\\\/credit=" + credit + "/\\\\/actype=1/\\\\/flag=1/\\\\/password=" + password + "/\\\\/cur=CNY\n"
	params := tool.EncryptDES([]byte(desKey), []byte(paramStr))
	secretKey := tool.MD5([]byte(params + md5Key))
	apiUrl := url + "/doBusiness.do?params=" + params + "&key=" + secretKey

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("AGZR-TransferConfirm-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	var respDataMap venuevo.AGZRResp
	err = xml.Unmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		return &venueResp
	}

	if respDataMap.Info == "0" {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}
	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueAGZR) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueAGZR) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueAGZR) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
