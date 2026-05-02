package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	WUG_Lang = "zh-CN"
)

var (
	WUGDZTransferMap = map[int]string{
		5: "系统错误",
	}
)

type VenueWUGDZ struct {
	VenueConfig conf.WUGDZ
}

func NewWUGDZ(venueConfig *conf.WUGDZ) IVenues {
	return &VenueWUGDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenueWUGDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	paramStr := fmt.Sprintf("host_id=%s&member_id=%s",
		v.VenueConfig.HostId, request.UserName)

	apiUrl := url + "/funds/createplayer/?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-CreateUser-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	respDataMap := map[string]interface{}{}
	err = tool.JsonUnmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	_, ok := respDataMap["status_code"]
	if !ok {
		global.G_LOG.Errorf("%s status_code is empty", apiLog)
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}

	var respData venuevo.WUGDZResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.StatusCode == 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueWUGDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	paramStr := fmt.Sprintf("host_id=%s&member_id=%s",
		v.VenueConfig.HostId, request.UserName)

	apiUrl := url + "/funds/getbalance/?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	respDataMap := map[string]interface{}{}
	err = tool.JsonUnmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	_, ok := respDataMap["status_code"]
	if !ok {
		global.G_LOG.Errorf("%s status_code is empty", apiLog)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}

	var respData venuevo.WUGDZResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// WUGDZ 金额单位分
	if respData.StatusCode == 0 {
		amount := float64(respData.Balance) / 100
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueWUGDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueWUGDZ) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueWUGDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	amount := int64(request.Amount * 100) // 单位为分

	paramStr := fmt.Sprintf("host_id=%s&member_id=%s&txn_id=%s&amount=%v",
		v.VenueConfig.HostId, request.UserName, request.OrderSn, amount)

	apiUrl := url + "/funds/deposit/?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-Deposit-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	respDataMap := map[string]interface{}{}
	err = tool.JsonUnmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	_, ok := respDataMap["status_code"]
	if !ok {
		global.G_LOG.Errorf("%s status_code is empty", apiLog)
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = respStr
		return &venueResp
	}

	var respData venuevo.WUGDZResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.StatusCode == 0 && amount > 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueWUGDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	amount := int64(request.Amount * 100) // 单位为分

	paramStr := fmt.Sprintf("host_id=%s&member_id=%s&txn_id=%s&amount=%v",
		v.VenueConfig.HostId, request.UserName, request.OrderSn, amount)

	apiUrl := url + "/funds/withdraw/?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-Withdraw-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	respDataMap := map[string]interface{}{}
	err = tool.JsonUnmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	_, ok := respDataMap["status_code"]
	if !ok {
		global.G_LOG.Errorf("%s status_code is empty", apiLog)
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = respStr
		return &venueResp
	}

	var respData venuevo.WUGDZResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.StatusCode == 0 && amount > 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

func (v VenueWUGDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	gameId := request.GameCode
	if gameId == "" {
		gameId = "KYS-H5-99999" // 默认进入 777 老虎机
	}

	paramStr := fmt.Sprintf("host_id=%s&game_id=%s&lang=%s&access_token=%s&return_url=%s",
		v.VenueConfig.HostId, gameId, WUG_Lang, request.Token, v.VenueConfig.ReturnUrl)

	apiUrl := url + "/launch/?" + paramStr
	//global.G_LOG.Infof(apiUrl)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-LoginGame-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)

	if err != nil {
		global.G_LOG.Errorf("login error :%v", apiLog)
	}
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = apiUrl
	return &venueResp
}

func (v VenueWUGDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	paramStr := fmt.Sprintf("host_id=%s&txn_id=%s",
		v.VenueConfig.HostId, request.OrderSn)

	apiUrl := url + "/funds/log/?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("WUGDZ-TransferConfirm-param:%s apiUrl:%s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	var respData []*venuevo.WUGDZTransferConfirmResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// 如果订单存在就是成功，没有处理中这个状态
	if len(respData) > 0 {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueWUGDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueWUGDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueWUGDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}
